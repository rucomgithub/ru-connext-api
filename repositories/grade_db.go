package repositories

import "fmt"

func (r *gradeRepoDB) GetGradeYear(std_code, year string) (*[]GradeRepo, error) {
	if std_code == "6299999991" {
		std_code = "6407501375"
		fmt.Printf("register: %s \n", std_code)
	}
	if std_code == "6299999992" {
		std_code = "6202408966"
		fmt.Printf("register: %s \n", std_code)
	}
	grade := []GradeRepo{}
	query := `select REGIS_YEAR,REGIS_SEMESTER,COURSE_NO,CREDIT,GRADE from gp000.ugb_master_grade 
				where std_code=:param1 and regis_year = :param2
				order by regis_year desc,regis_semester desc`

	err := r.oracle_db.Select(&grade, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &grade, nil
}

func (r *gradeRepoDB) GetGradeAll(std_code string) (*[]GradeRepo, error) {
	if std_code == "6299999991" {
		std_code = "6407501375"
		fmt.Printf("register: %s \n", std_code)
	}
	if std_code == "6299999992" {
		std_code = "5802031012"
		fmt.Printf("register: %s \n", std_code)
	}
	grade := []GradeRepo{}
	query := `select REGIS_YEAR,REGIS_SEMESTER,COURSE_NO,CREDIT,GRADE from gp000.ugb_master_grade 
				where std_code=:param1
				order by regis_year desc,regis_semester desc`

	err := r.oracle_db.Select(&grade, query, std_code)

	if err != nil {
		return nil, err
	}

	return &grade, nil
}
