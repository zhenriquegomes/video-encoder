package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhenriquegomes/video-encoder/domain"
)

func TestVideoIDIsNotAnUUID(t *testing.T) {
	video, err := domain.NewVideo("resource", "filepath")
	assert.NotNil(t, video)
	assert.Nil(t, err)

	video.ID = "nouuid"
	err = video.Validade()
	assert.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video, err := domain.NewVideo("resource", "filepath")

	assert.NotNil(t, video)
	assert.Nil(t, err)
}
