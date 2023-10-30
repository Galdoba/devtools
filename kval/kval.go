package main

import (
	"fmt"
	"os"
	"strconv"
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
		//newlist
		{
			Name:        "new",
			Usage:       "create new list of k-v pairs",
			UsageText:   "kval newlist [list1]...",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "args as names if new lists",
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
						return err
					}
					if present {
						return fmt.Errorf("list '%v' is present", arg)
					}
					kv, err := keyval.NewKVlist(arg)
					if err != nil {
						return err
					}
					msg = fmt.Sprintf("list '%v' created", arg)
					created++
					if c.Bool("verbose") {
						fmt.Println(msg)
					}
					if err := kv.Save(); err != nil {
						return err
					}
				}
				fmt.Printf("Lists created: %v\n", created)
				return nil
			},
		},
		//print
		{
			Name:        "print",
			Usage:       "print lists of k-v pairs",
			UsageText:   "kval print [list1]...",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Read",
			Action: func(c *cli.Context) error {
				fullTree := []string{}
				args := c.Args()
				switch args.Len() {
				case 0:
					fullTree = directory.Tree(keyval.MakePathJS(""))
				default:
					for _, arg := range args.Slice() {
						fullTree = append(fullTree, directory.Tree(keyval.MakePathJS(arg))...)
					}
				}
				for _, leaf := range fullTree {
					if keyval.KVlistPresent(leaf) {
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
					kv, err := keyval.Load(path)
					if err != nil {
						return err
					}

					text := fmt.Sprintf("Source: %v\n", kv.Path)
					for _, k := range kv.Keys() {
						vals, err := kv.GetAll(k)
						if err != nil {
							return err
						}
						text += fmt.Sprintf("%s: [", k)
						for _, v := range vals {
							text += `"` + v + `"` + ", "
						}
						text = strings.TrimSuffix(text, ", ")
						text += "]"
					}
					fmt.Printf("%s\n", text)
				}
				return nil
			},
		},
		//keys
		{
			Name:        "keys",
			Usage:       "print keys-only form kval file(s)",
			UsageText:   "kval keys [list]...",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Read",
			Action: func(c *cli.Context) error {
				fullTree := []string{}
				args := c.Args()
				switch args.Len() {
				case 0:
					fullTree = directory.Tree(keyval.MakePathJS(""))
				default:
					for _, arg := range args.Slice() {
						fullTree = append(fullTree, directory.Tree(keyval.MakePathJS(arg))...)
					}
				}
				for _, leaf := range fullTree {
					if keyval.KVlistPresent(leaf) {
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

					kv, err := keyval.Load(path)
					if err != nil {
						return err
					}
					text := fmt.Sprintf("Source: %v\n", kv.Path)
					for _, k := range kv.Keys() {

						text += fmt.Sprintf("%s\n", k)

					}
					fmt.Printf("%s", text)
				}
				return nil
			},
		},
		//stats
		{
			Name:        "stats",
			Usage:       "print info about database",
			UsageText:   "kval stats",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Delete",
			Action: func(c *cli.Context) error {
				fullTree := []string{}
				args := c.Args()
				switch args.Len() {
				case 0:
					fullTree = directory.Tree(keyval.MakePathJS(""))
				default:
					for _, arg := range args.Slice() {
						fullTree = append(fullTree, directory.Tree(keyval.MakePathJS(arg))...)
					}
				}

				for _, leaf := range fullTree {
					if keyval.KVlistPresent(leaf) {
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
				for _, loc := range listsClean {
					kv, err := keyval.Load(loc)
					if err != nil {
						return err
					}
					kvmap := kv.Data()
					for _, v := range kvmap {
						vals = vals + len(v)
						keys++
					}
				}
				report := ""
				report += fmt.Sprintf("Lists found     : %v\n", listsNum)
				report += fmt.Sprintf("Keys found      : %v\n", keys)
				report += fmt.Sprintf("Values found    : %v", vals)
				fmt.Println(report)
				return nil
			},
		},
		//write
		{
			Name:        "write",
			Usage:       "set/change k-v pair to database",
			UsageText:   "-location -key key1 {vals}...",
			Description: "Заменяет все значения для указанного ключа на значения указанные в аргументах",
			Category:    "Update",
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
				kv, err := keyval.Load(c.String("loc"))
				if err != nil {
					return err

				}
				vals := c.Args().Slice()
				if len(vals) == 0 {
					return fmt.Errorf("no arguments provided")
				}
				if err := kv.Set(c.String("k"), vals...); err != nil {
					return err
				}
				if err := kv.Save(); err != nil {
					return err
				}
				return nil
			},
		},
		//append
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
				&cli.BoolFlag{
					Name:     "unique",
					Category: "Args",
					Usage:    "add only unique",
					Required: false,
					Aliases: []string{
						"u",
					},
				},
			},
			Action: func(c *cli.Context) error {
				kv, err := keyval.Load(c.String("loc"))
				if err != nil {
					return err
				}
				vals := c.Args().Slice()
				for _, val := range vals {
					if err := kv.Add(c.String("k"), val, c.Bool("u")); err != nil {
						return err
					}
				}
				// if err := kv.Save(); err != nil {
				// 	return err
				// }
				return nil
			},
		},
		//read
		{
			Name:        "read",
			Usage:       "return list of all values for key",
			UsageText:   "-location -key key1 [index]...",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Принимает индексы (целые числа). Если индексов нет - выводит все значения",
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
				key := c.String("k")
				indexes := []int{}
				for n, val := range c.Args().Slice() {
					i, err := strconv.Atoi(val)
					if err != nil {
						return fmt.Errorf("incorrect index passed %v (%v): %v", n, val, err.Error())
					}
					indexes = append(indexes, i)
				}
				kv, err := keyval.Load(c.String("loc"))
				if err != nil {
					return err
				}

				one, err := kv.GetSingle(key)
				if err == nil {
					fmt.Println(one)
					return nil
				}

				many, err := kv.GetAll(key)
				if err != nil {
					return err
				}
				res := ""
				switch len(indexes) {
				case 0:
					for _, vl := range many {
						res += fmt.Sprintf("%v\n", vl)
					}
				default:
					vals, err := kv.GetByIndex(key, indexes...)
					if err != nil {
						return fmt.Errorf("incorrect index passed: %v", err.Error())
					}
					for _, v := range vals {
						res += v + "\n"
					}
				}
				fmt.Printf("%v", res)
				return nil
			},
		},
		//confirm
		{
			Name:        "confirm",
			Usage:       "return 1 if exist and 0 if not",
			UsageText:   "-location [-key]...",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Принимает индексы (целые числа). Если индексов нет - выводит все значения",
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
					Required: false,
					Aliases: []string{
						"k",
					},
				},
			},
			Action: func(c *cli.Context) error {
				key := c.String("k")

				kv, err := keyval.Load(c.String("loc"))
				if err != nil {
					fmt.Println("0")
					return nil
				}
				if key == "" {
					fmt.Println("1")
					return nil
				}

				one, err := kv.GetSingle(key)
				if err != nil {
					fmt.Println("0")
					return nil
				}
				out := ""
				switch one {
				case "":
					out = "0"
				default:
					out = "1"
				}
				fmt.Println(out)
				return nil
			},
		},
		//remove
		{
			Name:        "remove",
			Usage:       "remove specific value from k-v pair",
			UsageText:   "kval remove -from [list] -key [key] -val [val]",
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
					Aliases: []string{
						"v",
					},
				},
			},
			Action: func(c *cli.Context) error {
				key := c.String("k")
				kv, err := keyval.Load(c.String("loc"))
				if err != nil {
					return err
				}
				switch c.String("v") {
				case "":
					err := kv.RemoveByKey(key)
					if err != nil {
						return err
					}
				default:
					switch len(kv.Data()[key]) {
					case 1:
						err := kv.RemoveByKey(key)
						if err != nil {
							return err
						}
					default:
						err := kv.RemoveByVal(key, c.String("v"))
						if err != nil {
							return err
						}
					}
				}
				return kv.Save()
			},
		},
		//delete
		{
			Name:        "delete",
			Usage:       "delete list from database",
			UsageText:   "-location",
			Description: "TODO: подробное описание команды",
			ArgsUsage:   "Аргументов не имеет\nВ планах локальный режим и указание файла в который должен писаться отчет",
			Category:    "Control",
			Action: func(c *cli.Context) error {
				vals := c.Args().Slice()
				for _, val := range vals {
					path := keyval.MakePathJS(val)
					present, err := keyval.Present(path)
					if err != nil {
						return err
					}
					if !present {
						return fmt.Errorf("no list found")
					}
					f, _ := os.Stat(path)
					if f.IsDir() {
						return fmt.Errorf("'%v' is a directory", path)
					}
					if err := os.Remove(path); err != nil {
						return fmt.Errorf("can't remove %v: %v", path, err.Error())
					}
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
