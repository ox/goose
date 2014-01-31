package main

import (
 "bytes"
 "errors"
 "fmt"
 "log"
 "os/exec"
 "strings"
)

var createDBCmd = &Command{
 Name:    "create-db",
 Usage:   "",
 Summary: "Create the db for the given environement",
 Help:    ``,
 Run:     createDBRun,
}

func valueForOpenStr(openStr, key string) (string, error) {
 parts := strings.Split(openStr, " ")
 for _, option := range parts {
   if strings.HasPrefix(option, key+"=") {
     return strings.TrimPrefix(option, key+"="), nil
   }
 }
 return "", errors.New(fmt.Sprintf("could not find key %s", key))
}

func createDBRun(cmd *Command, args ...string) {
 conf, err := dbConfFromFlags()
 if err != nil {
   log.Fatal(err)
 }

 if conf.Driver.Name != "postgres" {
   log.Println("create-db is only available on postgres for now")
   return
 }

 dbName, err := valueForOpenStr(conf.Driver.OpenStr, "dbname")
 if err != nil {
   log.Fatalf("can't create unspecified dbname")
 }

 user, err := valueForOpenStr(conf.Driver.OpenStr, "user")
 if err != nil {
   log.Println("no user= specified, defaulting to current user")
 }

 execCmd := exec.Command("createdb", dbName)
 if user != "" {
   execCmd.Args = append(execCmd.Args, "-O", user)
 }

 var cmdOutput bytes.Buffer
 execCmd.Stderr = &cmdOutput
 err = execCmd.Run()
 if err != nil {
   log.Fatalf(cmdOutput.String())
 }
 log.Println("created", dbName)
}
