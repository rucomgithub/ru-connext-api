package students

func (s *studentServices) GetStudentAll() (*[]StudentResponse, error) {

	students, err := s.studentRepo.GetStudentAll()

	if err != nil {
		return nil, err
	}

	studentRec := []StudentResponse{}

	for _, c := range *students {
		studentRec = append(studentRec, StudentResponse{
			STUDENT_CODE: c.STD_CODE,
		})
	}

	return &studentRec, nil
}
