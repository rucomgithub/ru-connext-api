package mr30s

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func filterSetGroupByCourseNo(data []mr30Record, filterCourseNo func(string) bool) []mr30Record {

	fltd := make([]mr30Record, 0)

	for _, e := range data {

		if filterCourseNo(e.COURSE_NO) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}

func (mr30 *mr30Services) GetMr30(course_year, course_semester string) (*Mr30Response, error) {

	mr30Response := Mr30Response{
		COURSE_YEAR:     "",
		COURSE_SEMESTER: "",
		RECORD:          []mr30Record{},
	}

	key := "mr30::" + course_year + "/" + course_semester
	mr30Cache, err := mr30.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(mr30Cache), &mr30Response)
		fmt.Println("cache")
		return &mr30Response, nil
	}

	fmt.Println("database")

	mr30Repo, err := mr30.mr30Repo.GetMr30(course_year, course_semester)
	if err != nil {
		return &mr30Response, err
	}

	mr30Rec := []mr30Record{}
	for _, c := range *mr30Repo {

		mr30Rec = append(mr30Rec, mr30Record{
			ID:                   c.ID,
			COURSE_YEAR:          c.COURSE_YEAR,
			COURSE_SEMESTER:      c.COURSE_SEMESTER,
			COURSE_NO:            c.COURSE_NO,
			COURSE_METHOD:        c.COURSE_METHOD,
			COURSE_METHOD_NUMBER: c.COURSE_METHOD_NUMBER,
			DAY_CODE:             c.DAY_CODE,
			TIME_CODE:            c.TIME_CODE,
			ROOM_GROUP:           c.ROOM_GROUP,
			INSTR_GROUP:          c.INSTR_GROUP,
			COURSE_CREDIT:        c.COURSE_CREDIT,
			COURSE_METHOD_DETAIL: c.COURSE_METHOD_DETAIL,
			DAY_NAME_S:           c.DAY_NAME_S,
			TIME_PERIOD:          c.TIME_PERIOD,
			COURSE_ROOM:          c.COURSE_ROOM,
			COURSE_INSTRUCTOR:    c.COURSE_INSTRUCTOR,
			SHOW_RU30:            c.SHOW_RU30,
			COURSE_PR:            c.COURSE_PR,
			COURSE_COMMENT:       c.COURSE_COMMENT,
			COURSE_EXAMDATE:      c.COURSE_EXAMDATE,
		})
	}

	mr30Response = Mr30Response{
		COURSE_YEAR:     course_year,
		COURSE_SEMESTER: course_semester,
		RECORD:          mr30Rec,
	}

	if len(mr30Rec) != 0 {
		mr30JSON, _ := json.Marshal(&mr30Response)
		timeNow := time.Now()
		redisCacheMr30 := time.Unix(timeNow.Add(time.Hour*1).Unix(), 0)
		_ = mr30.redis_cache.Set(ctx, key, mr30JSON, redisCacheMr30.Sub(timeNow)).Err()
	}

	return &mr30Response, nil
}

func (mr30 *mr30Services) GetMr30Year() (*Mr30YearResponse, error) {

	mr30YearResponse := Mr30YearResponse{
		RECORDYEAR: []mr30YearRecord{},
	}

	key := "mr30::yearall"
	mr30Cache, err := mr30.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(mr30Cache), &mr30YearResponse)
		fmt.Println("cache")
		return &mr30YearResponse, nil
	}

	fmt.Println("database")

	mr30YearRepo, err := mr30.mr30Repo.GetMr30Year()
	if err != nil {
		return &mr30YearResponse, err
	}

	mr30YearRec := []mr30YearRecord{}

	for _, c := range *mr30YearRepo {

		mr30YearRec = append(mr30YearRec, mr30YearRecord{
			COURSE_YEAR:     c.COURSE_YEAR,
			COURSE_SEMESTER: c.COURSE_SEMESTER,
		})
	}

	mr30YearResponse = Mr30YearResponse{
		RECORDYEAR: mr30YearRec,
	}

	if len(mr30YearRec) != 0 {
		mr30JSON, _ := json.Marshal(&mr30YearResponse)
		timeNow := time.Now()
		redisCacheMr30 := time.Unix(timeNow.Add(time.Minute*1).Unix(), 0)
		_ = mr30.redis_cache.Set(ctx, key, mr30JSON, redisCacheMr30.Sub(timeNow)).Err()
	}

	return &mr30YearResponse, nil
}

func (mr30 *mr30Services) GetMr30Searching(course_year, course_semester, course_no string) (*Mr30Response, error) {

	mr30Response := Mr30Response{
		COURSE_YEAR:     "",
		COURSE_SEMESTER: "",
		RECORD:          []mr30Record{},
	}

	key := "mr30::" + course_year + "/" + course_semester
	mr30Cache, err := mr30.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(mr30Cache), &mr30Response)

		prefix := course_no
		res := filterSetGroupByCourseNo(mr30Response.RECORD, func(s string) bool {
			return strings.HasPrefix(s, strings.ToUpper(prefix))
		})

		mr30Response := Mr30Response{
			COURSE_YEAR:     course_year,
			COURSE_SEMESTER: course_semester,
			RECORD:          res,
		}

		return &mr30Response, nil
	}

	mr30Repo, err := mr30.mr30Repo.GetMr30(course_year, course_semester)
	if err != nil {
		return &mr30Response, err
	}

	mr30Rec := []mr30Record{}
	for _, c := range *mr30Repo {

		mr30Rec = append(mr30Rec, mr30Record{
			ID:                   c.ID,
			COURSE_YEAR:          c.COURSE_YEAR,
			COURSE_SEMESTER:      c.COURSE_SEMESTER,
			COURSE_NO:            c.COURSE_NO,
			COURSE_METHOD:        c.COURSE_METHOD,
			COURSE_METHOD_NUMBER: c.COURSE_METHOD_NUMBER,
			DAY_CODE:             c.DAY_CODE,
			TIME_CODE:            c.TIME_CODE,
			ROOM_GROUP:           c.ROOM_GROUP,
			INSTR_GROUP:          c.INSTR_GROUP,
			COURSE_CREDIT:        c.COURSE_CREDIT,
			COURSE_METHOD_DETAIL: c.COURSE_METHOD_DETAIL,
			DAY_NAME_S:           c.DAY_NAME_S,
			TIME_PERIOD:          c.TIME_PERIOD,
			COURSE_ROOM:          c.COURSE_ROOM,
			COURSE_INSTRUCTOR:    c.COURSE_INSTRUCTOR,
			SHOW_RU30:            c.SHOW_RU30,
			COURSE_PR:            c.COURSE_PR,
			COURSE_COMMENT:       c.COURSE_COMMENT,
			COURSE_EXAMDATE:      c.COURSE_EXAMDATE,
		})
	}

	prefix := course_no
	res := filterSetGroupByCourseNo(mr30Rec, func(s string) bool {
		return strings.HasPrefix(s, strings.ToUpper(prefix))
	})

	mr30Response = Mr30Response{
		COURSE_YEAR:     course_year,
		COURSE_SEMESTER: course_semester,
		RECORD:          res,
	}

	return &mr30Response, nil
}

func (mr30 *mr30Services) GetMr30Pagination(course_year, course_semester, limit, offset string) (*Mr30Response, error) {

	mr30Response := Mr30Response{
		COURSE_YEAR:     "",
		COURSE_SEMESTER: "",
		RECORD:          []mr30Record{},
	}

	mr30RecPage := []mr30Record{}

	key := "mr30::" + course_year + "/" + course_semester
	mr30Cache, err := mr30.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(mr30Cache), &mr30Response)

		start, _ := strconv.Atoi(offset)
		end, _ := strconv.Atoi(limit)
		for i := start; i < end; i++ {
			mr30RecPage = append(mr30RecPage, mr30Response.RECORD[i])
		}

		mr30ResponseRedisCach := Mr30Response{
			COURSE_YEAR:     mr30Response.COURSE_YEAR,
			COURSE_SEMESTER: mr30Response.COURSE_SEMESTER,
			RECORD:          mr30RecPage,
		}

		return &mr30ResponseRedisCach, nil
	}

	mr30Repo, err := mr30.mr30Repo.GetMr30(course_year, course_semester)
	if err != nil {
		return &mr30Response, err
	}

	mr30Rec := []mr30Record{}
	for _, c := range *mr30Repo {

		mr30Rec = append(mr30Rec, mr30Record{
			ID:                   c.ID,
			COURSE_YEAR:          c.COURSE_YEAR,
			COURSE_SEMESTER:      c.COURSE_SEMESTER,
			COURSE_NO:            c.COURSE_NO,
			COURSE_METHOD:        c.COURSE_METHOD,
			COURSE_METHOD_NUMBER: c.COURSE_METHOD_NUMBER,
			DAY_CODE:             c.DAY_CODE,
			TIME_CODE:            c.TIME_CODE,
			ROOM_GROUP:           c.ROOM_GROUP,
			INSTR_GROUP:          c.INSTR_GROUP,
			COURSE_CREDIT:        c.COURSE_CREDIT,
			COURSE_METHOD_DETAIL: c.COURSE_METHOD_DETAIL,
			DAY_NAME_S:           c.DAY_NAME_S,
			TIME_PERIOD:          c.TIME_PERIOD,
			COURSE_ROOM:          c.COURSE_ROOM,
			COURSE_INSTRUCTOR:    c.COURSE_INSTRUCTOR,
			SHOW_RU30:            c.SHOW_RU30,
			COURSE_PR:            c.COURSE_PR,
			COURSE_COMMENT:       c.COURSE_COMMENT,
			COURSE_EXAMDATE:      c.COURSE_EXAMDATE,
		})
	}

	mr30Response = Mr30Response{
		COURSE_YEAR:     course_year,
		COURSE_SEMESTER: course_semester,
		RECORD:          mr30Rec,
	}

	if len(mr30Rec) != 0 {
		mr30JSON, _ := json.Marshal(&mr30Response)
		timeNow := time.Now()
		redisCacheMr30 := time.Unix(timeNow.Add(time.Hour*1).Unix(), 0)
		_ = mr30.redis_cache.Set(ctx, key, mr30JSON, redisCacheMr30.Sub(timeNow)).Err()
	}

	start, _ := strconv.Atoi(offset)
	end, _ := strconv.Atoi(limit)
	for i := start; i < end; i++ {
		mr30RecPage = append(mr30RecPage, mr30Rec[i])
	}

	mr30Response = Mr30Response{
		COURSE_YEAR:     course_year,
		COURSE_SEMESTER: course_semester,
		RECORD:          mr30RecPage,
	}

	return &mr30Response, nil
}
