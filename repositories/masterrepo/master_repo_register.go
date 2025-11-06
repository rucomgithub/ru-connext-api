package masterrepo

import "fmt"

func (r *studentRepoDB) GetRegisterByYear(std_code, year string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}

	query := "SELECT YEAR,SEMESTER,STD_CODE,COURSE_NO,CREDIT FROM dbgmis00.vm_gstd_course WHERE STD_CODE = :param1 and YEAR = :param2 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db_dbg.Select(&register, query, std_code, year)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *studentRepoDB) GetRegisterAll(std_code string) (*[]RegisterRepo, error) {

	register := []RegisterRepo{}

	query := "SELECT YEAR,SEMESTER,STD_CODE,COURSE_NO,CREDIT FROM dbgmis00.vm_gstd_course WHERE STD_CODE = :param1 ORDER BY YEAR DESC, SEMESTER DESC"

	err := r.oracle_db_dbg.Select(&register, query, std_code)

	if err != nil {
		return nil, err
	}

	return &register, nil
}

func (r *studentRepoDB) GetRegisterFeeAll(std_code,role string) (*[]RegisterFeeRepo, error) {

	register := []RegisterFeeRepo{}

	query := `SELECT R.STD_CODE,R.YEAR,R.SEMESTER,R.TOTAL_AMOUNT,R.REGIS_TYPE,RT.REGIS_NAME
			FROM DBGRAD00.VM_RG_RECEIPT_EGRAD R  
			INNER JOIN DBGRAD00.VM_REGIS_TYPE_EGRAD RT ON R.REGIS_TYPE = RT.REGIS_TYPE 
			WHERE R.STD_CODE = :param1 AND R.YEAR IS NOT NULL AND R.SEMESTER IS NOT NULL ORDER BY R.YEAR,R.SEMESTER`

	if(role == "Doctor") {
		query =`SELECT R.STD_CODE,R.YEAR,R.SEMESTER,SUM(R.RECEIPT_AMOUNT) TOTAL_AMOUNT,R.RECEIPT_PAID_TYPE REGIS_TYPE ,R.RECEIPT_STATUS REGIS_NAME
				FROM D_RECEIPT R 
				GROUP BY R.STD_CODE,R.YEAR,R.SEMESTER,R.RECEIPT_PAID_TYPE,R.RECEIPT_STATUS
				HAVING R.STD_CODE = :param1  AND R.YEAR IS NOT NULL AND R.SEMESTER IS NOT NULL ORDER BY R.YEAR,R.SEMESTER,R.RECEIPT_PAID_TYPE,R.RECEIPT_STATUS`
	}
     fmt.Println(role,query)
	err := r.oracle_db_dbg.Select(&register, query, std_code)

	if err != nil {
		return nil, err
	}

	return &register, nil
}
