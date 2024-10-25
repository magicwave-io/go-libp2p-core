package routing

import (
	"context"
	"time"

	cid "github.com/ipfs/go-cid"
)

type ContentListing struct {
	Cid         cid.Cid
	Frequency   uint64
	LastUpdated time.Time
}

type FrequentContentProvider interface {
	GetMostFrequentContentAsync(ctx context.Context) <-chan ContentListing
}
