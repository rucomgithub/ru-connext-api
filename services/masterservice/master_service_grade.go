package masterservice

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (s *studentServices) GetGradeAll(std_code string) (*GradeResponse, error) {

	gradeResponse := GradeResponse{
		STD_CODE:       std_code,
		YEAR:           "",
		SUMMARY_CREDIT: 0,
		GPA:            0.0,
		GRADE:          []GradeResponseRepo{},
	}

	key := "v2-master-grade::" + std_code
	fmt.Println(key)
	gradeCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(gradeCache), &gradeResponse)
		fmt.Println("cache-master-grade")
		return &gradeResponse, nil
	}

	fmt.Println("database-master-grade")

	gradeRepo, err := s.studentRepo.GetGradeAll(std_code)

	if err != nil {
		log.Println(err.Error())
		return &gradeResponse, errs.NewUnExpectedError()
	}

	gpaRepo, err := s.studentRepo.GetGpaAll(std_code)

	if err != nil {
		log.Println(err.Error())
		return &gradeResponse, errs.NewUnExpectedError()
	}

	gradeRec := []GradeResponseRepo{}
	for _, c := range *gradeRepo {
		gradeRec = append(gradeRec, GradeResponseRepo{
			YEAR:      c.YEAR,
			SEMESTER:  c.SEMESTER,
			STD_CODE:  c.STD_CODE,
			COURSE_NO: c.COURSE_NO,
			CREDIT:    c.CREDIT,
			GRADE:     c.GRADE,
			COURSE_TYPE_NO : c.COURSE_TYPE_NO,
			THAI_DESCRIPTION : c.THAI_DESCRIPTION,
		})
	}

	if len(gradeRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	gradeResponse = GradeResponse{
		STD_CODE:       std_code,
		YEAR:           "",
		SUMMARY_CREDIT: gpaRepo.SUMMARY_CREDIT,
		GPA:            gpaRepo.GPA,
		GRADE:          gradeRec,
	}

	registerJSON, _ := json.Marshal(&gradeResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &gradeResponse, nil
}

func (s *studentServices) GetGradeByYear(std_code, year string) (*GradeResponse, error) {

	gradeResponse := GradeResponse{
		STD_CODE:       std_code,
		YEAR:           "",
		SUMMARY_CREDIT: 0,
		GPA:            0.0,
		GRADE:          []GradeResponseRepo{},
	}

	key := "v2-master-grade-year::" + std_code + "::" + year
	fmt.Println(key)
	gradeCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(gradeCache), &gradeResponse)
		fmt.Println("cache-master-grade-year")
		return &gradeResponse, nil
	}

	fmt.Println("database-master-grade-year")

	gradeRepo, err := s.studentRepo.GetGradeByYear(std_code, year)

	if err != nil {
		log.Println(err.Error()+"grade")
		return &gradeResponse, errs.NewUnExpectedError()
	}

	gpaRepo, err := s.studentRepo.GetGpaByYear(std_code, year)

	if err != nil {
		log.Println(err.Error()+"gpa")
		return &gradeResponse, errs.NewUnExpectedError()
	}

	gradeRec := []GradeResponseRepo{}
	for _, c := range *gradeRepo {
		gradeRec = append(gradeRec, GradeResponseRepo{
			YEAR:      c.YEAR,
			SEMESTER:  c.SEMESTER,
			STD_CODE:  c.STD_CODE,
			COURSE_NO: c.COURSE_NO,
			CREDIT:    c.CREDIT,
			GRADE:     c.GRADE,
		})
	}

	if len(gradeRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	gradeResponse = GradeResponse{
		STD_CODE:       std_code,
		YEAR:           year,
		SUMMARY_CREDIT: gpaRepo.SUMMARY_CREDIT,
		GPA:            gpaRepo.GPA,
		GRADE:          gradeRec,
	}

	registerJSON, _ := json.Marshal(&gradeResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*1).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &gradeResponse, nil
}
