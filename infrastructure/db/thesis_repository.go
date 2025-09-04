package db

import (
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"RU-Smart-Workspace/ru-smart-api/domain/repositories"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
)

type thesisJournalRepository struct {
	db *sqlx.DB
}

func NewThesisJournalRepository(db *sqlx.DB) repositories.ThesisJournalRepository {
	return &thesisJournalRepository{
		db: db,
	}
}

func (r *thesisJournalRepository) Create(ctx context.Context, thesisJournal *entities.ThesisJournal) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert Thesis Journal
	thesisQuery := `
        INSERT INTO EGRAD_THESIS (
            STD_CODE, PROGRAM, MAJOR, FACULTY,
            THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
            CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			SIMILARITY
        ) VALUES (
            :STD_CODE, :PROGRAM, :MAJOR, :FACULTY,
            :THESIS_TYPE, :THESIS_TITLE_THAI, :THESIS_TITLE_ENGLISH,
            :CREATED_AT, :UPDATED_AT,:CREATED_BY, :UPDATED_BY,
			:SIMILARITY
        )`

	_, err = tx.NamedExecContext(ctx, thesisQuery, thesisJournal)
	if err != nil {
		return fmt.Errorf("failed to insert thesis journal: %w", err)
	}

	// Insert publications
	if (thesisJournal.JournalPublication != nil && len(thesisJournal.JournalPublication) > 0) {
		publicationQuery := `
            INSERT INTO EGRAD_PUBLICATIONS (
                STD_CODE, TYPE, ARTICLE_TITLE, JOURNAL_NAME, COUNTRY, STATUS,
                YEAR, VOLUME, ISSUE, MONTH, PUBLISH_YEAR, PAGE_FROM, PAGE_TO,
                PUBLISH_LEVEL, TCI_GROUP, CREATED_AT
            ) VALUES (
                :STD_CODE, :TYPE, :ARTICLE_TITLE, :JOURNAL_NAME, :COUNTRY, :STATUS,
                :YEAR, :VOLUME, :ISSUE, :MONTH, :PUBLISH_YEAR, :PAGE_FROM, :PAGE_TO,
                :PUBLISH_LEVEL, :TCI_GROUP, :CREATED_AT
            )`

		for _, pub := range thesisJournal.JournalPublication {
			_, err = tx.NamedExecContext(ctx, publicationQuery, pub)
			if err != nil {
				return fmt.Errorf("failed to insert publication: %w", err)
			}
		}
	}

	// Insert conference presentation
	if thesisJournal.ConferencePresentation != nil {
		confQuery := `
            INSERT INTO EGRAD_CONFERENCE_PRESENTATIONS (
                STD_CODE, TYPE, ARTICLE_TITLE, CONFERENCE_NAME, CONFERENCE_DATE,
                ORGANIZER, LOCATION, COUNTRY, STATUS, PAGE_FROM, PAGE_TO, CREATED_AT
            ) VALUES (
                :STD_CODE, :TYPE, :ARTICLE_TITLE, :CONFERENCE_NAME, :CONFERENCE_DATE,
                :ORGANIZER, :LOCATION, :COUNTRY, :STATUS, :PAGE_FROM, :PAGE_TO, :CREATED_AT
            )`

		_, err = tx.NamedExecContext(ctx, confQuery, thesisJournal.ConferencePresentation)
		if err != nil {
			return fmt.Errorf("failed to insert conference presentation: %w", err)
		}
	}

	// Insert other publication
	if thesisJournal.OtherPublication != nil {
		otherQuery := `
            INSERT INTO EGRAD_OTHER_PUBLICATIONS (
                STD_CODE, ARTICLE_TITLE, SOURCE_TYPE, SOURCE_DETAIL, CREATED_AT
            ) VALUES (
                :STD_CODE, :ARTICLE_TITLE, :SOURCE_TYPE, :SOURCE_DETAIL, :CREATED_AT
            )`

		_, err = tx.NamedExecContext(ctx, otherQuery, thesisJournal.OtherPublication)
		if err != nil {
			return fmt.Errorf("failed to insert other publication: %w", err)
		}
	}

	return tx.Commit()
}

func (r *thesisJournalRepository) GetByID(ctx context.Context, id string) (*entities.ThesisJournal, error) {
	thesisJournal := &entities.ThesisJournal{}

	query := `
        SELECT STD_CODE, PROGRAM, MAJOR, FACULTY,
               THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
               CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			   SIMILARITY
        FROM EGRAD_THESIS WHERE STD_CODE = :1`

	err := r.db.GetContext(ctx, thesisJournal, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Thesis journal not found")
		}
		return nil, fmt.Errorf("failed to get thesis journal: %w", err)
	}

	// Load publications
	publications := []entities.JournalPublication{}
	pubQuery := `
        SELECT ID,STD_CODE, TYPE, ARTICLE_TITLE, JOURNAL_NAME, COUNTRY, STATUS,
               YEAR, VOLUME, ISSUE, MONTH, PUBLISH_YEAR, PAGE_FROM, PAGE_TO,
               PUBLISH_LEVEL, TCI_GROUP, CREATED_AT
        FROM EGRAD_PUBLICATIONS WHERE STD_CODE = :1`

	err = r.db.SelectContext(ctx, &publications, pubQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to get publications: %w", err)
	}
	thesisJournal.JournalPublication = publications

	// Load conference presentation
	confPres := &entities.ConferencePresentation{}
	confQuery := `
        SELECT STD_CODE, TYPE, ARTICLE_TITLE, CONFERENCE_NAME, CONFERENCE_DATE,
               ORGANIZER, LOCATION, COUNTRY, STATUS, PAGE_FROM, PAGE_TO, CREATED_AT
        FROM EGRAD_CONFERENCE_PRESENTATIONS WHERE STD_CODE = :1`

	err = r.db.GetContext(ctx, confPres, confQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to get conference presentation:::::::: %w", err)
	}
	if err != sql.ErrNoRows {
		thesisJournal.ConferencePresentation = confPres
	}

	// Load other publication
	otherPub := &entities.OtherPublication{}
	otherQuery := `
        SELECT STD_CODE, ARTICLE_TITLE, SOURCE_TYPE, SOURCE_DETAIL, CREATED_AT
        FROM EGRAD_OTHER_PUBLICATIONS WHERE STD_CODE = :1`

	err = r.db.GetContext(ctx, otherPub, otherQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to get other publication: %w", err)
	}
	if err != sql.ErrNoRows {
		thesisJournal.OtherPublication = otherPub
	}

	return thesisJournal, nil
}

func (r *thesisJournalRepository) GetByStudentID(ctx context.Context, studentID string) (*entities.ThesisJournal, error) {
	thesisJournal := &entities.ThesisJournal{}

	query := ` 
        SELECT STD_CODE, PROGRAM, MAJOR, FACULTY,
               THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
               CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			   SIMILARITY
        FROM EGRAD_THESIS WHERE STD_CODE = :1`

	err := r.db.GetContext(ctx, thesisJournal, query, studentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("thesis journal not found")
		}
		return nil, fmt.Errorf("failed to get thesis journal: %w", err)
	}

	return r.GetByID(ctx, thesisJournal.StudentID)
}

func (r *thesisJournalRepository) Update(ctx context.Context, thesisJournal *entities.ThesisJournal) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	thesis , err := r.GetByID(ctx, thesisJournal.StudentID)
	if err != nil {
		log.Print("Updating thesis journal error get Thesis: ", thesisJournal.StudentID)
	}

	// Update Thesis Journal
	log.Print("Updating thesis journal: ", thesisJournal.StudentID)
	thesisQuery := `UPDATE EGRAD_THESIS SET 
						PROGRAM = :PROGRAM,
						MAJOR = :MAJOR,
						FACULTY = :FACULTY,
						THESIS_TYPE = :THESIS_TYPE,
						THESIS_TITLE_THAI = :THESIS_TITLE_THAI,
						THESIS_TITLE_ENGLISH = :THESIS_TITLE_ENGLISH,
						UPDATED_AT = :UPDATED_AT,
						CREATED_BY = :CREATED_BY, 
						UPDATED_BY = :UPDATED_BY,
						SIMILARITY = :SIMILARITY
					WHERE STD_CODE = :STD_CODE`

	_, err = tx.NamedExecContext(ctx, thesisQuery, thesisJournal)
	if err != nil {
		return fmt.Errorf("failed to update thesis journal: %w", err)
	}

	// Insert publications
	log.Print(
		"Updating thesis journal publications for student ID: ", thesisJournal.StudentID,
		" with ", len(thesisJournal.JournalPublication), " publications",
	)
	if len(thesisJournal.JournalPublication) > 0 {
		publicationQuery := `
            UPDATE EGRAD_PUBLICATIONS 
			SET 
				TYPE = :TYPE, 
				ARTICLE_TITLE= :ARTICLE_TITLE, 
				JOURNAL_NAME= :JOURNAL_NAME, 
				COUNTRY= :COUNTRY, 
				STATUS= :STATUS,
                YEAR=:YEAR, 
				VOLUME= :VOLUME, 
				ISSUE= :ISSUE, 
				MONTH= :MONTH, 
				PUBLISH_YEAR= :PUBLISH_YEAR, 
				PAGE_FROM= :PAGE_FROM, 
				PAGE_TO= :PAGE_TO,
                PUBLISH_LEVEL=:PUBLISH_LEVEL, 
				TCI_GROUP= :TCI_GROUP
			WHERE ID = :ID`

		for _, pub := range thesisJournal.JournalPublication {
			log.Print(pub.Id)
			if (pub.Id == "" && len(thesis.JournalPublication) < 2) {
				log.Print(
					"Updating thesis journal publications for insert publication null ID: ", thesisJournal.StudentID, len(thesis.JournalPublication),
					" with ", len(thesisJournal.JournalPublication), " publications",
				)
				pub.StudentID = thesisJournal.StudentID
				publicationQuery := `
				INSERT INTO EGRAD_PUBLICATIONS (
					STD_CODE, TYPE, ARTICLE_TITLE, JOURNAL_NAME, COUNTRY, STATUS,
					YEAR, VOLUME, ISSUE, MONTH, PUBLISH_YEAR, PAGE_FROM, PAGE_TO,
					PUBLISH_LEVEL, TCI_GROUP
				) VALUES (
					:STD_CODE, :TYPE, :ARTICLE_TITLE, :JOURNAL_NAME, :COUNTRY, :STATUS,
					:YEAR, :VOLUME, :ISSUE, :MONTH, :PUBLISH_YEAR, :PAGE_FROM, :PAGE_TO,
					:PUBLISH_LEVEL, :TCI_GROUP
				)`
				_, err = tx.NamedExecContext(ctx, publicationQuery, pub)
				if err != nil {
					return fmt.Errorf("failed update publication to insert publication: %w", err)
				}
			}

			_, err = tx.NamedExecContext(ctx, publicationQuery, pub)
			if err != nil {
				return fmt.Errorf("failed to update publication: %w", err.Error())
			}
		}
	} else {
		log.Print(
			"Updating thesis publications for delete publications for student ID: ", thesisJournal.StudentID,
		)
		_, err = tx.ExecContext(ctx, "DELETE FROM EGRAD_PUBLICATIONS WHERE STD_CODE = :1", thesisJournal.StudentID)
		if err != nil {
			return fmt.Errorf("failed to delete publications: %w", err)
		}
	}

	// update conference presentation

	if thesisJournal.ConferencePresentation != nil {
		log.Print(
			"Updating thesis journal conference presentation for student ID: ", thesisJournal.StudentID,thesisJournal.ConferencePresentation,
		)
		confQuery := `UPDATE EGRAD_CONFERENCE_PRESENTATIONS
					SET
						TYPE = :TYPE,
						ARTICLE_TITLE = :ARTICLE_TITLE,
						CONFERENCE_NAME = :CONFERENCE_NAME,
						CONFERENCE_DATE = :CONFERENCE_DATE,
						ORGANIZER = :ORGANIZER,
						LOCATION = :LOCATION,
						COUNTRY = :COUNTRY,
						STATUS = :STATUS,
						PAGE_FROM = :PAGE_FROM,
						PAGE_TO = :PAGE_TO
					WHERE
						STD_CODE = :STD_CODE`

		res, err := tx.NamedExecContext(ctx, confQuery, thesisJournal.ConferencePresentation)
		if err != nil {

			return fmt.Errorf("failed to update conference presentation: %w", err)
		}

		rows, err := res.RowsAffected()

		if rows < 1 {
			confQuery := `
				INSERT INTO EGRAD_CONFERENCE_PRESENTATIONS (
					STD_CODE, TYPE, ARTICLE_TITLE, CONFERENCE_NAME, CONFERENCE_DATE,
					ORGANIZER, LOCATION, COUNTRY, STATUS, PAGE_FROM, PAGE_TO
				) VALUES (
					:STD_CODE, :TYPE, :ARTICLE_TITLE, :CONFERENCE_NAME, :CONFERENCE_DATE,
					:ORGANIZER, :LOCATION, :COUNTRY, :STATUS, :PAGE_FROM, :PAGE_TO
				)`

			_, err = tx.NamedExecContext(ctx, confQuery, thesisJournal.ConferencePresentation)
			if err != nil {
				return fmt.Errorf("failed to insert conference presentation: %w", err)
			}
		}
	} else {
		log.Print(
			"Updating thesis journal delete conference presentation for student ID: ", thesisJournal.StudentID,
		)
		_, err = tx.ExecContext(ctx, "DELETE FROM egrad_conference_presentations WHERE STD_CODE = :1", thesisJournal.StudentID)
		if err != nil {
			return fmt.Errorf("failed to delete conference presentations: %w", err)
		}
	}

	// update other publication
	
	if thesisJournal.OtherPublication != nil {
		log.Print(
			"Updating thesis journal other publication for student ID: ", thesisJournal.OtherPublication.StudentID,
			" with ", len(thesisJournal.OtherPublication.ArticleTitle), " OtherPublication",
		)
		otherQuery := `UPDATE EGRAD_OTHER_PUBLICATIONS
						SET
							ARTICLE_TITLE = :ARTICLE_TITLE,
							SOURCE_TYPE = :SOURCE_TYPE,
							SOURCE_DETAIL = :SOURCE_DETAIL
						WHERE
							STD_CODE = :STD_CODE`

		res, err := tx.NamedExecContext(ctx, otherQuery, thesisJournal.OtherPublication)
		if err != nil {
			return fmt.Errorf("failed to update other publication: %w", err)
		}

		rows, err := res.RowsAffected()

		if rows < 1 {
			otherQuery := `
				INSERT INTO EGRAD_OTHER_PUBLICATIONS (
					STD_CODE, ARTICLE_TITLE, SOURCE_TYPE, SOURCE_DETAIL
				) VALUES (
					:STD_CODE, :ARTICLE_TITLE, :SOURCE_TYPE, :SOURCE_DETAIL
				)`

			_, err = tx.NamedExecContext(ctx, otherQuery, thesisJournal.OtherPublication)
			if err != nil {
				return fmt.Errorf("failed to insert other publication: %w", err)
			}
		}
	} else {
		log.Print(
			"Updating thesis journal delete other publication for student ID: ", thesisJournal.StudentID,
		)
		_, err = tx.ExecContext(ctx, "DELETE FROM egrad_other_publications WHERE STD_CODE = :1", thesisJournal.StudentID)
		if err != nil {
			return fmt.Errorf("failed to delete other publications: %w", err)
		}
	}

	return tx.Commit()
}

func (r *thesisJournalRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Delete related records first
	_, err = tx.ExecContext(ctx, "DELETE FROM egrad_publications WHERE STD_CODE = :1", id)
	if err != nil {
		return fmt.Errorf("failed to delete publications: %w", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM egrad_conference_presentations WHERE STD_CODE = :1", id)
	if err != nil {
		return fmt.Errorf("failed to delete conference presentations: %w", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM egrad_other_publications WHERE STD_CODE = :1", id)
	if err != nil {
		return fmt.Errorf("failed to delete other publications: %w", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM egrad_thesis WHERE STD_CODE = :1", id)
	if err != nil {
		return fmt.Errorf("failed to delete thesis journal: %w", err)
	}

	return tx.Commit()
}

func (r *thesisJournalRepository) List(ctx context.Context, limit, offset int) ([]*entities.ThesisJournal, error) {
	thesisJournals := []*entities.ThesisJournal{}

	query := `
        SELECT STD_CODE, PROGRAM, MAJOR, FACULTY,
               THESIS_TYPE, THESIS_TITLE_THAI, THESIS_TITLE_ENGLISH,
               CREATED_AT, UPDATED_AT,CREATED_BY, UPDATED_BY,
			   SIMILARITY
        FROM EGRAD_THESIS
        ORDER BY CREATED_AT DESC
        OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	err := r.db.SelectContext(ctx, &thesisJournals, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list thesis journal: %w", err)
	}

	return thesisJournals, nil
}
