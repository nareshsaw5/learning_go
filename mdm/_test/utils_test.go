package test

import (
	"testing"

	"github.com/0101binarybard/mdm/utils"
	"github.com/stretchr/testify/assert"
)

func TestChangeName(t *testing.T) {

	newName := utils.ChangeName("1.1.1.1")
	assert.Equal(t, "new ip - 10.10.10.10", newName, "Should change new name")
}
