/*
 * Copyright (c) 2021-2031, 深圳市柏晓技术有限公司
 * All rights reserved.
 *
 */

package routers

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/penggy/EasyGoLib/db"

	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/penggy/EasyGoLib/utils"
	"github.com/penggy/cors"
	"github.com/penggy/sessions"
	validator "gopkg.in/go-playground/validator.v8"
)

/**
 * @apiDefine simpleSuccess
 * @apiSuccessExample 成功
 * HTTP/1.1 200 OK
 */

/**
 * @apiDefine authError
 * @apiErrorExample 认证失败
 * HTTP/1.1 401 access denied
 */

/**
 * @apiDefine pageParam
 * @apiParam {Number} start 分页开始,从零开始
 * @apiParam {Number} limit 分页大小
 * @apiParam {String} [sort] 排序字段
 * @apiParam {String=ascending,descending} [order] 排序顺序
 * @apiParam {String} [q] 查询参数
 */

/**
 * @apiDefine pageSuccess
 * @apiSuccess (200) {Number} total 总数
 * @apiSuccess (200) {Array} rows 分页数据
 */

/**
 * @apiDefine timeInfo
 * @apiSuccess (200) {String} rows.createAt 创建时间, YYYY-MM-DD HH:mm:ss
 * @apiSuccess (200) {String} rows.updateAt 结束时间, YYYY-MM-DD HH:mm:ss
 */

var Router *gin.Engine

func init() {
	mime.AddExtensionType(".svg", "image/svg+xml")
	mime.AddExtensionType(".m3u8", "application/vnd.apple.mpegurl")
	// mime.AddExtensionType(".m3u8", "application/x-mpegurl")
	mime.AddExtensionType(".ts", "video/mp2t")
	// prevent on Windows with Dreamware installed, modified registry .css -> application/x-css
	// see https://stackoverflow.com/questions/22839278/python-built-in-server-not-loading-css
	mime.AddExtensionType(".css", "text/css; charset=utf-8")

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = utils.GetLogWriter()
}

func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch err.Type {
			case gin.ErrorTypeBind:
				switch err.Err.(type) {
				case validator.ValidationErrors:
					errs := err.Err.(validator.ValidationErrors)
					for _, err := range errs {
						sec := utils.Conf().Section("localize")
						field := sec.Key(err.Field).MustString(err.Field)
						tag := sec.Key(err.Tag).MustString(err.Tag)
						c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("%s %s", field, tag))
						return
					}
				default:
					log.Println(err.Err.Error())
					c.AbortWithStatusJSON(http.StatusBadRequest, "Inner Error")
					return
				}
			}
		}
	}
}

func NeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sessions.Default(c).Get("uid") == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		c.Next()
	}
}

func Init() (err error) {

        go receiveMessage()
	Router = gin.New()
	pprof.Register(Router)
	// Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(Errors())
	Router.Use(cors.Default())

	store := sessions.NewGormStoreWithOptions(db.SQLite, sessions.GormStoreOptions{
		TableName: "t_sessions",
	}, []byte("柏晓技术@2022"))
	tokenTimeout := utils.Conf().Section("http").Key("token_timeout").MustInt(7 * 86400)
	store.Options(sessions.Options{HttpOnly: true, MaxAge: tokenTimeout, Path: "/"})
	sessionHandle := sessions.Sessions("token", store)

	{
		wwwDir := filepath.Join(utils.DataDir(), "www")
		Router.Use(static.Serve("/", static.LocalFile(wwwDir, true)))
	}

	{
		api := Router.Group("/api/v1").Use(sessionHandle)
		api.GET("/login", API.Login)
		api.GET("/userinfo", API.UserInfo)
		api.GET("/logout", API.Logout)
		api.GET("/defaultlogininfo", API.DefaultLoginInfo)
		api.GET("/modifypassword", NeedLogin(), API.ModifyPassword)
		api.GET("/serverinfo", API.GetServerInfo)
		api.GET("/restart", API.Restart)

                api.GET("/user/list", API.Userlist)
                api.GET("/user/defaultPwd", API.UserdefaultPwd)
                api.GET("/user/save", API.Usersave)
                api.GET("/user/remove", API.Userremove)
                api.GET("/user/resetPwd", API.Usersetpassword)
                api.GET("/deskcontrols", API.Deskcontrols)
                api.GET("/cloudesks", API.Cloudesks)
                api.GET("/windowrecord/querydaily",API.Windowrecord)    
                api.GET("/windowrecord/queryflags",API.Windowrecordflag)
                api.GET("/windowrecord/removedaily",API.Windowrecordremove)    
                api.GET("/tradidesk/tradidesklist",API.TradiDeskList)
                api.GET("/tradidesk/tradideskconfig",API.TradiDeskConfig)
                api.GET("/tradidesk/remove",API.TradiDeskRemove)

	}

	{

		mp4Path := utils.Conf().Section("base").Key("m3u8_dir_path").MustString("")
		if len(mp4Path) != 0 {
			Router.Use(static.Serve("/record", static.LocalFile(mp4Path, true)))
                        Router.Use(static.Serve("/operation", static.LocalFile(mp4Path, true)))
		}

	}

	return
}
