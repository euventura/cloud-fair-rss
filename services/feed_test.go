package services

import "testing"

func TestFetch(t *testing.T) {

	f := newFeed()
	f.fetch()

	t.Fail()

}
