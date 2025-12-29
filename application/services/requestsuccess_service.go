package services

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
	"fmt"
	"time"
)

func (s *thesisJournalService) GetRequestSuccessByID(ctx context.Context, id string) (*entities.RequestSuccess, error) {
	if id == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	return s.repo.GetRequestSuccessByID(ctx, id)
}

func (s *thesisJournalService) ListRequestSuccesss(ctx context.Context, limit, offset int) ([]*entities.RequestSuccess, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.ListRequestSuccess(ctx, limit, offset)
}

func (s *thesisJournalService) UpdateRequestSuccessStatus(ctx context.Context, id string) (*entities.RequestSuccess, error) {
	if id == "" {
		return nil, fmt.Errorf("student ID cannot be empty")
	}
	thesisRequestSuccess, err := s.repo.GetRequestSuccessByID(ctx, id)
	if err != nil {
		return nil, err
	}

	thesisRequestSuccess.MODIFIED = time.Now()
	if thesisRequestSuccess.STATUS == "requested" {
		thesisRequestSuccess.STATUS = "confirmed"
	} else {
		thesisRequestSuccess.STATUS = "requested"
	}

	err = s.repo.UpdateRequestSuccessStatus(ctx, thesisRequestSuccess)
	if err != nil {
		return nil, err
	}

	return s.repo.GetRequestSuccessByID(ctx, id)
}