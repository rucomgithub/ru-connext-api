package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

func (g *onDemandServices) OnDemand(onDemandRequest OnDemandRequest) (*OnDemandResponse, error) {

	return nil, nil
}

func (g *onDemandServices) OnDemandAll(requestBody OnDemandRequest) (*OnDemandResponse, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	onDemandResponse := OnDemandResponse{
		COURSE_NO: requestBody.COURSE_NO,
		YEAR:      requestBody.YEAR,
		SEMESTER:  requestBody.SEMESTER,
		RECORD:    []onDemandRecord{},
	}

	key := "ondemane::" + requestBody.COURSE_NO
	ondemandCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		//log.Println(err)
		_ = json.Unmarshal([]byte(ondemandCache), &onDemandResponse)
		fmt.Println("cache::" + key)
		return &onDemandResponse, nil
	}

	fmt.Println("database::" + key)

	ondemandrepo, err := g.onDemandRepo.GetOnDemandAll(requestBody.COURSE_NO, requestBody.YEAR, requestBody.SEMESTER)
	if err != nil {
		log.Println(err.Error())
		return &onDemandResponse, errs.NewUnExpectedError()
	}

	ondemandRec := []onDemandRecord{}
	for _, c := range *ondemandrepo {
		ondemandRec = append(ondemandRec, onDemandRecord{
			STUDY_SEMESTER: c.STUDY_SEMESTER,
			STUDY_YEAR:     c.STUDY_YEAR,
			COURSE_NO:      c.COURSE_NO,
			DAY_CODE:       c.DAY_CODE,
			TIME_CODE:      c.TIME_CODE,
			BUILDING_CODE:  c.BUILDING_CODE,
			ROOM_CODE:      c.ROOM_CODE,
		})
	}

	if len(ondemandRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลรายการ Course on Demand.")
	}

	onDemandResponse = OnDemandResponse{
		COURSE_NO: requestBody.COURSE_NO,
		YEAR:      requestBody.YEAR,
		SEMESTER:  requestBody.SEMESTER,
		COUNT:     len(ondemandRec),
		RECORD:    ondemandRec,
	}

	gradeJSON, _ := json.Marshal(&onDemandResponse)
	timeNow := time.Now()
	redisCachegrade := time.Unix(timeNow.Add(time.Second*10).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, gradeJSON, redisCachegrade.Sub(timeNow)).Err()

	return &onDemandResponse, nil
}
