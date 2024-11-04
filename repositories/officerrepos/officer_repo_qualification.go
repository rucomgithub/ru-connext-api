package officerrepos

import "fmt"

func (r *officerRepoDB) GetQualificationAll() (*[]Qualification, error) {

	qualifications := []Qualification{}

	sql := `select std_code,request_date,
	decode(operate_date,null,'-',operate_date) operate_date,
	decode(confirm_date,null,'-',confirm_date) confirm_date,
	status,created,modified,
	decode(description,null,'-',description) description,
	decode(cancel_date,null,'-',cancel_date) cancel_date
	from egrad_qualification order by request_date desc`

	err := r.oracle_db_dbg.Select(&qualifications, sql)
	if err != nil {
		return nil, err
	}

	return &qualifications, nil
}

func (r *officerRepoDB) GetQualification(std_code string) (*Qualification, error) {

	qualification := Qualification{}

	sql := `select std_code,request_date,
	decode(operate_date,null,'-',operate_date) operate_date,
	decode(confirm_date,null,'-',confirm_date) confirm_date,
	status,created,modified,
	decode(description,null,'-',description) description,
	decode(cancel_date,null,'-',cancel_date) cancel_date
	from egrad_qualification where std_code = :1`

	err := r.oracle_db_dbg.Get(&qualification, sql, std_code)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rows select: %s\n", std_code)

	return &qualification, nil
}

func (r *officerRepoDB) AddQualification(std_code string) error {

	result, err := r.oracle_db_dbg.Exec("INSERT INTO egrad_qualification (std_code,request_date,created,modified) VALUES (:1,sysdate,sysdate,sysdate)", std_code)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows insert affected: %d\n", rowsAffected)

	return nil
}

func (r *officerRepoDB) UpdateQualification(std_code, status, description string) (int64, error) {

	sql := `UPDATE egrad_qualification SET `
	switch status {
	case "request":
		{
			description = "-"
			sql += ` operate_date = cast(NULL AS TIMESTAMP WITH LOCAL TIME ZONE) , confirm_date = cast(NULL AS TIMESTAMP WITH LOCAL TIME ZONE) , cancel_date = cast(NULL AS TIMESTAMP WITH LOCAL TIME ZONE), status = :1 , modified = sysdate, description = :2 WHERE std_code = :3`
		}
	case "operate":
		sql += ` operate_date = sysdate ,status = :1 , modified = sysdate, description = :2 WHERE std_code = :3 and status = 'request'`
	case "confirm":
		sql += ` confirm_date = sysdate ,status = :1, modified = sysdate, description = :2 WHERE std_code = :3 and status = 'operate'`
	case "cancel":
		sql += ` cancel_date = sysdate ,status = :1, modified = sysdate, description = :2 WHERE std_code = :3`
	}

	result, err := r.oracle_db_dbg.Exec(sql, status, description, std_code)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	fmt.Printf("Rows update affected: %d\n", rowsAffected)

	return rowsAffected, nil
}
