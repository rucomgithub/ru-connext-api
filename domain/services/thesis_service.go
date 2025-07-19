package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
)

type ThesisJournalService interface {
	CreateThesisJournal(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	GetThesisJournal(ctx context.Context, id string) (*entities.ThesisJournal, error)
	GetThesisJournalByStudentID(ctx context.Context, studentID string) (*entities.ThesisJournal, error)
	UpdateThesisJournal(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	DeleteThesisJournal(ctx context.Context, id string) error
	ListThesisJournals(ctx context.Context, limit, offset int) ([]*entities.ThesisJournal, error)
}
