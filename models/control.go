package models

type Control struct {
	Id               string `gorm:"type:varchar(128);primary_key;unique"`
        Isusb           * bool    `gorm:"type:bool"`
        Isblue          * bool    `gorm:"type:bool"`
        Isscreen        * bool    `gorm:"type:bool"`
        Isprint         * bool    `gorm:"type:bool"`
        Screendesc       string `gorm:"type:varchar(128)"`
        Screenhostname   bool    `gorm:"type:bool"`
        Screenusername   bool    `gorm:"type:bool"`
        Screenmacaddr    bool    `gorm:"type:bool"`
        Printdesc       string `gorm:"type:varchar(128)"`
        Printhostname   bool    `gorm:"type:bool"`
        Printusername   bool    `gorm:"type:bool"`
        Printmacaddr    bool    `gorm:"type:bool"`

}
