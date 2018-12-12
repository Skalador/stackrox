package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	// general
	prometheus.MustRegister(
		panicCounter,
		processDedupeCacheHits,
		processDedupeCacheMisses,
		processEnrichmentHits,
		processEnrichmentDrops,
		processEnrichmentLRUCacheSize,
		sensorIndicatorChannelFullCounter,
		totalNetworkFlowsSentCounter,
		totalNetworkFlowsReceivedCounter,
		sensorEvents,
	)
}
