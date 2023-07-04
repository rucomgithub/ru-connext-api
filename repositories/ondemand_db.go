package repositories

import "fmt"


func (r *ondemandRepoDB) GetOndemandAll(course_no, semester,year string) (*OndemandRepo, error) {
//func (r *ondemandRepoDB) GetOndemandAll() (*[]OndemandRepo, error) {
	ondemand := OndemandRepo{}
	query := `select subject_code,subject_id,subject_NameEng,semester,year from  master_subject  where subject_id like ? and semester = ? and year = ?`
    coursePattern := "%"+course_no+"%"
	err := r.mysql_db.Get(&ondemand,query, coursePattern,semester,year)
    fmt.Println(ondemand.SUBJECT_CODE)
	if err != nil {
        //ondemand.TOTAL = 0
        // fmt.Println(ondemand)
        //กรณีหาไม่เจอ 404
        // ondemand.DETAIL = append(ondemand.DETAIL, 	OndemandSubjectCodeRepo  {
        //     AUDIO_ID   : "",
        //     SUBJECT_CODE :"",
        //     SUBJECT_ID	:"",
        //     AUDIO_SEC	:"",
        //     SEM      :"",
        //     YEAR         :"",
        //     AUDIO_CREATE         :"",
        //     AUDIO_STATUS         :"",
        //     AUDIO_TEACH         :"",
        //     AUDIO_COMMENT         :"",
    
        // })
        // return &ondemand, nil
		return nil, err
	}
    //ondemandSubjectCodeRepo := []OndemandSubjectCodeRepo{}

    detail := []OndemandSubjectCodeRepo{}
    query = `select audio_id,subject_code,subject_id,audio_sec,sem,year,audio_create,audio_status,audio_teach,audio_comment from  detail_audio  where subject_code = ? order by audio_id ASC`
    //coursePattern := "%"+course_no+"%"
    err = r.mysql_db.Select(&detail,query,ondemand.SUBJECT_CODE)
    // fmt.Println(subject_code)
    if err != nil {
        // fmt.Println(query)
        return nil, err
    }
    ondemand.DETAIL = detail
    // ondemandSubjectCodeRepo, err := r.GetOndemandSubjectCode(ondemand.SUBJECT_CODE)
    // for _,item := range *ondemandSubjectCodeRepo {
    //     ondemand.DETAIL = append(ondemand.DETAIL, 	OndemandSubjectCodeRepo  {
    //         AUDIO_ID   : item.AUDIO_ID,
    //         SUBJECT_CODE :item.SUBJECT_CODE,
    //         SUBJECT_ID	:item.SUBJECT_ID,
    //         AUDIO_SEC	:item.AUDIO_SEC,
    //         SEM      :item.SEM,
    //         YEAR         :item.YEAR,
    //         AUDIO_CREATE         :item.AUDIO_CREATE,
    //         AUDIO_STATUS         :item.AUDIO_STATUS,
    //         AUDIO_TEACH         :item.AUDIO_TEACH,
    //         AUDIO_COMMENT         :item.AUDIO_COMMENT,
    
    //     })
    // }
    //ondemand.TOTAL = len(ondemand.DETAIL)
    // defer rows.Close()
    // for rows.Next() {
    //     var data OndemandRepo
    //     err := rows.Scan(&data.SUBJECT_CODE,&data.SUBJECT_ID, &data.SUBJECT_NAME_ENG, &data.SEMESTER,&data.YEAR)
    //     if err != nil {
    //         return nil, err
    //     }
	// 	fmt.Println(data)
    //     ondemand = append(ondemand, data)
    // }
    // if err := rows.Err(); err != nil {
    //     return nil, err
    // }
    return &ondemand, nil
}

func (r *ondemandRepoDB) GetOndemandSubjectCode(subject_code string) (*[]OndemandSubjectCodeRepo, error) {
    //func (r *ondemandRepoDB) GetOndemandAll() (*[]OndemandRepo, error) {
        ondemand := []OndemandSubjectCodeRepo{}
        query := `select audio_id,subject_code,subject_id,audio_sec,sem,year,audio_create,audio_status,audio_teach,audio_comment from  detail_audio  where subject_code = ? order by audio_id ASC`
        //coursePattern := "%"+course_no+"%"
        rows,err := r.mysql_db.Query(query,subject_code)
        fmt.Println(subject_code)
        if err != nil {
            // fmt.Println(query)
            return nil, err
        }
    
        defer rows.Close()
        for rows.Next() {
            var data OndemandSubjectCodeRepo
            err := rows.Scan(&data.AUDIO_ID,&data.SUBJECT_CODE,&data.SUBJECT_ID, &data.AUDIO_SEC, &data.SEM, &data.YEAR, &data.AUDIO_CREATE, &data.AUDIO_STATUS, &data.AUDIO_TEACH,&data.AUDIO_COMMENT)
            if err != nil {
                return nil, err
            }
            // fmt.Println(data)
            ondemand = append(ondemand, data)
        }
        if err := rows.Err(); err != nil {
            return nil, err
        }
        return &ondemand, nil
    }