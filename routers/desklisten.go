/*
 * Copyright (c) 2021-2031, cbrdesk.com
 * All rights reserved.
 *
 */
package routers
import (
    "net"
    "encoding/csv"
    "fmt"
    "os"
    "log"
    "github.com/penggy/EasyGoLib/db"
    "github.com/deskvideosys/deskvideosys/models"
    "time"
    "math"
    //"strconv"
    "strings"
)
func receiveMessage() {
    netListen, err := net.Listen("tcp4", "0.0.0.0:12345")
    CheckError(err)
    defer netListen.Close()
    Log("Waiting for clients ...")
    for{
        conn, err := netListen.Accept()     
        if err != nil{
            continue        
        }
        Log(conn.RemoteAddr().String(), "tcp connect success")  
        go handleConnection(conn)
    }
}

func writerecord(oldtime time.Time,strpath string,strinfo string,strfile *string,iflag int) {
    var nowtime = time.Now() 
    duration := nowtime.Sub(oldtime)
    formatNow := nowtime.Format("2006-01-02 15:04:05")
    formatOld := oldtime.Format("2006-01-02 15:04:05")
    log.Println("writerecord is start");
    err := os.MkdirAll(strpath, 0766)
    strnow := strings.Split(formatNow, " ")
    strold := strings.Split(formatOld, " ")
    strdate := fmt.Sprintf("%04d%02d%02d",nowtime.Year(),nowtime.Month(),nowtime.Day())
    strtime := fmt.Sprintf("%02d%02d%02d",nowtime.Hour(),nowtime.Minute(),nowtime.Second())
    if strnow[0] != strold[0] || *strfile == "" {
        err := os.MkdirAll(strpath + "/" + strdate, 0766)
        *strfile = strpath + "/" + strdate + "/" + strtime + ".csv"
        if err != nil {
             log.Println(err)
        }
    }
    //strdura := duration.Hours() + ":" + duration.Minutes() + ":" + duration.Seconds()
    var allsec int = int(math.Floor(duration.Seconds()))
    strdura := fmt.Sprintf("%d", allsec)
    //strdura := fmt.Sprintf("%02d:%02d:%02d", ihour, imin, isec) 
    stritem := strings.Split(strinfo, ",")
    file, err := os.OpenFile(*strfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
    if err != nil {
        log.Println("open file is failed, err: ", err)
    }
    defer file.Close()
    w := csv.NewWriter(file)
    w.Write([]string{formatOld, stritem[0],stritem[1], formatNow, strdura})
    /*if iflag == 0 {
        w.Write([]string{formatOld, stritem[2],stritem[3], formatNow, strdura})
    }  else {
        nflag,_ := strconv.Atoi(stritem[4])
        if nflag != 0 { 
            w.Write([]string{formatOld, stritem[0],stritem[1], formatNow, stritem[4]})
        }
    }*/
    w.Flush()
}

func handleConnection(conn net.Conn) {
    var oldtime = time.Now()
    var strold string
    var startflag int = 0
    var fpath string
    var strfile string
    //var sysdeskid string
    buffer := make([]byte, 2048)        //建立一个slice
    for{
        n, err := conn.Read(buffer)     //读取客户端传来的内容
        if err != nil{
            writerecord(oldtime,fpath,strold,&strfile,0)    
            Log(conn.RemoteAddr().String(), "connection error: ", err)
            return      //当远程客户端连接发生错误（断开）后，终止此协程。
        } 
        strbuf := string(buffer[:n])
        log.Println(strbuf)
        deskc := "deskname:"
        deskn := strbuf[:9]
        if deskn == deskc {
           // var sysdeskid string
            var strloc string
            var strname string
            //sysdeskid = strbuf[9:]
            var desk models.Desk
            data := db.SQLite.First(&desk, "deskkey = ?", strbuf[9:])
            if data.RecordNotFound()  {
                   log.Println("return is not exist :"  + strbuf + " lower : " + strings.ToLower(strbuf[9:]))
                   return;
            } else {
                //sysdeskid = desk.Id
                strloc = desk.Username
                strname = desk.Deskname
                log.Println("pc desk : " + desk.Username)
            }  
         fpath = "/cbr/test/" + "/" + strloc + "/" + strname
       }  else {
          log.Println(strbuf + "  0:" + time.Now().Format("20060102 15:04:05"))
          if(startflag != 0) {
             writerecord(oldtime,fpath,strbuf,&strfile,1) 
             log.Println(strbuf + ":" + time.Now().Format("20060102 15:04:05"))
          }  else { startflag = startflag + 1 }
          strold = strbuf
          oldtime = time.Now()          
       }
    }
}
func Log(v ...interface{}) {
    log.Println(v...)
}
func CheckError(err error) {
    if err != nil{
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    }
}

