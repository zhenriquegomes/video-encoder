package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhenriquegomes/video-encoder/domain"
)

func TestJobIDIsNotAnUUID(t *testing.T) {
	job, err := domain.NewJob("output", "status")
	assert.NotNil(t, job)
	assert.Nil(t, err)

	job.ID = "nouuid"
	err = job.Validade()
	assert.Error(t, err)
}

func TestJobValidation(t *testing.T) {
	job, err := domain.NewJob("output", "status")

	assert.NotNil(t, job)
	assert.Nil(t, err)
}
