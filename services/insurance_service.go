package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

type Insurance struct {
	ID    string
	Title string
}

func GetInsuranceByID(id string, data []Insurance) string {
	for _, entry := range data {
		if entry.ID == id {
			return entry.Title
		}
	}
	return "Unknown Title"
}

func ParseDate(dateStr string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}
	return t
}

func CheckExpiry(typeinsurance, enddate, name string) string {
	today := time.Now()
	EndDate := ParseDate(enddate)

	duration := EndDate.Sub(today)
	daysLeft := int(duration.Hours() / 24)

	thaiEndDate := EndDate.AddDate(543, 0, 0)

	if daysLeft < 0 {
		return fmt.Sprintf("กรมธรรม์ประกันภัยแบบ %s ของ %s วันสิ้นสุด %s ซึ่งหมดอายุผ่านมาแล้ว %d วัน.", typeinsurance, name, thaiEndDate.Format("02/01/2006"), -daysLeft)
	} else if daysLeft == 0 {
		return fmt.Sprintf("กรมธรรม์ประกันภัยแบบ %s ของ %s วันสิ้นสุด %s ซึ่งหมดอายุในวันนี้.", typeinsurance, name, thaiEndDate.Format("02/01/2006"))
	} else {
		return fmt.Sprintf("กรมธรรม์ประกันภัยแบบ %s ของ %s วันสิ้นสุด %s อีก %d วันจะหมดอายุ.", typeinsurance, name, thaiEndDate.Format("02/01/2006"), daysLeft)
	}
}

func (g *InsuranceServices) GetInsuranceListAll(insuranceRequest InsuranceRequest) (*InsuranceResponse, error) {

	insurance := []Insurance{
		{ID: "GROUP", Title: "ประกันกลุ่ม"},
		{ID: "PERSONAL", Title: "ประกันบุคคล"},
	}

	if err := validator.Validate(insuranceRequest); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	insuranceResponse := InsuranceResponse{
		StudentCode: insuranceRequest.StudentCode,
	}

	key := "insurance::" + insuranceRequest.StudentCode
	insuranceCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(insuranceCache), &insuranceResponse)
		fmt.Println("insurance-cache")
		return &insuranceResponse, nil
	}

	fmt.Println("insurance-database")

	insuranceRepo, err := g.insuranceRepo.GetInsuranceListAll(insuranceRequest.StudentCode)

	if err != nil {
		return &insuranceResponse, err
	}

	insuranceRec := InsuranceRecord{}
	detail := []InsuranceRecord{}
	typeinsurance := ""

	for _, item := range *insuranceRepo {
		insuranceRec.StudentCode = item.StudentCode
		typeinsurance = GetInsuranceByID(item.TypeInsurance, insurance)
		detail = append(detail, InsuranceRecord{
			StudentCode:     item.StudentCode,
			PersonCode:      item.PersonCode,
			NameInsurance:   item.NameInsurance,
			StartDate:       item.StartDate,
			EndDate:         item.EndDate,
			StatusInsurance: item.StatusInsurance,
			TypeInsurance:   typeinsurance,
			YearMonth:       item.YearMonth,
			Expire:          CheckExpiry(typeinsurance, item.EndDate, item.NameInsurance),
		})
	}

	insuranceResponse = InsuranceResponse{
		StudentCode: insuranceRec.StudentCode,
		Detail:      detail,
		Total:       len(detail),
	}

	if len(detail) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลประกันของนักศึกษา " + insuranceRequest.StudentCode)
	}

	insuranceJSON, _ := json.Marshal(&insuranceResponse)
	timeNow := time.Now()
	redisCacheinsurance := time.Unix(timeNow.Add(time.Minute*30).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, insuranceJSON, redisCacheinsurance.Sub(timeNow)).Err()

	return &insuranceResponse, nil
}
