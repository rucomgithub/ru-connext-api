package masterrepo

func (r *studentRepoDB) GetRegisterByYear(std_code, year string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}

	query := "SELECT YEAR,SEMESTER,STD_CODE,COURSE_NO,CREDIT FROM dbgmis00.vm_gstd_course WHERE STD_CODE = :param1 and YEAR = :param2 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db_dbg.Select(&register, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *studentRepoDB) GetRegisterAll(std_code string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}

	query := "SELECT YEAR,SEMESTER,STD_CODE,COURSE_NO,CREDIT FROM dbgmis00.vm_gstd_course WHERE STD_CODE = :param1 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db_dbg.Select(&register, query, std_code)

	if err != nil {
		return nil, err
	}

	return &register, nil
}
