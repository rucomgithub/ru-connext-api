package studentr

func (r *studentRepoDB) GetStudentAll() (*[]StudentRepo, error) {

	students := []StudentRepo{}

	query := "SELECT STD_CODE FROM DBBACH00.UGB_REGIS_RU24 WHERE STD_CODE like '65%' and YEAR = 2565 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db.Select(&students, query)
	if err != nil {
		return nil, err
	}

	return &students, nil
}
