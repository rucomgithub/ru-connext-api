package studentr
import "fmt"
func (r *studentRepoDB) GetStudentProfile(studentCode string) (student *StudentProfileRepo, err error) {
	if studentCode == "6299999992" {
		studentCode = "6202408966"
		fmt.Printf("register: %s \n", studentCode)
	}
	student_info := StudentProfileRepo{}
	query := `SELECT STD_CODE,NAME_THAI,NAME_ENG,BIRTH_DATE,STD_STATUS_DESC_THAI,CITIZEN_ID,REGIONAL_NAME_THAI,STD_TYPE_DESC_THAI,FACULTY_NAME_THAI,MAJOR_NAME_THAI,WAIVED_NO,WAIVED_PAID,DECODE(WAIVED_TOTAL_CREDIT,null,0,WAIVED_TOTAL_CREDIT) AS WAIVED_TOTAL_CREDIT,CHK_CERT_NAME_THAI,PENAL_NAME_THAI,MOBILE_TELEPHONE,EMAIL_ADDRESS FROM DBBACH00.VM_STUDENT_PROFILE WHERE STD_CODE = :param1`

	err = r.oracle_db.Get(&student_info, query, studentCode)
	if err != nil {
		return nil, err
	}

	student = &student_info

	return student, nil
}
