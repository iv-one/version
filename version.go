package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	version "github.com/hashicorp/go-version"
)

func extractVersion(str string) (string, error) {
	re := regexp.MustCompile("(\\d+.[\\d\\.]+)")
	match := re.FindStringSubmatch(str)
	if len(match) == 0 {
		return "", fmt.Errorf("Can't extract version from %s", str)
	}
	return match[0], nil
}

func main() {
	constrains := strings.Join(os.Args[1:], "")
	reader := bufio.NewReader(os.Stdin)
	ver, _ := reader.ReadString('\n')

	checkVersion, err := extractVersion(ver)
	if err != nil {
		panic(err)
	}

	v1, err := version.NewVersion(checkVersion)
	if err != nil {
		panic(err)
	}

	constraints, err := version.NewConstraint(constrains)
	if err != nil {
		panic(err)
	}

	if constraints.Check(v1) {
		fmt.Printf("%s satisfies constraints %s", v1, constraints)
	} else {
		panic(fmt.Errorf("%s doesn't satisfies constraints %s", v1, constraints))
	}
}
