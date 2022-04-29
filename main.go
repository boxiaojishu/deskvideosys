/*
 * Copyright (c) 2021-2031, 深圳市柏晓技术有限公司
 * All rights reserved.
 *
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
        //"encoding/json"
        //"bytes"
        //"io/ioutil"
	"strings"
	"time"
        //"os/exec"
	//"github.com/penggy/EasyGoLib/db"
        //"github.com/robfig/cron"

	"github.com/deskvideosys/deskvideosys/models"
	"github.com/deskvideosys/deskvideosys/routers"
	figure "github.com/common-nighthawk/go-figure"
	"github.com/penggy/EasyGoLib/utils"
	"github.com/penggy/service"
)

var (
	gitCommitCode string
	buildDateTime string
)

type program struct {
	httpPort   int
	httpServer *http.Server
}

func (p *program) StopHTTP() (err error) {
	if p.httpServer == nil {
		err = fmt.Errorf("HTTP Server Not Found")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = p.httpServer.Shutdown(ctx); err != nil {
		return
	}
	return
}


func (p *program) StartHTTP() (err error) {
	p.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", p.httpPort),
		Handler:           routers.Router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	link := fmt.Sprintf("http://%s:%d", utils.LocalIP(), p.httpPort)
	log.Println("http server start -->", link)
	go func() {
		if err := p.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("start http server error", err)
		}
		log.Println("http server end")
	}()
	return
}


func (p *program) Start(s service.Service) (err error) {
	log.Println("********** START **********")
	if utils.IsPortInUse(p.httpPort) {
		err = fmt.Errorf("HTTP port[%d] In Use", p.httpPort)
		return
	}
	err = models.Init()
	if err != nil {
		return
	}
	err = routers.Init()
	if err != nil {
		return
	}
       	//p.StartRTSP()
	p.StartHTTP()

	if !utils.Debug {
		log.Println("log files -->", utils.LogDir())
		log.SetOutput(utils.GetLogWriter())
	}
	go func() {
		for range routers.API.RestartChan {
			p.StopHTTP()
			//p.StopRTSP()
			utils.ReloadConf()
			//p.StartRTSP()
			p.StartHTTP()
		}
	}()

	return
}

func (p *program) Stop(s service.Service) (err error) {
	defer log.Println("********** STOP **********")
	defer utils.CloseLogWriter()
	p.StopHTTP()
	models.Close()
	return
}

/*func startcron() {
        log.Println(" cron is start")
        c := cron.New()
        c.AddFunc("30 23 * * *", func() {
                log.Println("Run start cron...")
                timers.RmvVideoRecord()
        })
        c.Start()
        select{}

}*/

func main() {
	flag.StringVar(&utils.FlagVarConfFile, "config", "", "configure file path")
	flag.Parse()
	tail := flag.Args()

	// log
	log.SetPrefix("[DeskVideosys] ")
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Printf("git commit code:%s", gitCommitCode)
	log.Printf("build date:%s", buildDateTime)
	routers.BuildVersion = fmt.Sprintf("%s.%s", routers.BuildVersion, gitCommitCode)
	routers.BuildDateTime = buildDateTime
        //go startcron() 
     
 
      
	sec := utils.Conf().Section("service")
	svcConfig := &service.Config{
		Name:        sec.Key("name").MustString("DeskVideosys_Service"),
		DisplayName: sec.Key("display_name").MustString("DeskVideosys_Service"),
		Description: sec.Key("description").MustString("DeskVideosys_Service"),
	}

	httpPort := utils.Conf().Section("http").Key("port").MustInt(10008)
	p := &program{
		httpPort:   httpPort,
	}
	s, err := service.New(p, svcConfig)
	if err != nil {
		log.Println(err)
		utils.PauseExit()
	}
	if len(tail) > 0 {
		cmd := strings.ToLower(tail[0])
		if cmd == "install" || cmd == "stop" || cmd == "start" || cmd == "uninstall" {
			figure.NewFigure("DeskVideosys", "", false).Print()
			log.Println(svcConfig.Name, cmd, "...")
			if err = service.Control(s, cmd); err != nil {
				log.Println(err)
				utils.PauseExit()
			}
			log.Println(svcConfig.Name, cmd, "ok")
			return
		}
	}
	figure.NewFigure("DeskVideosys", "", false).Print()
	if err = s.Run(); err != nil {
		log.Println(err)
		utils.PauseExit()
	}
}
