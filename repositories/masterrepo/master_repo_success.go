package masterrepo

import (
	"fmt"
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
    NVL(ES.ID, -1) AS ID,
    NVL(ES.SUCCESS_YEAR, '-') AS SUCCESS_YEAR,
    NVL(ES.SUCCESS_SEMESTER, '-') AS SUCCESS_SEMESTER,
    NVL(ES.NAME_THAI_CONFIRM, '-') AS NAME_THAI_CONFIRM,
    NVL(ES.NAME_ENG_CONFIRM, '-') AS NAME_ENG_CONFIRM,
    NVL(ES.THESIS_THAI_CONFIRM, '-') AS THESIS_THAI_CONFIRM,
    NVL(ES.THESIS_ENG_CONFIRM, '-') AS THESIS_ENG_CONFIRM,
    NVL(ES.DEGREE_CONFIRM, '-') AS DEGREE_CONFIRM,
    NVL(ES.CHECKDEGREE, '-') AS CHECKDEGREE,
    NVL(ES.CHECKREGISTER, '-') AS CHECKREGISTER,
    NVL(ES.CHECKGPA, '-') AS CHECKGPA,
    NVL(ES.CHECKEXAM, '-') AS CHECKEXAM,
    NVL(TO_CHAR(ES.CREATED, 'YYYY-MM-DD HH24:MI:SS'), '-') AS CREATED,
    NVL(TO_CHAR(ES.MODIFIED, 'YYYY-MM-DD HH24:MI:SS'), '-') AS MODIFIED,
    NVL(ES.SUCCESS_CONFIRM, '-') AS SUCCESS_CONFIRM,
    NVL(ES.MAJOR_CONFIRM, '-') AS MAJOR_CONFIRM,
    NVL(ES.BIRTHDATE_CONFIRM, '-') AS BIRTHDATE_CONFIRM,
    NVL(DECODE(T.THAI_TITLE, NULL, '-', T.THAI_TITLE), '-') AS THESIS_THAI_TITLE,
    NVL(DECODE(T.ENG_TITLE, NULL, '-', T.ENG_TITLE), '-') AS THESIS_ENG_TITLE,
    NVL(T.THESIS_NAME, '-') AS THESIS_THESIS_NAME,
    NVL(T.THESIS_TYPE, '-') AS THESIS_THESIS_TYPE 
    from VM_STUDENT_S VSS 
    left join vt_thesis T on vss.std_code = t.std_code
    left join egrad_success ES on vss.std_code = es.std_code
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
