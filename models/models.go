package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"time"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-file-manager/config"
)

var Gorm map[string]*gorm.DB

func init() {
	Gorm = make(map[string]*gorm.DB)
}

// 初始化Gorm
func NewDB(dbname string) {

	var orm *gorm.DB
	var err error

	//默认配置
	//config.Config.SetDefault(dbname, map[string]interface{}{
	//	"dbHost":          "172.17.20.149",
	//	"dbName":          "ecrf",
	//	"dbUser":          "root",
	//	"dbPasswd":        "rcroot",
	//	"dbPort":          3307,
	//	"dbIdleconns_max": 0,
	//	"dbOpenconns_max": 20,
	//	"dbType":          "mysql",
	//})
	//dbHost := config.Config.GetString(dbname + ".dbHost")
	//dbName := config.Config.GetString(dbname + ".dbName")
	//dbUser := config.Config.GetString(dbname + ".dbUser")
	//dbPasswd := config.Config.GetString(dbname + ".dbPasswd")
	//dbPort := config.Config.GetString(dbname + ".dbPort")
	//dbType := config.Config.GetString(dbname + ".dbType")
	//
	//connectString := dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	//开启sql调试模式
	//GDB.LogMode(true)
	//for orm, err = gorm.Open(dbType, connectString); err != nil; {
	//	fmt.Println("数据库连接异常! 5秒重试")
	//	time.Sleep(5 * time.Second)
	//	orm, err = gorm.Open(dbType, connectString)
	//}

	for orm, err = gorm.Open("sqlite3", "/tmp/gorm.db"); err != nil; {
		fmt.Println("数据库连接异常! 5秒重试")
		time.Sleep(5 * time.Second)
		orm, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	}

	//for orm, err = gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword"); err != nil; {
	//	fmt.Println("数据库连接异常! 5秒重试")
	//	time.Sleep(5 * time.Second)
	//	orm, err = gorm.Open(dbType, connectString)
	//}

	orm.SingularTable(true)

	//连接池的空闲数大小
	orm.DB().SetMaxIdleConns(config.Config.GetInt(dbname + ".idleconns_max"))
	//最大打开连接数
	orm.DB().SetMaxIdleConns(config.Config.GetInt(dbname + ".openconns_max"))

	logDbMode := config.Config.GetString("system.logDbMode")
	orm.LogMode(logDbMode == "true")

	Gorm[dbname] = orm
	//defer Gorm[dbname].Close()
}

// 通过名称获取Gorm实例
func GetORMByName(dbname string) *gorm.DB {

	return Gorm[dbname]
}

// 获取默认的Gorm实例
func GetORM() *gorm.DB {

	return Gorm["dbdefault"]
}

type Base struct {
}

func (base *Base) Create(data interface{}) *gorm.DB {
	db := GetORM()
	return db.Create(data)
}

func (base *Base) Save(data interface{}) *gorm.DB {
	db := GetORM()
	return db.Save(data)
}

func (base *Base) Delete(data interface{}) *gorm.DB {
	db := GetORM()
	return db.Delete(data)
}

//
//func (cb *Base) Bind(i interface{}, c echo.Context) (err error) {
//	// 你也许会用到默认的绑定器
//	db := new(echo.DefaultBinder)
//	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
//		log.Println(22)
//		return
//	}
//
//	log.Println(11)
//	// 做你自己的实现
//
//	return
//}
