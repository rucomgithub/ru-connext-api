package masterservice

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (s *studentServices) GetRegisterAll(std_code string) (*RegisterResponse, error) {

	registerResponse := RegisterResponse{
		STD_CODE: std_code,
		YEAR:     "",
		REGISTER: []RegisterResponseRepo{},
	}

	key := "master-register::" + std_code
	fmt.Println(key)
	registerCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(registerCache), &registerResponse)
		fmt.Println("cache-master-register")
		return &registerResponse, nil
	}

	fmt.Println("database-master-register")

	registerRepo, err := s.studentRepo.GetRegisterAll(std_code)

	if err != nil {
		log.Println(err.Error())
		return &registerResponse, errs.NewUnExpectedError()
	}

	registerRec := []RegisterResponseRepo{}
	for _, c := range *registerRepo {
		registerRec = append(registerRec, RegisterResponseRepo{
			YEAR:      c.YEAR,
			SEMESTER:  c.SEMESTER,
			STD_CODE:  c.STD_CODE,
			COURSE_NO: c.COURSE_NO,
			CREDIT:    c.CREDIT,
		})
	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	registerResponse = RegisterResponse{
		STD_CODE: std_code,
		YEAR:     "",
		REGISTER: registerRec,
	}

	registerJSON, _ := json.Marshal(&registerResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &registerResponse, nil
}

func (s *studentServices) GetRegisterByYear(std_code, year string) (*RegisterResponse, error) {

	registerResponse := RegisterResponse{
		STD_CODE: std_code,
		YEAR:     year,
		REGISTER: []RegisterResponseRepo{},
	}

	key := "master-register::" + std_code + "::" + year
	fmt.Println(key)
	registerCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(registerCache), &registerResponse)
		fmt.Println("cache-master-register-year")
		return &registerResponse, nil
	}

	fmt.Println("database-master-register-year")

	registerRepo, err := s.studentRepo.GetRegisterByYear(std_code, year)

	if err != nil {
		log.Println(err.Error())
		return &registerResponse, errs.NewUnExpectedError()
	}

	registerRec := []RegisterResponseRepo{}
	for _, c := range *registerRepo {
		registerRec = append(registerRec, RegisterResponseRepo{
			YEAR:      c.YEAR,
			SEMESTER:  c.SEMESTER,
			STD_CODE:  c.STD_CODE,
			COURSE_NO: c.COURSE_NO,
			CREDIT:    c.CREDIT,
		})
	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	registerResponse = RegisterResponse{
		STD_CODE: std_code,
		YEAR:     year,
		REGISTER: registerRec,
	}

	registerJSON, _ := json.Marshal(&registerResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &registerResponse, nil
}
