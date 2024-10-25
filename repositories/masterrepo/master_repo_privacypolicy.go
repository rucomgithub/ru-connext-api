package masterrepo

import "fmt"

func (r *studentRepoDB) AddPrivacyPolicy(std_code, version string) error {

	result, err := r.oracle_db_dbg.Exec("INSERT INTO b_personal (std_code,version,created,modified) VALUES (:1,:2,sysdate,sysdate)", std_code, version)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	if err != nil {
		return err
	}

	return nil
}

func (r *studentRepoDB) UpdatePrivacyPolicy(std_code, version string) error {

	result, err := r.oracle_db_dbg.Exec("update b_personal set version = :1, modified = sysdate where std_code = :2", version, std_code)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Rows affected: %d\n", rowsAffected)

	if err != nil {
		return err
	}

	return nil
}

func (r *studentRepoDB) GetPrivacyPolicy(std_code string) (*PrivacyPolicy, error) {

	privacy := PrivacyPolicy{}

	query := "SELECT STD_CODE,VERSION,CREATED,MODIFIED FROM dbgmis00.B_PERSONAL WHERE STD_CODE = :param1"

	err := r.oracle_db_dbg.Get(&privacy, query, std_code)

	if err != nil {
		return nil, err
	}

	return &privacy, nil
}
