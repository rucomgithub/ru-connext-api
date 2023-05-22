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

func (m *registerRepoInterfaceMock) GetRegisterAll(std_code, year string) (*[]RegisterRepo, error) {
	args := m.Called(std_code, year)
	return args.Get(0).(*[]RegisterRepo), args.Error(1)
}

func (m *registerRepoInterfaceMock) GetListYearAll(std_code string) (*[]YearRepo, error) {
	return nil, nil
}
func (m *registerRepoInterfaceMock) GetListYearSemesterAll(std_code string) (*[]YearSemesterRepo, error) {
	return nil, nil
}
func (m *registerRepoInterfaceMock) GetScheduleAll(year, semester, studentCode string) (*[]ScheduleRepo, error) {
	return nil, nil
}
