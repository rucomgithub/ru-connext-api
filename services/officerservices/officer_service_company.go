package officerservices
import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"
)
func (s *officerServices) GetCommpanyList(std_code string) (*[]CompanyResponse, int, error) {

	companyResponses := []CompanyResponse{}

	key := "v2-company::all"
	fmt.Println(key)
	companyCache, err := s.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(companyCache), &companyResponses)
		fmt.Println("cache-company")
		return &companyResponses, 0, nil
	}

	fmt.Println("database-company")

	companyRepo, err := s.officerRepo.GetCompanyList(std_code)

	if err != nil {
		log.Println(err.Error())
		return &companyResponses, 0, errs.NewUnExpectedError()
	}

	companyRec := []CompanyResponse{}

	for _, c := range *companyRepo {
		companyRec = append(companyRec, CompanyResponse{
			STD_CODE: c.STD_CODE,
			EMAIL   : c.EMAIL,
			FULLNAME: c.FULLNAME,
			COMPANY : c.COMPANY,
			CREATED : c.CREATED,
			MODIFIED: c.MODIFIED,
		})
	}

	if len(companyRec) < 1 {
		return nil, 0, errs.NewNotFoundError("ไม่พบข้อมูล บริษัท/ห้างร้าน/หน่วยงาน/ส่วนราชการ")
	}

	companyJSON, _ := json.Marshal(&companyRec)
	timeNow := time.Now()
	redisCacheregister := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = s.redis_cache.Set(ctx, key, companyJSON, redisCacheregister.Sub(timeNow)).Err() 

	return &companyRec, len(companyRec), nil
}
