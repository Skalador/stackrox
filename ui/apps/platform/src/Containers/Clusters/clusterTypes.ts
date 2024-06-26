export type SensorHealthStatus = 'HEALTHY' | 'UNHEALTHY' | 'DEGRADED' | 'UNINITIALIZED';

export type ClusterHealthItemStatus =
    | 'HEALTHY'
    | 'UNHEALTHY'
    | 'DEGRADED'
    | 'UNINITIALIZED'
    | 'UNAVAILABLE';

export type ClusterHealthStatus = {
    admissionControlHealthStatus?: ClusterHealthItemStatus;
    admissionControlHealthInfo?: {
        totalDesiredPods: number;
        totalReadyPods: number;
        statusErrors: string[];
    };
    collectorHealthStatus?: ClusterHealthItemStatus;
    collectorHealthInfo?: {
        version: string;
        totalDesiredPods: number;
        totalReadyPods: number;
        totalRegisteredNodes: number;
        statusErrors: string[];
    };
    scannerHealthStatus?: ClusterHealthItemStatus;
    scannerHealthInfo?: {
        totalDesiredAnalyzerPods: number;
        totalReadyAnalyzerPods: number;
        totalDesiredDbPods: number;
        totalReadyDbPods: number;
        statusErrors: string[];
    };
    sensorHealthStatus: SensorHealthStatus;
    overallHealthStatus: SensorHealthStatus;
    healthInfoComplete: boolean;
    lastContact: string; // ISO 8601
};

export type ClusterHealthItem = 'collector' | 'sensor' | 'admissionControl' | 'scanner';

export type SensorUpgradeStatus = {
    upgradability: string;
    upgradabilityStatusReason: string;
    mostRecentProcess: {
        active: boolean;
        progress: {
            upgradeState: string;
            upgradeStatusDetail: string;
        };
        type: string;
    };
};

export type DynamicConfig = {
    registryOverride: string;
    admissionControllerConfig: {
        disableBypass: boolean;
        enabled: boolean;
        enforceOnUpdates: boolean;
        scanInline: boolean;
        timeoutSeconds: number;
    };
};

export type CentralEnv = {
    kernelSupportAvailable?: boolean;
    successfullyFetched?: boolean;
};

export type CertExpiryStatus = {
    sensorCertExpiry: string; // ISO 8601
    sensorCertNotBefore: string; // ISO 8601
};

export type ClusterStatus = {
    sensorVersion: string;
    providerMetadata: {
        region: string;
    };
    orchestratorMetadata: {
        version: string;
        buildDate: string;
    };
    upgradeStatus: SensorUpgradeStatus;
    certExpiryStatus: CertExpiryStatus;
};
