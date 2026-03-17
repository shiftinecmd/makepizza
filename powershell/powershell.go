package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/shiftinecmd/makepizza/common"
)

func ParseFromPowerShell(raw []string) (contents string, args common.Arguments) {
	// This command assumes that the program name (os.Args[0]) has been emitted from `raw`.
	parsed := map[string]string{}
	vLongArgs, vShortArgs, mapping := common.GetValidArgs(common.PowerShell)
	vLongArgsLower := []string{}
	var currentKey *string = nil

	for _, val := range vLongArgs {
		vLongArgsLower = append(vLongArgsLower, strings.ToLower(val))
	}

	// STAGE 1: Classify and validate raw flags and values into the `parsed` variable.
	for _, val := range raw {
		if strings.HasPrefix(val, "-") {
			// PowerShell short or long form
			sanitized := val[1:]

			for _, flag := range strings.Split(sanitized, "") {
				flagLower := strings.ToLower(flag)

				if !slices.Contains(vShortArgs, flagLower) && !slices.Contains(vLongArgsLower, flagLower) {
					log.Fatalf("Invalid PowerShell parameter: -%s", flag)
				}
				currentKey = &flagLower
				if _, ok := parsed[flagLower]; ok {
					log.Fatalf("PowerShell parameter -%s declared multiple times", flag)
				} else {
					parsed[flagLower] = ""
				}
			}
		} else if currentKey == nil {
			if len(contents) == 0 {
				contents = val
			} else {
				contents += " " + val
			}
		} else {
			if _, ok := parsed[*currentKey]; ok {
				if len(parsed[*currentKey]) == 0 {
					parsed[*currentKey] = val
				} else {
					parsed[*currentKey] += " " + val
				}
			} else {
				log.Fatalf("INTERNAL ERROR: currentKey %s points to non-existent map value", *currentKey)
			}
		}
	}

	// STAGE 2: Verify that the given value data types are correct.
	args = common.ParseArguments(parsed, mapping)

	return contents, args
}

func PowerShellHelp() {
	fmt.Println("FIXME: Add help page")
}

func PowerShellVersion() {
	fmt.Println("`Make-Pizza` for PowerShell conventions, version 1.0.0")
	common.PrintLicense()
}

func main() {
	rawArgs := os.Args[1:]

	if len(rawArgs) == 0 {
		log.Fatal("No arguments or instructions supplied. Use `Make-Pizza -h` or `Make-Pizza -Help` for more info.")
		os.Exit(1)
	}

	contents, args := ParseFromPowerShell(rawArgs)

	// Prioritize display help or version before making
	if args.Help != nil && *args.Help {
		PowerShellHelp()
		os.Exit(0)
	} else if args.Version != nil && *args.Version {
		PowerShellVersion()
		os.Exit(0)
	} else {
		os.Exit(int(common.MakePizza(contents, args)))
	}
}
