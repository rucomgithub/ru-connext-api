package masterrepo

import (
	"fmt"
)

func (r *studentRepoDB) GetStudentProfile(studentCode string) (student *StudentProfileRepo, err error) {

	student_info := StudentProfileRepo{}

	query := `SELECT S.STD_CODE,P.PRENAME_THAI_S||S.NAME_THAI NAME_THAI,P.PRENAME_ENG_S||S.NAME_ENG NAME_ENG,
			TO_CHAR (S.BIRTH_DATE,'FmDD Month YYYY','nls_calendar=''Thai Buddha''') BIRTH_DATE, ST.STATUS_NAME_THAI STD_STATUS_DESC_THAI,
			S.PERSONAL_ID CITIZEN_ID,S.REGINAL_NO
          || ' = '
          || CASE
                WHEN S.REGINAL_NO = 0
                THEN
                   R.REGINAL_NAME
                WHEN S.REGINAL_NO BETWEEN 1 AND 6
                THEN
                      'สาขาวิทยบริการ จังหวัด'
                   || R.REGINAL_NAME
                WHEN S.REGINAL_NO BETWEEN 8 AND 49
                THEN
                      'สาขาวิทยบริการ จังหวัด'
                   || R.REGINAL_NAME
                WHEN S.REGINAL_NO BETWEEN 50 AND 90
                THEN
                      'สาขาวิทยบริการ จังหวัด'
                   || R.REGINAL_NAME
                WHEN S.REGINAL_NO = 99
                THEN
                   'สาขาวิทยบริการต่างประเทศ'
                WHEN S.REGINAL_NO = 7
                THEN
                   'วิทยาเขตบางนา'
             END
             REGIONAL_NAME_THAI,STD_TYPE  STD_TYPE_DESC_THAI, M.THAI_NAME MAJOR_NAME_THAI,M.FACULTY_NO FACULTY_NAME_THAI,
             VC.THAI_NAME,VC.ENG_NAME,VC.THAI_DEGREE,VC.ENG_DEGREE,VC.THAI_MAJOR,
             NVL(A.E_MAIL,'-') EMAIL_ADDRESS,
             NVL(A.MOBILE_TELEPHONE,'-') MOBILE_TELEPHONE,
             DECODE(S.DEGREE_NAME,'ปริญญาโท','Master','Doctor') AS ROLE
         FROM DBGMIS00.VM_STUDENT S 
             LEFT JOIN DBGMIS00.VM_PRENAME P ON S.PRENAME_NO = P.PRENAME_NO
             LEFT JOIN DBGMIS00.VM_REGINAL_CENTER R ON S.REGINAL_NO = R.REGINAL_NO
             LEFT JOIN DBGRAD00.VM_MAJOR M ON S.MAJOR_NO = M.MAJOR_NO
             LEFT JOIN DBGRAD00.VM_STUDENT_S ST ON S.STD_CODE = ST.STD_CODE
             LEFT JOIN VM_STUDENT_ADDRESS A ON S.STD_CODE = A.STD_CODE 
             LEFT JOIN VM_CURRICULUM VC ON S.MAJOR_NO = VC.MAJOR_NO AND  S.CURR_NO = VC.CURR_NO
			WHERE S.STD_CODE = :param1`

	if studentCode == "5424101212" || studentCode == "5519860048" {
		query = `SELECT S.STD_CODE,'รักราม เรียนดี' NAME_THAI,'Rakram Reandee' NAME_ENG,
			TO_CHAR (S.BIRTH_DATE,'FmDD Month YYYY','nls_calendar=''Thai Buddha''') BIRTH_DATE, ST.STATUS_NAME_THAI STD_STATUS_DESC_THAI,
			S.PERSONAL_ID CITIZEN_ID,S.REGINAL_NO
          || ' = '
          || CASE
                WHEN S.REGINAL_NO = 0
                THEN
                   R.REGINAL_NAME
                WHEN S.REGINAL_NO BETWEEN 1 AND 6
                THEN
                      'สาขาวิทยบริการ จังหวัด'
                   || R.REGINAL_NAME
                WHEN S.REGINAL_NO BETWEEN 8 AND 49
                THEN
                      'สาขาวิทยบริการ จังหวัด'
                   || R.REGINAL_NAME
                WHEN S.REGINAL_NO BETWEEN 50 AND 90
                THEN
                      'สาขาวิทยบริการ จังหวัด'
                   || R.REGINAL_NAME
                WHEN S.REGINAL_NO = 99
                THEN
                   'สาขาวิทยบริการต่างประเทศ'
                WHEN S.REGINAL_NO = 7
                THEN
                   'วิทยาเขตบางนา'
             END
             REGIONAL_NAME_THAI,STD_TYPE  STD_TYPE_DESC_THAI, M.THAI_NAME MAJOR_NAME_THAI,M.FACULTY_NO FACULTY_NAME_THAI,
             VC.THAI_NAME,VC.ENG_NAME,VC.THAI_DEGREE,VC.ENG_DEGREE,VC.THAI_MAJOR,
             NVL(A.E_MAIL,'-') EMAIL_ADDRESS,
             NVL(A.MOBILE_TELEPHONE,'-') MOBILE_TELEPHONE,
             DECODE(S.DEGREE_NAME,'ปริญญาโท','Master','Doctor') AS ROLE 
         FROM DBGMIS00.VM_STUDENT S 
             LEFT JOIN DBGMIS00.VM_PRENAME P ON S.PRENAME_NO = P.PRENAME_NO
             LEFT JOIN DBGMIS00.VM_REGINAL_CENTER R ON S.REGINAL_NO = R.REGINAL_NO
             LEFT JOIN DBGRAD00.VM_MAJOR M ON S.MAJOR_NO = M.MAJOR_NO
             LEFT JOIN DBGRAD00.VM_STUDENT_S ST ON S.STD_CODE = ST.STD_CODE
             LEFT JOIN VM_STUDENT_ADDRESS A ON S.STD_CODE = A.STD_CODE 
             LEFT JOIN VM_CURRICULUM VC ON S.MAJOR_NO = VC.MAJOR_NO AND  S.CURR_NO = VC.CURR_NO
			WHERE S.STD_CODE = :param1`
	}

	fmt.Printf("profile: %s \n", studentCode)

	err = r.oracle_db_dbg.Get(&student_info, query, studentCode)

	if err == nil {
		student = &student_info
		return student, nil
	}

	return nil, err
}
