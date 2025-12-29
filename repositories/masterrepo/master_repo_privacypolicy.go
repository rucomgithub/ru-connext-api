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

func (r *studentRepoDB) UpdatePrivacyPolicy(std_code, version, status string) error {

	result, err := r.oracle_db_dbg.Exec("update b_personal set status = :1, modified = sysdate where std_code = :2 and version = :3", status, std_code, version)
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

func (r *studentRepoDB) GetPrivacyPolicy(std_code, version string) (*PrivacyPolicy, error) {

	privacy := PrivacyPolicy{}

	query := "SELECT STD_CODE,VERSION,STATUS,CREATED,MODIFIED FROM dbgmis00.B_PERSONAL WHERE STD_CODE = :param1 and VERSION = :param2"

	err := r.oracle_db_dbg.Get(&privacy, query, std_code, version)

	if err != nil {
		return nil, err
	}

	return &privacy, nil
}
