package entities

import (
	"time"
)

type PublicationType string

const (
	Doctoral    PublicationType = "doctoral"
	Master      PublicationType = "master"
	Thesis      PublicationType = "thesis"
	Independent PublicationType = "independent"
)

type Publication struct {
	StudentCode            string    `json:"std_code" db:"STD_CODE" validate:"min=10,max=10,regexp=^[0-9]{10}$"`
	ThesisTitle            string    `json:"thesis_title" db:"THESIS_TITLE" validate:"nonzero"`
	ThesisTitleEng         string    `json:"thesis_title_eng" db:"THESIS_TITLE_ENG" validate:"nonzero"`
	PublicationType        string    `json:"publication_type" db:"PUBLICATION_TYPE" validate:"nonzero"`
	Article1Journal        int       `json:"article1_journal" db:"ARTICLE1_JOURNAL" validate:"nonzero"`
	Article1Ejournal       int       `json:"article1_ejournal" db:"ARTICLE1_EJOURNAL" validate:"nonzero"`
	Article1Title          string    `json:"article1_title" db:"ARTICLE1_TITLE" validate:"nonzero"`
	Journal1Name           string    `json:"journal1_name" db:"JOURNAL1_NAME" validate:"nonzero"`
	Journal1Country        string    `json:"journal1_country" db:"JOURNAL1_COUNTRY" validate:"nonzero"`
	Journal1Accepted       int       `json:"journal1_accepted" db:"JOURNAL1_ACCEPTED" validate:"nonzero"`
	Journal1Published      int       `json:"journal1_published" db:"JOURNAL1_PUBLISHED" validate:"nonzero"`
	Journal1Volume         string    `json:"journal1_volume" db:"JOURNAL1_VOLUME" validate:"nonzero"`
	Journal1Issue          string    `json:"journal1_issue" db:"JOURNAL1_ISSUE" validate:"nonzero"`
	Journal1Month          string    `json:"journal1_month" db:"JOURNAL1_MONTH" validate:"nonzero"`
	Journal1Year           int       `json:"journal1_year" db:"JOURNAL1_YEAR" validate:"nonzero"`
	Journal1PageFrom       int       `json:"journal1_page_from" db:"JOURNAL1_PAGE_FROM" validate:"nonzero"`
	Journal1PageTo         int       `json:"journal1_page_to" db:"JOURNAL1_PAGE_TO" validate:"nonzero"`
	Journal1Level          int       `json:"journal1_level" db:"JOURNAL1_LEVEL" validate:"nonzero"`
	Journal1LevelTCI       int       `json:"journal1_level_tci" db:"JOURNAL1_LEVEL_TCI" validate:"nonzero"`
	Journal1TCIGroup       string    `json:"journal1_tci_group" db:"JOURNAL1_TCI_GROUP" validate:"nonzero"`
	Article2Journal        int       `json:"article2_journal" db:"ARTICLE2_JOURNAL"`
	Article2Ejournal       int       `json:"article2_ejournal" db:"ARTICLE2_EJOURNAL"`
	Article2Title          string    `json:"article2_title" db:"ARTICLE2_TITLE"`
	Journal2Name           string    `json:"journal2_name" db:"JOURNAL2_NAME"`
	Journal2Country        string    `json:"journal2_country" db:"JOURNAL2_COUNTRY"`
	Journal2Accepted       int       `json:"journal2_accepted" db:"JOURNAL2_ACCEPTED"`
	Journal2Published      int       `json:"journal2_published" db:"JOURNAL2_PUBLISHED"`
	Journal2Volume         string    `json:"journal2_volume" db:"JOURNAL2_VOLUME"`
	Journal2Issue          string    `json:"journal2_issue" db:"JOURNAL2_ISSUE"`
	Journal2Month          string    `json:"journal2_month" db:"JOURNAL2_MONTH"`
	Journal2Year           int       `json:"journal2_year" db:"JOURNAL2_YEAR"`
	Journal2PageFrom       int       `json:"journal2_page_from" db:"JOURNAL2_PAGE_FROM"`
	Journal2PageTo         int       `json:"journal2_page_to" db:"JOURNAL2_PAGE_TO"`
	Journal2Level          int       `json:"journal2_level" db:"JOURNAL2_LEVEL"`
	Journal2LevelTCI       int       `json:"journal2_level_tci" db:"JOURNAL2_LEVEL_TCI"`
	Journal2TCIGroup       string    `json:"journal2_tci_group" db:"JOURNAL2_TCI_GROUP"`
	ConfNational           int       `json:"conf_national" db:"CONF_NATIONAL"`
	ConfInterNational      int       `json:"conf_international" db:"CONF_INTERNATIONAL"`
	ConfArticleTitle       string    `json:"conf_article_title" db:"CONF_ARTICLE_TITLE"`
	ConfName               string    `json:"conf_name" db:"CONF_NAME"`
	ConfDate               time.Time `json:"conf_date" db:"CONF_DATE"`
	ConfOrganizer          string    `json:"conf_organizer" db:"CONF_ORGANIZER"`
	ConfLocation           string    `json:"conf_location" db:"CONF_LOCATION"`
	ConfCountry            string    `json:"conf_country" db:"CONF_COUNTRY"`
	ConfAccepted           int       `json:"conf_accepted" db:"CONF_ACCEPTED"`
	ConfPresented          int       `json:"conf_presented" db:"CONF_PRESENTED"`
	ConfPageFrom           int       `json:"conf_page_from" db:"CONF_PAGE_FROM"`
	ConfPageTo             int       `json:"conf_page_to" db:"CONF_PAGE_TO"`
	ConfLevel              string    `json:"conf_level" db:"CONF_LEVEL"`
	OtherArticleTitle      string    `json:"other_article_title" db:"OTHER_ARTICLE_TITLE"`
	OtherSourceOnline      int       `json:"other_source_online" db:"OTHER_SOURCE_ONLINE"`
	OtherSourceOnlineTitle string    `json:"other_source_online_title" db:"OTHER_SOURCE_ONLINE_TITLE"`
	OtherSource            int       `json:"other_source" db:"OTHER_SOURCE"`
	OtherSourceTitle       string    `json:"other_source_title" db:"OTHER_SOURCE_TITLE"`
	CreatedAt              time.Time `json:"created_at" db:"CREATED_AT"`
	UpdatedAt              time.Time `json:"updated_at" db:"UPDATED_AT"`
}
