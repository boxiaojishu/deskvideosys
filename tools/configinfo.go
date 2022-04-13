package tools

import (
    "github.com/penggy/EasyGoLib/utils"
)

var _storepath string
var _genvideoip string
var _clouddeskip string
var _analypath string

func init() {
   _storepath = utils.Conf().Section("base").Key("m3u8_dir_path").MustString("") 
   _genvideoip = utils.Conf().Section("base").Key("genvideoip").MustString("")
   _clouddeskip = utils.Conf().Section("base").Key("clouddeskip").MustString("")  
   _analypath = utils.Conf().Section("base").Key("desk_analy_path").MustString("")
}

func GetStorePath() string {
        return _storepath
}

func GetGenVideoIp() string {
     return _genvideoip
}

func GetCloudDeskIp() string {
     return _clouddeskip
}

func GetDeskAnalyPath() string {
     return _analypath
}



