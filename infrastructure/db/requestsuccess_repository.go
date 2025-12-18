package db

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)


func (r *thesisJournalRepository) GetRequestSuccessByID(ctx context.Context, id string) (*entities.RequestSuccess, error) {
	requestSuccess := &entities.RequestSuccess{}

	query := `
		SELECT
			STD_CODE,
			SUCCESS_YEAR,
			SUCCESS_SEMESTER,
			NAME_THAI,
			NAME_ENG ,
			THESIS_THAI,
			THESIS_ENG,
			DEGREE,
			REGISTRATION,
			GRADES,
			ADDRESS,
			CREATED,
			MODIFIED
        FROM EGRAD_REQUEST_SUCCESS WHERE STD_CODE = :1`

	err := r.db.GetContext(ctx, requestSuccess, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Student request success not found")
		}
		return nil, fmt.Errorf("failed to get student request success: %w", err)
	}

	return requestSuccess, nil
}

func (r *thesisJournalRepository) ListRequestSuccess(ctx context.Context, limit, offset int) ([]*entities.RequestSuccess, error) {
	requestSuccess := []*entities.RequestSuccess{}

	query := `
		SELECT
			STD_CODE,
			SUCCESS_YEAR,
			SUCCESS_SEMESTER,
			NAME_THAI,
			NAME_ENG ,
			THESIS_THAI,
			THESIS_ENG,
			DEGREE,
			REGISTRATION,
			GRADES,
			ADDRESS,
			CREATED,
			MODIFIED
			FROM EGRAD_REQUEST_SUCCESS
        ORDER BY CREATED DESC
        OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	err := r.db.SelectContext(ctx, &requestSuccess, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list student request success: %w", err)
	}

	return requestSuccess, nil
}
