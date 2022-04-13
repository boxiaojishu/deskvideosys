package tools

import (
        "fmt"
        "strconv"
)


/*
时间常量
*/
const (

    //定义每分钟的秒数
    SecondsPerMinute = 60
    //定义每小时的秒数
    SecondsPerHour = SecondsPerMinute * 60
    //定义每天的秒数
    SecondsPerDay = SecondsPerHour * 24
    
    MinutePerHour = 60
 
)

/*
时间转换函数
*/
func ResolveTime(seconds int) (day int, hour int, minute int) {
    //每分钟秒数
    minute = seconds / SecondsPerMinute
    //每小时秒数
    hour = seconds / SecondsPerHour
    //每天秒数
    day = seconds / SecondsPerDay
    return
}



func ConvertTime(seconds int)(hour int, minute int,second int){
    dmin := seconds / SecondsPerMinute
    second = seconds % SecondsPerMinute
   
    minute = dmin % MinutePerHour

   hour = dmin / MinutePerHour
   return
}

func IntConvertTime(seconds int)(strtime string){
     intsec := seconds //, _ := strconv.Atoi(seconds)
     hour,minute,secound := ConvertTime(intsec)
     strtime= fmt.Sprintf("%6d小时%02d分%02d秒",hour,minute,secound)
     return
}


func StrConvertTime(seconds string)(strtime string){
     intsec, _ := strconv.Atoi(seconds)
     hour,minute,secound := ConvertTime(intsec)
     strtime= fmt.Sprintf("%6d小时%02d分%02d秒",hour,minute,secound)  
     return
}

