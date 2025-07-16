package repositories

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
)

type PublicationRepository interface {
	Create(publication *entities.Publication) error
	GetByID(studentID string) (*entities.Publication, error)
	Update(publication *entities.Publication) error
	Delete(studentID string) error
	List(offset, limit int) ([]*entities.Publication, error)
	GetByStudentID(studentID string) ([]*entities.Publication, error)
}