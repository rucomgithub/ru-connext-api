package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

func (g *rotcsServices) GetRotcsRegister(requestBody RotcsRequest) (*RotcsRegisterResponse, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	rotcsResponse := RotcsRegisterResponse{
		StudentCode: requestBody.StudentCode,
		Detail:      []RotcsDetailRecord{},
		Total:       0,
	}

	key := "rotcs-register:::" + requestBody.StudentCode
	rotcsCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(rotcsCache), &rotcsResponse)
		fmt.Println("rotcs-cache")
		return &rotcsResponse, nil
	}

	fmt.Println("rotcs-register-database")
	rotcsRepo, err := g.rotcsRepo.GetRotcsRegister(requestBody.StudentCode)
	if err != nil {
		return &rotcsResponse, err
	}

	rotcsRec := RotcsRegisterRecord{}
	detail := []RotcsDetailRecord{}

	for _, item := range *rotcsRepo {
		rotcsRec.StudentCode = item.StudentCode
		detail = append(detail, RotcsDetailRecord{
			LayerArmy:    item.LayerArmy,
			LocationArmy: item.LocationArmy,
			YearReport:   item.YearReport,
			LayerReport:  item.LayerReport,
			TypeReport:   item.TypeReport,
			Status:       item.Status,
		})
	}

	if len(detail) < 1 {
		return &rotcsResponse, errs.NewNotFoundError("ไม่พบข้อมูลการฝึกเรียนวิชาทหารของ " + requestBody.StudentCode)
	}

	rotcsResponse = RotcsRegisterResponse{
		StudentCode: rotcsRec.StudentCode,
		Detail:      detail,
		Total:       len(detail),
	}
	
	if rotcsRec.StudentCode == "6410022734" {
		rotcsResponse.StudentCode = "6299999991"
	}
	if rotcsRec.StudentCode == "6406006293" {
		rotcsResponse.StudentCode = "6299999992"
	}

	rotcsJSON, _ := json.Marshal(&rotcsResponse)
	timeNow := time.Now()
	redisCacherotcs := time.Unix(timeNow.Add(time.Second*5).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, rotcsJSON, redisCacherotcs.Sub(timeNow)).Err()

	return &rotcsResponse, nil
}

func (g *rotcsServices) GetRotcsExtend(requestBody RotcsRequest) (*RotcsExtendResponse, error) {

	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	rotcsExtendResponse := RotcsExtendResponse{
		StudentCode: requestBody.StudentCode,
		Detail:      []RotcsExtendDetailResponse{},
		Total:       0,
	}

	key := "rotcs-extend:::" + requestBody.StudentCode
	rotcsCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(rotcsCache), &rotcsExtendResponse)
		fmt.Println("rotcs-extend-cache")
		return &rotcsExtendResponse, nil
	}

	fmt.Println("rotcs-extend-database")
	rotcsExtend, err := g.rotcsRepo.GetRotcsExtend(requestBody.StudentCode)

	if err != nil {
		return &rotcsExtendResponse, errs.NewNotFoundError("ไม่พบข้อมูลผ่อนผันหรือรักษาสิทธิ์การเกณฑ์ทหารของ " + requestBody.StudentCode)
	}

	detail := []RotcsExtendDetailResponse{}

	for _, item := range rotcsExtend.Detail {
		detail = append(detail, RotcsExtendDetailResponse{
			Id:               item.Id,
			RegisterYear:     item.RegisterYear,
			RegisterSemester: item.RegisterSemester,
			Credit:           item.Credit,
			Created:          item.Created,
			Modified:         item.Modified,
		})
	}

	if len(detail) < 1 {
		return &rotcsExtendResponse, errs.NewNotFoundError("ไม่พบข้อมูลรายละเอียดรายการผ่อนผันหรือรักษาสิทธิ์การเกณฑ์ทหารของ " + requestBody.StudentCode)
	}

	rotcsExtendResponse = RotcsExtendResponse{
		StudentCode: rotcsExtend.StudentCode,
		ExtendYear:  rotcsExtend.ExtendYear,
		Code9:       rotcsExtend.Code9,
		Option1:     rotcsExtend.Option1,
		Option2:     rotcsExtend.Option2,
		Option3:     rotcsExtend.Option3,
		Option4:     rotcsExtend.Option4,
		Option5:     rotcsExtend.Option5,
		Option6:     rotcsExtend.Option6,
		Option7:     rotcsExtend.Option7,
		Option8:     rotcsExtend.Option8,
		Option9:     rotcsExtend.Option9,
		OptionOther: rotcsExtend.OptionOther,
		Detail:      detail,
		Total:       len(detail),
	}

	if rotcsExtendResponse.StudentCode == "6401002222" {
		rotcsExtendResponse.StudentCode = "6299999991"
	}
	if rotcsExtendResponse.StudentCode == "6401008344" {
		rotcsExtendResponse.StudentCode = "6299999992"
	}

	rotcsJSON, _ := json.Marshal(&rotcsExtendResponse)
	timeNow := time.Now()
	redisCacherotcs := time.Unix(timeNow.Add(time.Second*5).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, rotcsJSON, redisCacherotcs.Sub(timeNow)).Err()

	return &rotcsExtendResponse, nil
}
