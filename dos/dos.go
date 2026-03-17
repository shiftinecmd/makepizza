package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/shiftinecmd/makepizza/common"
)

func ParseFromDos(raw []string) (contents string, args common.Arguments) {
	// This command assumes that the program name (os.Args[0]) has been emitted from `raw`.
	parsed := map[string]string{}
	_, vArgs, mapping := common.GetValidArgs(common.Dos)
	var currentKey *string = nil

	// STAGE 1: Classify and validate raw flags and values into the `parsed` variable.
	for _, val := range raw {
		if strings.HasPrefix(val, "/") {
			sanitized := ""
			draft := ""
			if strings.Contains(val, "=") {
				// Separate key and value, if necessary
				splits := strings.SplitN(val, "=", 1)
				sanitized = strings.ToLower(splits[0][1:])
				draft = splits[1]
			} else if strings.Contains(val, ":") {
				// Separate key and value, if necessary
				splits := strings.SplitN(val, ":", 1)
				sanitized = strings.ToLower(splits[0][1:])
				draft = splits[1]
			} else {
				sanitized = val[1:]
			}

			if !slices.Contains(vArgs, sanitized) {
				log.Fatalf("Invalid DOS parameter: /%s", sanitized)
			}
			currentKey = &sanitized

			if _, ok := parsed[sanitized]; ok {
				log.Fatalf("DOS parameter /%s declared multiple times", val)
			} else {
				parsed[sanitized] = draft
			}
		} else {
			if currentKey != nil && len(contents) > 0 {
				log.Fatalf("The DOS pizza content cannot be split in the middle of parameters")
			}

			currentKey = nil
			if len(contents) == 0 {
				contents = val
			} else {
				contents += " " + val
			}
		}
	}

	// STAGE 2: Verify that the given value data types are correct.
	args = common.ParseArguments(parsed, mapping)

	return contents, args
}

func DosHelp() {
	fmt.Println("FIXME: Add help page")
}

func DosVersion() {
	fmt.Println("`mkpizza` for DOS/Batch conventions, version 1.0.0")
	common.PrintLicense()
}

func main() {
	rawArgs := os.Args[1:]

	if len(rawArgs) == 0 {
		log.Fatal("No arguments or instructions supplied. Use `mkpizza /?` for more info.")
		os.Exit(1)
	}

	contents, args := ParseFromDos(rawArgs)

	// Prioritize display help or version before making
	if args.Help != nil && *args.Help {
		DosHelp()
		os.Exit(0)
	} else if args.Version != nil && *args.Version {
		DosVersion()
		os.Exit(0)
	} else {
		os.Exit(int(common.MakePizza(contents, args)))
	}
}
