package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/domain/repositories"
	"RU-Smart-Workspace/ru-smart-api/domain/services"
	"RU-Smart-Workspace/ru-smart-api/errs"
)

type publicationService struct {
	repo repositories.PublicationRepository
}

func NewPublicationService(repo repositories.PublicationRepository) services.PublicationService {
	return &publicationService{repo: repo}
}

func (s *publicationService) CreatePublication(publication *entities.Publication) error {
	_, err := s.repo.GetByID(publication.StudentCode)
	if err == nil {
		return errs.NewBadRequestError("เนื่องจากพบรายการ " + publication.StudentCode)
	}
	return s.repo.Create(publication)
}

func (s *publicationService) GetPublication(id string) (*entities.Publication, error) {
	return s.repo.GetByID(id)
}

func (s *publicationService) UpdatePublication(publication *entities.Publication) error {
	_, err := s.repo.GetByID(publication.StudentCode)
	if err != nil {
		return errs.NewBadRequestError("เนื่องจากไม่พบรายการ " + publication.StudentCode)
	}
	return s.repo.Update(publication)
}

func (s *publicationService) DeletePublication(studentID string) error {
	_, err := s.repo.GetByID(studentID)
	if err != nil {
		return errs.NewBadRequestError("เนื่องจากไม่พบรายการ " + studentID)
	}
	return s.repo.Delete(studentID)
}

func (s *publicationService) ListPublications(offset, limit int) ([]*entities.Publication, error) {
	return s.repo.List(offset, limit)
}

func (s *publicationService) GetStudentPublications(studentID string) ([]*entities.Publication, error) {
	return s.repo.GetByStudentID(studentID)
}
