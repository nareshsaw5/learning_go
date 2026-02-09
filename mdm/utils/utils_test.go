package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeName(t *testing.T) {

	newName := ChangeName("1.1.1.1")
	assert.Equal(t, "new ip - 10.10.10.10", newName, "Should change new name")
}
