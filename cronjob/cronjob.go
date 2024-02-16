package cronjob

import (
	"api-openlive/blockchain"
	"api-openlive/utils"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

type CronJob struct {
	c     *cron.Cron
	sync  *blockchain.Synchronization
	mutex *sync.RWMutex
}

func NewCronJob(sync *blockchain.Synchronization, mutex *sync.RWMutex) *CronJob {
	cron := CronJob{
		c:     cron.New(),
		sync:  sync,
		mutex: mutex,
	}
	return &cron
}

func (c CronJob) StartCronJob() {
	c.getListCronJob()
	c.c.Start()
}

func (c CronJob) getListCronJob() {
	c.c.AddFunc("@every 0h01m00s", c.updateTransactionProcessing)
	c.c.AddFunc("@every 1h0m0s", c.updateConvertRate)
}

func (c CronJob) updateTransactionProcessing() {
	c.sync.StartSync(1)
}

func (c CronJob) updateConvertRate() {
	convertRate, err := utils.GetExchangeRates("tether", 1)
	if err != nil {
		log.Println("error get convert rate")
	}

	utils.SetRate(convertRate, c.mutex)
}
