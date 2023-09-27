//package main

package tasks

import (
	"swap/blockchain"
	"swap/db"
	"swap/services"
	"swap/tasks/service"
	"github.com/jasonlvhit/gocron"
	"time"
)

func Start() {
	db.InitMysql()

	services.InitMainNet()

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(30).Seconds().From(gocron.NextTick()).Do(service.PollBlockTask)
	_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(service.PollBlockTask)
	_ = s.Every(1).Day().From(gocron.NextTick()).From(gocron.NextTick()).Do(service.GetPioneerRechargeWeight)
	_ = s.Every(4).Hours().From(gocron.NextTick()).From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTask)
	<-s.Start()

}
