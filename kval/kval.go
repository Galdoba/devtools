package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/devtools/directory"
	"github.com/Galdoba/devtools/keyval"
	"github.com/urfave/cli/v2"
)

const (
	program = "kval"
)

func main() {
	warnings := []string{}
	lists := []string{}
	loc := ""
	//dv := []string{"buffer"}
	app := cli.NewApp()
	app.Version = "v 0.0.1"
	app.Name = program
	app.Usage = "TODO: USAGE"
	app.Description = "Manager for collections of key-val pairs"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "verbose",
		},

		//TODO: обдумать замену флагов книга\глава\страница на флаг локация
		//логика: программа редактирует одну страницу за раз (возможна редакция нескольких ключей на одной странице)
	}
	//ДО НАЧАЛА ДЕЙСТВИЯ
	app.Before = func(c *cli.Context) error {
		loc = c.String("from")

		//проверяем аргументы

		return nil
	}

	//ПО ОКОНЧАНИЮ ДЕЙСТВИЯ
	app.After = func(c *cli.Context) error {
		for _, warn := range warnings {
			fmt.Println("warning:", warn)
		}
		return nil
	}
	app.Commands = []*cli.Command{
		{
			Name:        "newlist",
			Usage:       "kval newlist [list1]...",
			UsageText:   "create new list of k-v pairs",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "args as names if new lists",
			Category:    "Create",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "force_new",
					Aliases: []string{"fn"},
				},
			},
			Action: func(c *cli.Context) error {
				args := c.Args()
				if len(args.Slice()) == 0 {
					return fmt.Errorf("no arguments provided")
				}
				msg := ""
				created := 0
				for _, arg := range args.Slice() {
					present, err := keyval.Present(arg)
					if err != nil {
						warnings = append(warnings, err.Error())
						continue
					}
					if present {
						if !c.Bool("fn") {
							warnings = append(warnings, fmt.Sprintf("list %v is present", arg))
							continue
						}
						err = keyval.DeleteCollection(arg)
						if err != nil {
							return err
						}
					}
					_, err = keyval.NewCollection(arg)
					if err != nil {
						warnings = append(warnings, err.Error())
						continue
					}
					msg = fmt.Sprintf("list '%v' created", arg)
					created++

					if c.Bool("verbose") {
						fmt.Println(msg)
					}
				}

				fmt.Printf("Lists created: %v\n", created)
				return nil
			},
		},
		{
			Name:        "print",
			Usage:       "kval print [list1]...",
			UsageText:   "print list of k-v pairs",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Info",
			Action: func(c *cli.Context) error {
				fullTree := []string{}
				args := c.Args()
				switch args.Len() {
				case 0:
					fullTree = directory.Tree(keyval.MakePath(""))
				default:
					for _, arg := range args.Slice() {
						fullTree = append(fullTree, directory.Tree(keyval.MakePath(arg))...)
					}
				}
				for _, leaf := range fullTree {
					if strings.HasSuffix(leaf, ".kv") {
						lists = append(lists, leaf)
					}
				}
				listsClean := []string{}
			loop1:
				for _, list := range lists {
					for _, clean := range listsClean {
						if list == clean {
							continue loop1
						}
					}
					listsClean = append(listsClean, list)
				}
				if loc != "" && len(listsClean) == 0 {
					return fmt.Errorf("no lists found in %v", loc)
				}
				for _, path := range listsClean {
					kv, err := keyval.LoadCollection(path)
					if err != nil {
						return err
					}
					prntArg := false
					keys, vals := kv.List()
					for i, k := range keys {
						if k != "" {
							if !prntArg {
								prntArg = true
								fmt.Println(path)
							}
							fmt.Printf("%v ==> %v\n", k, vals[i])
						}
					}
				}
				return nil
			},
		},
		{
			Name:        "stats",
			Usage:       "kval stats",
			UsageText:   "print info about database",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Info",
			Action: func(c *cli.Context) error {
				fullTree := []string{}
				args := c.Args()
				switch args.Len() {
				case 0:
					fullTree = directory.Tree(keyval.MakePath(""))
				default:
					for _, arg := range args.Slice() {
						fullTree = append(fullTree, directory.Tree(keyval.MakePath(arg))...)
					}
				}

				for _, leaf := range fullTree {
					if strings.HasSuffix(leaf, ".kv") {
						lists = append(lists, leaf)
					}
				}
				listsClean := []string{}
			loop1:
				for _, list := range lists {
					for _, clean := range listsClean {
						if list == clean {
							continue loop1
						}
					}
					listsClean = append(listsClean, list)
				}
				if loc != "" && len(listsClean) == 0 {
					return fmt.Errorf("no lists found in %v", loc)
				}
				listsNum := len(listsClean)
				keys := 0
				vals := 0
				badVal := 0
				for _, loc := range listsClean {
					kvmap := keyval.MapCollection(loc)
					for _, v := range kvmap {
						switch v {
						case "":
							badVal++
						default:
							vals++
						}
						keys++
					}
				}
				report := ""
				report += fmt.Sprintf("Lists found     : %v\n", listsNum)
				report += fmt.Sprintf("Keys found      : %v\n", keys)
				report += fmt.Sprintf("Values found    : %v\n", vals)
				pst := float64(int((float64(vals)/float64(keys))*10000)) / 100
				if badVal == 0 {
					pst = 100.0
				}
				report += fmt.Sprintf("Database Health : %v", pst) + " %\n"
				fmt.Println(report)
				return nil
			},
		},
		{
			Name:        "write",
			Usage:       "set/change k-v pair to database",
			UsageText:   "-location -key key1 -val value1",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Control",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "location",
					Category: "Context",
					Usage:    "set book/chapter/list with one argument",
					Required: true,
					Aliases: []string{
						"to",
						"from",
						"page",
						"loc",
					},
				},
				&cli.StringFlag{
					Name:     "key",
					Category: "Args",
					Usage:    "set key argument",
					Required: true,
					Aliases: []string{
						"k",
					},
				},
				&cli.StringFlag{
					Name:     "value",
					Category: "Args",
					Usage:    "set value argument",
					Required: true,
					Aliases: []string{
						"v",
					},
				},
			},
			Action: func(c *cli.Context) error {
				kv, err := keyval.LoadCollection(c.String("loc"))
				if err != nil {
					return err

				}
				kv.Set(c.String("k"), c.String("v"))
				if keyval.SaveCollection(kv) != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:        "append",
			Usage:       "adds value to k-v pair in database",
			UsageText:   "-location -key key1 -val value1",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Control",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "location",
					Category: "Context",
					Usage:    "set book/chapter/list with one argument",
					Required: true,
					Aliases: []string{
						"to",
						"from",
						"loc",
					},
				},
				&cli.StringFlag{
					Name:     "key",
					Category: "Args",
					Usage:    "set key argument",
					Required: true,
					Aliases: []string{
						"k",
					},
				},
				&cli.StringFlag{
					Name:     "value",
					Category: "Args",
					Usage:    "set value argument",
					Required: true,
					Aliases: []string{
						"v",
					},
				},
			},
			Action: func(c *cli.Context) error {
				kv, err := keyval.LoadCollection(c.String("loc"))
				if err != nil {
					return err

				}
				val := kv.Get(c.String("k"))
				nval := ""
				switch val {
				default:
					valSl := keyval.SliceValues(kv.Get(c.String("k")))
					for _, vl := range valSl {
						if vl == c.String("v") {
							return nil
						}
					}
					nval = val + keyval.KVUnitSep + c.String("v")
				case "":
					nval = c.String("v")
				}
				kv.Set(c.String("k"), nval)
				if keyval.SaveCollection(kv) != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:        "read",
			Usage:       "return list of all values for key",
			UsageText:   "-location -key key1",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Control",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "location",
					Category: "Context",
					Usage:    "set book/chapter/list with one argument",
					Required: true,
					Aliases: []string{
						"to",
						"from",
						"page",
						"loc",
					},
				},
				&cli.StringFlag{
					Name:     "key",
					Category: "Args",
					Usage:    "set key argument",
					Required: true,
					Aliases: []string{
						"k",
					},
				},
			},
			Action: func(c *cli.Context) error {
				kv, err := keyval.LoadCollection(c.String("loc"))
				if err != nil {
					return err

				}
				val := kv.Get(c.String("k"))
				vals := keyval.SliceValues(val)
				for _, vl := range vals {
					if strings.Contains(vl, "\n") {
						switch {
						case strings.HasPrefix(vl, `"`) && strings.HasSuffix(vl, `"`):
							vl = "`" + vl + "`"
						default:
							vl = `"` + vl + `"`

						}
					}
					fmt.Println(vl)
				}

				return nil
			},
		},
	}

	args := os.Args
	if err := app.Run(args); err != nil {
		fmt.Printf("%v returned error: %v\n", program, err.Error())
	}

}
