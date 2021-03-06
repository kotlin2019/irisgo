/**
*@Author: haoxiongxiao
*@Date: 2019/4/3
*@Description: CREATE GO FILE commands
 */
package api_template

import (
	"github.com/xhaoxiong/irisgo/utils"
	"log"
	"os"
	"path"
	"strings"
)

var main = `package main

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	
	"{{.Appname}}/web/middleware"
	"{{.Appname}}/config"
	"{{.Appname}}/models"
	"{{.Appname}}/route"
)

var (
	cfg = pflag.StringP("config", "c", "", "./config.yaml")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	models.DB.Init()
	app := newApp()

	route.InitRouter(app)

	app.Run(iris.Addr(viper.GetString("addr")))
}

func newApp() *iris.Application {
	app := iris.New()
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	//	AllowCredentials: true,
	//	AllowedHeaders:   []string{"*"},
	//})
	//
	//app.Use(crs) //
	//app.StaticWeb("/assets", "./web/views/admin/assets")
	//app.RegisterView(iris.HTML("./web/views/admin", ".html"))
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.GetJWT().Serve)//是否启用jwt中间件
	app.Use(middleware.LogMiddle)//是否启用logrus中间件
	app.Configure(iris.WithOptimizations)

	return app
}

`

func CreatedApp(appPath, appName string) {
	log.Println("Creating application...")
	os.MkdirAll(appName, 0755)
	os.Mkdir(path.Join(appName, "conf"), 0755)
	os.Mkdir(path.Join(appName, "config"), 0755)
	os.Mkdir(path.Join(appName, "models"), 0755)
	os.Mkdir(path.Join(appName, "route"), 0755)
	os.Mkdir(path.Join(appName, "repositories"), 0755)
	os.Mkdir(path.Join(appName, "service"), 0755)
	os.MkdirAll(path.Join(appName, "/web/controllers"), 0755)
	os.MkdirAll(path.Join(appName, "/web/middleware"), 0755)
	utils.WriteToFile(path.Join(appName, "conf", "config.yaml"), conf)
	utils.WriteToFile(path.Join(appName, "config", "config.go"), config)
	utils.WriteToFile(path.Join(appName, "models", "init.go"), mysqlInit)
	utils.WriteToFile(path.Join(appName, "models", "log.go"), traceLogger)
	utils.WriteToFile(path.Join(appName, "service", "TestService.go"), strings.Replace(service, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "repositories", "TestRepo.go"), strings.Replace(repositories, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "route", "route.go"), strings.Replace(route, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "/web/controllers", "TestController.go"), controllers)
	utils.WriteToFile(path.Join(appName, "/web/controllers", "Common.go"), common)
	utils.WriteToFile(path.Join(appName, "/web/middleware", "jwt.go"), strings.Replace(jwt, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "/web/middleware", "logrus.go"), strings.Replace(logrus, "{{.Appname}}", appName, -1))
	utils.WriteToFile(path.Join(appName, "main.go"), strings.Replace(main, "{{.Appname}}", appName, -1))
	log.Println("new application successfully created!")
}
