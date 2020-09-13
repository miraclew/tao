package main

import (
	"fmt"
	"github.com/miraclew/tao/tools/tao/generator"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	e, err := generator.NewEngine()
	if err != nil {
		fmt.Printf("create engine error: %s\n", err)
		return
	}
	exists, err := e.Workspace.TemplatesExists()
	if err != nil {
		fmt.Printf("check template dir error: %s\n", err)
	}
	if !exists {
		err = e.UpdateTemplates()
		if err != nil {
			fmt.Printf("update templates error: %s\n", err)
		}
	}

	app := cli.NewApp()
	app.Name = "Tao"
	app.Usage = "Tao 工具集"
	app.Version = "0.1.0"
	//app.Action =
	app.Commands = []*cli.Command{
		{
			Name:  "tpl",
			Usage: "获取/更新模板文件",
			Action: func(context *cli.Context) error {
				return e.UpdateTemplates()
			},
			SkipFlagParsing: true,
		},
		{
			Name:  "proto",
			Usage: "创建新Proto文件",
			Action: func(context *cli.Context) error {
				return e.GenerateProto()
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "资源名称",
					Required: true,
				},
				&cli.StringFlag{
					Name:    "out",
					Aliases: []string{"o"},
					Usage:   "文件输出路径",
				},
			},
		},
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "创建API文件",
			Action: func(context *cli.Context) error {
				return e.GenerateAPI(context.Args().First())
			},
			SkipFlagParsing: true,
		},
		{
			Name:    "svc",
			Aliases: []string{"s"},
			Usage:   "创建service文件",
			Action: func(context *cli.Context) error {
				return e.GenerateService(context.Args().First(), context.Bool("default"))
			},
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "default",
					Aliases: []string{"d"},
					Usage:   "使用默认实现模板",
				},
			},
		},
		{
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "创建Repo文件",
			Action: func(context *cli.Context) error {
				return e.GenerateRepo()
			},
			SkipFlagParsing: true,
		},
		{
			Name:    "locator",
			Aliases: []string{"l"},
			Usage:   "创建locator文件",
			Action: func(context *cli.Context) error {
				return e.GenerateLocator()
			},
			SkipFlagParsing: true,
		},
		{
			Name:    "doc",
			Aliases: []string{"d"},
			Usage:   "创建文档文件",
			Action: func(context *cli.Context) error {
				return e.GenerateOpenAPIV3()
			},
			SkipFlagParsing: true,
		},
		{
			Name:    "sql",
			Aliases: []string{"d"},
			Usage:   "创建SQL schema文件",
			Action: func(context *cli.Context) error {
				return e.GenerateSql()
			},
			SkipFlagParsing: true,
		},
		{
			Name:  "kotlin",
			Usage: "创建Kotlin客户端",
			Action: func(context *cli.Context) error {
				return e.GenerateKotlin(context.Args().First())
			},
			SkipFlagParsing: true,
		},
		{
			Name:  "dart",
			Usage: "创建Dart客户端",
			Action: func(context *cli.Context) error {
				return e.GenerateDart()
			},
			SkipFlagParsing: true,
		},
		{
			Name:  "swift",
			Usage: "创建Swift客户端",
			Action: func(context *cli.Context) error {
				return e.GenerateSwift(context.Args().First())
			},
			SkipFlagParsing: true,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
