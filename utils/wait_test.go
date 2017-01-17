package utils

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
)

func TestWaitForSpecificOrError(t *testing.T){

	waitInterval := 100*time.Millisecond
	timeout := 1000*time.Millisecond
	times := 0
	err := WaitForSpecificOrError(func() (bool, error) {
		times += 1
		println("times", times)
		if times == 3 {
			return true, nil
		}
		return false, nil
	}, timeout, waitInterval)
	assert.NoError(t, err)
	assert.Equal(t, 3, times)

	times = 0
	err = WaitForSpecificOrError(func() (bool, error) {
		times += 1
		println("times", times)
		if times == 3 {
			return false, errors.New("error")
		}
		return false, nil
	}, timeout, waitInterval)
	assert.Error(t, err)
	assert.Equal(t, 3, times)

	times = 0
	err = WaitForSpecificOrError(func() (bool, error) {
		times += 1
		println("times", times)
		return false, nil
	}, timeout, waitInterval)
	assert.Error(t, err)
	assert.Equal(t, 10, times)
}