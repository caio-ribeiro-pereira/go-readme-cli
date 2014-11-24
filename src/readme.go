package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "github.com/codegangsta/cli"
)

func isConfirmed(value string) bool {
  return value == "y" || value == "Y" || value == "yes"
}

func check(err error) {
  if err != nil {
    log.Fatal(err)
    panic(err)
  }
}

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
        in := bufio.NewReader(os.Stdin)
        fmt.Println("Project's name:")
        prjName, _ := in.ReadString('\n')
        fmt.Println("Project's short description:")
        prjDesc, _ := in.ReadString('\n')
        fmt.Println("Author's name:")
        authorName, _ := in.ReadString('\n')
        fmt.Println("Author's e-mail:")
        authorEmail, _ := in.ReadString('\n')
        fmt.Println("This project has setup's instruction? (yes)")
        prjSetup, _ := in.ReadString('\n')
        fmt.Println("This project has development's instruction? (yes)")
        prjDev, _ := in.ReadString('\n')
        fmt.Println("This project has test's instruction? (yes)")
        prjTest, _ := in.ReadString('\n')
        fmt.Println("License's name:")
        prjLicense, _ := in.ReadString('\n')
        file, err := os.Create("README.md")
        check(err)
        out := bufio.NewWriter(file)
        out.WriteString("# "+ prjName +"\n\n")
        out.WriteString(prjDesc +"\n\n")
        out.WriteString("## About\n\nWrite here about what is your project...\n\n")
        if isConfirmed(prjSetup) {
          out.WriteString("## Instalation\n\nWrite here about how install your project  \n`write here some code if is necessary!`\n\n")
        }
        if isConfirmed(prjDev) {
          out.WriteString("## Development\n\nWrite here about how setup your project environment for development purposes  \n`write here some code if is necessary!`\n\n")
        }
        if isConfirmed(prjTest) {
          out.WriteString("## Running test\n\nWrite here about how execute test tasks into your project  \n`write here some code if is necessary!`\n\n")
        }
        out.WriteString("## Author\n\n"+ authorName +" <"+ authorEmail +">  \nLicensed by "+ prjLicense)
        out.Flush()
      },
    },
  }
  app.Run(os.Args)
}