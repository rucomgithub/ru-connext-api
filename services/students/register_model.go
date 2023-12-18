package students

type (
	RegisterAllRequest struct {
		STD_CODE string `json:"std_code"`
		YEAR     string `json:"year"`
	}

	RegisterAllResponse struct {
		STD_CODE string `json:"std_code"`
		YEAR     string `json:"year"`
		RECORD   []registerAllRecord
	}

	registerAllRecord struct {
		YEAR      string `json:"year"`
		SEMESTER  string `json:"semester"`
		COURSE_NO string `json:"course_no"`
		CREDIT    string `json:"credit"`
	}
)
