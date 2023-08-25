package services

import (
	"RU-Smart-Workspace/ru-smart-api/errs"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/validator.v2"
)
type Data struct {
	ID    string
	Title string
	Code  string
}
func GetTitleByID(id string, data []Data) string {
	for _, entry := range data {
		if entry.ID == id {
			return entry.Title
		}
	}
	return "Unknown Title"
}
func (g *ondemandServices) GetOndemandAll(requestBody OndemandRequest) (*OndemandResponse, error) {

	data := []Data{
		{ID: "0", Title: "งดบรรยาย", Code: "เปิด"},
		{ID: "1", Title: "บรรยาย", Code: "เปิด"},
		{ID: "2", Title: "วีดีโอขัดข้อง กำลังรอดำเนินการแก้ไข", Code: "เปิด"},
		{ID: "4", Title: "หยุดนักขัตฤกษ์งดบรรยาย", Code: "เปิด"},
		{ID: "5", Title: "ช่วงสอบงดบรรยาย", Code: "เปิด"},
		{ID: "6", Title: "ช่วงรับปริญญางดบรรยาย", Code: "เปิด"},
		{ID: "7", Title: "จบการบรรยาย", Code: "เปิด"},
		{ID: "8", Title: "อาจารย์ย้ายห้องบรรยายโปรดติดต่อคณะ", Code: "เปิด"},
		{ID: "9", Title: "ปิดคอร์ส", Code: "เปิด"},
		{ID: "10", Title: "ไม่มีไฟล์วีดีโอ รอตรวจสอบและดำเนินการแก้ไข", Code: "เปิด"},
		{ID: "11", Title: "ไม่มีนักศึกษาเข้าเรียน", Code: "เปิด"},
		{ID: "12", Title: "ไม่มีการบรรยาย เนื่องจากไม่มีนักศึกษาเข้าเรียน", Code: "เปิด"},
		{ID: "13", Title: "ไม่มีเสียง รอตรวจสอบและดำเนินการแก้ไข", Code: "เปิด"},
		{ID: "14", Title: "ไม่มีภาพ รอตวรจสอบและดำเนินการแก้ไข", Code: "เปิด"},
		{ID: "15", Title: "จากการตรวจสอบ ไม่พบวีดีโอต้นฉบับ ขออภัยไว้ ณ ที่นี้ด้วยครับ", Code: "เปิด"},
		{ID: "16", Title: "บรรยายในชั้นเรียนไม่มีการบันทึกเทป (โปรดติดต่อคณะหรืออาจารย์ผู้สอน)", Code: "เปิด"},
		{ID: "3", Title: "รอไฟล์จากทางคณะ", Code: "เปิด"},
		{ID: "18", Title: "วีดีโอสอนขัดข้อง อาจารย์เเจ้งใช้เทป S/63", Code: "เปิด"},
		{ID: "19", Title: "วีดีโอการสอนการบันทึกขัดข้อง", Code: "เปิด"},
		{ID: "99", Title: "วีดีโอการสอนการบันทึกขัดข้อง", Code: "เปิด"},
		{ID: "98", Title: "อาจารย์แจ้งใช้วิดีโอ ภาคฤดูร้อน 2564", Code: "เปิด"},
		{ID: "97", Title: "อาจารย์แจ้งใช้วิดีโอ ภาคเรียนที่ 1/2565", Code: "เปิด"},
		{ID: "96", Title: "อาจารย์แจ้งใช้วิดีโอ ภาคเรียนที่ 2/2565", Code: "เปิด"},
		// Add more data entries here
	}
	
	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	ondemandResponse := OndemandResponse{
		SUBJECT_ID: requestBody.SUBJECT_ID,
		SEMESTER:	requestBody.SEMESTER,
		YEAR:		requestBody.YEAR,
		RECORD:   	ondemandRecord{},
	}


	key := "ondemand::" + requestBody.SUBJECT_ID + "-" + requestBody.SEMESTER+ "-" + requestBody.YEAR
	ondemandCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(ondemandCache), &ondemandResponse)
		fmt.Println("cache")
		return &ondemandResponse, nil
	}

	fmt.Println("database")

	ondemandRepo, err := g.ondemandRepo.GetOndemandAll(requestBody.SUBJECT_ID, requestBody.SEMESTER,requestBody.YEAR)
	if err != nil {
		log.Println(err.Error()+"no have subject")
		return &ondemandResponse,nil
	}

	ondemandRec := ondemandRecord{}
	
		ondemandRec = ondemandRecord{
			SUBJECT_CODE:	ondemandRepo.SUBJECT_CODE,
			SUBJECT_ID:     ondemandRepo.SUBJECT_ID,
			SUBJECT_NAME_ENG: ondemandRepo.SUBJECT_NAME_ENG,
			SEMESTER:      ondemandRepo.SEMESTER,
			YEAR:         ondemandRepo.YEAR,
			TOTAL:	0,
			
		}
		
		for _,item := range ondemandRepo.DETAIL{
			fmt.Println(item.SUBJECT_CODE)
			
			ondemandRec.DETAIL =  append(ondemandRec.DETAIL, 	ondemandSubjectCodeRecord  {
				AUDIO_ID   : item.AUDIO_ID,
				SUBJECT_CODE :item.SUBJECT_CODE,
				SUBJECT_ID	:item.SUBJECT_ID,
				AUDIO_SEC	:item.AUDIO_SEC,
				SEM      :item.SEM,
				YEAR         :item.YEAR,
				AUDIO_CREATE         :item.AUDIO_CREATE,
				AUDIO_STATUS         :GetTitleByID(item.AUDIO_STATUS,data),
				AUDIO_TEACH         :item.AUDIO_TEACH,
				AUDIO_COMMENT         :item.AUDIO_COMMENT,
		
			})
		}
		ondemandRec.TOTAL=len(ondemandRec.DETAIL)

	// if len(ondemandRec) < 1 {
	// 	return nil, errs.NewNotFoundError("ไม่พบข้อมูลผล")
	// }

	ondemandResponse = OndemandResponse{
		SUBJECT_ID:requestBody.SUBJECT_ID,
		SEMESTER:requestBody.SEMESTER,
		YEAR:requestBody.YEAR,
		RECORD:ondemandRec,
	}

	ondemandJSON, _ := json.Marshal(&ondemandResponse)
	timeNow := time.Now()
	redisCacheondemand := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, ondemandJSON, redisCacheondemand.Sub(timeNow)).Err()

	return &ondemandResponse, nil
}


func (g *ondemandServices) GetOndemandSubjectCode(requestBody OndemandSubjectCodeRequest) (*OndemandSubjectCodeResponse, error) {


	if err := validator.Validate(requestBody); err != nil {
		log.Println(err)
		return nil, errs.NewBadRequestError(err.Error())
	}

	ondemandSubjectCodeResponse := OndemandSubjectCodeResponse{
		SUBJECT_CODE: requestBody.SUBJECT_CODE,
		RECORD:   	[]ondemandSubjectCodeRecord{},
	}


	key := "ondemand::" + requestBody.SUBJECT_CODE
	ondemandCache, err := g.redis_cache.Get(ctx, key).Result()
	if err == nil {
		log.Println(err)
		_ = json.Unmarshal([]byte(ondemandCache), &ondemandSubjectCodeResponse)
		fmt.Println("cache")
		return &ondemandSubjectCodeResponse, nil
	}

	fmt.Println("database")

	ondemandRepo, err := g.ondemandRepo.GetOndemandSubjectCode(requestBody.SUBJECT_CODE)
	if err != nil {
		log.Println(err.Error())
		return &ondemandSubjectCodeResponse, errs.NewUnExpectedError()
	}

	ondemandSubjectCodeRec := []ondemandSubjectCodeRecord{}
	for _, c := range *ondemandRepo {
		ondemandSubjectCodeRec = append(ondemandSubjectCodeRec, ondemandSubjectCodeRecord{
			AUDIO_ID:     c.AUDIO_ID,
			SUBJECT_CODE:     c.SUBJECT_CODE,
			SUBJECT_ID:	c.SUBJECT_ID,
			AUDIO_SEC:	c.AUDIO_SEC,
			SEM:      c.SEM,
			YEAR:     c.YEAR,
			AUDIO_CREATE:  c.AUDIO_CREATE,
			AUDIO_STATUS:  c.AUDIO_STATUS,
			AUDIO_TEACH:   c.AUDIO_TEACH,
			AUDIO_COMMENT:  c.AUDIO_COMMENT,

		})
	}
	

	if len(ondemandSubjectCodeRec) < 1 {
		return nil, errs.NewNotFoundError("ไม่พบข้อมูลผล")
	}

	ondemandSubjectCodeResponse = OndemandSubjectCodeResponse{
		SUBJECT_CODE:requestBody.SUBJECT_CODE,
		RECORD:ondemandSubjectCodeRec,
	}

	ondemandJSON, _ := json.Marshal(&ondemandSubjectCodeResponse)
	timeNow := time.Now()
	redisCacheondemand := time.Unix(timeNow.Add(time.Second*30).Unix(), 0)
	_ = g.redis_cache.Set(ctx, key, ondemandJSON, redisCacheondemand.Sub(timeNow)).Err()

	return &ondemandSubjectCodeResponse, nil
}

