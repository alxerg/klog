package commands

import (
	"klog/cli"
	"klog/datetime"
	"time"
)

func Start(env cli.Environment, args []string) int {
	now := time.Now()
	today, _ := datetime.CreateDateFromTime(now)
	wd, err := env.Store.Get(today)
	if err != nil {
		// todo create new
		return 182763
	}
	nowTime, _ := datetime.CreateTimeFromTime(now)
	wd.AddOpenRange(nowTime)
	env.Store.Save(wd)
	return cli.OK
}
