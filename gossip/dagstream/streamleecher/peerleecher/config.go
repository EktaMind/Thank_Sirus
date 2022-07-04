package peerleecher

import (
	"time"

	"github.com/EktaMind/Thank_Sirus/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}
