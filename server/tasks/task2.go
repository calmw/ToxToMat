// //package main
package tasks

//
////
//import (
//	"swap/db"
//	"swap/services"
//	"swap/tasks/service"
//	"github.com/jasonlvhit/gocron"
//	"time"
//)
//
//func Start() {
//	db.InitMysql()
//
//	services.InitMainNet()
//
//	s := gocron.NewScheduler()
//	s.ChangeLoc(time.UTC)
//	//_ = s.Every(30).Seconds().From(gocron.NextTick()).Do(service.PollBlockTask)
//	_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(service.PollBlockTask)
//	_ = s.Every(1).Day().From(gocron.NextTick()).Do(service.GetPioneerRechargeWeight)
//	<-s.Start()
//
//}
