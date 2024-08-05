package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zhenriquegomes/video-encoder/domain"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validade()
	assert.Error(t, err)
}

func TestVideoIDIsNotAnUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "nouuid"
	video.ResourceID = "resource"
	video.FilePath = "filepath"
	video.CreatedAt = time.Now()

	err := video.Validade()
	assert.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.New().String()
	video.ResourceID = "resource"
	video.FilePath = "filepath"
	video.CreatedAt = time.Now()

	err := video.Validade()
	assert.Nil(t, err)
}
