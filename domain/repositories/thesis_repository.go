package repositories

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
)

type ThesisJournalRepository interface {
	Create(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	GetByID(ctx context.Context, id string) (*entities.ThesisJournal, error)
	GetByStudentID(ctx context.Context, studentID string) (*entities.ThesisJournal, error)
	Update(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*entities.ThesisJournal, error)
}
