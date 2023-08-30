package repositories

func (r *scholarshipRepoDB) GetScholarshipAll(std_code string) (*[]ScholarshipRepo, error) {
	scholarship := []ScholarshipRepo{}
	query := `SELECT SC.ID,SC.STD_CODE,SC.SCHOLARSHIP_YEAR,SC.SCHOLARSHIP_SEMESTER,SC.SCHOLARSHIP_TYPE,SC.AMOUNT,SC.CREATED,SC.MODIFIED,SC.USERNAME,SUB.SUBSIDY_NO,SUB.SUBSIDY_NAME_THAI 
	FROM SCHOLARSHIP SC INNER JOIN SUBSIDY SUB ON SC.SCHOLARSHIP_TYPE = SUB.SUBSIDY_NO WHERE SC.STD_CODE = :param1 `
	err := r.oracle_db.Select(&scholarship, query, std_code)

	if err != nil {
		return nil, err
	}

	return &scholarship, nil
}
