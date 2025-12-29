package officerrepos

func (r *officerRepoDB) GetCompanyList(std_code string) (*[]Company, error) {

	companys := []Company{}

	sql := `select std_code,email,
	fullname,company,created,modified
	from egrad_company where std_code = :std_code order by modified desc`

	err := r.oracle_db_dbg.Select(&companys, sql, std_code)
	if err != nil {
		return nil, err
	}

	return &companys, nil
}