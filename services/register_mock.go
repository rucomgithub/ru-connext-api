package services

import "github.com/stretchr/testify/mock"

type registerServicesMock struct {
	mock.Mock
}

func NewRegisterServicesMock() *registerServicesMock {
	return &registerServicesMock{}
}

func (m *registerServicesMock) GetRegister(requestBody RegisterRequest) (*RegisterResponse, error) {
	args := m.Called(requestBody)
	return args.Get(0).(*RegisterResponse), args.Error(1)
}

func (m *registerServicesMock) GetRegisterYear(std_code string) (*RegisterYearResponse, error) {
	return nil, nil
}
func (m *registerServicesMock) GetRegisterGroupYearSemester(std_code string) (*RegisterYearSemesterResponse, error) {
	return nil, nil
}
func (m *registerServicesMock) GetRegisterMr30(std_code string, registerMr30Request RegisterMr30Request) (*RegisterMr30Response, error) {
	return nil, nil
}
func (m *registerServicesMock) GetRegisterMr30Latest(std_code string) (*RegisterMr30Response, error) {
	return nil, nil
}
