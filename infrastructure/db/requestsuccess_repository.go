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

	query := `SELECT  
			VSS.ENROLL_YEAR,
			VSS.ENROLL_SEMESTER,
			VSS.STD_CODE,
			VSS.PRENAME_THAI_S,
			VSS.PRENAME_ENG_S,
			VSS.FIRST_NAME,
			VSS.LAST_NAME,
			VSS.FIRST_NAME_ENG,
			VSS.LAST_NAME_ENG,
			VSS.THAI_NAME,
			VSS.PLAN_NO,
			VSS.SEX,
			VSS.REGINAL_NAME,
			VSS.SUBSIDY_NAME,
			VSS.STATUS_NAME_THAI,
			VSS.BIRTH_DATE,
			VSS.STD_ADDR,
			VSS.ADDR_TEL,
			VSS.JOB_POSITION,
			VSS.STD_OFFICE,
			VSS.OFFICE_TEL,
			VSS.DEGREE_NAME,
			VSS.BSC_DEGREE_NO,
			VSS.BSC_DEGREE_THAI_NAME,
			VSS.BSC_INSTITUTE_NO,
			VSS.INSTITUTE_THAI_NAME,
			VSS.CK_CERT_NO,
			VSS.CHK_CERT_NAME_THAI,
			NVL(ES.SUCCESS_YEAR, '-') AS SUCCESS_YEAR,
			NVL(ES.SUCCESS_SEMESTER, '-') AS SUCCESS_SEMESTER,
			NVL(ES.NAME_THAI, '-') AS NAME_THAI,
			NVL(ES.NAME_ENG, '-') AS NAME_ENG,
			NVL(ES.DEGREE, '-') AS DEGREE,
			NVL(ES.THESIS_THAI, '-') AS THESIS_THAI,
			NVL(ES.THESIS_ENG, '-') AS THESIS_ENG,
			NVL(ES.REGISTRATION, '-') AS REGISTRATION,
			NVL(ES.GRADES, '-') AS GRADES,
			NVL(ES.ADDRESS, '-') AS ADDRESS,
			NVL(TO_CHAR(ES.CREATED, 'YYYY-MM-DD HH24:MI:SS'), '-') AS CREATED,
			NVL(TO_CHAR(ES.MODIFIED, 'YYYY-MM-DD HH24:MI:SS'), '-') AS MODIFIED,
			NVL(DECODE(T.THESIS_TITLE_THAI, NULL, '-', T.THESIS_TITLE_THAI), '-') AS THESIS_THAI_TITLE,
			NVL(DECODE(T.THESIS_TITLE_ENGLISH, NULL, '-', T.THESIS_TITLE_ENGLISH), '-') AS THESIS_ENG_TITLE,
			NVL(T.THESIS_TYPE, '-') AS THESIS_TYPE,
			NVL(T.SIMILARITY, -1) AS SIMILARITY 
			from egrad_request_success es
			inner join VM_STUDENT_S VSS on vss.std_code = es.std_code
			inner join egrad_thesis_similarity t on vss.std_code = t.std_code
        ORDER BY es.CREATED DESC
        OFFSET :1 ROWS FETCH NEXT :2 ROWS ONLY`

	err := r.db.SelectContext(ctx, &requestSuccess, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list student request success: %w", err)
	}

	return requestSuccess, nil
}
