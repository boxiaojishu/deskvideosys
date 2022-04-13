package routers

import (
        "log"
        "time"
        "github.com/deskvideosys/deskvideosys/models"
        "github.com/gin-gonic/gin"
        //"github.com/penggy/sessions"
        "github.com/penggy/EasyGoLib/db"
        "github.com/penggy/EasyGoLib/utils"
 )

type userrecordst struct {
    UserCount    int `json:"userCount"`
    UserList    interface{} `json:"userList"`
}

type userform struct {
    Start     int    `form:"start"`
    Limit     int    `form:"limit"`
    Q         string `form:"q"`
    Sort      string `form:"sort"`
    Order     string `form:"order"`
}


func (h *APIHandler) Userlist(c *gin.Context) {
  /*  userfm := &userform{}
        //form := utils.NewPageForm()
        if err := c.Bind(userfm); err != nil {
                return
        }
*/
    datas := make([]interface{}, 0) 
    var users []models.User
        log.Println("parent node")
        db.SQLite.Find(&users)
        if err := db.SQLite.Find(&users).Error; err != nil {
            log.Printf("find stream err:%v", err)
            return
        }

        for i := len(users) - 1; i > -1; i-- {
            v := users[i]
            log.Println(mapToJson(v))
            //var isenable bool
            
                datas = append(datas, map[string]interface{}{
                        "ID":        v.ID,
                        "Username":       v.Username,
                        "Role":      v.Role,
                })
        }
         var analys userrecordst
        analys.UserCount = len(users)
        analys.UserList = datas
        //pr := utils.NewPageResult(datas)
        //if userfm.Sort != "" {
        //        pr.Sort(userfm.Sort, userfm.Order)
       // }
        //pr.Slice(windowfm.Start, windowfm.Limit)
        //log.Println(mapToJson(records))
       // log.Println(analys.Name)
        //log.Println(mapToJson(pr.Rows))
        c.IndentedJSON(200, analys)

    
}

func (h *APIHandler) UserdefaultPwd(c *gin.Context) {

       c.IndentedJSON(200, "123456")
}

type saveform struct {
    Id       string `form:"id"`
    Username     string    `form:"username"`
    Role     string    `form:"role"`
    Enable         bool `form:"enable"`
}

func (h *APIHandler) Usersave(c *gin.Context) {
    userfm := &saveform{}
        //form := utils.NewPageForm()
        if err := c.Bind(userfm); err != nil {
                return
        }
      log.Println(userfm.Id)
      log.Println(userfm.Username)
      log.Println(userfm.Role)
      pmd5 := utils.MD5("123456")
      if userfm.Id == ""  {
          log.Println("user is not exist")
          //sess := sessions.Default(c)
          //uname := sess.Get("uname")
          tempcont :=   &models.User{Username:userfm.Username,Role:userfm.Role,Password:pmd5}
          //tempcont :=   &models.User{Username:userfm.Username,Password:pmd5,}
          db.SQLite.Create(tempcont)
      } else {
          log.Println("user is exist")
          var user models.User
          data := db.SQLite.First(&user, "id = ?", userfm.Id)
          if data.RecordNotFound()  {
                 log.Println(userfm.Id + " is not find")
          } else {
              db.SQLite.Model(&user).Updates(models.User{Username:userfm.Username,Role:userfm.Role}) 
          }
      }

}

type upasswdform struct {
    Id     string    `form:"id"`
    Password  string `form:"password"`
}

func (h *APIHandler) Usersetpassword(c *gin.Context) {
    upfm := &upasswdform{}
        //form := utils.NewPageForm()
        if err := c.Bind(upfm); err != nil {
                return
        }

        var user models.User
        db.SQLite.First(&user,"id = ?", upfm.Id)
        db.SQLite.Model(&user).Update("password", utils.MD5(upfm.Password))
        db.SQLite.Model(&user).Update("updatedat", time.Now().Format("2006-01-02 15:04:05"))
      log.Println(upfm.Id)
      log.Println(upfm.Password)
      

}

type uremoveform struct {
    Id     string    `form:"id"`
}

func (h *APIHandler) Userremove(c *gin.Context) {
     removefm  := &uremoveform{}
      if err := c.Bind(removefm); err != nil {
                return
        }
        log.Println(removefm.Id)
      var user models.User
        user.ID = removefm.Id
        db.SQLite.Delete(user)        

}

