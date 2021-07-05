package extensions

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/go-logr/logr"
	"github.com/joelanford/helm-operator/pkg/extensions"
	"github.com/pkg/errors"
	centralv1Alpha1 "github.com/stackrox/rox/operator/api/central/v1alpha1"
	utils "github.com/stackrox/rox/operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreV1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/utils/pointer"
)

const (
	defaultPVCName = "stackrox-db"
)

var (
	defaultPVCSize       = resource.MustParse("100Gi")
	errMultipleOwnedPVCs = errors.New("operator is only allowed to have 1 owned PVC")
)

// ReconcilePVCExtension reconciles PVCs created by the operator
func ReconcilePVCExtension(k8sClient kubernetes.Interface) extensions.ReconcileExtension {
	return wrapExtension(reconcilePVC, k8sClient)
}

func reconcilePVC(ctx context.Context, central *centralv1Alpha1.Central, k8sClient kubernetes.Interface, _ func(statusFunc updateStatusFunc), log logr.Logger) error {
	ext := reconcilePVCExtensionRun{
		ctx:        ctx,
		namespace:  central.GetNamespace(),
		pvcClient:  k8sClient.CoreV1().PersistentVolumeClaims(central.GetNamespace()),
		centralObj: central,
		log:        log,
	}

	return ext.Execute()
}

type reconcilePVCExtensionRun struct {
	ctx        context.Context
	namespace  string
	pvcClient  coreV1.PersistentVolumeClaimInterface
	centralObj *centralv1Alpha1.Central
	log        logr.Logger
}

func (r *reconcilePVCExtensionRun) Execute() error {
	if r.centralObj.DeletionTimestamp != nil {
		return r.handleDelete()
	}

	if r.centralObj.Spec.Central.GetHostPath() != "" {
		if r.centralObj.Spec.Central.GetPersistentVolumeClaim() != nil {
			return errors.New("invalid persistence configuration, either hostPath oder persistentVolumeClaim must be set, not both")
		}
		return nil
	}

	pvcConfig := r.centralObj.Spec.Central.GetPersistentVolumeClaim()
	if pvcConfig == nil {
		pvcConfig = &centralv1Alpha1.PersistentVolumeClaim{}
	}
	claimName := pointer.StringPtrDerefOr(pvcConfig.ClaimName, defaultPVCName)

	pvc, err := r.pvcClient.Get(r.ctx, claimName, metav1.GetOptions{})
	if err != nil {
		if !apiErrors.IsNotFound(err) {
			return errors.Wrapf(err, "fetching referenced %s pvc", claimName)
		}
		pvc = nil
	}

	ownedPVC, err := r.getUniqueOwnedPVCs()
	if err != nil {
		return err
	}
	if ownedPVC != nil {
		if ownedPVC.GetName() != claimName && pvc == nil {
			return errors.Errorf(
				"Could not create PVC %q because the operator can only manage 1 PVC for Central. To fix this either reference a manually created PVC or remove the OwnerReference of the %q PVC.", claimName, ownedPVC.GetName())
		}
	}

	// The reconciliation loop should fail if a PVC should be reconciled which is not owned by the operator.
	if pvc != nil && !metav1.IsControlledBy(pvc, r.centralObj) {
		if pvcConfig.StorageClassName != nil || !pvcConfig.Size.IsZero() {
			err := errors.Errorf("Failed reconciling PVC %q. Please remove the storageClassName and size properties from your spec, or change the name to allow the operator to create a new one with a different name.", claimName)
			r.log.Error(err, "failed reconciling PVC")
			return err
		}
		return nil
	}

	if pvc == nil {
		return r.handleCreate(claimName, pvcConfig)
	}

	return r.handleReconcile(pvc, pvcConfig)
}

func (r *reconcilePVCExtensionRun) handleDelete() error {
	ownedPVCs, err := r.getOwnedPVC()
	if err != nil {
		return errors.Wrap(err, "fetching operator owned PVCs")
	}

	for _, ownedPVC := range ownedPVCs {
		utils.RemoveOwnerRef(ownedPVC, r.centralObj)
		r.log.Info(fmt.Sprintf("removed owner reference from %q", ownedPVC.GetName()))

		if _, err := r.pvcClient.Update(r.ctx, ownedPVC, metav1.UpdateOptions{}); err != nil {
			return errors.Wrapf(err, "removing OwnerReference from %s pvc", ownedPVC.GetName())
		}
	}
	return nil
}

func (r *reconcilePVCExtensionRun) handleCreate(claimName string, pvcConfig *centralv1Alpha1.PersistentVolumeClaim) error {
	newPVC := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: claimName,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(r.centralObj, r.centralObj.GroupVersionKind()),
			},
		},

		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: getResourceQuantityOr(pvcConfig.Size, defaultPVCSize),
				},
			},
			StorageClassName: pvcConfig.StorageClassName,
		},
	}

	if _, err := r.pvcClient.Create(r.ctx, newPVC, metav1.CreateOptions{}); err != nil {
		return errors.Wrapf(err, "creating new %s pvc", claimName)
	}
	return nil
}

func (r *reconcilePVCExtensionRun) handleReconcile(existingPVC *corev1.PersistentVolumeClaim, pvcConfig *centralv1Alpha1.PersistentVolumeClaim) error {
	shouldUpdate := false

	if !pvcConfig.Size.IsZero() {
		existingPVC.Spec.Resources.Requests = corev1.ResourceList{
			corev1.ResourceStorage: pvcConfig.Size,
		}
		shouldUpdate = true
	}

	if pointer.StringPtrDerefOr(pvcConfig.StorageClassName, "") != "" && pvcConfig.StorageClassName != existingPVC.Spec.StorageClassName {
		existingPVC.Spec.StorageClassName = pvcConfig.StorageClassName
		shouldUpdate = true
	}

	if shouldUpdate {
		if _, err := r.pvcClient.Update(r.ctx, existingPVC, metav1.UpdateOptions{}); err != nil {
			return errors.Wrapf(err, "updating %s pvc", existingPVC.GetName())
		}
	}
	return nil
}

func getResourceQuantityOr(q resource.Quantity, d resource.Quantity) resource.Quantity {
	if q.IsZero() {
		return d
	}
	return q
}

func (r *reconcilePVCExtensionRun) getOwnedPVC() ([]*corev1.PersistentVolumeClaim, error) {
	pvcList, err := r.pvcClient.List(r.ctx, metav1.ListOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "receiving list PVC list for %s %s", r.centralObj.GroupVersionKind(), r.centralObj.GetName())
	}

	var ownedPVCs []*corev1.PersistentVolumeClaim
	for _, item := range pvcList.Items {
		if metav1.IsControlledBy(&item, r.centralObj) {
			tmp := item
			ownedPVCs = append(ownedPVCs, &tmp)
		}
	}

	return ownedPVCs, nil
}

func (r *reconcilePVCExtensionRun) getUniqueOwnedPVCs() (*corev1.PersistentVolumeClaim, error) {
	pvcList, err := r.getOwnedPVC()
	if err != nil {
		return nil, err
	}

	// If no previously created managed PVC was found everything is ok.
	if len(pvcList) == 0 {
		return nil, nil
	}

	if len(pvcList) > 1 {
		var names []string
		for _, item := range pvcList {
			names = append(names, item.GetName())
		}
		sort.Strings(names)
		return nil, errors.Wrapf(errMultipleOwnedPVCs,
			"multiple owned PVCs were found, please remove not used ones or delete their OwnerReferences. Found PVCs: %s", strings.Join(names, ", "))
	}

	return pvcList[0], nil
}
