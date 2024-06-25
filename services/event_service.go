package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

// type Insurance struct {
// 	ID    string
// 	Title string
// }

// func GetInsuranceByID(id string, data []Insurance) string {
// 	for _, entry := range data {
// 		if entry.ID == id {
// 			return entry.Title
// 		}
// 	}
// 	return "Unknown Title"
// }

// func ParseDate(dateStr string) time.Time {
// 	layout := "2006-01-02"
// 	t, err := time.Parse(layout, dateStr)
// 	if err != nil {
// 		fmt.Println("Error parsing date:", err)
// 	}
// 	return t
// }

// func CheckExpiry(typeinsurance, enddate, name string) string {
// 	today := time.Now()
// 	EndDate := ParseDate(enddate)

// 	duration := EndDate.Sub(today)
// 	daysLeft := int(duration.Hours() / 24)

// 	thaiEndDate := EndDate.AddDate(543, 0, 0)

// 	if daysLeft < 0 {
// 		return fmt.Sprintf("กรมธรรม์ประกันภัยแบบ %s ของ %s วันสิ้นสุด %s ซึ่งหมดอายุผ่านมาแล้ว %d วัน.", typeinsurance, name, thaiEndDate.Format("02/01/2006"), -daysLeft)
// 	} else if daysLeft == 0 {
// 		return fmt.Sprintf("กรมธรรม์ประกันภัยแบบ %s ของ %s วันสิ้นสุด %s ซึ่งหมดอายุในวันนี้.", typeinsurance, name, thaiEndDate.Format("02/01/2006"))
// 	} else {
// 		return fmt.Sprintf("กรมธรรม์ประกันภัยแบบ %s ของ %s วันสิ้นสุด %s อีก %d วันจะหมดอายุ.", typeinsurance, name, thaiEndDate.Format("02/01/2006"), daysLeft)
// 	}
// }

func (g *EventServices) GetEventListAll(eventRequest EventRequest) (*EventResponse, error) {

	// insurance := []Insurance{
	// 	{ID: "GROUP", Title: "ประกันกลุ่ม"},
	// 	{ID: "PERSONAL", Title: "ประกันบุคคล"},
	// }

	if err := validator.Validate(eventRequest); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	eventResponse := EventResponse{
		StdID: eventRequest.StdID,
	}

	key := "event::" + eventRequest.StdID
	eventCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(eventCache), &eventResponse)
		fmt.Println("event-cache")
		return &eventResponse, nil
	}

	fmt.Println("event-database")

	eventRepo, err := g.eventRepo.GetEventListAll(eventRequest.StdID)

	if err != nil {
		return &eventResponse, err
	}

	eventRec := EventRecord{}
	detail := []EventRecord{}

	for _, item := range *eventRepo {
		eventRec.StdID = item.StdID
		detail = append(detail, EventRecord{
			StdID:    item.StdID,
			Title:    item.Title,
			Time:     item.Time,
			TypeName: item.TypeName,
			Club:     item.Club,
			Semester: item.Semester,
			Year:     item.Year,
		})
	}

	eventResponse = EventResponse{
		StdID:  eventRec.StdID,
		Detail: detail,
		Total:  len(detail),
	}

	if len(detail) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลกิจกรรมนักศึกษา")
	}

	eventJSON, _ := json.Marshal(&eventResponse)
	timeNow := time.Now()
	redisCacheEvent := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, eventJSON, redisCacheEvent.Sub(timeNow)).Err()

	return &eventResponse, nil
}
