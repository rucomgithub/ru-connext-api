package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

func (scholarShip *scholarShipServices) GetScholarshipAll(requestBody ScholarShipRequest) (*ScholarShipResponse, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	if requestBody.STD_CODE == "" {
		return nil, errs.NewBadRequestError("ระบุรหัสนักศึกษาให้ถูกต้อง")
	}

	scholarShipResponse := ScholarShipResponse{
		STD_CODE: requestBody.STD_CODE,
		RECORD:   []scholarShipRecord{},
	}
	log.Println(requestBody.STD_CODE)
	key := "scholarShip::" + requestBody.STD_CODE
	scholarShipCache, err := scholarShip.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(scholarShipCache), &scholarShipResponse)
		fmt.Println("cache-grade")
		return &scholarShipResponse, nil
	}

	fmt.Println("database-sholarShip")

	scholarShipRepo, err := scholarShip.scholarShipRepo.GetScholarshipAll(requestBody.STD_CODE)
	if err != nil {
		log.Println(err.Error())
		return &scholarShipResponse, errs.NewUnExpectedError()
	}

	scholarShipRec := []scholarShipRecord{}
	fmt.Println(scholarShipRec)
	for _, c := range *scholarShipRepo {
		scholarShipRec = append(scholarShipRec, scholarShipRecord{
			ID:                  c.ID,
			StudentCode:         c.StudentCode,
			ScholarshipYear:     c.ScholarshipYear,
			ScholarshipSemester: c.ScholarshipSemester,
			ScholarshipType:     c.ScholarshipType,
			Amount:              c.Amount,
			Created:             c.Created,
			Modified:            c.Modified,
			Username:            c.Username,
			SubsidyNumber:       c.SubsidyNumber,
			SubsidyNameThai:     c.SubsidyNameThai,
		})
	}

	if len(scholarShipRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลผลการศึกษา")
	}

	scholarShipResponse = ScholarShipResponse{
		STD_CODE: requestBody.STD_CODE,
		Total:    len(scholarShipRec),
		RECORD:   scholarShipRec,
	}

	sholarShipJSON, _ := json.Marshal(&scholarShipResponse)
	timeNow := time.Now()
	redisCacheScholarShip := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = scholarShip.redis_cache.Set(ctx, key, sholarShipJSON, redisCacheScholarShip.Sub(timeNow)).Err()

	return &scholarShipResponse, nil
}
