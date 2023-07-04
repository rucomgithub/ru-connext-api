package repositories

func (r *registerRepoDB) GetRegister(std_code, year string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}
	query := "SELECT YEAR,SEMESTER,COURSE_NO,STD_CODE,CREDIT FROM DBBACH00.UGB_REGIS_RU24 WHERE STD_CODE = :param1 and YEAR = :param2 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db.Select(&register, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *registerRepoDB) GetRegisterYearList(std_code string, year string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}
	query := "SELECT YEAR,SEMESTER,COURSE_NO,STD_CODE,CREDIT FROM DBBACH00.UGB_REGIS_RU24 WHERE STD_CODE = :param1 and YEAR = :param2 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db.Select(&register, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *registerRepoDB) GetRegisterYear(std_code string) (*[]RegisterYearRepo, error) {

	register := []RegisterYearRepo{}
	query := "SELECT YEAR FROM (SELECT YEAR FROM DBBACH00.UGB_REGIS_RU24 WHERE STD_CODE = :param1) GROUP BY YEAR ORDER BY 1 DESC"

	err := r.oracle_db.Select(&register, query, std_code)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *registerRepoDB) GetRegisterGroupYearSemester(std_code string) (*[]RegisterYearSemesterRepo, error) {

	register := []RegisterYearSemesterRepo{}
	query := "SELECT YEAR,SEMESTER FROM (SELECT YEAR,SEMESTER FROM DBBACH00.UGB_REGIS_RU24 WHERE STD_CODE = :param1 and SEMESTER in (1,2,3)) GROUP BY YEAR,SEMESTER ORDER BY 1 DESC , 2 DESC"

	err := r.oracle_db.Select(&register, query, std_code)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *registerRepoDB) GetRegisterMr30(year, semester, studentCode string) (*[]RegisterMr30Repo, error) {

	register := []RegisterMr30Repo{}

	query := `select b.course_no, a.id, a.course_year, a.course_semester, a.course_no, a.course_method, a.course_method_number, a.day_code, a.time_code, a.room_group, a.instr_group, 
a.course_method_detail, a.day_name_s, a.time_period, a.course_room, a.course_instructor, a.show_ru30, a.course_credit, a.course_pr, a.course_comment, a.exam_time as course_examdate 
FROM (SELECT DISTINCT a.study_year || a.study_semester || a.course_no || a.course_method || a.course_method_number || a.day_code || a.time_code || a.room_group || a.instr_group 
|| ex.exam_date || ex.exam_period id, a.study_year course_year, a.study_semester course_semester, a.course_no, a.course_method, a.course_method_number, a.day_code, a.time_code, a.room_group, a.instr_group, 
DECODE (a.course_method, '1', 'SEC.', '2', 'LEC.', '3', 'LAB.', '4', 'VDO.', '') || '' || DECODE (a.course_method_number, 0, '', a.course_method_number) course_method_detail, 
DECODE (c.day_name_s, '-', '', c.day_name_s) day_name_s, DECODE (b.time_start, NULL, b.time_ru30, b.time_start || '-' || b.time_end) time_period, 
DECODE (d.broom, ' 0000000000', '', d.broom) course_room, e.instructor_name_ru30 course_instructor, cr.show_ru30, cr.credit || ' CR.' course_credit, pr.pr_course_name course_pr, 
cm.course_comment course_comment, ex.exam_time FROM ugb_ru30_daytime a, ugb_time_schedule b, ugb_day_schedule c, ugb_course_comment cm, ugb_pr_course pr, 
(SELECT a.course_no, a.credit, a.show_ru30, a.declare_year FROM ugb_course a WHERE a.declare_year = (SELECT MAX (declare_year) FROM ugb_course WHERE a.course_no = course_no)) cr, 
(  SELECT room_group, SUBSTR (MAX (broom), 2) broom FROM (    SELECT room_group, SYS_CONNECT_BY_PATH (broom, ', ') broom 
FROM (SELECT room_group, broom, ROW_NUMBER () OVER (PARTITION BY room_group ORDER BY ROWNUM) rn FROM (SELECT room_group, TRIM (building_code) || '' || TRIM (room_code) broom FROM ugb_ru30_room 
WHERE STUDY_SEMESTER = :param1 AND STUDY_YEAR = :param2)) CONNECT BY room_group = PRIOR room_group AND rn = PRIOR rn + 1 START WITH rn = 1) GROUP BY room_group) d, 
(  SELECT instr_group, SUBSTR (MAX (instructor_name_ru30), 3) instructor_name_ru30 FROM (    SELECT instr_group, SYS_CONNECT_BY_PATH (instructor_name_ru30, ' , ') instructor_name_ru30 
FROM (SELECT instr_group, instructor_name_ru30, ROW_NUMBER () OVER (PARTITION BY instr_group ORDER BY ROWNUM) rn FROM (SELECT a.instr_group, a.instructor_code, rk.RANK_NAME_ENG_S || b.instructor_name_ru30 instructor_name_ru30 
FROM ugb_ru30_instructor a, ugb_instructor b, ugb_rank rk 
WHERE study_semester = :param3 AND study_year = :param4 AND a.instructor_code = b.instructor_code AND b.rank_no = rk.rank_no(+))) 
CONNECT BY instr_group = PRIOR instr_group AND rn = PRIOR rn + 1 START WITH rn = 1) 
GROUP BY instr_group) e, ( (SELECT course_no, TO_CHAR (exam_date, 'DDMMYYYY') exam_date, trim(exam_period) exam_period, TO_CHAR (exam_date, 'FmDD Mon YYYY', 'nls_calendar=''Thai Buddha''') || ' ' ||trim(exam_period) exam_time 
FROM ugb_exam 
WHERE course_semester = :param5 AND course_year = :param6 UNION SELECT DISTINCT course_no, TO_CHAR (exam_date, ''), trim(exam_period) exam_period, 'คณะจัดสอบเอง' exam_time FROM ugb_hour_ru30 
WHERE study_semester = :param7  AND study_year = :param8 AND flag_exam <> 0)) ex 
WHERE     a.study_semester = :param9  AND a.study_year = :param10 AND a.time_code = b.time_code AND a.day_code = c.day_code 
AND a.room_group = d.room_group 
AND a.instr_group = e.instr_group AND a.course_no = cr.course_no AND a.course_no = pr.course_no(+) AND a.course_no = cm.course_no(+) AND a.course_no = ex.course_no 
ORDER BY a.course_no, a.course_method, a.course_method_number ) a ,(select course_no from ugb_regis_ru24  
where semester=:param11 and year=:param12 and  std_code= :param13 
) b where a.course_no = b.course_no`

	err := r.oracle_db.Select(&register, query, semester, year, semester, year, semester, year, semester, year, semester, year, semester, year, studentCode)
	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *registerRepoDB) GetYearSemesterLatest() (student *YearSemesterRepo, err error) {

	yearsemester := YearSemesterRepo{}
	query := `SELECT CURRENT_YEAR YEAR, CURRENT_SEMESTER SEMESTER FROM CURRENT_REGISTER WHERE ROWNUM = 1`

	err = r.oracle_db.Get(&yearsemester, query)
	if err != nil {
		return nil, err
	}

	student = &yearsemester

	return &yearsemester, nil
}
