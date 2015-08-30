package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

type LICENSE struct {
	Name     string
	Year     int
	Format   string
	Resource string
}

func NewLICENSE(name string, year int, format string) (*LICENSE, error) {
	format = strings.ToLower(format)
	resource := "template/" + format + ".tmpl"

	formats := []string{
		"mit",
		"bsd2",
		"bsd3",
		"gpl",
		"lgpl",
		"asl",
		"cpl",
	}

	for idx, f := range formats {
		// LICENSEのフォーマットが見つかったら抜ける
		if f == format {
			break
		}

		// 見つからずに最後まで到達したらエラーを返す
		if idx == len(formats)-1 {
			err := fmt.Errorf("Unknown license format.")

			return nil, err
		}
	}

	license := &LICENSE{Name: name, Year: year, Format: format, Resource: resource}
	return license, nil
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		n string
		f string
		y int
		o string
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&n, "n", "Name", "name")
	flags.StringVar(&f, "f", "", "format")
	flags.IntVar(&y, "y", time.Now().Year(), "year")
	flags.StringVar(&o, "o", "STDOUT", "output")
	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// ライセンス形式の指定は必須
	if f == "" {
		fmt.Println("Please set license format. Ex. 'license -f gpl'")
		return ExitCodeError
	}
	// バージョン表示
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	license, err := NewLICENSE(n, y, f)
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	// テンプレート読み込み
	tmpl, err := Asset(license.Resource)
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	// byte => string に変換する
	tmplStr := string(tmpl[:])

	t := template.New("t")
	template.Must(t.Parse(tmplStr))

	// ファイルの書き込み
	if o == "STDOUT" {
		t.Execute(os.Stdout, license)
		fmt.Println("where")
	} else {
		file, err := os.Create(o)
		if err != nil {
			fmt.Println("Can't create a new file.")
			return ExitCodeError
		}
		t.Execute(file, license)
	}

	return ExitCodeOK
}
