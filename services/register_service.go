package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

func (register *registerServices) GetRegisterMr30Latest(std_code string) (*RegisterMr30Response, error) {

	yearSemesterRepo, err := register.registerRepo.GetRegisterGroupYearSemester(std_code)
	if err != nil {
		log.Println(err.Error())
		return nil, errs.NewUnExpectedError()
	}

	yearSemesterRec := []registerYearSemesterRecord{}
	for _, c := range *yearSemesterRepo {
		yearSemesterRec = append(yearSemesterRec, registerYearSemesterRecord{
			YEAR:     c.YEAR,
			SEMESTER: c.SEMESTER,
		})
	}

	registerMr30Response := RegisterMr30Response{
		YEAR:     yearSemesterRec[0].YEAR,
		SEMESTER: yearSemesterRec[0].SEMESTER,
		RECORD:   []RegisterMr30Record{},
	}

	key := "registerMr30lates::" + std_code + yearSemesterRec[0].YEAR + yearSemesterRec[0].SEMESTER
	registerCache, err := register.redis_cache.Get(ctx, key).Result()
	if err == nil {

		_ = json.Unmarshal([]byte(registerCache), &registerMr30Response)
		return &registerMr30Response, nil
	}

	fmt.Println("database")

	registerRepo, err := register.registerRepo.GetRegisterMr30(yearSemesterRec[0].YEAR, yearSemesterRec[0].SEMESTER, std_code)

	if err != nil {
		log.Println(err.Error())
		return &registerMr30Response, errs.NewUnExpectedError()
	}

	registerRec := []RegisterMr30Record{}
	for _, r := range *registerRepo {

		registerRec = append(registerRec, RegisterMr30Record{
			ID:                   r.ID,
			COURSE_YEAR:          r.COURSE_YEAR,
			COURSE_SEMESTER:      r.COURSE_SEMESTER,
			COURSE_NO:            r.COURSE_NO,
			COURSE_METHOD:        r.COURSE_METHOD,
			COURSE_METHOD_NUMBER: r.COURSE_METHOD_NUMBER,
			DAY_CODE:             r.DAY_CODE,
			TIME_CODE:            r.TIME_CODE,
			ROOM_GROUP:           r.ROOM_GROUP,
			INSTR_GROUP:          r.INSTR_GROUP,
			COURSE_METHOD_DETAIL: r.COURSE_METHOD_DETAIL,
			DAY_NAME_S:           r.DAY_NAME_S,
			TIME_PERIOD:          r.TIME_PERIOD,
			COURSE_ROOM:          r.COURSE_ROOM,
			COURSE_INSTRUCTOR:    r.COURSE_INSTRUCTOR,
			SHOW_RU30:            r.SHOW_RU30,
			COURSE_CREDIT:        r.COURSE_CREDIT,
			COURSE_PR:            r.COURSE_PR,
			COURSE_COMMENT:       r.COURSE_COMMENT,
			COURSE_EXAMDATE:      r.COURSE_EXAMDATE,
		})

	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	registerMr30Response = RegisterMr30Response{
		YEAR:     yearSemesterRec[0].YEAR,
		SEMESTER: yearSemesterRec[0].SEMESTER,
		RECORD:   registerRec,
	}

	registerJSON, _ := json.Marshal(registerMr30Response)
	timeNow := time.Now()
	redisCache := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = register.redis_cache.Set(ctx, key, registerJSON, redisCache.Sub(timeNow)).Err()

	return &registerMr30Response, nil
}

func (register *registerServices) GetRegisterMr30(std_code string, requestBody RegisterMr30Request) (*RegisterMr30Response, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	registerMr30Response := RegisterMr30Response{
		YEAR:     requestBody.YEAR,
		SEMESTER: requestBody.SEMESTER,
		RECORD:   []RegisterMr30Record{},
	}

	key := "registerMr30::" + std_code + requestBody.YEAR + requestBody.SEMESTER
	registerCache, err := register.redis_cache.Get(ctx, key).Result()
	if err == nil {

		_ = json.Unmarshal([]byte(registerCache), &registerMr30Response)
		return &registerMr30Response, nil
	}

	fmt.Println("database")

	registerRepo, err := register.registerRepo.GetRegisterMr30(requestBody.YEAR, requestBody.SEMESTER, std_code)

	if err != nil {
		log.Println(err.Error())
		return &registerMr30Response, errs.NewUnExpectedError()
	}

	registerRec := []RegisterMr30Record{}
	for _, r := range *registerRepo {

		registerRec = append(registerRec, RegisterMr30Record{
			ID:                   r.ID,
			COURSE_YEAR:          r.COURSE_YEAR,
			COURSE_SEMESTER:      r.COURSE_SEMESTER,
			COURSE_NO:            r.COURSE_NO,
			COURSE_METHOD:        r.COURSE_METHOD,
			COURSE_METHOD_NUMBER: r.COURSE_METHOD_NUMBER,
			DAY_CODE:             r.DAY_CODE,
			TIME_CODE:            r.TIME_CODE,
			ROOM_GROUP:           r.ROOM_GROUP,
			INSTR_GROUP:          r.INSTR_GROUP,
			COURSE_METHOD_DETAIL: r.COURSE_METHOD_DETAIL,
			DAY_NAME_S:           r.DAY_NAME_S,
			TIME_PERIOD:          r.TIME_PERIOD,
			COURSE_ROOM:          r.COURSE_ROOM,
			COURSE_INSTRUCTOR:    r.COURSE_INSTRUCTOR,
			SHOW_RU30:            r.SHOW_RU30,
			COURSE_CREDIT:        r.COURSE_CREDIT,
			COURSE_PR:            r.COURSE_PR,
			COURSE_COMMENT:       r.COURSE_COMMENT,
			COURSE_EXAMDATE:      r.COURSE_EXAMDATE,
		})

	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	registerMr30Response = RegisterMr30Response{
		YEAR:     requestBody.YEAR,
		SEMESTER: requestBody.SEMESTER,
		RECORD:   registerRec,
	}

	registerJSON, _ := json.Marshal(registerMr30Response)
	timeNow := time.Now()
	redisCache := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = register.redis_cache.Set(ctx, key, registerJSON, redisCache.Sub(timeNow)).Err()

	return &registerMr30Response, nil
}

func (register *registerServices) GetRegister(requestBody RegisterRequest) (*RegisterResponse, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	registerResponse := RegisterResponse{
		STD_CODE: requestBody.STD_CODE,
		YEAR:     requestBody.YEAR,
		RECORD:   []registerRecord{},
	}

	key := "register::" + requestBody.STD_CODE + "-" + requestBody.YEAR
	registerCache, err := register.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(registerCache), &registerResponse)
		fmt.Println("cache")
		return &registerResponse, nil
	}

	fmt.Println("database")

	registerRepo, err := register.registerRepo.GetRegister(requestBody.STD_CODE, requestBody.YEAR)
	if err != nil {
		log.Println(err.Error())
		return &registerResponse, errs.NewUnExpectedError()
	}

	registerRec := []registerRecord{}
	for _, c := range *registerRepo {
		registerRec = append(registerRec, registerRecord{
			YEAR:      c.YEAR,
			SEMESTER:  c.SEMESTER,
			COURSE_NO: c.COURSE_NO,
			CREDIT:    c.CREDIT,
		})
	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลการลงทะเบียนเรียน")
	}

	registerResponse = RegisterResponse{
		STD_CODE: requestBody.STD_CODE,
		YEAR:     requestBody.YEAR,
		RECORD:   registerRec,
	}

	registerJSON, _ := json.Marshal(&registerResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = register.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &registerResponse, nil
}

func (register *registerServices) GetRegisterYear(std_code string) (*RegisterYearResponse, error) {

	if std_code == "" {
		return nil, errs.NewBadRequestError("ระบุรหัสนักศึกษาให้ถูกต้อง")
	}

	registerYearResponse := RegisterYearResponse{
		STD_CODE: std_code,
		RECORD:   []registerYearRecord{},
	}

	key := "register-year::" + std_code
	registerCache, err := register.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(registerCache), &registerYearResponse)
		fmt.Println("cache")
		return &registerYearResponse, nil
	}

	fmt.Println("database")

	registerRepo, err := register.registerRepo.GetRegisterYear(std_code)
	if err != nil {
		log.Println(err.Error())
		return nil, errs.NewUnExpectedError()
	}

	registerRec := []registerYearRecord{}
	for _, c := range *registerRepo {
		registerRec = append(registerRec, registerYearRecord{
			YEAR: c.YEAR,
		})
	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลปีการศึกษา")
	}

	registerYearResponse = RegisterYearResponse{
		STD_CODE: std_code,
		RECORD:   registerRec,
	}

	registerJSON, _ := json.Marshal(&registerYearResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = register.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &registerYearResponse, nil
}

func (register *registerServices) GetRegisterGroupYearSemester(std_code string) (*RegisterYearSemesterResponse, error) {

	if std_code == "" {
		return nil, errs.NewBadRequestError("ระบุรหัสนักศึกษาให้ถูกต้อง")
	}

	registerYearSemesterResponse := RegisterYearSemesterResponse{
		STD_CODE: std_code,
		RECORD:   []registerYearSemesterRecord{},
	}

	key := "register-year-semester::" + std_code
	registerCache, err := register.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(registerCache), &registerYearSemesterResponse)
		fmt.Println("cache")
		return &registerYearSemesterResponse, nil
	}

	fmt.Println("database")

	registerRepo, err := register.registerRepo.GetRegisterGroupYearSemester(std_code)
	if err != nil {
		log.Println(err.Error())
		return nil, errs.NewUnExpectedError()
	}

	registerRec := []registerYearSemesterRecord{}
	for _, c := range *registerRepo {
		registerRec = append(registerRec, registerYearSemesterRecord{
			YEAR:     c.YEAR,
			SEMESTER: c.SEMESTER,
		})
	}

	if len(registerRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลปีการศึกษา")
	}

	registerYearSemesterResponse = RegisterYearSemesterResponse{
		STD_CODE: std_code,
		RECORD:   registerRec,
	}

	registerJSON, _ := json.Marshal(&registerYearSemesterResponse)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = register.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()

	return &registerYearSemesterResponse, nil
}
