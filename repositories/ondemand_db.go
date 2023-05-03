package repositories

func (r *onDemandRepoDB) GetOnDemand(std_code, year string) (*[]OnDemandRepo, error) {

	// grade := []GradeRepo{}
	// query := `select REGIS_YEAR,REGIS_SEMESTER,COURSE_NO,CREDIT,GRADE from gp000.ugb_master_grade
	// 			where std_code=:param1 and regis_year = :param2
	// 			order by regis_year desc,regis_semester desc`

	// err := r.oracle_db.Select(&grade, query, std_code, year)

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (r *onDemandRepoDB) GetOnDemandAll(course_no, year, semester string) (*[]OnDemandRepo, error) {

	ondemands := []OnDemandRepo{}

	query := `SELECT a.subject_code,a.subject_id,a.subject_NameThai,a.subject_NameEng,a.subject_Credit,a.semester,a.year FROM master_subject as a INNER JOIN detail_audio as b ON a.subject_code=b.subject_code WHERE a.subject_id='ENG1001' AND a.semester='2' AND a.year='65' AND a.status='1' ORDER BY a.year DESC,a.subject_id ASC ,a.semester DESC,b.audio_sec ASC`

	results, err := r.mysql_db.Query(query)

	if err != nil {
		print(err)
		return nil, err
	}

	defer results.Close()

	for results.Next() {
		var od OnDemandRepo
		// for each row, scan the result into our tag composite object
		err = results.Scan(&od.STUDY_SEMESTER, &od.STUDY_YEAR, &od.COURSE_NO, &od.DAY_CODE, &od.TIME_CODE, &od.BUILDING_CODE, &od.ROOM_CODE)
		if err != nil {
			print(err)
			return nil, err
		}

		ondemands = append(ondemands, od)

	}

	return &ondemands, nil
}
