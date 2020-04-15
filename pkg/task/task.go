package task

import (
	"feedProject/pkg/service"
	"fmt"
	"github.com/robfig/cron"
)

func Start()  {
	task := cron.New()

	_ = task.AddFunc("*/10 * * * * ?", func() {
		fmt.Printf("SyncQueue:%v\n", service.Feed().SyncQueue())
	})

	task.Run()
}
