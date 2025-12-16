package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/domain/repositories"
	"RU-Smart-Workspace/ru-smart-api/domain/services"
	"context"
	"fmt"
	"time"
)

type thesisJournalService struct {
	repo repositories.ThesisJournalRepository
}

func NewThesisJournalService(repo repositories.ThesisJournalRepository) services.ThesisJournalService {
	return &thesisJournalService{
		repo: repo,
	}
}

func (s *thesisJournalService) CreateThesisJournal(ctx context.Context, thesisJournal *entities.ThesisJournal) error {
	thesisJournal.CreatedAt = time.Now()
	thesisJournal.UpdatedAt = time.Now()

	thesisJournal.CreatedBy = thesisJournal.StudentID
	thesisJournal.UpdatedBy = thesisJournal.StudentID

	// Generate UUIDs for publications
	for i := range thesisJournal.JournalPublication {
		thesisJournal.JournalPublication[i].StudentID = thesisJournal.StudentID
		thesisJournal.JournalPublication[i].CreatedAt = time.Now()
	}

	// Generate UUID for conference presentation
	if thesisJournal.ConferencePresentation != nil {
		thesisJournal.ConferencePresentation.StudentID = thesisJournal.StudentID
		thesisJournal.ConferencePresentation.CreatedAt = time.Now()
	}

	// Generate UUID for other publication
	if thesisJournal.OtherPublication != nil {
		thesisJournal.OtherPublication.StudentID = thesisJournal.StudentID
		thesisJournal.OtherPublication.CreatedAt = time.Now()
	}

	return s.repo.Create(ctx, thesisJournal)
}

func (s *thesisJournalService) GetThesisJournal(ctx context.Context, id string) (*entities.ThesisJournal, error) {
	if id == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *thesisJournalService) UpdateThesisJournalStatus(ctx context.Context, id string) (*entities.ThesisJournal, error) {
	if id == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	thesisJournal, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	thesisJournal.UpdatedAt = time.Now()
	if thesisJournal.STATUS == "requested" {
		thesisJournal.STATUS = "confirmed"
	} else {
		thesisJournal.STATUS = "requested"
	}

	err = s.repo.Update(ctx, thesisJournal)
	if err != nil {
		return nil, err
	}
	return thesisJournal, nil
}

func (s *thesisJournalService) GetJournalByValidateID(ctx context.Context, studentID string) (*entities.ThesisJournal, error) {
	if studentID == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	return s.repo.GetJournalByValidateID(ctx, studentID) 
}

func (s *thesisJournalService) UpdateThesisJournal(ctx context.Context, thesisjournal *entities.ThesisJournal) error {
	if thesisjournal.StudentID == "" {
		return fmt.Errorf("student ID cannot be empty")
	}
	thesisjournal.UpdatedAt = time.Now()
	return s.repo.Update(ctx, thesisjournal)
}

func (s *thesisJournalService) DeleteThesisJournal(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("student ID cannot be empty")
	}
	return s.repo.Delete(ctx, id)
}

func (s *thesisJournalService) ListThesisJournals(ctx context.Context, limit, offset int) ([]*entities.ThesisJournal, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 { 
		offset = 0
	}
	return s.repo.List(ctx, limit, offset)
}

func (s *thesisJournalService) CreateThesisSimilarity(ctx context.Context, thesisJournal *entities.ThesisSimilarity) error {
	thesisJournal.CreatedAt = time.Now()
	thesisJournal.UpdatedAt = time.Now()

	thesisJournal.CreatedBy = thesisJournal.StudentID
	thesisJournal.UpdatedBy = thesisJournal.StudentID

	return s.repo.CreateSimilarity(ctx, thesisJournal)
}

func (s *thesisJournalService) UpdateThesisSimilarity(ctx context.Context, thesisjournal *entities.ThesisSimilarity) error {
	if thesisjournal.StudentID == "" {
		return fmt.Errorf("student ID cannot be empty")
	}
	thesisjournal.UpdatedAt = time.Now()
	return s.repo.UpdateSimilarity(ctx, thesisjournal)
}

func (s *thesisJournalService) GetSimilarityByStudentID(ctx context.Context, studentID string) (*entities.ThesisSimilarity, error) {
	if studentID == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	return s.repo.GetSimilarityByID(ctx, studentID)
}

func (s *thesisJournalService) ListThesisSimilaritys(ctx context.Context, limit, offset int) ([]*entities.ThesisSimilarity, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.ListSimilarity(ctx, limit, offset)
}

func (s *thesisJournalService) DeleteThesisSimilarity(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("student ID cannot be empty")
	}
	return s.repo.DeleteSimilarity(ctx, id)
}
