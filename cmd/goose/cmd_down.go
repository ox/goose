package main

import (
	"github.com/ox/goose/lib/goose"
	"log"
)

var downTo *int64

var downCmd = &Command{
	Name:    "down",
	Usage:   "[-to=]",
	Summary: "Roll back the version by 1, or to a target migration",
	Help:    ``,
	Run:     downRun,
}

func init() {
	downTo = downCmd.Flag.Int64("to", -1, "the target migration to roll back to")
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

	if *downTo < 0 || *downTo >= current {
		*downTo, err = goose.GetPreviousDBVersion(conf.MigrationsDir, current)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = goose.RunMigrations(conf, conf.MigrationsDir, *downTo); err != nil {
		log.Fatal(err)
	}
}
