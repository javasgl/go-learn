package search

type defaultMatcher struct{}

func init() {

	var matcher defaultMatcher
	Register("default", defaultMatcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
