package tags

import (
	messages "github.com/cucumber/messages/go/v21"
)

// ApplyTagFilter will apply a filter string on the
// array of pickles and returned the filtered list.
func ApplyTagFilter(filter string, pickles []*messages.Pickle) []*messages.Pickle {
	_ = "STUB: not implemented"
	return nil
}

// Based on http://behat.readthedocs.org/en/v2.5/guides/6.cli.html#gherkin-filters
func match(filters string, tags []*messages.PickleTag) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func matchAnd(filter string, tags []*messages.PickleTag) bool {
	_ = "STUB: not implemented"
	return false
}

func contains(tags []*messages.PickleTag, tag string) bool { _ = "STUB: not implemented"; return false }
