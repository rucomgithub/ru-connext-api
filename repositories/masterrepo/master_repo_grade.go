package masterrepo

func (r *studentRepoDB) GetGradeByYear(std_code, year string) (*[]GradeRepo, error) {

	grades := []GradeRepo{}

	query := "SELECT YEAR,SEMESTER,STD_CODE,COURSE_NO,CREDIT,GRADE FROM dbgmis00.vm_gstd_course WHERE STD_CODE = :param1 and YEAR = :param2 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db_dbg.Select(&grades, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &grades, nil
}

func (r *studentRepoDB) GetGradeAll(std_code string) (*[]GradeRepo, error) {

	grades := []GradeRepo{}

	query := `SELECT GC.YEAR,GC.SEMESTER,GC.STD_CODE,GC.COURSE_NO,GC.CREDIT,GC.GRADE ,NVL(C.COURSE_TYPE_NO,-1) COURSE_TYPE_NO, NVL(C.THAI_DESCRIPTION,'-') THAI_DESCRIPTION
    FROM  VM_GSTD_COURSE GC
    LEFT JOIN (
    SELECT A.COURSE_NO,T.COURSE_TYPE_NO,T.THAI_DESCRIPTION
    FROM B_COURSE A
    INNER JOIN (
        SELECT COURSE_NO, MAX(YEAR) YEAR
        FROM B_COURSE
        GROUP BY COURSE_NO
    ) B ON A.COURSE_NO = B.COURSE_NO AND A.YEAR = B.YEAR
    INNER JOIN B_COURSE_TYPE T ON A.COURSE_TYPE_NO = T.COURSE_TYPE_NO
    ) C ON GC.COURSE_NO = C.COURSE_NO
    WHERE GC.STD_CODE = :param1 ORDER BY C.COURSE_TYPE_NO,GC.YEAR,GC.SEMESTER`

	err := r.oracle_db_dbg.Select(&grades, query, std_code)

	if err != nil {
		return nil, err
	}

	return &grades, nil
}

func (r *studentRepoDB) GetGpaAll(std_code string) (*GPARepo, error) {

	gpa := GPARepo{}

	query := `select sum(decode(b.grade_point,0,decode(b.grade,'F',a.credit,0),a.credit)) credit
			,trunc(decode(sum(a.credit),0,0,sum(a.credit*b.grade_point)/ decode(sum(decode(b.grade_point,0,decode(b.grade,'F',a.credit,0),a.credit)),0,1,sum(decode(b.grade_point,0,decode(b.grade,'F',a.credit,0),a.credit))) ),2) GPA
			from  VM_GSTD_COURSE a, VM_G_GRADE b  where ( a.std_code = :param1) and a.grade = b.grade`

	err := r.oracle_db_dbg.Get(&gpa, query, std_code)

	if err != nil {
		return nil, err
	}

	return &gpa, nil
}

func (r *studentRepoDB) GetGpaByYear(std_code, year string) (*GPARepo, error) {

	gpa := GPARepo{}

	query := `select sum(decode(b.grade_point,0,decode(b.grade,'F',a.credit,0),a.credit)) credit
			,trunc(decode(sum(a.credit),0,0,sum(a.credit*b.grade_point)/ decode(sum(decode(b.grade_point,0,decode(b.grade,'F',a.credit,0),a.credit)),0,1,sum(decode(b.grade_point,0,decode(b.grade,'F',a.credit,0),a.credit))) ),2) GPA
			from  VM_GSTD_COURSE a, VM_G_GRADE b  where ( a.std_code = :param1 and a.year = :param2) and a.grade = b.grade`

	err := r.oracle_db_dbg.Get(&gpa, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &gpa, nil
}
