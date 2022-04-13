package models

type Desk struct {
	Id               string `gorm:"type:varchar(256);primary_key;unique"`
        Deskname           string `gorm:"type:varchar(128)"`
        Deskloc        string `gorm:"type:varchar(128)"`
        Username          string `gorm:"type:varchar(128)"`
        Userpwd           string `gorm:"type:varchar(128)"`
	Deskip        string `gorm:"type:varchar(128)"`
        Deskmac        string `gorm:"type:varchar(128)"`
        Deskdesc       string `gorm:"type:varchar(256)"`
        Deskkey        string `gorm:"type:varchar(256)"`
        IsRecord          bool    `gorm:"type:bool"`
        Savedays          int `gorm:"type:int"`
}
