package main

import (
	"github.com/ox/goose/lib/goose"
	"log"
)

var upCmd = &Command{
	Name:    "up",
	Usage:   "[-to=]",
	Summary: "Migrate the DB to the most recent version available, or until a target version",
	Help:    ``,
	Run:     upRun,
}

func init() {
	upCmd.Flag.Int64("to", -1, "the migration to migrate to, default: latest")
}

func upRun(cmd *Command, args ...string) {
	to := cmd.GetFlagValue("to").(int64)

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	recent, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	if to > recent || to < 0 {
		to = recent
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, to); err != nil {
		log.Fatal(err)
	}
}
