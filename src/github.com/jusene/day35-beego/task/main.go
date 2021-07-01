package main

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
	"time"
)

func main() {
	// CREATE A TASK
	tk1 := task.NewTask("tk1", "*/3 * * * * *", func(ctx context.Context) error {
		logs.Info("tk1")
		return nil
	})

	// CHECK TASK
	err := tk1.Run(context.Background())
	if err != nil {
		logs.Error(err)
	}

	// ADD TASK TO GLOBAL TODOLIST
	task.AddTask("tk1", tk1)

	// STAT TASKS
	task.StartTask()

	// wait
	time.Sleep(12 * time.Second)
	defer task.StopTask()

}
