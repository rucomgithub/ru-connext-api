package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

func (g *gradeServices) GradeYear(requestBody GradeRequest) (*GradeResponse, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	gradeResponse := GradeResponse{
		STD_CODE: requestBody.STD_CODE,
		YEAR:     requestBody.YEAR,
		RECORD:   []gradeRecord{},
	}

	key := "grade::" + requestBody.STD_CODE + "-" + requestBody.YEAR
	gradeCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(gradeCache), &gradeResponse)
		fmt.Println("cache-grade")
		return &gradeResponse, nil
	}

	fmt.Println("database-grade")

	gradeRepo, err := g.gradeRepo.GetGradeYear(requestBody.STD_CODE, requestBody.YEAR)
	if err != nil {
		log.Println(err.Error())
		return &gradeResponse, errs.NewUnExpectedError()
	}

	gradeRec := []gradeRecord{}
	for _, c := range *gradeRepo {
		gradeRec = append(gradeRec, gradeRecord{
			REGIS_YEAR:     c.REGIS_YEAR,
			REGIS_SEMESTER: c.REGIS_SEMESTER,
			COURSE_NO:      c.COURSE_NO,
			CREDIT:         c.CREDIT,
			GRADE:          c.GRADE,
		})
	}

	if len(gradeRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลผลการศึกษา")
	}

	gradeResponse = GradeResponse{
		STD_CODE: requestBody.STD_CODE,
		YEAR:     requestBody.YEAR,
		RECORD:   gradeRec,
	}

	gradeJSON, _ := json.Marshal(&gradeResponse)
	timeNow := time.Now()
	redisCachegrade := time.Unix(timeNow.Add(time.Minute*30).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, gradeJSON, redisCachegrade.Sub(timeNow)).Err()

	return &gradeResponse, nil
}

func (g *gradeServices) GradeAll(std_code string) (*GradeResponse, error) {

	if std_code == "" {
		return nil, errs.NewBadRequestError("ระบุรหัสนักศึกษาให้ถูกต้อง")
	}

	gradeResponse := GradeResponse{
		STD_CODE: std_code,
		YEAR:     "",
		RECORD:   []gradeRecord{},
	}

	key := "grade::" + std_code
	gradeCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(gradeCache), &gradeResponse)
		fmt.Println("cache-grade")
		return &gradeResponse, nil
	}

	fmt.Println("database-grade")

	gradeRepo, err := g.gradeRepo.GetGradeAll(std_code)
	if err != nil {
		log.Println(err.Error())
		return &gradeResponse, errs.NewUnExpectedError()
	}

	gradeRec := []gradeRecord{}
	for _, c := range *gradeRepo {
		gradeRec = append(gradeRec, gradeRecord{
			REGIS_YEAR:     c.REGIS_YEAR,
			REGIS_SEMESTER: c.REGIS_SEMESTER,
			COURSE_NO:      c.COURSE_NO,
			CREDIT:         c.CREDIT,
			GRADE:          c.GRADE,
		})
	}

	if len(gradeRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลผลการศึกษา")
	}

	gradeResponse = GradeResponse{
		STD_CODE: std_code,
		YEAR:     "",
		RECORD:   gradeRec,
	}

	gradeJSON, _ := json.Marshal(&gradeResponse)
	timeNow := time.Now()
	redisCachegrade := time.Unix(timeNow.Add(time.Minute*30).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, gradeJSON, redisCachegrade.Sub(timeNow)).Err()

	return &gradeResponse, nil
}
