package routers

import (
        "github.com/deskvideosys/deskvideosys/tools"
        "github.com/gin-gonic/gin"
        "log"
        "os"
        "io"
        "net/http"
        "strings"
        "bufio"
        "io/ioutil"
        "github.com/penggy/EasyGoLib/utils"
        //"time"
)


type Windowrecordst struct {
    Name    string `json:"name"`
    List    interface{} `json:"list"`
}

func NewwindowResult(records interface{}) *Windowrecordst{
                return &Windowrecordst{
                         Name:"test",
                   List:[]interface{}{records},

                }
}

type Windowform struct {
    User      string `form:"user"`
    Virname   string `form:"virname"`
    Period    string `form:"period"`
    Start     int    `form:"start"`
    Limit     int    `form:"limit"`
    Q         string `form:"q"`
    Sort      string `form:"sort"`
    Order     string `form:"order"`
}

/*func filefindline(filename string,linfo string) bool {
    fd, err := os.OpenFile(filename, os.O_RDWR, 0666)
        defer fd.Close()
        if err != nil {
               log.Println(err)
                return false;
        }

    scanner := bufio.NewScanner(fd)
    //Reading lines
    for scanner.Scan() {
        line := scanner.Text()
         if strings.Index(line,linfo) == 0 {
            return true
         }
    }
    return false
}*/

func (h *APIHandler) Windowrecord(c *gin.Context) {

        windowfm := &Windowform{}
        //form := utils.NewPageForm()
        if err := c.Bind(windowfm); err != nil {
                return
        }
        records := make([]interface{}, 0)
        rootpath := "/cbr/test/"
        userpath :=  windowfm.User + "/" + windowfm.Virname + "/" + windowfm.Period
        recordpath2 := rootpath + "/" + userpath 
        //importpath := rootpath + "/" + userpath + "/important.txt"        

        log.Println("it is hello:")
        log.Println(windowfm.User)
         log.Println(windowfm.Start)
        log.Println(windowfm.Limit)
       
        files, err := ioutil.ReadDir(recordpath2)
        if err != nil {
          log.Println(err.Error())
          c.AbortWithStatusJSON(http.StatusBadRequest, "no records")
          return
        }
 
        for _, file := range files {
           //fmt.Println(file.Name())
           recordpath := recordpath2 + "/" + file.Name()
        


        f, err := os.Open(recordpath)
        if err != nil {
           log.Println(err.Error())
           c.AbortWithStatusJSON(http.StatusBadRequest, "no records")
           return
        }
        //建立缓冲区，把文件内容放到缓冲区中
        buf := bufio.NewReader(f)
        for {
            //遇到\n结束读取
            videoinfo, errR := buf.ReadBytes('\n')
            if errR != nil {
                if errR == io.EOF {
                    break
                }
                log.Println(errR.Error())
                c.AbortWithStatusJSON(http.StatusBadRequest, "no records")
                return
            }
            //fmt.Println(string(b))
            videoinfo2 := strings.Replace(string(videoinfo), "\n", "", -1)
            info := strings.Split(videoinfo2, ",") 
            if(len(info) < 5){
                continue 
            }
            /*timeTemplate4 := "15:04:05"
            importinfo := filefindline(importpath,info[0])
            stamp1, _ := time.ParseInLocation(timeTemplate4, info[0], time.Local)
            stamp2, _ := time.ParseInLocation(timeTemplate4, info[2], time.Local)
            log.Println(info[2])  
            log.Println(stamp1)
            vpath := "/record/"  + windowfm.User + "/" + windowfm.Virname + "/" +  info[1]
            log.Println(stamp2)
            new1 := stamp2.Sub(stamp1)
            log.Println(new1.String())*/
            //log.Println(strings.Replace(info[4], "\r", "", -1) )
            records = append(records, map[string]interface{}{
                        "user":        windowfm.User,
                        "virname":       windowfm.Virname,
                        "name":      "liqin",
                        "location":    info[1],
                        "startAt": info[0],
                        "winname": info[1],
                        "appname": info[2],
                        "duration": tools.StrConvertTime(strings.Replace(info[4], "\r", "", -1)),
                        /*"outBytes":  pusher.OutBytes(),
                        "startAt":   utils.DateTime(pusher.StartAt()),
                        "onlines":   len(pusher.GetPlayers()),*/
                })
        }
        }
        // desks := NewdeskResult(records)
        /*desks :=Deskrecordst{
                   Name:"test",
                   List:[]interface{}{records},
                  }*/
        var windows Windowrecordst
        windows.Name = "test"
        windows.List = records
        pr := utils.NewPageResult(records)
        if windowfm.Sort != "" {
                pr.Sort(windowfm.Sort, windowfm.Order)
        }
        //pr.Slice(windowfm.Start, windowfm.Limit)
        //log.Println(mapToJson(records))
        log.Println(windows.Name)
        log.Println(mapToJson(pr.Rows))
        c.IndentedJSON(200, pr)


     
}


/*func FlagAllByVieo(monthflag map[string][]byte,  fileDir string)  {
    files, _ := ioutil.ReadDir(fileDir)
    for _, onefile := range files {
           log.Println(onefile.Name())
          if len(onefile.Name()) != 8 {
             continue
           }
           monthi := onefile.Name()[0:6]
          if( onefile.IsDir() ){
                      
                     datei := onefile.Name()[6:8]
                     i,_ := strconv.Atoi(datei)
                     if(monthflag[monthi] == nil){
                         var chars = []byte("0000000000000000000000000000000000")
                         monthflag[monthi] = chars
                      }                  

                     monthflag[monthi][i-1]='1'
               
          }
    }
}

type flagform struct {
    User      string `form:"user"`
    Virname   string `form:"virname"`
}

type flagdata struct {
       Data interface{} `json:"data"`
    }
*/
func (h *APIHandler) Windowrecordflag(c *gin.Context) {

        deskflag := &flagform{}
        //form := utils.NewPageForm()
        if err := c.Bind(deskflag); err != nil {
                return
        }
        log.Println(deskflag.User)
        log.Println(deskflag.Virname)
        var testd map[string][]byte
        testd = make(map[string][]byte, 0)
        rootpath := "/cbr/test/"
        storepath :=rootpath + "/" + deskflag.User + "/" + deskflag.Virname + "/"
       // storepath :="/opt/store/liwei/test10/"
        FlagAllByVieo(testd,storepath);
        var datas map[string]interface{}
        datas = make(map[string]interface{}, 0)
        for strmonth := range testd {
            datas[strmonth] = string(testd[strmonth])
            log.Println(datas[strmonth])
        }
        
        /*//datas := map[string]interface{}{
             datas := map[string][]byte{
            // "202010": "11111101000000110110100101001",
        //}    
       */

        log.Println(deskflag.User)
        //gin.H{"message": "hello world"}
       c.IndentedJSON(200, &datas)
      // c.IndentedJSON(200, flagdata{Data:&datas})   
}

/*type rcrmvform struct {
    User      string `form:"user"`
    Virname   string `form:"virname"`
    Period    string `form:"period"`
}
*/

func (h *APIHandler) Windowrecordremove(c *gin.Context) {
        windowrcrmv := &rcrmvform{}
        //form := utils.NewPageForm()
        if err := c.Bind(windowrcrmv); err != nil {
                return
        }
        rootpath := "/cbr/test/"

        storepath := rootpath + "/" + windowrcrmv.User + "/" + windowrcrmv.Virname + "/" + windowrcrmv.Period + "/"
        log.Println(windowrcrmv.User)
        log.Println(windowrcrmv.Virname)
        log.Println(windowrcrmv.Period)

        os.RemoveAll(storepath)
}



