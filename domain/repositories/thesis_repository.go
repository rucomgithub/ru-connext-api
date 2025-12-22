package repositories

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
)

type ThesisJournalRepository interface {
	Create(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	GetByID(ctx context.Context, id string) (*entities.ThesisJournal, error)
	GetJournalByValidateID(ctx context.Context, studentID string) (*entities.ThesisJournal, error)
	Update(ctx context.Context, thesisjournal *entities.ThesisJournal) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*entities.ThesisJournal, error)

	CreateSimilarity(ctx context.Context, thesisjournal *entities.ThesisSimilarity) error
	UpdateSimilarity(ctx context.Context, thesisjournal *entities.ThesisSimilarity) error
	GetSimilarityByID(ctx context.Context, studentID string) (*entities.ThesisSimilarity, error)
	ListSimilarity(ctx context.Context, limit, offset int) ([]*entities.ThesisSimilarity, error)
	DeleteSimilarity(ctx context.Context, id string) error

	ListRequestSuccess(ctx context.Context, limit, offset int) ([]*entities.RequestSuccess, error)
	GetRequestSuccessByID(ctx context.Context, id string) (*entities.RequestSuccess, error)
	UpdateRequestSuccessStatus(ctx context.Context, thesisjournal *entities.RequestSuccess) error
}
