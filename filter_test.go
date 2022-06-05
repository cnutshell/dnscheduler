package dnscheduler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	dnStores := []*DNStore{
		NewDNStore(0, 10),
		NewDNStore(1, 10),
		NewDNStore(10, 10),
		NewDNStore(5, 5),
	}

	fullFilter := &filterOutFull{}
	var candidates []*DNStore
	for _, store := range dnStores {
		if fullFilter.Filter(store) {
			continue
		}
		candidates = append(candidates, store)
	}
	require.Equal(t, 2, len(candidates))

	emptyFilter := &filterOutEmpty{}
	candidates = candidates[:0]
	for _, store := range dnStores {
		if emptyFilter.Filter(store) {
			continue
		}
		candidates = append(candidates, store)
	}
	require.Equal(t, 3, len(candidates))
}
