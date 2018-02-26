package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"

	version "github.com/hashicorp/go-version"
)

var (
	booleanMode bool
)

func processError(err error) {
	if err != nil {
		if booleanMode {
			fmt.Print("false")
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
}

// ExtractVersion parse string and extract version number
func ExtractVersion(str string) (string, error) {
	patterns := [...]string{"version (\\d+.\\d+.d\\+)", "version (\\d[\\d\\.]*)", "(\\d[\\d\\.]*)"}
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(str)
		if len(match) != 0 {
			return match[1], nil
		}
	}

	return "", fmt.Errorf("Can't extract version from \"%s\"", str)
}

func readFromPipe() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return "", fmt.Errorf("The command is intended to work with pipes\nUsage: go version | version -b \">=1.9\"")
	}

	reader := bufio.NewReader(os.Stdin)
	var result []string

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		result = append(result, fmt.Sprintf("%c", input))
	}

	return strings.Join(result, ""), nil
}

func main() {
	log.SetFlags(0)

	app := cli.NewApp()
	app.Name = "version"
	app.Usage = "CLI command to verify versions and version constraints."
	app.Version = "1.0.4"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ivan Diachenko",
			Email: "ivan.dyachenko@gmail.com",
		},
	}
	app.UsageText = "version [global options] constraints [version]"
	app.Commands = []cli.Command{
		{
			Name:    "parse",
			Aliases: []string{"p"},
			Usage:   "parse version from arguments or pipe",
			Action: func(c *cli.Context) error {
				var versionString string
				var err error

				if c.NArg() > 0 {
					versionString = strings.Join(c.Args(), "")
				} else {
					versionString, err = readFromPipe()
					processError(err)
				}

				parsedVersion, err := ExtractVersion(versionString)
				processError(err)

				fmt.Println(parsedVersion)
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		var versionString string
		var err error

		if c.NArg() > 1 {
			versionString = strings.Join(c.Args().Tail(), "")
		} else {
			versionString, err = readFromPipe()
			processError(err)
		}

		checkVersion, err := ExtractVersion(versionString)
		processError(err)

		ver, err := version.NewVersion(checkVersion)
		processError(err)

		constraints, err := version.NewConstraint(c.Args().First())
		processError(err)

		if constraints.Check(ver) {
			if !booleanMode {
				log.Printf("%s satisfies constraints %s", ver, constraints)
			} else {
				fmt.Print("true")
			}
		} else {
			processError(fmt.Errorf("%s doesn't satisfies constraints %s", ver, constraints))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "boolean, b",
			Usage:       "boolean mode return '1' or '0' and always exit with 0",
			Destination: &booleanMode,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
