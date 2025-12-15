package masterservice

import (
	"RU-Smart-Workspace/ru-smart-api/middlewares"
	"RU-Smart-Workspace/ru-smart-api/domain/entities"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (s *studentServices) GetStudentSuccessCheck(token string) (studentSuccessResponse *StudentSuccessService, err error) {
	student := StudentSuccessService{}

	_, err = middlewares.VerifyCertificateToken("accessToken", token, s.redis_cache)
	if err != nil {
		err = errors.New("ข้อมูลตรวจสอบ Certificate หมดอายุ.")
		return nil, err
	}

	fmt.Println(token)

	claim, err := middlewares.GetCertificateClaims(token)

	if err != nil {
		err = errors.New("ข้อมูลตรวจสอบ Certificate ไม่ถูกต้อง.")
		return nil, err
	}

	studentCode := claim.StudentCode

	sp, err := s.studentRepo.GetStudentSuccess(studentCode)

	if err != nil {
		err = errors.New("ไม่พบข้อมูล Certificate.")
		return &student, err
	}

	student = StudentSuccessService{
		STD_CODE        : sp.STD_CODE,
		NAME_THAI       : sp.NAME_THAI,
		NAME_ENG        : sp.NAME_ENG,
		YEAR            : sp.YEAR,
		SEMESTER        : sp.SEMESTER,
		CURR_NAME       : sp.CURR_NAME,
		CURR_ENG   		: sp.CURR_ENG,
		THAI_NAME       : sp.THAI_NAME,
		ENG_NAME        : sp.ENG_NAME,
		MAJOR_NAME      : sp.MAJOR_NAME,
		MAJOR_ENG       : sp.MAJOR_ENG,
		MAIN_MAJOR_THAI : sp.MAIN_MAJOR_THAI,
		MAIN_MAJOR_ENG  : sp.MAIN_MAJOR_ENG,
		PLAN            : sp.PLAN,
		GPA             : sp.GPA,
		CONFERENCE_NO   : sp.CONFERENCE_NO,
		SERIAL_NO       : sp.SERIAL_NO,
		CONFERENCE_DATE : sp.CONFERENCE_DATE,
		ADMIT_DATE  	: sp.ADMIT_DATE,
		ADMIT_DATE_EN  	: sp.ADMIT_DATE_EN,
		GRADUATED_DATE  : sp.GRADUATED_DATE,
		GRADUATED_DATE_EN  : sp.GRADUATED_DATE_EN,
		CONFIRM_DATE    : sp.CONFIRM_DATE,
		MOBILE    : sp.MOBILE,
		EMAIL    : sp.EMAIL,
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
		STD_CODE        : sp.STD_CODE,
		NAME_THAI       : sp.NAME_THAI,
		NAME_ENG        : sp.NAME_ENG,
		YEAR            : sp.YEAR,
		SEMESTER        : sp.SEMESTER,
		CURR_NAME       : sp.CURR_NAME,
		CURR_ENG   		: sp.CURR_ENG,
		THAI_NAME       : sp.THAI_NAME,
		ENG_NAME        : sp.ENG_NAME,
		MAJOR_NAME      : sp.MAJOR_NAME,
		MAJOR_ENG       : sp.MAJOR_ENG,
		MAIN_MAJOR_THAI : sp.MAIN_MAJOR_THAI,
		MAIN_MAJOR_ENG  : sp.MAIN_MAJOR_ENG,
		PLAN            : sp.PLAN,
		GPA             : sp.GPA,
		CONFERENCE_NO   : sp.CONFERENCE_NO,
		SERIAL_NO       : sp.SERIAL_NO,
		CONFERENCE_DATE : sp.CONFERENCE_DATE,
		ADMIT_DATE  	: sp.ADMIT_DATE,
		ADMIT_DATE_EN  	: sp.ADMIT_DATE_EN,
		GRADUATED_DATE  : sp.GRADUATED_DATE,
		GRADUATED_DATE_EN  : sp.GRADUATED_DATE_EN,
		CONFIRM_DATE    : sp.CONFIRM_DATE,
		MOBILE    : sp.MOBILE,
		EMAIL    : sp.EMAIL,
	}

	studentSuccessResponse = &student

	studentProfileJSON, _ := json.Marshal(studentSuccessResponse)
	timeNow := time.Now()
	redisCacheStudentProfile := time.Unix(timeNow.Add(time.Second*10).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, studentProfileJSON, redisCacheStudentProfile.Sub(timeNow)).Err()

	return studentSuccessResponse, nil
}

func (s *studentServices) GetStudentRequestSuccess(studentCode string) (studentSuccessResponse *StudentRequestSuccessService, err error) {

	student := StudentRequestSuccessService{}

	key := studentCode + "::requestsuccess"
	studentCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(studentCache), &student)
		return &student, nil
	}

	sp, err := s.studentRepo.GetStudentRequestSuccess(studentCode)

	if err != nil {
		return &student, err
	}

	student = StudentRequestSuccessService {
		ENROLL_YEAR            : sp.ENROLL_YEAR,
		ENROLL_SEMESTER        : sp.ENROLL_SEMESTER,
		STD_CODE               : sp.STD_CODE,
		PRENAME_THAI_S         : sp.PRENAME_THAI_S,
		PRENAME_ENG_S          : sp.PRENAME_ENG_S,
		FIRST_NAME             : sp.FIRST_NAME,
		LAST_NAME              : sp.LAST_NAME,
		FIRST_NAME_ENG         : sp.FIRST_NAME_ENG,
		LAST_NAME_ENG          : sp.LAST_NAME_ENG,
		THAI_NAME              : sp.THAI_NAME,
		PLAN_NO                : sp.PLAN_NO,
		SEX                    : sp.SEX,
		REGINAL_NAME           : sp.REGINAL_NAME,
		SUBSIDY_NAME           : sp.SUBSIDY_NAME,
		STATUS_NAME_THAI       : sp.STATUS_NAME_THAI,
		BIRTH_DATE             : sp.BIRTH_DATE,
		STD_ADDR               : sp.STD_ADDR,
		ADDR_TEL               : sp.ADDR_TEL,
		JOB_POSITION           : sp.JOB_POSITION,
		STD_OFFICE             : sp.STD_OFFICE,
		OFFICE_TEL             : sp.OFFICE_TEL,
		DEGREE_NAME            : sp.DEGREE_NAME,
		BSC_DEGREE_NO          : sp.BSC_DEGREE_NO,
		BSC_DEGREE_THAI_NAME   : sp.BSC_DEGREE_THAI_NAME,
		BSC_INSTITUTE_NO       : sp.BSC_INSTITUTE_NO,
		INSTITUTE_THAI_NAME    : sp.INSTITUTE_THAI_NAME,
		CK_CERT_NO             : sp.CK_CERT_NO,
		CHK_CERT_NAME_THAI     : sp.CHK_CERT_NAME_THAI,
		SUCCESS_YEAR           : sp.SUCCESS_YEAR,
		SUCCESS_SEMESTER       : sp.SUCCESS_SEMESTER,
		NAME_THAI      			: sp.NAME_THAI,
		NAME_ENG       			: sp.NAME_ENG,
		THESIS_THAI    			: sp.THESIS_THAI,
		THESIS_ENG     			: sp.THESIS_ENG,
		DEGREE         			: sp.DEGREE,
		GRADES            		: sp.GRADES,
		REGISTRATION          	: sp.REGISTRATION,
		ADDRESS               	: sp.ADDRESS,
		CREATED                : sp.CREATED,
		MODIFIED               : sp.MODIFIED,
		THESIS_THAI_TITLE      : sp.THESIS_THAI_TITLE,
		THESIS_ENG_TITLE       : sp.THESIS_ENG_TITLE,
		THESIS_TYPE     		: sp.THESIS_TYPE,
		SIMILARITY     			: sp.SIMILARITY,
	}	

	studentSuccessResponse = &student

	studentProfileJSON, _ := json.Marshal(studentSuccessResponse)
	timeNow := time.Now()
	redisCacheStudentProfile := time.Unix(timeNow.Add(time.Second*10).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, studentProfileJSON, redisCacheStudentProfile.Sub(timeNow)).Err()

	return studentSuccessResponse, nil
}

func (s *studentServices) AddRequestSuccess(request *entities.RequestSuccess) error {

	err := s.studentRepo.CreateRequestSuccess(request)

	if err != nil {
		return err
	}

	return nil
}

func (s *studentServices) EditRequestSuccess(request *entities.RequestSuccess) error {

	err := s.studentRepo.UpdateRequestSuccess(request)

	if err != nil {
		return err
	}

	return nil
}

