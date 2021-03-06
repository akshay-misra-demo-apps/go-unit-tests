package add

import (
	"testing"

	"github.com/misraak/app/api/add"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, add.Add(1, 1, 1))
	assert.Equal(t, 3, add.Add(1, -1, 3))
	assert.Equal(t, 2, add.Add(1/10, 1, 1))
	assert.Equal(t, -1, add.Add(0, 0, -1))
	assert.Equal(t, 3, add.Add(1%2, 1, 1))
	assert.Equal(t, 3, add.Add(1, 1, 1))
}

func TestAddFail(t *testing.T) {
	assert.NotEqual(t, 5, add.Add(1, 1, 1))
}

func TestAddAll(t *testing.T) {
	assert.Equal(t, 5, add.AddAll([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, 3, add.AddAll([]int{1, -1, 3}))
	assert.Equal(t, 2, add.AddAll([]int{1 / 10, 1, 1}))
	assert.Equal(t, 0, add.AddAll([]int{0, 1, -1}))
	assert.Equal(t, 1, add.AddAll([]int{1 % 2, 1, -1}))
	assert.Equal(t, 12, add.AddAll([]int{1, 1, 10}))
}

func TestAddAllFail(t *testing.T) {
	assert.NotEqual(t, 5, add.AddAll([]int{1, 1, 1}))
}

func TestSumPositive(t *testing.T) {
	assert.Equal(t, 5, add.SumPositive([]int{1, 1, 1, 1, 1}))
	assert.Equal(t, -1, add.SumPositive([]int{1, -1, 3}))
	assert.Equal(t, 2, add.SumPositive([]int{1 / 10, 1, 1}))
	assert.Equal(t, -1, add.SumPositive([]int{0, 1, -1}))
	assert.Equal(t, -1, add.SumPositive([]int{1 % 2, 1, -1}))
	assert.Equal(t, 12, add.SumPositive([]int{1, 1, 10}))
}

func TestSumPositiveFail(t *testing.T) {
	assert.NotEqual(t, 1, add.SumPositive([]int{1, 1, -1}))
}
