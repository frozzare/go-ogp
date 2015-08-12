package ogp

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestFetch(t *testing.T) {
	list := Fetch("http://ogp.me")
	assert.Equal(t, "Open Graph protocol", list["title"])
	assert.Equal(t, "website", list["type"])
	assert.Equal(t, "http://ogp.me/", list["url"])
}
