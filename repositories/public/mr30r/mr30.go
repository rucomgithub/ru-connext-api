package mr30r

func (r *mr30RepoDB) GetMr30(course_year, course_semester string) (*[]Mr30Repo, error) {

	mr30_info := []Mr30Repo{}
	query := "SELECT distinct a.study_year||a.study_semester||a.course_no||a.course_method||a.course_method_number||a.day_code||a.time_code||a.room_group||a.instr_group||ex.exam_date||trim(ex.exam_period) id, a.study_year course_year, a.study_semester course_semester, a.course_no, a.course_method, a.course_method_number, a.day_code, a.time_code, a.room_group, a.instr_group, DECODE (a.course_method, '1', 'SEC.', '2', 'LEC.', '3', 'LAB.', '4', 'VDO.', '') || '' || DECODE (a.course_method_number, 0, '', a.course_method_number) course_method_detail, DECODE (c.day_name_s, '-', '', c.day_name_s) day_name_s, DECODE (b.time_start, NULL, b.time_ru30, b.time_start || '-' || b.time_end) time_period, DECODE (d.broom, ' 0000000000', '', replace(d.broom,'/','-')) course_room, e.instructor_name_ru30 course_instructor, cr.show_ru30, cr.credit course_credit, pr.pr_course_name course_pr, cm.course_comment course_comment, ex.course_examdate FROM ugb_ru30_daytime a, ugb_time_schedule b, ugb_day_schedule c, ugb_course_comment cm, ugb_pr_course pr, (SELECT a.course_no, a.credit, a.show_ru30, a.declare_year FROM ugb_course a WHERE a.declare_year = (SELECT MAX (declare_year) FROM ugb_course WHERE a.course_no = course_no)) cr, ( SELECT room_group, SUBSTR (MAX (broom), 2) broom FROM ( SELECT room_group, SYS_CONNECT_BY_PATH (broom, ', ') broom FROM (SELECT room_group, broom, ROW_NUMBER () OVER (PARTITION BY room_group ORDER BY ROWNUM) rn FROM (SELECT room_group, TRIM (building_code) || '' || TRIM (room_code) broom FROM ugb_ru30_room WHERE STUDY_YEAR = :param1 AND STUDY_SEMESTER = :param2)) CONNECT BY room_group = PRIOR room_group AND rn = PRIOR rn + 1 START WITH rn = 1) GROUP BY room_group) d, (  SELECT instr_group, SUBSTR (MAX (instructor_name_ru30), 3) instructor_name_ru30 FROM ( SELECT instr_group, SYS_CONNECT_BY_PATH (instructor_name_ru30, ' , ') instructor_name_ru30 FROM (SELECT instr_group, instructor_name_ru30, ROW_NUMBER () OVER (PARTITION BY instr_group ORDER BY ROWNUM) rn FROM (SELECT a.instr_group, a.instructor_code, rk.RANK_NAME_ENG_S || b.instructor_name_ru30 instructor_name_ru30 FROM ugb_ru30_instructor a, ugb_instructor b, ugb_rank rk WHERE study_year = :param3 AND study_semester = :param4 AND a.instructor_code = b.instructor_code AND b.rank_no = rk.rank_no(+))) CONNECT BY instr_group = PRIOR instr_group AND rn = PRIOR rn + 1 START WITH rn = 1) GROUP BY instr_group) e, ( (SELECT course_no,to_char(exam_date,'DDMMYYYY') exam_date,trim(exam_period) exam_period, TO_CHAR (exam_date, 'FmDD Mon YYYY', 'nls_calendar=''Thai Buddha''') || ' ' || trim(exam_period) course_examdate FROM ugb_exam WHERE course_year = :param5 AND course_semester = :param6 UNION SELECT DISTINCT course_no,to_char(exam_date,''),exam_period, 'คณะจัดสอบเอง' course_examdate FROM ugb_hour_ru30 WHERE study_year = :param7 AND study_semester = :param8 AND flag_exam <> 0)) ex WHERE a.study_year = :param9 AND a.study_semester = :param10 AND a.time_code = b.time_code AND a.day_code = c.day_code AND a.room_group = d.room_group AND a.instr_group = e.instr_group AND a.course_no = cr.course_no AND a.course_no = pr.course_no(+) AND a.course_no = cm.course_no(+) AND a.course_no = ex.course_no ORDER BY a.course_no, a.course_method, a.course_method_number"

	err := r.oracle_db.Select(&mr30_info, query, course_year, course_semester, course_year, course_semester, course_year, course_semester, course_year, course_semester, course_year, course_semester)

	if err != nil {
		return nil, err
	}

	return &mr30_info, nil
}

func (r *mr30RepoDB) GetMr30Year() (*[]Mr30YearRepo, error) {

	mr30_info := []Mr30YearRepo{}
	query := "SELECT course_year,course_semester from (select a.study_year course_year, a.study_semester course_semester FROM ugb_ru30_daytime a group by a.study_year, a.study_semester order by 1 desc , 2 desc) where rownum < 4"

	err := r.oracle_db.Select(&mr30_info, query)

	if err != nil {
		return nil, err
	}

	return &mr30_info, nil
}
