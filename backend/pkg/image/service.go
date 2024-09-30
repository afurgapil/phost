package image

import (
	"fmt"

	"github.com/afurgapil/phost/backend/pkg/entities"
)

type Service interface {
	CreateImage(image entities.Image) (entities.Image, error)
	GetImageByID(id int) (entities.Image, error)
	DeleteImage(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateImage(image entities.Image) (entities.Image, error) {
	if image.ID != 0 || image.Value == "" {
		return entities.Image{}, fmt.Errorf("image ID and name cannot be empty")
	}
	return s.repo.CreateImage(image)
}

func (s *service) GetImageByID(id int) (entities.Image, error) {
	return s.repo.GetImageByID(id)
}

func (s *service) DeleteImage(id int) error {
	return s.repo.DeleteImage(id)
}
