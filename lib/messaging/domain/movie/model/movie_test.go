package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	now := time.Now()
	actual, _ := NewMovie("title", "url", &now)
	assert.Equal(t, "title", actual.Title())
	assert.Equal(t, "url", actual.URL())
	//assert.Equal(t, true, actual.IsPublishedToday(0))
	//assert.Equal(t, false, actual.IsPublishedToday(30))
}

func TestNewMovieNotURL(t *testing.T) {
	now := time.Now()
	_, err := NewMovie("title", "", &now)
	assert.EqualError(t, err, "URLは必須")
}
