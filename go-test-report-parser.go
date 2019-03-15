package main

import (
	"flag"
	"fmt"
	"github.com/jstemmer/go-junit-report/parser"
	"github.com/xiaosongluo/go-test-report-parser/formatter"
	_ "github.com/xiaosongluo/go-test-report-parser/formatter/junit"
	_ "github.com/xiaosongluo/go-test-report-parser/formatter/markdownFunction"
	"os"
)

var (
	packageName   string
	formatterName string
	goVersionFlag string
	setExitCode   bool
)

func init() {
	flag.StringVar(&packageName, "package-name", "", "specify a package name (compiled test have no package name in output)")
	flag.StringVar(&formatterName, "formatter-name", "JUnitFormatter", "specify a formatter name")
	flag.StringVar(&goVersionFlag, "go-version", "", "specify the value to use for the go.version property in the generated XML")
	flag.BoolVar(&setExitCode, "set-exit-code", false, "set exit code to 1 if tests failed")
}

func main() {
	flag.Parse()

	if flag.NArg() != 0 {
		fmt.Println("go-junit-report does not accept positional arguments")
		os.Exit(1)
	}

	// Read input
	report, err := parser.Parse(os.Stdin, packageName)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		os.Exit(1)
	}

	// Output
	output := formatter.GetAllFormatter()[formatterName]
	if output != nil {
		err = output.Formatter(report, os.Stdout)
		if err != nil {
			fmt.Printf("Error Output: %s\n", err)
			os.Exit(1)
		}
	}

	if setExitCode && report.Failures() > 0 {
		os.Exit(1)
	}
}
