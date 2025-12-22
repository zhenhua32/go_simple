package main

import (
	"flag"
	"fmt"
)

type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

func (c *Config) Setup() {
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it default is empty")
	flag.StringVar(&c.subject, "s", "", "subject is a string, it default is empty (shorthand)")
	flag.BoolVar(&c.isAwesome, "is_awesome", false, "is it awesome? default is false")
	flag.IntVar(&c.howAwesome, "how_awesome", 10, "how awesome out of 10?")
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is not awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d/10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
