package dnscheduler

import (
	"math"
)

const (
	CEILING uint8 = math.MaxUint8
)

// DNScore scores DN store
type DNScore interface {
	Type() string
	Score(store *DNStore) uint8
	Ceiling() uint8
}

// scoreForLeast priority DN store with least DN shard
type scoreForLeast struct {
}

func (s *scoreForLeast) Type() string {
	return "least-dn-first"
}

func (S *scoreForLeast) Score(score *DNStore) uint8 {
	return CEILING - score.DNCount
}

func (s *scoreForLeast) Ceiling() uint8 {
	return CEILING
}

// scoreForMost priority DN store with the most DN shard
type scoreForMost struct {
}

func (s *scoreForMost) Type() string {
	return "most-dn-first"
}

func (S *scoreForMost) Score(score *DNStore) uint8 {
	return score.DNCount
}

func (s *scoreForMost) Ceiling() uint8 {
	return CEILING
}

// constructScores constructs proper scores
func constructScores(cmd DNCommand) []DNScore {
	var scores []DNScore
	switch cmd {
	case TerminateDN:
		scores = []DNScore{
			&scoreForMost{},
		}
	case LaunchDN:
		scores = []DNScore{
			&scoreForLeast{},
		}
	default:
		panic("unexpected command")
	}
	return scores
}
