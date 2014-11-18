package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "readme"
  app.Usage = "A CLI tool that generates a minimal README for your projects."
  app.Commands = []cli.Command{
    {
      Name: "init",
      ShortName: "i",
      Usage: "Generates README.md based on some questions you will answer",
      Action: func(c *cli.Context) {
        fmt.Println("Project's name:")
        fmt.Println("Project's description:")
        fmt.Println("Project's site:")
        fmt.Println("Author's name:")
        fmt.Println("Author's e-mail:")
        fmt.Println("This project has instalation's instruction? (yes)")
        fmt.Println("This project has development's instruction? (yes)")
        fmt.Println("This project has test's instruction? (yes)")
        fmt.Println("License's name:")
      },
    },
  }
  app.Run(os.Args)
}