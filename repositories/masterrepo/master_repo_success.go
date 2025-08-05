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
