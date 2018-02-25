package main

import (
	"bufio"
	"fmt"
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

	return "", fmt.Errorf("Can't extract version from %s", str)
}

func main() {
	log.SetFlags(0)

	app := cli.NewApp()
	app.Name = "version"
	app.Usage = "CLI command to verify versions and version constraints."
	app.Version = "1.0.3"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ivan Diachenko",
			Email: "ivan.dyachenko@gmail.com",
		},
	}
	app.UsageText = "version [global options] constraints [version]"
	app.Action = func(c *cli.Context) error {
		var versionString string
		constrains := c.Args().First()
		if c.NArg() > 1 {
			versionString = strings.Join(c.Args().Tail(), "")
		} else {
			reader := bufio.NewReader(os.Stdin)
			versionString, _ = reader.ReadString('\n')
			versionString = strings.Trim(versionString, "\n ")
		}

		checkVersion, err := ExtractVersion(versionString)
		processError(err)

		ver, err := version.NewVersion(checkVersion)
		processError(err)

		constraints, err := version.NewConstraint(constrains)
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
