package officerrepos

func (r *officerRepoDB) GetDateString(str string) string {
	length := len(str)
	switch length {
		case 4 :
			return `'YYYY'`
		case 7 :
			return `'YYYY-MM'`
		case 10 :
			return `'YYYY-MM-DD'`
		default:
			return `'YYYY-MM-DD'`
	}
}

func (r *officerRepoDB) FindReport(startdate,enddate string) ([]map[string]interface{},error) {

	dateString := r.GetDateString(startdate)

	 sql := `SELECT ROW_NUMBER() OVER (ORDER BY date_report) AS id, date_report, 
	SUM(CASE WHEN status = 'request' THEN count_report ELSE 0 END) AS request,
	SUM(CASE WHEN status = 'operate' THEN count_report ELSE 0 END) AS operate,
	SUM(CASE WHEN status = 'confirm' THEN count_report ELSE 0 END) AS confirm
	from (select TO_CHAR(g.MODIFIED,` + dateString + `) DATE_REPORT,g.status,count(g.status) COUNT_REPORT 
	from egrad_qualification g
	where TO_CHAR(g.MODIFIED,` + dateString + `) between :1 and :2 
	group by g.status,TO_CHAR(g.MODIFIED,` + dateString + `))
	GROUP BY date_report`

		rows, err := r.oracle_db_dbg.Query(sql,startdate,enddate)

		defer rows.Close()

		if err != nil {
			return nil, err
		}

		// ดึงชื่อคอลัมน์
		columns, err := rows.Columns()
		if err != nil {
			return nil , err
		}



		var reports []map[string]interface{}

		for rows.Next() {
			// สร้าง slice ของ interface{} ขนาดเท่ากับจำนวนคอลัมน์
			values := make([]interface{}, len(columns))
			// สร้าง slice ของ pointers สำหรับ scan
			scanArgs := make([]interface{}, len(columns))
			for i := range values {
				scanArgs[i] = &values[i]
			}

			// สแกนข้อมูล
			err := rows.Scan(scanArgs...)
			if err != nil {
				return nil,err
			}

			// สร้าง map สำหรับเก็บข้อมูล
			report := make(map[string]interface{})
			for i, col := range columns {
				if values[i] != nil {
					// เก็บข้อมูลลงใน map โดยใช้ชื่อคอลัมน์เป็นคีย์
					report[col] = values[i]
					// for _, fee := range fees {
					// 	if(strings.Contains(col,fee.FEE_NO)){
					// 		report[col+"_name"] = fee.FEE_NAME
					// 	}
					// }
				} else {
					report[col] = nil
				}
			}

			reports = append(reports, report)
		}

		if err = rows.Err(); err != nil {
			return nil,err
		}

	return reports, nil
}