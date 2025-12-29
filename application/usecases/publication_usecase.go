package usecases

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/domain/services"

	"gopkg.in/validator.v2"
)

type PublicationUseCase struct {
	service services.PublicationService
}

func NewPublicationUseCase(service services.PublicationService) *PublicationUseCase {
	return &PublicationUseCase{service: service}
}

func (uc *PublicationUseCase) CreatePublication(publication *entities.Publication) error {
	if err := validator.Validate(publication); err != nil {
		return err
	}
	return uc.service.CreatePublication(publication)
}

func (uc *PublicationUseCase) GetPublication(id string) (*entities.Publication, error) {
	return uc.service.GetPublication(id)
}

func (uc *PublicationUseCase) UpdatePublication(publication *entities.Publication) error {
	if err := validator.Validate(publication); err != nil {
		return err
	}
	return uc.service.UpdatePublication(publication)
}

func (uc *PublicationUseCase) DeletePublication(studentID string) error {
	return uc.service.DeletePublication(studentID)
}

func (uc *PublicationUseCase) ListPublications(offset, limit int) ([]*entities.Publication, error) {
	return uc.service.ListPublications(offset, limit)
}

func (uc *PublicationUseCase) GetStudentPublications(studentID string) ([]*entities.Publication, error) {
	return uc.service.GetStudentPublications(studentID)
}
