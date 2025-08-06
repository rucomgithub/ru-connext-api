package entities

import (
	"time"
)

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
