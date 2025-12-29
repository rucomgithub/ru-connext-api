package masterservice

func (s *studentServices) GetQualification(std_code string) (*QualificationResponse, error) {

	qf, err := s.studentRepo.GetQualification(std_code)

	if err != nil {
		return nil, err
	}

	qualification := QualificationResponse{
		STD_CODE:     qf.STD_CODE,
		REQUEST_DATE: qf.REQUEST_DATE,
		OPERATE_DATE: qf.OPERATE_DATE,
		CONFIRM_DATE: qf.CONFIRM_DATE,
		CANCEL_DATE:  qf.CANCEL_DATE,
		STATUS:       qf.STATUS,
		CREATED:      qf.CREATED,
		MODIFIED:     qf.MODIFIED,
		DESCRIPTION:  qf.DESCRIPTION,
	}

	return &qualification, nil
}

func (s *studentServices) AddQualification(std_code string) (*QualificationResponse, error) {

	err := s.studentRepo.AddQualification(std_code)

	if err != nil {
		return nil, err
	}

	qf, err := s.GetQualification(std_code)
	if err != nil {
		return nil, err
	}

	qualification := QualificationResponse{
		STD_CODE:     qf.STD_CODE,
		REQUEST_DATE: qf.REQUEST_DATE,
		OPERATE_DATE: qf.OPERATE_DATE,
		CONFIRM_DATE: qf.CONFIRM_DATE,
		CANCEL_DATE:  qf.CANCEL_DATE,
		STATUS:       qf.STATUS,
		CREATED:      qf.CREATED,
		MODIFIED:     qf.MODIFIED,
		DESCRIPTION:  qf.DESCRIPTION,
	}

	return &qualification, nil
}
