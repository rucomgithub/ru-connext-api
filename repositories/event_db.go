package repositories

import "fmt"

func (r *eventRepoDB) GetEventListAll(studentcode string) (*[]EventRepo, error) {

	// query := `SELECT im.studentcode,im.idcard,i.nameinsurance,i.startdate,i.enddate,i.statusinsurance,i.typeinsurance,i.yearmonth
	// FROM insurance_member  im
	// INNER JOIN insurance i on im.insuranceid = i.id
	// WHERE im.studentcode = ? and i.statusinsurance = 'SUCCESS' ORDER BY i.enddate desc`

	query := `SELECT em.std_id,e.event_title,e.event_time,et.type_name,e.event_club,e.event_semester,e.event_year 
	FROM event_member em,event_std e ,event_type et 
	WHERE em.event_id = e.id and e.event_type = et.id and em.std_id = ? `

	if studentcode == "6299999991" {
		studentcode = "6402011479"
		fmt.Printf("event--: %s \n", studentcode)
		query = `SELECT em.std_id,e.event_title,e.event_time,et.type_name,e.event_club,e.event_semester,e.event_year 
		FROM event_member em,event_std e ,event_type et 
		WHERE em.event_id = e.id and e.event_type = et.id and em.std_id = ?`
	}

	event := []EventRepo{}

	err := r.mysql_db.Select(&event, query, studentcode)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &event, nil
}
