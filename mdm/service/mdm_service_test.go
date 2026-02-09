package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDeviceName(t *testing.T) {
	deviceName := GetDeviceName("Test")
	assert.Equal(t, "1.1.1.1", deviceName, "The two words should be the same.")
}
