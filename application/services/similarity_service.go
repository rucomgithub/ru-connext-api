package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
	"fmt"
	"time"
)

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

func (s *thesisJournalService) UpdateThesisSimilarityStatus(ctx context.Context, id string) (*entities.ThesisSimilarity, error) {
	if id == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	thesisJournalSimilarity, err := s.repo.GetSimilarityByID(ctx, id)
	if err != nil {
		return nil, err
	}
	thesisJournalSimilarity.UpdatedAt = time.Now()
	if thesisJournalSimilarity.Status == "requested" {
		thesisJournalSimilarity.Status = "confirmed"
	} else {
		thesisJournalSimilarity.Status = "requested"
	}

	err = s.repo.UpdateSimilarity(ctx, thesisJournalSimilarity)
	if err != nil {
		return nil, err
	}
	return thesisJournalSimilarity, nil
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
