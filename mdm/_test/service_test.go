package test

import (
	"testing"

	"github.com/0101binarybard/mdm/service"
	"github.com/stretchr/testify/assert"
)

func TestGetDeviceName(t *testing.T) {
	deviceName := service.GetDeviceName("Test")
	assert.Equal(t, "1.1.1.1", deviceName, "The two words should be the same.")
}
