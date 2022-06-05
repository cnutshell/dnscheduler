package dnscheduler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScheduleDN(t *testing.T) {
	stores := []*DNStore{
		NewDNStore(1, 2),
		NewDNStore(2, 2),
		NewDNStore(0, 2),
		NewDNStore(4, 5),
	}

	op, err := ScheduleDN(7, stores)
	require.NoError(t, err)
	require.Nil(t, op)

	op, err = ScheduleDN(6, stores)
	require.NoError(t, err)
	require.Equal(t, op.DNCommand, TerminateDN)
	require.Equal(t, int(op.DNStore.DNCount), 4)
	require.Equal(t, int(op.DNStore.Capacity), 5)

	op, err = ScheduleDN(8, stores)
	require.NoError(t, err)
	require.Equal(t, op.DNCommand, LaunchDN)
	require.Equal(t, int(op.DNStore.DNCount), 0)
	require.Equal(t, int(op.DNStore.Capacity), 2)
}
