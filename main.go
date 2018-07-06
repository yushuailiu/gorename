package main

import (
	"github.com/urfave/cli"
	"go/token"
	"go/parser"
	"fmt"
	"golang.org/x/tools/go/ast/astutil"
	"strings"
	"bytes"
	"go/printer"
	"io/ioutil"
	"os"
	"path/filepath"
	"path"
	"github.com/fatih/color"
)

func main()  {
	app := cli.NewApp()
	app.Name = "gorename"
	app.Usage = "Rename golang package"
	app.Version = "0.0.1"
	app.ArgsUsage = "[source file or directory path] [old package name] [new package name]"
	app.Author = "YuShuai Liu <admin@liuyushuai.com>"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "source, s",
			Value: "./",
			Usage: "source package path or file path",
		},
	}

	app.Action = func(c *cli.Context) {
		source := c.String("source")
		from := c.Args().Get(0)
		to := c.Args().Get(1)

		fileInfo, err := os.Stat(source)
		if err != nil {
			cli.NewExitError("source is not a directory or file", -1)
			return
		}

		if from == "" || to == "" {
			cli.ShowAppHelp(c)
			return
		}

		fmt.Println(color.GreenString("[INFO]"), "start update import ", from, " to ", to)

		var exitError *cli.ExitError

		if !fileInfo.IsDir() {
			exitError = ProcessFile(source, from ,to)
		} else {
			exitError = ProcessDir(source, from, to , c)
		}
		if exitError != nil {
			fmt.Println(color.YellowString("[WARN]"), exitError.Error())
		} else {
			fmt.Println(color.GreenString("[INFO] success!"))
		}
	}
	app.Run(os.Args)
}

func ProcessDir(dir string, from string, to string, c *cli.Context) *cli.ExitError {

	answer := ""
	absDir,_ := filepath.Abs(dir)
	fmt.Print(color.YellowString("[WARNING] "), "Rename the package ", color.YellowString(from),
		" to ", color.YellowString(to), " which files in directory of ",
			color.YellowString(absDir), " ? Yes (Y) or No (N):")
	fmt.Scanln(&answer)


	if answer != "Yes" && answer != "Y" {
		return cli.NewExitError("Do nothing", 0)
	}

	filepath.Walk(dir, func(filepath string, info os.FileInfo, err error) error {
		if path.Ext(filepath) == ".go" {
			ProcessFile(filepath, from ,to)
		}
		return nil
	})
	return nil
}

func ProcessFile(filePath string, from string, to string) *cli.ExitError {
	fSet := token.NewFileSet()

	file, err := parser.ParseFile(fSet, filePath, nil, 0)

	if err != nil {
		fmt.Println(err)
	}

	imports := astutil.Imports(fSet, file)

	changeNum := 0


	for _, tempPackage := range imports {
		for _, mImport := range tempPackage {
			importString := strings.TrimSuffix(strings.TrimPrefix(mImport.Path.Value, "\""), "\"")

			if strings.HasPrefix(importString, from) {
				changeNum ++

				replacePackage := strings.Replace(importString, from , to, -1)


				if mImport.Name != nil && len(mImport.Name.Name) > 0 {
					astutil.DeleteNamedImport(fSet, file, mImport.Name.Name, importString)
					astutil.AddNamedImport(fSet, file, mImport.Name.Name, replacePackage)
				} else {
					astutil.DeleteImport(fSet, file, importString)
					astutil.AddImport(fSet, file, replacePackage)
				}
			}
			
		}
	}

	if changeNum > 0 {
		var outputBuffer bytes.Buffer

		printer.Fprint(&outputBuffer, fSet, file)

		ioutil.WriteFile(filePath, outputBuffer.Bytes(), os.ModePerm)

		change := "change"

		if changeNum > 1 {
			change += "s"
		}

		fmt.Println(color.GreenString("[INFO]"), changeNum, change, "in file", filePath)
	}

	return nil
}


