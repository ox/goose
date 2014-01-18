package main

import (
	"github.com/ox/goose/lib/goose"
	"log"
)

// the target version to migrate down to
var upTarget *int64

var upCmd = &Command{
	Name:    "up",
	Usage:   "[--target=]",
	Summary: "Migrate the DB to the most recent version available, or until a target version",
	Help:    `up extended help here...`,
	Run:     upRun,
}

func init() {
	upTarget = upCmd.Flag.Int64("target", -1, "the target migration to migrate to, default: latest")
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

	if *upTarget > recent || *upTarget < 0 {
		*upTarget = recent
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, *upTarget); err != nil {
		log.Fatal(err)
	}
}
