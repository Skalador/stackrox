import React from 'react';
import { Link } from 'react-router-dom';
import { ExpandableRowContent, Table, Tbody, Td, Th, Thead, Tr } from '@patternfly/react-table';
import { gql } from '@apollo/client';
import { min } from 'date-fns';

import useSet from 'hooks/useSet';
import { UseURLSortResult } from 'hooks/useURLSort';
import VulnerabilitySeverityIconText from 'Components/PatternFly/IconText/VulnerabilitySeverityIconText';
import { VulnerabilitySeverity, VulnerabilityState } from 'types/cve.proto';
import VulnerabilityFixableIconText from 'Components/PatternFly/IconText/VulnerabilityFixableIconText';
import { DynamicColumnIcon } from 'Components/DynamicIcon';
import DateDistance from 'Components/DateDistance';
import { getWorkloadEntityPagePath } from '../../utils/searchUtils';

import EmptyTableResults from '../components/EmptyTableResults';
import DeploymentComponentVulnerabilitiesTable, {
    DeploymentComponentVulnerability,
    ImageMetadataContext,
    deploymentComponentVulnerabilitiesFragment,
} from './DeploymentComponentVulnerabilitiesTable';
import {
    getIsSomeVulnerabilityFixable,
    getHighestVulnerabilitySeverity,
} from '../../utils/vulnerabilityUtils';
import PendingExceptionLabelLayout from '../components/PendingExceptionLabelLayout';
import PartialCVEDataAlert from '../../components/PartialCVEDataAlert';

export const deploymentWithVulnerabilitiesFragment = gql`
    ${deploymentComponentVulnerabilitiesFragment}
    fragment DeploymentWithVulnerabilities on Deployment {
        id
        images(query: $query) {
            ...ImageMetadataContext
        }
        imageVulnerabilities(query: $query, pagination: $pagination) {
            vulnerabilityId: id
            cve
            operatingSystem
            summary
            pendingExceptionCount: exceptionCount(requestStatus: $statusesForExceptionCount)
            images(query: $query) {
                imageId: id
                imageComponents(query: $query) {
                    ...DeploymentComponentVulnerabilities
                }
            }
        }
    }
`;

export type DeploymentWithVulnerabilities = {
    id: string;
    images: ImageMetadataContext[];
    imageVulnerabilities: {
        vulnerabilityId: string;
        cve: string;
        operatingSystem: string;
        summary: string;
        pendingExceptionCount: number;
        images: {
            imageId: string;
            imageComponents: DeploymentComponentVulnerability[];
        }[];
    }[];
};

type DeploymentVulnerabilityImageMapping = {
    imageMetadataContext: ImageMetadataContext;
    componentVulnerabilities: DeploymentComponentVulnerability[];
};

function formatVulnerabilityData(deployment: DeploymentWithVulnerabilities): {
    vulnerabilityId: string;
    cve: string;
    operatingSystem: string;
    severity: VulnerabilitySeverity;
    isFixable: boolean;
    discoveredAtImage: Date | null;
    summary: string;
    affectedComponentsText: string;
    images: DeploymentVulnerabilityImageMapping[];
    pendingExceptionCount: number;
}[] {
    // Create a map of image ID to image metadata for easy lookup
    // We use 'Partial' here because there is no guarantee that the image will be found
    const imageMap: Partial<Record<string, ImageMetadataContext>> = {};
    deployment.images.forEach((image) => {
        imageMap[image.id] = image;
    });

    return deployment.imageVulnerabilities.map((vulnerability) => {
        const { vulnerabilityId, cve, operatingSystem, summary, images, pendingExceptionCount } =
            vulnerability;
        // Severity, Fixability, and Discovered date are all based on the aggregate value of all components
        const allVulnerableComponents = vulnerability.images.flatMap((img) => img.imageComponents);
        const allVulnerabilities = allVulnerableComponents.flatMap((c) => c.imageVulnerabilities);
        const highestVulnSeverity = getHighestVulnerabilitySeverity(allVulnerabilities);
        const isFixableInDeployment = getIsSomeVulnerabilityFixable(allVulnerabilities);
        const allDiscoveredDates = allVulnerableComponents
            .flatMap((c) => c.imageVulnerabilities.map((v) => v.discoveredAtImage))
            .filter((d): d is string => d !== null);
        const oldestDiscoveredVulnDate = min(...allDiscoveredDates);
        // TODO This logic is used in many places, could extract to a util
        const uniqueComponents = new Set(allVulnerableComponents.map((c) => c.name));
        const affectedComponentsText =
            uniqueComponents.size === 1
                ? uniqueComponents.values().next().value
                : `${uniqueComponents.size} components`;

        const vulnerabilityImages = images
            .map((img) => ({
                imageMetadataContext: imageMap[img.imageId],
                componentVulnerabilities: img.imageComponents,
            }))
            // filter out values where the vulnerability->image mapping is missing
            .filter(
                (vulnImageMap): vulnImageMap is DeploymentVulnerabilityImageMapping =>
                    !!vulnImageMap.imageMetadataContext
            );

        return {
            vulnerabilityId,
            cve,
            operatingSystem,
            severity: highestVulnSeverity,
            isFixable: isFixableInDeployment,
            discoveredAtImage: oldestDiscoveredVulnDate,
            summary,
            affectedComponentsText,
            images: vulnerabilityImages,
            pendingExceptionCount,
        };
    });
}

export type DeploymentVulnerabilitiesTableProps = {
    deployment: DeploymentWithVulnerabilities;
    getSortParams: UseURLSortResult['getSortParams'];
    isFiltered: boolean;
    vulnerabilityState: VulnerabilityState | undefined; // TODO Make Required when the ROX_VULN_MGMT_UNIFIED_CVE_DEFERRAL feature flag is removed
};

function DeploymentVulnerabilitiesTable({
    deployment,
    getSortParams,
    isFiltered,
    vulnerabilityState,
}: DeploymentVulnerabilitiesTableProps) {
    const expandedRowSet = useSet<string>();

    const vulnerabilities = formatVulnerabilityData(deployment);

    return (
        <Table variant="compact">
            <Thead noWrap>
                <Tr>
                    <Th>{/* Header for expanded column */}</Th>
                    <Th sort={getSortParams('CVE')}>CVE</Th>
                    <Th>OS</Th>
                    <Th>CVE severity</Th>
                    <Th>
                        CVE status
                        {isFiltered && <DynamicColumnIcon />}
                    </Th>
                    <Th>
                        Affected components
                        {isFiltered && <DynamicColumnIcon />}
                    </Th>
                    <Th>First discovered</Th>
                </Tr>
            </Thead>
            {vulnerabilities.length === 0 && <EmptyTableResults colSpan={7} />}
            {vulnerabilities.map((vulnerability, rowIndex) => {
                const {
                    vulnerabilityId,
                    cve,
                    operatingSystem,
                    severity,
                    summary,
                    isFixable,
                    images,
                    affectedComponentsText,
                    discoveredAtImage,
                    pendingExceptionCount,
                } = vulnerability;
                const isExpanded = expandedRowSet.has(vulnerabilityId);

                return (
                    <Tbody key={vulnerabilityId} isExpanded={isExpanded}>
                        <Tr>
                            <Td
                                expand={{
                                    rowIndex,
                                    isExpanded,
                                    onToggle: () => expandedRowSet.toggle(vulnerabilityId),
                                }}
                            />
                            <Td dataLabel="CVE" modifier="nowrap">
                                <PendingExceptionLabelLayout
                                    hasPendingException={pendingExceptionCount > 0}
                                    cve={cve}
                                    vulnerabilityState={vulnerabilityState}
                                >
                                    <Link
                                        to={getWorkloadEntityPagePath(
                                            'CVE',
                                            cve,
                                            vulnerabilityState
                                        )}
                                    >
                                        {cve}
                                    </Link>
                                </PendingExceptionLabelLayout>
                            </Td>
                            <Td modifier="nowrap" dataLabel="OS">
                                {operatingSystem}
                            </Td>
                            <Td modifier="nowrap" dataLabel="Severity">
                                <VulnerabilitySeverityIconText severity={severity} />
                            </Td>
                            <Td modifier="nowrap" dataLabel="CVE Status">
                                <VulnerabilityFixableIconText isFixable={isFixable} />
                            </Td>
                            <Td dataLabel="Affected components">{affectedComponentsText}</Td>
                            <Td modifier="nowrap" dataLabel="First discovered">
                                <DateDistance date={discoveredAtImage} />
                            </Td>
                        </Tr>
                        <Tr isExpanded={isExpanded}>
                            <Td />
                            <Td colSpan={6}>
                                <ExpandableRowContent>
                                    {summary && images.length > 0 ? (
                                        <>
                                            <p className="pf-v5-u-mb-md">{summary}</p>
                                            <DeploymentComponentVulnerabilitiesTable
                                                images={images}
                                                cve={cve}
                                                vulnerabilityState={vulnerabilityState}
                                            />
                                        </>
                                    ) : (
                                        <PartialCVEDataAlert />
                                    )}
                                </ExpandableRowContent>
                            </Td>
                        </Tr>
                    </Tbody>
                );
            })}
        </Table>
    );
}

export default DeploymentVulnerabilitiesTable;
