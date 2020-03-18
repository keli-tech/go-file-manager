package crontab

import "github.com/robfig/cron"

func Init() {
	//todo
	c := cron.New()
	//todo  调整为每天0点30分0秒  "0 30 0 * * *"
	//c.AddFunc("0 30 0 * * *", ChangeStatus)

	c.Start()
}
