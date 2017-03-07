package main

import (
	"flag"
)

type CLIParser struct {
	ServerMode bool
	IP         string
	Port       string
}

func parseCLI() {
	clip = &CLIParser{}

	flag.BoolVar(&clip.ServerMode, "s", false, "run in server (bank) mode")
	flag.StringVar(&clip.IP, "i", IP, "listen on this IP address")
	flag.StringVar(&clip.Port, "p", PORT, "listen on this port")

	flag.Parse()
}
