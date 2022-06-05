package dnscheduler

import (
	"fmt"
)

// DNCommand is excuted by DN store
type DNCommand uint32

const (
	LaunchDN DNCommand = iota
	TerminateDN
)

// DNOperator indicates operation for DN store
type DNOperator struct {
	DNStore   *DNStore
	DNCommand DNCommand
}

// DNStore has statistics for DN store
type DNStore struct {
	DNCount  uint8
	Capacity uint8
}

func NewDNStore(count, capacity uint8) *DNStore {
	return &DNStore{
		DNCount:  count,
		Capacity: capacity,
	}
}

// ScheduleDN keeps the expected number of DN shard
func ScheduleDN(expected uint32, stores []*DNStore) (*DNOperator, error) {
	var totalDN uint32
	for _, store := range stores {
		totalDN += uint32(store.DNCount)
	}

	if expected == totalDN {
		return nil, nil
	}

	cmd := LaunchDN
	if totalDN > expected {
		cmd = TerminateDN
	}

	// filter first
	candidates := filterDNStore(stores, constructFilters(cmd))

	// select the most appropriate DN store
	target, err := selectDNStore(candidates, constructScores(cmd))
	if err != nil {
		return nil, err
	}

	return &DNOperator{
		DNStore:   target,
		DNCommand: cmd,
	}, nil
}

func filterDNStore(stores []*DNStore, filters []DNFilter) []*DNStore {
	var candidates []*DNStore
	for _, store := range stores {
		for _, filter := range filters {
			if filter.Filter(store) {
				continue
			}
			candidates = append(candidates, store)
		}
	}
	return candidates
}

func selectDNStore(stores []*DNStore, scores []DNScore) (*DNStore, error) {
	if len(stores) == 0 {
		return nil, fmt.Errorf("empty list of dn store")
	}

	selected := stores[0]
	maxScore := normalize(selected, scores)
	for i := 1; i < len(stores); i++ {
		if maxScore < normalize(stores[i], scores) {
			selected = stores[i]
		}
	}

	return selected, nil
}

// normalize makes score unified
func normalize(store *DNStore, scores []DNScore) uint32 {
	var grade float32
	for _, score := range scores {
		grade += float32(score.Score(store)) / float32(score.Ceiling()) * 100
	}
	return uint32(grade)
}
