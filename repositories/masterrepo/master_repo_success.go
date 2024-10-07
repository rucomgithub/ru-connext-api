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
	NVL(g.MAJOR_NAME,'-') MAJOR_NAME,
	NVL(g.PLAN,'-') PLAN,
	g.CONFERENCE_NO,
	g.SERIAL_NO,
	TO_CHAR(TO_DATE(g.conference_date,'DD/MM/YYYY'),'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') conference_date,
	TO_CHAR(TO_DATE(g.graduated_date,'DD/MM/YYYY'),'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') graduated_date,
	TO_CHAR(TO_DATE(g.conference_date,'DD/MM/YYYY')+7,'FMDD MONTH YYYY','NLS_CALENDAR=''THAI BUDDHA'' NLS_DATE_LANGUAGE=THAI') confirm_date, 
	nvl(nvl(g.main_major_thai,g.major_name) , '-') major_name_thai 
	from  vm_graduate g 
	where  std_code = :param1`

	fmt.Printf("success: %s \n", studentCode)

	err = r.oracle_db_dbg.Get(&student_info, query, studentCode)

	if err == nil {
		student = &student_info
		return student, nil
	}

	return nil, err
}
