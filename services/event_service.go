package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)

func ChangeSemester(semester string) string {
	if semester == "3" {
		return "ฤดูร้อน"
	} else {
		return semester
	}
}

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

	key := "v2-event::" + eventRequest.StdID
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
			Semester: ChangeSemester(item.Semester),
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
