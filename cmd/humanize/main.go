package main

import (
	"flag"
	"fmt"

	humanhash "github.com/bibenga/humanhash-go"
	"github.com/google/uuid"
)

var (
	value  string
	format string
)

func init() {
	flag.StringVar(&value, "value", "", "value")
	flag.StringVar(&format, "format", "uuid", "the format of the value [uuid, text]")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if value == "" {
		_, humanized, err := humanhash.NewUuid()
		if err != nil {
			panic(err)
		}
		fmt.Println(humanized)
	} else {
		if format == "uuid" {
			u, err := uuid.Parse(value)
			if err != nil {
				panic(err)
			}
			humanized, err := humanhash.Humanize(u[:])
			if err != nil {
				panic(err)
			}
			fmt.Println(humanized)
		} else {
			humanized, err := humanhash.Humanize([]byte(value))
			if err != nil {
				panic(err)
			}
			fmt.Println(humanized)
		}
	}
}

func usage() {
	fmt.Println("Usage:")

	order := []string{"value", "format"}
	for _, name := range order {
		flagEntry := flag.CommandLine.Lookup(name)
		fmt.Printf("  -%s\n", flagEntry.Name)
		fmt.Printf("\t%s\n", flagEntry.Usage)
	}

	fmt.Println()
	fmt.Println(`Example commands:

	$ humanize -value=938AEF34-912E-4729-8C30-7DC347504688
	$ humanize -value=123456 -format=text
	`)
}
