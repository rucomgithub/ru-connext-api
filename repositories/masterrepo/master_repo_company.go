package masterrepo

import "fmt"

func (r *studentRepoDB) GetCommpany(std_code,email string) (*Company, error) {

	company := Company{}

	sql := `select std_code,email,fullname,company,created,modified from egrad_company where std_code = :1 and email = :2`

	err := r.oracle_db_dbg.Get(&company, sql, std_code, email)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Rows select: %s\n", std_code)

	return &company, nil
}

func (r *studentRepoDB) AddCommpany(std_code,email,fullname,company string) error {

	fmt.Println(std_code,email,fullname,company)

	result, err := r.oracle_db_dbg.Exec("INSERT INTO egrad_company (std_code,email,fullname,company,created,modified) VALUES (:1,:2,:3,:4,sysdate,sysdate)", std_code,email,fullname,company)
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