# deskvideosys终端上网行为管理服务器


## 主要功能特点

- 基于Golang开发维护；

- 支持Windows、Linux、macOS平台；

- 支持用户管理；

- 终端自动匹配，不固定IP,通过机器码匹配；

- 支持终端信息采集

-支持收集后的终端信息查询

- 支持桌面终端配置；

- Web后台管理；

- 分布式负载均衡；


## 安装部署



- [下载解压 版本包](http://www.cbrdesk.com/page74)

- 直接运行(Windows)

    deskvideosys.exe
    
    以 `Ctrl + C` 停止服务

- 以服务启动(Windows)

    ServiceInstall-deskvideosys.exe
    
    以 ServiceUninstall-deskvideosys.exe 卸载 deskvideosys 服务

- 直接运行(Linux/macOS)

		cd deskvideosys
		./deskvideosys
		# Ctrl + C


- 查看界面
	
	打开浏览器输入 [http://localhost:10008](http://localhost:10008), 进入控制页面,默认用户名密码是admin/admin


## 效果图


## 二次开发

### 准备工具

        # go tools
        go get -u -v github.com/kardianos/govendor
        go get -u -v github.com/penggy/gobuild

        # npm tools
        npm i -g apidoc
        npm i -g rimraf


### 编译命令

- 获取代码

        cd $GOPATH/src/github.com
        mkdir deskvideosys && cd deskvideosys
        git clone https://github.com/boxiaojishu/deskvideosys.git --depth=1 deskvideosys
        cd deskvideosys

- 以开发模式运行

        npm run dev

- 以开发模式运行前端 Run as dev mode

        npm run dev:www       

- 编译前端  Build www

        cd web_src && npm i
        cd ..
        npm run build:www

- 编译 Windows 版本 Build windows version

        npm run build:win

- 编译 Linux/macOS 版本 (在 bash 环境下执行) Build linux/macOS version

        npm run build:lin       

- 清理编译文件 Clean

        npm run clean 

- 打包 Pack

        # install pack
        npm i -g @penggy/pack

        # for windows
        npm run build:win
        pack zip

        # for linux/macOS
        npm run build:lin
        pack tar

        # for clean
        pack clean


## 技术支持

- 邮件：[liqingbai@163.com](mailto:liqingbai@163.com) 

- QQ交流：**3316954823**

- 柏晓技术终端上网行为管理服务器是属于柏晓技术团队的免费产品，大家免费使用，同时，柏晓技术团队也能提供相应的收费技术咨询、技术服务和技术定制，谢谢大家支持！


## 获取更多信息

**deskvideosys**开源项目：[www.cbrdesk.com](http://www.cbrdesk.com)

Copyright &copy; boxiaozhishu Team 2020-2022

