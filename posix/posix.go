package posix

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/shiftinecmd/makepizza/common"
)

func ParseFromPosix(raw []string) (contents string, args common.Arguments) {
	// This command assumes that the program name (os.Args[0]) has been emitted from `raw`.
	parsed := map[string]string{}
	vLongArgs, vShortArgs, mapping := common.GetValidArgs(common.Posix)
	var currentKey *string = nil

	// STAGE 1: Classify and validate raw flags and values into the `parsed` variable.
	for _, val := range raw {
		if strings.HasPrefix(val, "--") {
			// POSIX long form
			if !slices.Contains(vLongArgs, val[2:]) {
				log.Fatalf("Invalid POSIX long parameter: %s", val)
			}

			sanitized := val[2:]
			currentKey = &sanitized

			if _, ok := parsed[sanitized]; ok {
				log.Fatalf("POSIX long parameter %s declared multiple times", val)
			} else {
				parsed[sanitized] = ""
			}
		} else if strings.HasPrefix(val, "-") {
			// POSIX short form
			sanitized := val[1:]

			for _, flag := range strings.Split(sanitized, "") {
				if !slices.Contains(vShortArgs, flag) {
					log.Fatalf("Invalid POSIX short parameter: -%s", flag)
				}
				currentKey = &flag
				if _, ok := parsed[flag]; ok {
					log.Fatalf("POSIX short parameter -%s declared multiple times", flag)
				} else {
					parsed[flag] = ""
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

func PosixHelp() {
	fmt.Println("FIXME: Add help page")
}

func PosixVersion() {
	fmt.Println("`makepizza` for POSIX/UNIX conventions, version 1.0.0")
	common.PrintLicense()
}

func main() {
	rawArgs := os.Args[1:]

	if len(rawArgs) == 0 {
		log.Fatal("No arguments or instructions supplied. Use `makepizza -h` or `makepizza --help` for more info.")
		os.Exit(1)
	}

	contents, args := ParseFromPosix(rawArgs)

	// Prioritize display help or version before making
	if args.Help != nil && *args.Help {
		PosixHelp()
		os.Exit(0)
	} else if args.Version != nil && *args.Version {
		PosixVersion()
		os.Exit(0)
	} else {
		os.Exit(int(common.MakePizza(contents, args)))
	}
}
