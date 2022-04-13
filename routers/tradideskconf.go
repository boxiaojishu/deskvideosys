package routers

import (
        "github.com/gin-gonic/gin"
        "log"
        "fmt"
        "github.com/deskvideosys/deskvideosys/tools"
        "github.com/penggy/EasyGoLib/utils"
        "github.com/penggy/EasyGoLib/db"
        "github.com/deskvideosys/deskvideosys/models"
//        "strconv"
)

func (h *APIHandler) TradiDeskList(c *gin.Context) {

        form := utils.NewPageForm()
        if err := c.Bind(form); err != nil {
                return
        }
        tradidesks := make([]interface{}, 0) 
        
        var desks []models.Desk
        db.SQLite.Find(&desks)
        if err := db.SQLite.Find(&desks).Error; err != nil {
            log.Printf("find stream err:%v", err)
            return
        }
        for i := len(desks) - 1; i > -1; i-- {
            v := desks[i]
            log.Println(mapToJson(v))
            tradidesks = append(tradidesks,v)
        }
        log.Println(mapToJson(tradidesks))
        pr := utils.NewPageResult(tradidesks)
        if form.Sort != "" {
                pr.Sort(form.Sort, form.Order)
        }
        pr.Slice(form.Start, form.Limit)
        c.IndentedJSON(200, pr)

}

func (h *APIHandler) TradiDeskConfig(c *gin.Context) {
        type Form struct {
                Deskname         string `form:"deskname" binding:"required"`
                Deskloc          string `form:"deskloc" binding:"required"`
                Username          string `form:"username"`
                Deskdesc           string `form:"deskdesc"`
                Deskkey            string `form:"deskkey"`
                IsRecord          bool   `form:"isRecord"`
                Savedays       int    `form:"savedays"`
        }
        var form Form
        err := c.Bind(&form)
        if err != nil {
                log.Printf("Pull to push err:%v", err)
                return
        }
        log.Println("TradiDeskConfig is start")
        genuuid := tools.UniqueId()
        log.Println(form.IsRecord)
        log.Println(form.Savedays)

        var desk = models.Desk{
                Id:               genuuid,
                Deskname:          form.Deskname,
                Deskloc:          form.Deskloc,
                Username:          form.Username,
                Deskdesc:           form.Deskdesc,
                Deskkey:           form.Deskkey,
                IsRecord:          form.IsRecord,
                Savedays:          form.Savedays,
        }
        if db.SQLite.Where(&models.Desk{Deskname: form.Deskname}).First(&models.Desk{}).RecordNotFound() {
                db.SQLite.Create(&desk)
        } else {
                db.SQLite.Save(&desk)
        }
        c.IndentedJSON(200, "1")
}

func (h *APIHandler) TradiDeskRemove(c *gin.Context) {
        type Form struct {
                Id string `form:"id" binding:"required"`
        }
        var form Form
        err := c.Bind(&form)
        if err != nil {
                log.Printf("stop pull to push err:%v", err)
                return
        }
        var desk models.Desk
        desk.Id = form.Id
        db.SQLite.Delete(desk)
        c.IndentedJSON(200, fmt.Sprintf("Id[%s] delete success", form.Id))
}


