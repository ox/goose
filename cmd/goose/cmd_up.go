package main

import (
	"github.com/ox/goose/lib/goose"
	"log"
)

// the target version to migrate down to
var upTo *int64

var upCmd = &Command{
	Name:    "up",
	Usage:   "[-to=]",
	Summary: "Migrate the DB to the most recent version available, or until a target version",
	Help:    ``,
	Run:     upRun,
}

func init() {
	upTo = upCmd.Flag.Int64("to", -1, "the target migration to migrate to, default: latest")
}

func upRun(cmd *Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	recent, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	if *upTo > recent || *upTo < 0 {
		*upTo = recent
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, *upTo); err != nil {
		log.Fatal(err)
	}
}
