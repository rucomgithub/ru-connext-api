package masterrepo

import (
	"fmt"
    "RU-Smart-Workspace/ru-smart-api/domain/entities"
)

func (r *studentRepoDB) GetStudentSuccess(studentCode string) (student *StudentSuccessRepo, err error) { 

	student_info := StudentSuccessRepo{}

	query := `select g.STD_CODE,
    g.name_thai,
    g.name_eng,
    g.year,
    g.semester,
    g.CURR_NAME,
    NVL(g.CURR_NAME,'-') CURR_NAME,
    NVL(g.CURR_ENG,'-') CURR_ENG,
    g.THAI_NAME,
    g.ENG_NAME,
    NVL(g.MAJOR_NAME,'-') MAJOR_NAME,
    NVL(g.MAJOR_ENG,'-') MAJOR_ENG,
    NVL(g.MAIN_MAJOR_THAI,'-') MAIN_MAJOR_THAI,
    NVL(g.MAIN_MAJOR_ENG,'-') MAIN_MAJOR_ENG,
    g.GPA,
    NVL(g.PLAN,'-') PLAN,
    g.CONFERENCE_NO,
    g.SERIAL_NO,
    TO_CHAR(TO_DATE(g.conference_date,'DD/MM/YYYY'),'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') conference_date,
    TO_CHAR(TO_DATE(g.admit_date,'DD/MM/YYYY'),'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') admit_date,
    UPPER(TO_CHAR(TO_DATE(g.admit_date, 'DD/MM/YYYY'), 'Month')) || TO_CHAR(TO_DATE(g.admit_date, 'DD/MM/YYYY'), 'DD, YYYY') AS admit_date_en,
    TO_CHAR(TO_DATE(g.graduated_date,'DD/MM/YYYY'),'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') graduated_date,
    UPPER(TO_CHAR(TO_DATE(g.graduated_date, 'DD/MM/YYYY'), 'Month')) || TO_CHAR(TO_DATE(g.graduated_date, 'DD/MM/YYYY'), 'DD, YYYY') AS graduated_date_en,
    TO_CHAR(TO_DATE(g.conference_date,'DD/MM/YYYY')+7,'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') confirm_date,
    NVL(a.E_MAIL,'-') EMAIL,
    NVL(a.MOBILE_TELEPHONE,'-') MOBILE
    from  vm_graduate g
    left join vm_student_address a on g.std_code = a.std_code 
    where  g.std_code = :param1`

	fmt.Printf("success: %s \n", studentCode)

	err = r.oracle_db_dbg.Get(&student_info, query, studentCode)

	if err == nil {
		student = &student_info
		return student, nil
	}

	return nil, err
}

// db is *sqlx.DB
func (r *studentRepoDB) CreateRequestSuccess(row *entities.RequestSuccessRepo) error {
	 query := `
            INSERT INTO DBGMIS00.EGRAD_REQUEST_SUCCESS
            (STD_CODE, SUCCESS_YEAR, SUCCESS_SEMESTER,
            NAME_THAI, NAME_ENG, DEGREE,
            THESIS_THAI, THESIS_ENG, REGISTRATION, GRADES, ADDRESS)
            VALUES
            (:STD_CODE, :SUCCESS_YEAR, :SUCCESS_SEMESTER,
            :NAME_THAI, :NAME_ENG, :DEGREE,
            :THESIS_THAI, :THESIS_ENG, :REGISTRATION, :GRADES, :ADDRESS)
            `
    result, err := r.oracle_db_dbg.NamedExec(query, row)
     
    if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	if err != nil {
		return err
	}
    return err
}

// db is *sqlx.DB
func (r *studentRepoDB) UpdateRequestSuccess(row *entities.RequestSuccessRepo) error {
    query := `
           UPDATE DBGMIS00.EGRAD_REQUEST_SUCCESS 
           SET SUCCESS_YEAR = :SUCCESS_YEAR, SUCCESS_SEMESTER = :SUCCESS_SEMESTER,
           NAME_THAI = :NAME_THAI, NAME_ENG = :NAME_ENG, DEGREE = :DEGREE,
           THESIS_THAI = :THESIS_THAI, THESIS_ENG = :THESIS_ENG, 
           REGISTRATION = :REGISTRATION, GRADES = :GRADES, ADDRESS = :ADDRESS
           WHERE STD_CODE = :STD_CODE
           `
    result , err := r.oracle_db_dbg.NamedExec(query, row)

    if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	if err != nil {
		return err
	}

   return err
}

func (r *studentRepoDB) GetStudentRequestSuccess(studentCode string) (student *StudentRequestSuccessRepo, err error) { 

	student_info := StudentRequestSuccessRepo{}

	query := `select  
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
    from VM_STUDENT_S VSS 
    left join egrad_thesis_similarity t on vss.std_code = t.std_code
    left join egrad_request_success ES on vss.std_code = es.std_code
    where VSS.STD_CODE = :param1`

	fmt.Printf("requestsuccess: %s \n", studentCode)

	err = r.oracle_db_dbg.Get(&student_info, query, studentCode)

	if err == nil {
		student = &student_info
		return student, nil
	}

    fmt.Printf("error requestsuccess: %s \n", err.Error())

	return nil, err
}