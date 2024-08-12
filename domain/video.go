package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo(resouceID string, filePath string) (*Video, error) {
	video := Video{
		ResourceID: resouceID,
		FilePath:   filePath,
	}
	video.prepare()
	err := video.Validade()
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (video *Video) prepare() {
	video.ID = uuid.NewString()
	video.CreatedAt = time.Now()
}

func (video *Video) Validade() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
