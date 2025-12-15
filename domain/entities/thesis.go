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
	Similarity             float64                 `json:"similarity" db:"SIMILARITY"`
	STATUS                 string                  `json:"status" db:"STATUS"`
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

// DB model — ใช้กับ sqlx / database/sql
type RequestSuccess struct {
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
