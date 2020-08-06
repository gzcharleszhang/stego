package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainFn(t *testing.T) {
	called := false
	executeCmd = func() {
		called = true
	}
	main()
	assert.True(t, called)
}
