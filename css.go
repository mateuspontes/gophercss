package css

import (
	"fmt"
	"io"
	"os"
)

var defaultBuffer io.Writer = os.Stdout

type Style map[string]interface{}

type subStyle struct {
	property string
	style    Style
}

func Set(property string, style Style) {
	path := []string{property}
	printProperty(path, style)
}

func printProperty(path []string, style Style) {
	for _, item := range path {
		fmt.Fprint(defaultBuffer, item+" ")
	}

	fmt.Fprint(defaultBuffer, "{ ")

	subStyles := []subStyle{}

	for subProperty, value := range style {
		switch v := value.(type) {
		case Style:
			subStyles = append(subStyles, subStyle{subProperty, v})
			continue

		case string:
			fmt.Fprintf(defaultBuffer, "%s: %q; ", subProperty, v)

		default:
			fmt.Fprintf(defaultBuffer, "%s: %v; ", subProperty, v)
		}
	}

	fmt.Fprint(defaultBuffer, "}\n")

	for _, subStyle := range subStyles {
		subPath := path[:]
		subPath = append(subPath, subStyle.property)
		printProperty(subPath, subStyle.style)
	}
}

func setOutput(output io.Writer) {
	defaultBuffer = output
}
