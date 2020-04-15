package task

import (
	"feedProject/pkg/service"
	"fmt"
	"github.com/robfig/cron"
)

func Start()  {
	task := cron.New()

	_ = task.AddFunc("*/1 * * * * ?", func() {
		fmt.Println("SyncQueue:%v", service.Feed().SyncQueue())
	})

	task.Run()
}
