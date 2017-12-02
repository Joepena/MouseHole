package grifts

import (
	"github.com/joepena/mouse_hole/models"
	"github.com/markbates/grift/grift"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

var _ = grift.Namespace("test", func() {

	grift.Desc("events", "creates a database with a capped events collection")
	grift.Add("events", func(c *grift.Context) error {
		models.GetDBInstance().CreateCappedCollection("test_application", "events", 1000000000)
		return nil
	})

	grift.Desc("auth_seed", "creates a database with a seeded test users")
	grift.Add("auth_seed", func(c *grift.Context) error {
		dB := models.GetDBInstance()
		users := []models.User{
			{
				ID:       "aaa.bbbb.ccc",
				Email:    "xxx@xxx.com",
				Password: "xxx",
				DBName:   "app1",
			},
			{
				ID:       "bbb.ddd.zzz",
				Email:    "yyy@xxx.com",
				Password: "xxx",
				DBName:   "app2",
			},
			{
				ID:       "vvv.rrrr.yyy",
				Email:    "zzz@xxx.com",
				Password: "xxx",
				DBName:   "app3",
			},
		}

		for _, user := range users {
			err := user.Create(dB)
			if err != nil {
				log.WithField("Function", "grift/auth_seed").Error(err)
			} else {
				log.WithField("Function", "grift/auth_seed").Infof("%v inserted successfully!", user)
			}
		}

		return nil
	})
})
