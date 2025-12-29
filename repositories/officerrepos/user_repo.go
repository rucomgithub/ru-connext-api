package officerrepos

import "fmt"

func (r *officerRepoDB) GetUserLogin(username string) (*UserLoginRepo, error) {

	userrole := UserLoginRepo{}

	query := `select username,role,key,created,modified, 1 as status from userrole where username = :param1`

	fmt.Printf("officer: %s \n", username)

	err := r.oracle_db_dbg.Get(&userrole, query, username)

	if err != nil {
		return nil, err

	}

	fmt.Println(userrole)

	return &userrole, nil
}
