package officerservices

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (s *officerServices) GetQualificationAll() (*[]QualificationResponse, int, error) {

	qualificationResponses := []QualificationResponse{}

	key := "qualification::all"
	fmt.Println(key)
	qualificationCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(qualificationCache), &qualificationResponses)
		fmt.Println("cache-qualification")
		return &qualificationResponses, 0, nil
	}

	fmt.Println("database-qualification")

	qualificationRepo, err := s.officerRepo.GetQualificationAll()

	if err != nil {
		log.Println(err.Error())
		return &qualificationResponses, 0, errs.NewUnExpectedError()
	}

	qualificationRec := []QualificationResponse{}

	for _, c := range *qualificationRepo {
		qualificationRec = append(qualificationRec, QualificationResponse{
			STD_CODE:     c.STD_CODE,
			REQUEST_DATE: c.REQUEST_DATE,
			OPERATE_DATE: c.OPERATE_DATE,
			CONFIRM_DATE: c.CONFIRM_DATE,
			CANCEL_DATE:  c.CANCEL_DATE,
			STATUS:       c.STATUS,
			CREATED:      c.CREATED,
			MODIFIED:     c.MODIFIED,
			DESCRIPTION:  c.DESCRIPTION,
		})
	}

	if len(qualificationRec) < 1 {
		return nil, 0, errs.NewNotFoundError("ไม่พบข้อมูลการยื่นขอเอกสาร")
	}

	qualificationJSON, _ := json.Marshal(&qualificationRec)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, qualificationJSON, redisCacheregister.Sub(timeNow)).Err()

	return &qualificationRec, len(qualificationRec), nil
}

func (s *officerServices) GetQualification(std_code string) (*QualificationResponse, error) {

	qf, err := s.officerRepo.GetQualification(std_code)

	if err != nil {
		return nil, err
	}

	qualification := QualificationResponse{
		STD_CODE:     qf.STD_CODE,
		REQUEST_DATE: qf.REQUEST_DATE,
		OPERATE_DATE: qf.OPERATE_DATE,
		CONFIRM_DATE: qf.CONFIRM_DATE,
		CANCEL_DATE:  qf.CANCEL_DATE,
		STATUS:       qf.STATUS,
		CREATED:      qf.CREATED,
		MODIFIED:     qf.MODIFIED,
		DESCRIPTION:  qf.DESCRIPTION,
	}

	return &qualification, nil
}

func (s *officerServices) UpdateQualification(std_code, status, description string) (*QualificationResponse, int64, error) {

	if status == "confirm" {
		description = "อนุมัติเอกสารแล้ว"
	}

	rowsAffected, err := s.officerRepo.UpdateQualification(std_code, status, description)

	if err != nil {
		return nil, 0, err
	}

	qf, err := s.officerRepo.GetQualification(std_code)
	if err != nil {
		return nil, 0, err
	}

	qualification := QualificationResponse{
		STD_CODE:     qf.STD_CODE,
		REQUEST_DATE: qf.REQUEST_DATE,
		OPERATE_DATE: qf.OPERATE_DATE,
		CONFIRM_DATE: qf.CONFIRM_DATE,
		CANCEL_DATE:  qf.CANCEL_DATE,
		STATUS:       qf.STATUS,
		CREATED:      qf.CREATED,
		MODIFIED:     qf.MODIFIED,
		DESCRIPTION:  qf.DESCRIPTION,
	}

	return &qualification, rowsAffected, nil
}
