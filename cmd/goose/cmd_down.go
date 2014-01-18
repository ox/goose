package main

import (
	"github.com/ox/goose/lib/goose"
	"log"
)

var downTarget *int64

var downCmd = &Command{
	Name:    "down",
	Usage:   "[--target=]",
	Summary: "Roll back the version by 1, or to a target migration",
	Help:    `down extended help here...`,
	Run:     downRun,
}

func init() {
	downTarget = downCmd.Flag.Int64("target", -1, "the target migration to roll back to")
}

func downRun(cmd *Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	current, err := goose.GetDBVersion(conf)
	if err != nil {
		log.Fatal(err)
	}

	if *downTarget < 0 || *downTarget >= current {
		*downTarget, err = goose.GetPreviousDBVersion(conf.MigrationsDir, current)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = goose.RunMigrations(conf, conf.MigrationsDir, *downTarget); err != nil {
		log.Fatal(err)
	}
}
