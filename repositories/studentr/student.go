package studentr

func (r *studentRepoDB) GetStudentAll() (*[]StudentRepo, error) {

	students := []StudentRepo{}

	query := "select STD_CODE from DBBACH00.VM_STUDENT_PROFILE WHERE std_code like '6553000%'"

	err := r.oracle_db.Select(&students, query)
	if err != nil {
		return nil, err
	}

	return &students, nil
}
