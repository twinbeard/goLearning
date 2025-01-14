package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "1 + 1 should be 2") // 2 is the expected result
}
