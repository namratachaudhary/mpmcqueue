package mpmcq

import (
	"testing"

	"gotest.tools/assert"
)

func TestMPMCQueueInsert(t *testing.T) {
	mq := NewMPMCQueue(10)
	assert.Equal(t, uint64(16), mq.Capacity())

	err := mq.Put(1)
	assert.NilError(t, err, nil)

	result, err := mq.Pop()
	assert.NilError(t, err, nil)

	assert.Equal(t, 1, result)
}

func TestMPMCQueueInsertMultiple(t *testing.T) {
	var err error

	mq := NewMPMCQueue(10)

	err = mq.Put(1)
	assert.NilError(t, err, nil)

	err = mq.Put(3)
	assert.NilError(t, err, nil)

	result, err := mq.Pop()
	assert.NilError(t, err, nil)

	assert.Equal(t, 1, result)

	result, err = mq.Pop()
	assert.NilError(t, err, nil)

	assert.Equal(t, 3, result)
}

func TestMPMCQueuePopEmpty(t *testing.T) {
	mq := NewMPMCQueue(10)

	_, err := mq.tryPop(1)
	assert.Equal(t, ErrTimeout, err)
}
