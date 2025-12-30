package students

import (
	"encoding/json"
	"fmt"
	"time"
)

func (s *studentServices) GetRegister(studentCode, courseYear, courseSemester string) (*RegisterResponse, error) {

	register := &RegisterResponse{
		STUDENT_CODE:    studentCode,
		COURSE_YEAR:     courseYear,
		COURSE_SEMESTER: courseSemester,
		REGISTER:        []RegisterResponseFromDB{},
	}

	key := studentCode + "::register"
	registerCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {

		_ = json.Unmarshal([]byte(registerCache), &register)
		return register, nil
	}

	registerRepo, err := s.studentRepo.GetRegister(studentCode, courseYear, courseSemester)
	if err != nil {
		return register, err
	}

	registerTemp := []RegisterResponseFromDB{}
	for _, r := range *registerRepo {

		registerTemp = append(registerTemp, RegisterResponseFromDB{
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

	register = &RegisterResponse{
		STUDENT_CODE:    studentCode,
		COURSE_YEAR:     courseYear,
		COURSE_SEMESTER: courseSemester,
		REGISTER:        registerTemp,
	}

	registerJSON, _ := json.Marshal(register)
	timeNow := time.Now()
	redisCache := time.Unix(timeNow.Add(time.Minute*20).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, registerJSON, redisCache.Sub(timeNow)).Err()

	return register, nil
}

func (s *studentServices) GetRegisterAll(std_code string, year string) (*RegisterAllResponse, error) {

	registerAllResponse := RegisterAllResponse{
		STD_CODE: std_code,
		YEAR:     year,
		RECORD:   []registerAllRecord{},
	}

	key := "v2-register::" + std_code + "-" + year
	registerCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(registerCache), &registerAllResponse)
		fmt.Println("cache")
		return &registerAllResponse, nil
	}

	fmt.Println("database")

	registerAllRepo, err := s.studentRepo.GetRegisterAll(std_code, year)
	if err != nil {
		return &registerAllResponse, err
	}

	registerRec := []registerAllRecord{}
	for _, c := range *registerAllRepo {
		registerRec = append(registerRec, registerAllRecord{
			YEAR:      c.YEAR,
			SEMESTER:  c.SEMESTER,
			COURSE_NO: c.COURSE_NO,
			CREDIT:    c.CREDIT,
		})
	}

	registerAllResponse = RegisterAllResponse{
		STD_CODE: std_code,
		YEAR:     year,
		RECORD:   registerRec,
	}

	if len(registerRec) != 0 {
		registerJSON, _ := json.Marshal(&registerAllResponse)
		timeNow := time.Now()
		redisCacheregister := time.Unix(timeNow.Add(time.Minute*2).Unix(), 0)
		_ = s.redis_cache.Set(ctx, key, registerJSON, redisCacheregister.Sub(timeNow)).Err()
	}

	fmt.Println(registerRec)

	return &registerAllResponse, nil
}
