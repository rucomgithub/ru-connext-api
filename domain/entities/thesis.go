package entities

import (
	"time"
)

type ThesisSimilarity struct {
	StudentID          string    `json:"studentId" db:"STD_CODE"`
	Program            string    `json:"program" db:"PROGRAM"`
	Major              string    `json:"major" db:"MAJOR"`
	Faculty            string    `json:"faculty" db:"FACULTY"`
	ThesisType         string    `json:"thesisType" db:"THESIS_TYPE"`
	ThesisTitleThai    string    `json:"thesisTitleThai" db:"THESIS_TITLE_THAI"`
	ThesisTitleEnglish string    `json:"thesisTitleEnglish" db:"THESIS_TITLE_ENGLISH"`
	Similarity         float64   `json:"similarity" db:"SIMILARITY"`
	CreatedAt          time.Time `json:"createdAt" db:"CREATED_AT"`
	UpdatedAt          time.Time `json:"updatedAt" db:"UPDATED_AT"`
	CreatedBy          string    `json:"createdBy" db:"CREATED_BY"`
	UpdatedBy          string    `json:"updatedBy" db:"UPDATED_BY"`
	Status             string    `json:"status" db:"STATUS"`
}

type ThesisJournal struct {
	StudentID              string                  `json:"studentId" db:"STD_CODE"`
	Program                string                  `json:"program" db:"PROGRAM"`
	Major                  string                  `json:"major" db:"MAJOR"`
	Faculty                string                  `json:"faculty" db:"FACULTY"`
	ThesisType             string                  `json:"thesisType" db:"THESIS_TYPE"`
	ThesisTitleThai        string                  `json:"thesisTitleThai" db:"THESIS_TITLE_THAI"`
	ThesisTitleEnglish     string                  `json:"thesisTitleEnglish" db:"THESIS_TITLE_ENGLISH"`
	JournalPublication     []JournalPublication    `json:"publications"`
	ConferencePresentation *ConferencePresentation `json:"conferencePresentation"`
	OtherPublication       *OtherPublication       `json:"otherPublication"`
	CreatedAt              time.Time               `json:"createdAt" db:"CREATED_AT"`
	UpdatedAt              time.Time               `json:"updatedAt" db:"UPDATED_AT"`
	CreatedBy              string                  `json:"createdBy" db:"CREATED_BY"`
	UpdatedBy              string                  `json:"updatedBy" db:"UPDATED_BY"`
	Status                 string                  `json:"status" db:"STATUS"`
}

type JournalPublication struct {
	Id           string    `json:"id" db:"ID"`
	StudentID    string    `json:"studentId" db:"STD_CODE"`
	Type         string    `json:"type" db:"TYPE"`
	ArticleTitle string    `json:"articleTitle" db:"ARTICLE_TITLE"`
	JournalName  string    `json:"journalName" db:"JOURNAL_NAME"`
	Country      string    `json:"country" db:"COUNTRY"`
	Status       string    `json:"status" db:"STATUS"`
	Year         string    `json:"year" db:"YEAR"`
	Volume       string    `json:"volume" db:"VOLUME"`
	Issue        string    `json:"issue" db:"ISSUE"`
	Month        string    `json:"month" db:"MONTH"`
	PublishYear  string    `json:"publishYear" db:"PUBLISH_YEAR"`
	PageFrom     int       `json:"pageFrom" db:"PAGE_FROM"`
	PageTo       int       `json:"pageTo" db:"PAGE_TO"`
	PublishLevel string    `json:"level" db:"PUBLISH_LEVEL"`
	TCIGroup     string    `json:"tciGroup" db:"TCI_GROUP"`
	CreatedAt    time.Time `json:"createdAt" db:"CREATED_AT"`
}

type ConferencePresentation struct {
	StudentID      string    `json:"studentId" db:"STD_CODE"`
	Type           string    `json:"type" db:"TYPE"`
	ArticleTitle   string    `json:"articleTitle" db:"ARTICLE_TITLE"`
	ConferenceName string    `json:"conferenceName" db:"CONFERENCE_NAME"`
	ConferenceDate string    `json:"conferenceDate" db:"CONFERENCE_DATE"`
	Organizer      string    `json:"organizer" db:"ORGANIZER"`
	Location       string    `json:"location" db:"LOCATION"`
	Country        string    `json:"country" db:"COUNTRY"`
	Status         string    `json:"status" db:"STATUS"`
	PageFrom       int       `json:"pageFrom" db:"PAGE_FROM"`
	PageTo         int       `json:"pageTo" db:"PAGE_TO"`
	CreatedAt      time.Time `json:"createdAt" db:"CREATED_AT"`
}

type OtherPublication struct {
	StudentID    string    `json:"studentId" db:"STD_CODE"`
	ArticleTitle string    `json:"articleTitle" db:"ARTICLE_TITLE"`
	SourceType   string    `json:"sourceType" db:"SOURCE_TYPE"`
	SourceDetail string    `json:"sourceDetail" db:"SOURCE_DETAIL"`
	CreatedAt    time.Time `json:"createdAt" db:"CREATED_AT"`
}

type RequestSuccess struct {
	ENROLL_YEAR          string `db:"ENROLL_YEAR" json:"enrollYear"`
	ENROLL_SEMESTER      string `db:"ENROLL_SEMESTER" json:"enrollSemester"`
	STD_CODE             string `db:"STD_CODE" json:"stdCode"`
	PRENAME_THAI_S       string `db:"PRENAME_THAI_S" json:"prenameThai"`
	PRENAME_ENG_S        string `db:"PRENAME_ENG_S" json:"prenameEng"`
	FIRST_NAME           string `db:"FIRST_NAME" json:"firstName"`
	LAST_NAME            string `db:"LAST_NAME" json:"lastName"`
	FIRST_NAME_ENG       string `db:"FIRST_NAME_ENG" json:"firstNameEng"`
	LAST_NAME_ENG        string `db:"LAST_NAME_ENG" json:"lastNameEng"`
	THAI_NAME            string `db:"THAI_NAME" json:"thaiName"`
	PLAN_NO              string `db:"PLAN_NO" json:"planNo"` 
	SEX                  string `db:"SEX" json:"sex"`
	REGINAL_NAME         string `db:"REGINAL_NAME" json:"reginalName"`
	SUBSIDY_NAME         string `db:"SUBSIDY_NAME" json:"subsidyName"`
	STATUS_NAME_THAI     string `db:"STATUS_NAME_THAI" json:"statusNameThai"`
	BIRTH_DATE           string `db:"BIRTH_DATE" json:"birthDate"`
	STD_ADDR             string `db:"STD_ADDR" json:"studentAddress"`
	ADDR_TEL             string `db:"ADDR_TEL" json:"addressTel"`
	JOB_POSITION         string `db:"JOB_POSITION" json:"jobPosition"`
	STD_OFFICE           string `db:"STD_OFFICE" json:"studentOffice"`
	OFFICE_TEL           string `db:"OFFICE_TEL" json:"officeTel"`
	DEGREE_NAME          string `db:"DEGREE_NAME" json:"degreeName"`
	BSC_DEGREE_NO        string `db:"BSC_DEGREE_NO" json:"bscDegreeNo"`
	BSC_DEGREE_THAI_NAME string `db:"BSC_DEGREE_THAI_NAME" json:"bscDegreeThaiName"`
	BSC_INSTITUTE_NO     string `db:"BSC_INSTITUTE_NO" json:"bscInstituteNo"`
	INSTITUTE_THAI_NAME  string `db:"INSTITUTE_THAI_NAME" json:"instituteThaiName"`
	CK_CERT_NO           string `db:"CK_CERT_NO" json:"ckCertNo"`
	CHK_CERT_NAME_THAI   string `db:"CHK_CERT_NAME_THAI" json:"chkCertNameThai"`
	SUCCESS_YEAR         string `db:"SUCCESS_YEAR" json:"SUCCESS_YEAR"`
	SUCCESS_SEMESTER     string `db:"SUCCESS_SEMESTER" json:"SUCCESS_SEMESTER"`
	NAME_THAI            string `db:"NAME_THAI" json:"nameThai"`
	NAME_ENG             string `db:"NAME_ENG" json:"nameEng"`
	THESIS_THAI          string `db:"THESIS_THAI" json:"thesisThai"`
	THESIS_ENG           string `db:"THESIS_ENG" json:"thesisEng"`
	DEGREE               string `db:"DEGREE" json:"degree"`
	REGISTRATION         string `db:"REGISTRATION" json:"registration"`
	GRADES               string `db:"GRADES" json:"grades"`
	ADDRESS              string `db:"ADDRESS" json:"address"`
	CREATED              string `db:"CREATED" json:"created"`
	MODIFIED             string `db:"MODIFIED" json:"modified"`
	THESIS_THAI_TITLE    string `db:"THESIS_THAI_TITLE" json:"thesisThaiTitle"`
	THESIS_ENG_TITLE     string `db:"THESIS_ENG_TITLE" json:"thesisEngTitle"`
	THESIS_TYPE          string `db:"THESIS_TYPE" json:"thesisType"`
	SIMILARITY           string `db:"SIMILARITY" json:"similarity"`
}

type RequestSuccessRepo struct {
	STD_CODE         string `json:"STD_CODE" validate:"min=9,max=10,regexp=^[0-9]" db:"STD_CODE"`
	SUCCESS_YEAR     string `json:"SUCCESS_YEAR" validate:"min=4,max=4,regexp=^[0-9]" db:"SUCCESS_YEAR"`
	SUCCESS_SEMESTER string `json:"SUCCESS_SEMESTER" validate:"min=1,max=1,regexp=^[0-9]" db:"SUCCESS_SEMESTER"`
	NAME_THAI        string `json:"NAME_THAI" db:"NAME_THAI"`
	NAME_ENG         string `json:"NAME_ENG" db:"NAME_ENG"`
	DEGREE           string `json:"DEGREE" db:"DEGREE"`
	THESIS_THAI      string `json:"THESIS_THAI" db:"THESIS_THAI"`
	THESIS_ENG       string `json:"THESIS_ENG" db:"THESIS_ENG"`
	REGISTRATION     string `json:"REGISTRATION" db:"REGISTRATION"`
	GRADES           string `json:"GRADES" db:"GRADES"`
	ADDRESS          string `json:"ADDRESS" db:"ADDRESS"`
	CREATED          string `json:"CREATED" db:"CREATED"`
	MODIFIED         string `json:"MODIFIED" db:"MODIFIED"`
}

