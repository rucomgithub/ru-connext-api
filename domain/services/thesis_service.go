package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
)

type ThesisJournalService interface {
	CreateThesisJournal(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	GetThesisJournal(ctx context.Context, id string) (*entities.ThesisJournal, error)
	UpdateThesisJournalStatus(ctx context.Context, id string) (*entities.ThesisJournal, error)
	GetJournalByValidateID(ctx context.Context, studentID string) (*entities.ThesisJournal, error)
	UpdateThesisJournal(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	DeleteThesisJournal(ctx context.Context, id string) error
	ListThesisJournals(ctx context.Context, limit, offset int) ([]*entities.ThesisJournal, error)

	CreateThesisSimilarity(ctx context.Context, thesisjournal *entities.ThesisSimilarity) error
	UpdateThesisSimilarity(ctx context.Context, thesisjournal *entities.ThesisSimilarity) error
	UpdateThesisSimilarityStatus(ctx context.Context, id string) (*entities.ThesisSimilarity, error)
	GetSimilarityByStudentID(ctx context.Context, studentID string) (*entities.ThesisSimilarity, error)
	ListThesisSimilaritys(ctx context.Context, limit, offset int) ([]*entities.ThesisSimilarity, error)
	DeleteThesisSimilarity(ctx context.Context, id string) error
}
