package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-file-manager/config"
	"go-file-manager/crontab"
	"go-file-manager/handler"
	"html/template"
	"io"
	"log"
	"net/http"

	//"go-file-manager/handler/admin"
	//"go-file-manager/handler/patient"
	"go-file-manager/handler/assets"
	"go-file-manager/models"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// todo
func main() {

	// Create a new instance of Echo
	e := echo.New()

	flag.String("conf", "local", "--conf=dev or --conf=prod")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	confName := viper.GetString("conf")

	runPath := getRunPath()
	os.Mkdir(runPath+"/logs", os.ModePerm)
	config.NewConfig(runPath+"/config", confName)
	models.NewDB("dbdefault")
	crontab.Init()

	gzipLevel := config.Config.GetInt("system.gzipLevel")
	if gzipLevel > 0 {
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: gzipLevel,
		}))
	}

	assetsPath := config.Config.GetString("file.assetsPath")
	e.Static("/static", runPath+"/"+assetsPath)
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	//t := &Template{
	//    templates: template.Must(template.ParseGlob("views/*.html")),
	//}
	//e.Renderer = t

	accessLogPath := config.Config.GetString("file.accessLogPath")
	accessLogPath = runPath + "/" + accessLogPath
	accessLogFile, err := os.OpenFile(accessLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer accessLogFile.Close()
	if err != nil {
		log.Fatalln("open accessLogFile error !")
	}
	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: accessLogFile,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// e.File("/", "public/index.html")
	e.GET("/", handler.Index)
	e.GET("/api/assets", assets.Index)
	e.DELETE("/api/assets/:assetID", assets.Delete)

	port := config.Config.GetString("system.port")
	// Start as a web server

	runLogPath := config.Config.GetString("file.runLogPath")
	runLogPath = runPath + "/" + runLogPath
	runLogFile, err := os.OpenFile(runLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer runLogFile.Close()
	if err != nil {
		log.Fatalln("open runLogFile error !")
	}
	e.Logger.SetOutput(runLogFile)
	if err := e.Start(port); err != nil {
		log.Println(err)
	}
}

func getRunPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
