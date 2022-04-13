package tools

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
     "github.com/penggy/EasyGoLib/utils"
)

var _db *gorm.DB

func init() {
    var err error
    deskip := utils.Conf().Section("rtsp").Key("clouddeskip").MustString("")
    dataaddr := "root:uroot012@(" + deskip + ":3306)/ovirt_development?charset=utf8&parseTime=True&loc=Local"
    _db, err = gorm.Open("mysql", dataaddr)
    if err != nil {
        return;
        //panic(err)
    }
    _db.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
    _db.DB().SetMaxIdleConns(20)   //连接池最大允许的空闲连接数
   _db.SingularTable(true)
}

func GetDB() *gorm.DB {
	return _db
}






