package repositories

import "fmt"

func (r *rotcsRepoDB) GetRotcsRegister(std_code string) (*[]RotcsRegisterRepo, error) {
	if std_code == "6299999991" {
		std_code = "6405000982"
		fmt.Printf("register: %s \n", std_code)
	}
	if std_code == "6299999992" {
		std_code = "6406600012"
		fmt.Printf("register: %s \n", std_code)
	}
	register := []RotcsRegisterRepo{}
	query := `SELECT report.studentCode,layerArmy,locationArmy, yearReport, layerReport,
				CASE WHEN typeReport = "P" THEN "เลื่อนชั้น" 
				WHEN typeReport = "R" THEN "ซ้ำชั้น" 
				WHEN typeReport = "W" THEN "รอรับสิทธิ" 
				WHEN typeReport = "M" THEN "โอนย้ายสถานศึกษาวิชาทหาร" 
				ELSE "-" 
				END typeReport, 
				CASE 
				WHEN status = "CONFIRM" THEN "ยืนยัน" ELSE "-" 
				END status
				FROM report inner join student on report.studentCode = student.studentCode 
				inner join faculty on report.facultyId = faculty.id 
				WHERE report.studentCode = ? ORDER BY report.yearReport asc, report.layerArmy asc`
	err := r.mysql_db.Select(&register, query, std_code)
	if err != nil {
		return nil, err
	}
	return &register, nil
}

func (r *rotcsRepoDB) GetRotcsExtend(std_code string) (*RotcsExtendRepo, error) {
	if std_code == "6299999991" {
		std_code = "6401002222"
		fmt.Printf("register: %s \n", std_code)
	}
	if std_code == "6299999992" {
		std_code = "6401008344"
		fmt.Printf("register: %s \n", std_code)
	}
	extend := RotcsExtendRepo{}
	query := `SELECT studentCode, extendYear,
	CONCAT("ใบสำคัญ สด.9 เลขที่ : ",code9) code9,
	CASE 
	WHEN option1 = "1" THEN "สำเนา สด.9 จำนวน 2 ฉบับ : มี" ELSE "สำเนา สด.9 จำนวน 2 ฉบับ : ไม่มี" 
	END option1,
	CASE 
	WHEN option2 = "1" THEN "สำเนาทะเบียนบ้าน จำนวน 2 ฉบับ : มี" ELSE "สำเนาทะเบียนบ้าน จำนวน 2 ฉบับ : ไม่มี" 
	END option2,
	CASE 
	WHEN option3 = "1" THEN "สำเนาหมายเรียก(สด.35) จำนวน 2 ฉบับ : มี" ELSE "สำเนาหมายเรียก(สด.35) จำนวน 2 ฉบับ : ไม่มี" 
	END option3,
	CASE 
	WHEN option4 = "1" THEN "สำเนาบัตรนักศึกษา จำนวน 2 ฉบับ : มี" ELSE "สำเนาบัตรนักศึกษา จำนวน 2 ฉบับ : ไม่มี" 
	END option4,
	CASE 
	WHEN option5 = "1" THEN "ใบรับรองผลการศึกษา 9 หน่วยกิต จำนวน 2 ฉบับ : มี" ELSE "ใบรับรองผลการศึกษา 9 หน่วยกิต จำนวน 2 ฉบับ : ไม่มี" 
	END option5,
	CASE 
	WHEN option6 = "1" THEN "หนังสือรับรองการเป็นนักศึกษา จำนวน 2 ฉบับ : มี" ELSE "หนังสือรับรองการเป็นนักศึกษา จำนวน 2 ฉบับ : ไม่มี" 
	END option6,
	CASE 
	WHEN option7 = "1" THEN "สำเนาบัตรประชาชน จำนวน 2 ฉบับ : มี" ELSE "สำเนาบัตรประชาชน จำนวน 2 ฉบับ : ไม่มี" 
	END option7,
	CASE 
	WHEN option8 = "1" THEN "ใบเสร็จลงทะเบียนเรียนภาคปัจจุบัน จำนวน 2 ฉบับ : มี" ELSE "ใบเสร็จลงทะเบียนเรียนภาคปัจจุบัน จำนวน 2 ฉบับ : ไม่มี" 
	END option8,
	CASE 
	WHEN option9 = "1" THEN "อื่นๆ : มี" ELSE "อื่นๆ : ไม่มี" 
	END option9,
	CASE 
	WHEN option9 = "1" THEN concat("กรณีระบุอื่นๆ : ",optionOther) ELSE "กรณีระบุอื่นๆ : -" 
	END optionOther
	FROM extend inner join faculty on extend.facultyId = faculty.id 
	WHERE extend.studentCode = ? ORDER BY extend.studentCode ASC`
	err := r.mysql_db.Get(&extend, query, std_code)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	detail := []RotcsExtendDetailRepo{}

	query = `SELECT id,registerYear,registerSemester,credit,created,modified
	FROM extendDetail
	WHERE extendDetail.studentCode = ? ORDER BY extendDetail.registerYear Desc,registerYear,registerSemester Desc`
	err = r.mysql_db.Select(&detail, query, std_code)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	extend.Detail = detail

	return &extend, nil
}
