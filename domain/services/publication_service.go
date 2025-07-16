package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
)

type PublicationService interface {
	CreatePublication(publication *entities.Publication) error
	GetPublication(studentID string) (*entities.Publication, error)
	UpdatePublication(publication *entities.Publication) error
	DeletePublication(studentID string) error
	ListPublications(offset, limit int) ([]*entities.Publication, error)
	GetStudentPublications(studentID string) ([]*entities.Publication, error)
}