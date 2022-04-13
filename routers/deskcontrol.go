package routers

import (
        "github.com/gin-gonic/gin"
        "log"
        "github.com/penggy/EasyGoLib/utils"
        "github.com/penggy/EasyGoLib/db"
        "github.com/deskvideosys/deskvideosys/models"
)


type controlresult struct {
    Virid    string `json:"Virid"`
    Username string `json:"Username"`
    Virname string `json:"Virname"`
    Isblue bool `json:"Isblue"`
    Isusb bool `json:"Isusb"`
    Isscreen bool `json:"Isscreen"`
    Isprint bool `json:"Isprint"`
    screendesc string `json:"Screendesc"`
    printdesc string `json:"Printdesc"`
}

func (h *APIHandler) Deskcontrols(c *gin.Context) {
        log.Println("cloud desks is start0")
        form := utils.NewPageForm()
        if err := c.Bind(form); err != nil {
                return
        }
        log.Println("cloud desks is start1")
        //var controls []models.Control
        // db.SQLite.Find(&controls)
        //hostname := utils.GetRequestHostname(c.Request)
        virtuals := make([]interface{}, 0)
        gettradidesk(&virtuals)
        //gettradidesk(virtuals)
        //log.Println(mapToJson(virtuals))
        //log.Println(mapToJson(results))
        pr := utils.NewPageResult(virtuals)
        if form.Sort != "" {
                pr.Sort(form.Sort, form.Order)
        }
        pr.Slice(form.Start, form.Limit)
        log.Println(mapToJson(pr.Rows))
        c.IndentedJSON(200, pr)


}

func gettradidesk(virtuals * []interface{}) int {
    var desks []models.Desk
    db.SQLite.Find(&desks)
    if err := db.SQLite.Find(&desks).Error; err != nil {
        log.Printf("find desk err:%v", err)
        return -1
    }
    for i := len(desks) - 1; i > -1; i-- {
        v := desks[i]
        /*if form.Q != "" && !strings.Contains(strings.ToLower(v.Username), strings.ToLower(form.Q)) {
            continue
        }*/

        var control models.Control
        data := db.SQLite.First(&control, "id = ?", v.Id)
        if data.RecordNotFound()  {
            log.Println("control is not find")
            *virtuals = append(*virtuals, map[string]interface{}{
                        "Virid":        v.Id,
                        "Username":       v.Username,
                        "Virname":      v.Deskname,
                        "Isusb":    false,
                        "Isblue": false,
                        "Isscreen": false,
                        "Isprint": false,
                         "Screendesc" :"",
                         "Printdesc" :"",
                        /*"inBytes":   pusher.InBytes(),
                        "outBytes":  pusher.OutBytes(),
                        "startAt":   utils.DateTime(pusher.StartAt()),
                        "onlines":   len(pusher.GetPlayers()),*/
                })

        } else {
                 log.Println("is usb:")
                      var isusb,isblue,isscreen,isprint bool
                      log.Println(control.Isusb)
                      if control.Isusb == nil {
                         isusb = false
                      } else {
                          isusb = *control.Isusb
                      }
                      if control.Isblue == nil {
                         isblue = false
                      } else {
                          isblue = *control.Isblue
                      }
                      if control.Isscreen == nil {
                         isscreen = false
                      } else {
                          isscreen = *control.Isscreen
                      }
                      if control.Isprint == nil {
                         isprint = false
                      } else {
                          isprint = *control.Isprint
                      }
                      *virtuals = append(*virtuals, map[string]interface{}{
                        "Virid":        v.Id,
                        "Username":       v.Username,
                        "Virname":      v.Deskname,
                        "Screendesc" :control.Screendesc,
                        "Printdesc" :control.Printdesc,
                        "Isusb":    isusb,
                        "Isblue": isblue,
                        "Isscreen": isscreen,
                        "Isprint": isprint,
                        /*"inBytes":   pusher.InBytes(),*/
                })


        }
        log.Println(mapToJson(v))
        //videos = append(videos,v)
    }
 
return 0;
}





