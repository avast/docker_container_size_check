package main

import (
	"github.com/alecthomas/units"
	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevel(t *testing.T) {
	container := types.Container{}
	container.Labels = make(map[string]string)
	container.Labels["TEST"] = "10MB"
	container.Labels["INVALID"] = "10XX"

	assert.Equal(t, int64(10), level(container, 10, ""), "global level setting")
	assert.Equal(t, int64(10*units.Mebibyte), level(container, 10, "TEST"), "override level setting")
	assert.Equal(t, int64(10), level(container, 10, "INVALID"), "invalid override units")
}