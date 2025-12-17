package db

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func (r *thesisJournalRepository) CreateSimilarity(ctx context.Context, thesisJournal *entities.ThesisSimilarity) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert Thesis Journal
	thesisQuery := `
        INSERT INTO EGRAD_THESIS_SIMILARITY (
            STD_CODE, PROGRAM, MAJOR, FACULTY,
            THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
            CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			SIMILARITY,STATUS
        ) VALUES (
            :STD_CODE, :PROGRAM, :MAJOR, :FACULTY,
            :THESIS_TYPE, :THESIS_TITLE_THAI, :THESIS_TITLE_ENGLISH,
            :CREATED_AT, :UPDATED_AT,:CREATED_BY, :UPDATED_BY,
			:SIMILARITY, :STATUS
        )`

	_, err = tx.NamedExecContext(ctx, thesisQuery, thesisJournal)
	if err != nil {
		return fmt.Errorf("failed to insert thesis similarity: %w", err)
	}

	return tx.Commit()
}

func (r *thesisJournalRepository) UpdateSimilarity(ctx context.Context, thesisJournal *entities.ThesisSimilarity) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	thesis, err := r.GetSimilarityByID(ctx, thesisJournal.StudentID)
	if err != nil {
		log.Print("Updating thesis similarity error get row: ", thesisJournal.StudentID)
	}

	// Update Thesis Journal
	log.Print("Updating thesis journal: ", thesis.StudentID)

	thesisQuery := `UPDATE EGRAD_THESIS_SIMILARITY SET 
						PROGRAM = :PROGRAM,
						MAJOR = :MAJOR,
						FACULTY = :FACULTY,
						THESIS_TYPE = :THESIS_TYPE,
						THESIS_TITLE_THAI = :THESIS_TITLE_THAI,
						THESIS_TITLE_ENGLISH = :THESIS_TITLE_ENGLISH,
						UPDATED_AT = :UPDATED_AT,
						CREATED_BY = :CREATED_BY, 
						UPDATED_BY = :UPDATED_BY,
						SIMILARITY = :SIMILARITY, 
						STATUS = :STATUS
					WHERE STD_CODE = :STD_CODE`

	_, err = tx.NamedExecContext(ctx, thesisQuery, thesisJournal)
	if err != nil {
		return fmt.Errorf("failed to update thesis similarity: %w", err)
	}

	return tx.Commit()
}

func (r *thesisJournalRepository) GetSimilarityByID(ctx context.Context, id string) (*entities.ThesisSimilarity, error) {
	thesisSimilarity := &entities.ThesisSimilarity{}

	query := `
        SELECT STD_CODE, PROGRAM, MAJOR, FACULTY,
               THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
               CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			   SIMILARITY, STATUS
        FROM EGRAD_THESIS_SIMILARITY WHERE STD_CODE = :1`

	err := r.db.GetContext(ctx, thesisSimilarity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Thesis similarity not found")
		}
		return nil, fmt.Errorf("failed to get thesis similarity: %w", err)
	}

	return thesisSimilarity, nil
}

func (r *thesisJournalRepository) ListSimilarity(ctx context.Context, limit, offset int) ([]*entities.ThesisSimilarity, error) {
	thesisSimilaritys := []*entities.ThesisSimilarity{}

	query := `
        SELECT STD_CODE, PROGRAM, MAJOR, FACULTY,
               THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
               CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			   SIMILARITY, STATUS
        FROM EGRAD_THESIS_SIMILARITY 
        ORDER BY CREATED_AT DESC
        OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	err := r.db.SelectContext(ctx, &thesisSimilaritys, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list thesis similarity: %w", err)
	}

	return thesisSimilaritys, nil
}

func (r *thesisJournalRepository) DeleteSimilarity(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "DELETE FROM EGRAD_THESIS_SIMILARITY WHERE STD_CODE = :1", id)
	if err != nil {
		return fmt.Errorf("failed to delete similarity journal: %w", err)
	}

	return tx.Commit()
}
