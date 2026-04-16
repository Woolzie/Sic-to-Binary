package utils

import(
	"flag"
)

func ParseCli(){
	flagSic := flag.Bool("sic", false, "start program to use sic format by default")
	flag.Parse()

	SicEnabled = *flagSic
}
