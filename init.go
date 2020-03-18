package main

import (
	"database/sql"
	"flag"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-file-manager/config"
	"go-file-manager/models"
	"log"
)

func main() {

	flag.String("conf", "dev", "--conf=dev or --conf=prod")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	confName := viper.GetString("conf")

	config.NewConfig("config", confName)
	models.NewDB("dbdefault")

	//dbname := "dbdefault"
	//dbHost := config.Config.GetString(dbname + ".dbHost")
	//dbName := config.Config.GetString(dbname + ".dbName")
	//dbUser := config.Config.GetString(dbname + ".dbUser")
	//dbPasswd := config.Config.GetString(dbname + ".dbPasswd")
	//dbPort := config.Config.GetString(dbname + ".dbPort")
	//dbType := config.Config.GetString(dbname + ".dbType")
	//connectString := dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	//db, err := gorm.Open(dbType, connectString)
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		log.Print(err)
	} else {
		if false {
			db.DropTable(&models.Assets{})
		}
		//db.CreateTable(&models.Assets{})

		assets := models.Assets{
			Name:     "东风破.mp3",
			Path:     "/叶惠美",
			FullPath: "/叶惠美/东风破.mp3",
			Type:     "mp3",
			Size:     "",
			Status: sql.NullBool{
				false,
				true,
			},
		}
		assets.Create(&assets)

		//newAdmin("admin001", "admin@ecrf.com", "admin001", "上海1")
		//newAdmin("test001", "test001@ecrf.com", "test001", "上海2")
	}
}
