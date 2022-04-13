package routers

import (
        "github.com/gin-gonic/gin"
        "encoding/json"
        "log"
        "os"
        "strings"
        "bufio"
        "io/ioutil"
       "strconv"
        "github.com/penggy/EasyGoLib/utils"
)

type result struct {
    Virid    string `json:"Virid"`
    Username string `json:"Username"`
    Virname string `json:"Virname"`
    Isrecord bool `json:"Isrecord"`
    Isrc     int    
    Savedays int
}

func mapToJson(result interface{}) string {
    // map转 json str
    jsonBytes, _ := json.Marshal(result)
    jsonStr := string(jsonBytes)
    return jsonStr
}

func (h *APIHandler) Cloudesks(c *gin.Context) {
        log.Println("cloud desks is start0")
        form := utils.NewPageForm()
        if err := c.Bind(form); err != nil {
                return
        }
        log.Println("cloud desks is start1")
        virtuals := make([]interface{}, 0)
        gettradidesk(&virtuals)
        pr := utils.NewPageResult(virtuals)
        if form.Sort != "" {
                pr.Sort(form.Sort, form.Order)
        }
        pr.Slice(form.Start, form.Limit)
        log.Println(mapToJson(pr.Rows))
        c.IndentedJSON(200, pr)


}


type Deskrecordst struct {
    Name    string `json:"name"`
    List    interface{} `json:"list"`
}

func NewdeskResult(records interface{}) *Deskrecordst{
                return &Deskrecordst{
                         Name:"test",
                   List:[]interface{}{records},

                }
}

type Deskform struct {
    User      string `form:"user"`
    Virname   string `form:"virname"`
    Period    string `form:"period"`
    Sort      string `form:"sort"`
    Order     string `form:"order"`
}

func filefindline(filename string,linfo string) bool {
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
}


type flagform struct {
    User      string `form:"user"`
    Virname   string `form:"virname"`
}

func FlagAllByVieo(monthflag map[string][]byte,  fileDir string)  {
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

type rcrmvform struct {
    User      string `form:"user"`
    Virname   string `form:"virname"`
    Period    string `form:"period"`
}

