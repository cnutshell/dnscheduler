package dnscheduler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScore(t *testing.T) {
	dnStores := []*DNStore{
		NewDNStore(0, 10),
		NewDNStore(2, 10),
		NewDNStore(3, 10),
	}

	leastScorer := &scoreForLeast{}
	selected := dnStores[0]
	for i := 1; i < len(dnStores); i++ {
		if leastScorer.Score(selected) <
			leastScorer.Score(dnStores[i]) {
			selected = dnStores[i]
		}
	}
	require.Equal(t, int(selected.DNCount), 0)

	mostScorer := &scoreForMost{}
	selected = dnStores[0]
	for i := 1; i < len(dnStores); i++ {
		if mostScorer.Score(selected) <
			leastScorer.Score(dnStores[i]) {
			selected = dnStores[i]
		}
	}
	require.Equal(t, int(selected.DNCount), 3)
}
