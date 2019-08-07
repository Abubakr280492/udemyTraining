package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Action = func(c *cli.Context) error {
    fmt.Println("boom! I say!")
    return nil
  }
<<<<<<< HEAD
ls
=======

>>>>>>> 466b5f809418bf709baabb8ddcc83b362ff4c2e9
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
