package main

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

func (s *httpServer) Cron() {
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", func() {
		err := s.store.Users.Update_UserLimitReset(context.Background())
		if err != nil {
			s.logger.Error(err.Error())
		}
	})

	if err != nil {
		s.logger.Error(err.Error())
	}

	c.Start()
	fmt.Printf("[INFO] Cron is started...\n")

	select {}
}
