package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationService_Create(t *testing.T) {
	assert.Nil(t, Relation().Create(1, 3))
	assert.Nil(t, Relation().Create(2, 3))
	assert.Nil(t, Relation().Create(2, 1))
}
