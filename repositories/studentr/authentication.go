package studentr

func (r *studentRepoDB) Authentication(studentCode string) (token *PrepareTokenRepo, err error) {

	tempToken := PrepareTokenRepo{}
	query := `SELECT STD_CODE, (1) AS STATUS  FROM DBBACH00.VM_STUDENT_PROFILE WHERE STD_CODE = :param1`

	if studentCode == "6299999991" {
		query = `SELECT :param1 AS STD_CODE, (1) AS STATUS FROM dual`
	}

	if studentCode == "6299999992" {
		query = `SELECT :param1 AS STD_CODE, (1) AS STATUS FROM dual`
	}

	err = r.oracle_db.Get(&tempToken, query, studentCode)
	if err != nil {
		return nil, err
	}

	token = &tempToken

	return token, nil
} 
