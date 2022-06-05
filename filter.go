package dnscheduler

// DNFilter filter unexpected DN store
type DNFilter interface {
	Type() string
	Filter(store *DNStore) bool
}

// filterOutFull filter out full DN store
type filterOutFull struct {
}

func (f *filterOutFull) Type() string {
	return "filter-full-out"
}

func (f *filterOutFull) Filter(store *DNStore) bool {
	return store.DNCount >= store.Capacity
}

// filterOutEmpty filter out empty DN store
type filterOutEmpty struct {
}

func (f *filterOutEmpty) Type() string {
	return "filter-empty-out"
}

func (f *filterOutEmpty) Filter(store *DNStore) bool {
	return store.DNCount == 0
}

// constructFilters constructs proper filters
func constructFilters(cmd DNCommand) []DNFilter {
	var filters []DNFilter
	switch cmd {
	case TerminateDN:
		filters = []DNFilter{
			&filterOutEmpty{},
		}
	case LaunchDN:
		filters = []DNFilter{
			&filterOutFull{},
		}
	default:
		panic("unexpected command")
	}
	return filters
}
