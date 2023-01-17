package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidHttpUrl(t *testing.T) {
	assert.True(t, IsValidHttpUrl("https://example.org"))
	assert.True(t, IsValidHttpUrl("https://example.org/?example=true"))
	assert.True(t, IsValidHttpUrl("http://example.org"))
	assert.True(t, IsValidHttpUrl("http://example.org/?example=true"))
	assert.False(t, IsValidHttpUrl("https://"))
	assert.False(t, IsValidHttpUrl("http://"))
	assert.False(t, IsValidHttpUrl(""))
	assert.False(t, IsValidHttpUrl("http:///"))
	assert.False(t, IsValidHttpUrl("https:///"))
}

func TestRemoveParamsFromUrl(t *testing.T) {
	u, err := StripQueryParamFromUrl("https://example.org/example?a=b&c=d")
	assert.Nil(t, err)
	assert.Equal(t, "https://example.org/example", u)
	u, err = StripQueryParamFromUrl("https://example.org/?a=b&c=d")
	assert.Nil(t, err)
	assert.Equal(t, "https://example.org/", u)
	u, err = StripQueryParamFromUrl("https://example.org/example")
	assert.Nil(t, err)
	assert.Equal(t, "https://example.org/example", u)
	u, err = StripQueryParamFromUrl("https://example.org/example?")
	assert.Nil(t, err)
	assert.Equal(t, "https://example.org/example", u)
}
