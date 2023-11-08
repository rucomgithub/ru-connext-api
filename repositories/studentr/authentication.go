package studentr

import "fmt"

func (r *studentRepoDB) Authentication(studentCode string) (token *PrepareTokenRepo, err error) {

	tempToken := PrepareTokenRepo{}

	query := `SELECT STD_CODE, (1) AS STATUS,'Bachelor' AS ROLE FROM DBBACH00.VM_STUDENT_PROFILE WHERE STD_CODE = :param1`
	fmt.Println("Bachelor")
	err = r.oracle_db.Get(&tempToken, query, studentCode)
	if err == nil {
		token = &tempToken
		return token, err
	}

	query = `SELECT STD_CODE, (1) AS STATUS, DECODE(DEGREE_NAME,'ปริญญาโท','Master','Doctor') AS ROLE FROM DBGMIS00.VM_STUDENT WHERE STD_CODE = :param1`
	fmt.Println("Master")

	err = r.oracle_db_dbg.Get(&tempToken, query, studentCode)
	if err == nil {
		token = &tempToken
		return token, err
	}

	return nil, err
}
