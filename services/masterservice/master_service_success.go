package masterservice

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (s *studentServices) GetStudentSuccessCheck(token string) (studentSuccessResponse *StudentSuccessService, err error) {
	student := StudentSuccessService{}

	_, err = middlewares.VerifyCertificateToken("accessToken", token, s.redis_cache)
	if err != nil {
		err = errors.New("Token Certificate หมดอายุ.")
		return nil, err
	}

	fmt.Println(token)

	claim, err := middlewares.GetCertificateClaims(token)

	if err != nil {
		err = errors.New("Token Certificate ไม่ดูกต้อง.")
		return nil, err
	}

	studentCode := claim.StudentCode

	sp, err := s.studentRepo.GetStudentSuccess(studentCode)

	if err != nil {
		err = errors.New("ไม่พบข้อมูล Certificate.")
		return &student, err
	}

	student = StudentSuccessService{
		STD_CODE:        sp.STD_CODE,
		NAME_THAI:       sp.NAME_THAI,
		NAME_ENG:        sp.NAME_ENG,
		YEAR:            sp.YEAR,
		SEMESTER:        sp.SEMESTER,
		CURR_NAME:       sp.CURR_NAME,
		MAJOR_NAME_THAI: sp.MAJOR_NAME_THAI,
		MAJOR_NAME:      sp.MAJOR_NAME,
		PLAN:            sp.PLAN,
		CONFERENCE_NO:   sp.CONFERENCE_NO,
		SERIAL_NO:       sp.SERIAL_NO,
		CONFERENCE_DATE: sp.CONFERENCE_DATE,
		GRADUATED_DATE:  sp.GRADUATED_DATE,
		CONFIRM_DATE:    sp.CONFIRM_DATE,
	}

	return &student, nil
}

func (s *studentServices) GetStudentSuccess(studentCode string) (studentSuccessResponse *StudentSuccessService, err error) {

	student := StudentSuccessService{}

	key := studentCode + "::success"
	studentCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(studentCache), &student)
		return &student, nil
	}

	sp, err := s.studentRepo.GetStudentSuccess(studentCode)

	if err != nil {
		return &student, err
	}

	student = StudentSuccessService{
		STD_CODE:        sp.STD_CODE,
		NAME_THAI:       sp.NAME_THAI,
		NAME_ENG:        sp.NAME_ENG,
		YEAR:            sp.YEAR,
		SEMESTER:        sp.SEMESTER,
		CURR_NAME:       sp.CURR_NAME,
		MAJOR_NAME_THAI: sp.MAJOR_NAME_THAI,
		MAJOR_NAME:      sp.MAJOR_NAME,
		PLAN:            sp.PLAN,
		CONFERENCE_NO:   sp.CONFERENCE_NO,
		SERIAL_NO:       sp.SERIAL_NO,
		CONFERENCE_DATE: sp.CONFERENCE_DATE,
		GRADUATED_DATE:  sp.GRADUATED_DATE,
		CONFIRM_DATE:    sp.CONFIRM_DATE,
	}

	studentSuccessResponse = &student

	studentProfileJSON, _ := json.Marshal(studentSuccessResponse)
	timeNow := time.Now()
	redisCacheStudentProfile := time.Unix(timeNow.Add(time.Second*10).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, studentProfileJSON, redisCacheStudentProfile.Sub(timeNow)).Err()

	return studentSuccessResponse, nil
}
