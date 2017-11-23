package main

import "github.com/urfave/cli"

func checkIT(c *cli.Context) {
	check(c.Args()[0])
}

func check(s string) bool {
	return true
}
