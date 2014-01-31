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
	upCmd.Flag.Int64("to", -1,
		"the migration to migrate to, default: latest")
	upCmd.Flag.Int64("target", -1,
		"the id of the migration that you want to run")
	upCmd.Flag.Bool("force", false,
		"the migration will be reported as a success, even if it fails")
}

func upRun(cmd *Command, args ...string) {
	to := cmd.GetFlagValue("to").(int64)
	target := cmd.GetFlagValue("target").(int64)
	force := cmd.GetFlagValue("force").(bool)

	log.Println(to, target, force)

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
