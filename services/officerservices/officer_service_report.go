package officerservices

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (s *officerServices) GetReport(reportRequest *ReportRequest) (*ReportResponse, error) {

	reportResponse := ReportResponse{
		StartDate : reportRequest.StartDate,
		EndDate   : reportRequest.EndDate,
	}

	key := "v2-report::" + reportRequest.StartDate + reportRequest.EndDate
	reportCache, err := s.redis_cache.Get(ctx,key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(reportCache), &reportResponse)
		fmt.Println("report-cache")
		return &reportResponse, nil
	}

	fmt.Println("report-database")

	reportsRepo , err := s.officerRepo.FindReport(reportRequest.StartDate,reportRequest.EndDate)

	if err != nil {
		return &reportResponse, err
	}

	var reports []map[string]interface{}

	 for _, item := range reportsRepo {
        // วนลูปผ่าน key-value pair ใน map
		reports = append(reports, item)
    }

	// for key, item := range reportsRepo {
	// 	fmt.Println(key)
	// 	reports = append(reports, item)
	// }

	reportResponse = ReportResponse{
		StartDate :reportRequest.StartDate,
		EndDate  : reportRequest.EndDate,
		Reports:      reports,
		Count:       len(reports),
	}

	if len(reports) < 1 {
		errStr := fmt.Sprintf("ไม่พบข้อมูลรายงาน %s ถึง %s",reportRequest.StartDate,reportRequest.EndDate)
		return &reportResponse, errs.NewNotFoundError(errStr)
	}

	reportsJSON, _ := json.Marshal(&reportResponse)
	timeNow := time.Now()
	redisCachereport := time.Unix(timeNow.Add(time.Second * 5).Unix(), 0)
	_ = s.redis_cache.Set(ctx,key, reportsJSON, redisCachereport.Sub(timeNow)).Err()

	return &reportResponse, nil
}