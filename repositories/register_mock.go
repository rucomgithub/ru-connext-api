package repositories

import "github.com/stretchr/testify/mock"

type registerRepoInterfaceMock struct {
	mock.Mock
}

type RegisterRecord struct {
	YEAR      string `json:"year"`
	SEMESTER  string `json:"semester"`
	COURSE_NO string `json:"course_no"`
	CREDIT    string `json:"credit"`
}

func NewRegisterRepoMock() *registerRepoInterfaceMock {
	return &registerRepoInterfaceMock{}
}

func (m *registerRepoInterfaceMock) GetRegister(std_code, year string) (*[]RegisterRepo, error) {
	args := m.Called(std_code, year)
	return args.Get(0).(*[]RegisterRepo), args.Error(1)
}

func (m *registerRepoInterfaceMock) GetRegisterYear(std_code string) (*[]RegisterYearRepo, error) {
	return nil, nil
}
func (m *registerRepoInterfaceMock) GetRegisterGroupYearSemester(std_code string) (*[]RegisterYearSemesterRepo, error) {
	return nil, nil
}
func (m *registerRepoInterfaceMock) GetRegisterMr30(year, semester, studentCode string) (*[]RegisterMr30Repo, error) {
	return nil, nil
}
