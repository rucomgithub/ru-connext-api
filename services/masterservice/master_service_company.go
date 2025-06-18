package masterservice

func (s *studentServices) GetCommpany(std_code,email string) (*CompanyResponse, error) {

	qf, err := s.studentRepo.GetCommpany(std_code, email)

	if err != nil {
		return nil, err
	}

	company := CompanyResponse{
		STD_CODE:     qf.STD_CODE,
		EMAIL:     		qf.EMAIL,
		FULLNAME:     qf.FULLNAME,
		COMPANY:     qf.COMPANY,
		CREATED:      qf.CREATED,
		MODIFIED:     qf.MODIFIED,
	}

	return &company, nil
}

func (s *studentServices) AddCommpany(request CompanyRequest) (*CompanyResponse, error) {
	company := CompanyResponse{}
	err := s.studentRepo.AddCommpany(request.STD_CODE, request.EMAIL, request.FULLNAME, request.COMPANY)

	if err != nil {
		qf, err := s.GetCommpany(request.STD_CODE, request.EMAIL)
		if err == nil {
			company = CompanyResponse{
				STD_CODE:     qf.STD_CODE,
				EMAIL:     		qf.EMAIL,
				FULLNAME:     qf.FULLNAME,
				COMPANY:     qf.COMPANY,
				CREATED:      qf.CREATED,
				MODIFIED:     qf.MODIFIED,
			}

			return &company, err
		}
		return nil, err
	}

	qf, err := s.GetCommpany(request.STD_CODE, request.EMAIL)
	if err != nil {
		return nil, err
	}

	company = CompanyResponse{
		STD_CODE:     qf.STD_CODE,
		EMAIL:     		qf.EMAIL,
		FULLNAME:     qf.FULLNAME,
		COMPANY:     qf.COMPANY,
		CREATED:      qf.CREATED,
		MODIFIED:     qf.MODIFIED,
	}

	return &company, nil
}
