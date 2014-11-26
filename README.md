# Readme Generator

A CLI tool written in Go which generates a common README.md for your project.

## About

This tool will ask some basic informations about your project and in the end it will generate a README.md file.

## How to use

#### Instalation

``` bash
go get github.com/caio-ribeiro-pereira/goreadme
```

*Obs.:* After install it, the `goreadme` command will be available to use, so just run the command below answering all following questions. 

#### Running

``` bash
goreadme
```

if you wanna ignore some default template instructions, you can choose some of these flags:

``` bash
goreadme --ignore-setup, --igs # ignore setup's instructions
goreadme --ignore-dev, --igd   # ignore development's instructions
goreadme --ignore-test, --igt  # ignore test's instructions
```

## Author

Caio Ribeiro Pereira <caio.ribeiro.pereira@gmail.com>  
MIT License: http://caio-ribeiro-pereira.mit-license.org
