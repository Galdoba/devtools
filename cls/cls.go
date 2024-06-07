package cls

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

const (
	//levels
	LV_DEBUG  = iota //пишем начало/конец функций и значения переменных
	LV_INFO          //пишем общая информация
	LV_REPORT        //пишем завершение процессов
	LV_ERROR         //пишем ошибки приводящие к рестарту процесса
	LV_FATAL         //пишем ошибки приводящие концу программы
	lv_none          //не пишем
)

/*
cls.Fprintf()
cls.Errorf() error
cls.Info()
cls.Report()


*/

// cls - Combined Log Structure
type cls struct {
	levelControl map[int][]*log.Logger
	errorPrinter *log.Logger
	printerMode  bool
	files        []string
}

func New() *cls {
	ls := cls{}
	ls.errorPrinter = log.New(os.Stderr, "", log.Lshortfile)
	ls.levelControl = make(map[int][]*log.Logger)
	return &ls
}

func writeMsg(logger *log.Logger, msg string) {
	logger.Println(msg)
}

func (l *cls) AddFile(path string, lv int, prefix string, flags int) error {
	for _, fl := range l.files {
		if path == fl {
			return fmt.Errorf("cls can't add file '%v': was already added", path)
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("cls can't add file: %v", err.Error())
	}
	// defer f.Close()
	for i := lv; i < lv_none; i++ {
		logger := log.New(f, leveltag(i)+prefix, flags)
		l.levelControl[i] = append(l.levelControl[i], logger)
	}
	l.files = append(l.files, path)
	return nil
}

func leveltag(lv int) string {
	switch lv {
	default:
		return ""

	case LV_DEBUG:
		return "DEBUG: "
	case LV_INFO:
		return "INFO:  "
	case LV_ERROR:
		return "ERROR: "
	case LV_FATAL:
		return "FATAL: "
	}
}

func (c *cls) Debug(msg string) {
	if loggers, ok := c.levelControl[LV_DEBUG]; ok {
		for _, logger := range loggers {
			writeMsg(logger, msg)
		}
	}
}

func (c *cls) Info(msg string) {
	if loggers, ok := c.levelControl[LV_INFO]; ok {
		for _, logger := range loggers {
			writeMsg(logger, msg)
		}
	}
}

func (c *cls) Error(msg string) {
	msgCol := color.RedString(msg)
	c.errorPrinter.Output(2, msgCol)
	if loggers, ok := c.levelControl[LV_ERROR]; ok {
		for _, logger := range loggers {
			writeMsg(logger, msg)
		}
	}
}

func (c *cls) Fatal(msg string) {
	msgCol := color.RedString(msg)
	c.errorPrinter.Output(2, msgCol)
	if loggers, ok := c.levelControl[LV_FATAL]; ok {
		for _, logger := range loggers {
			writeMsg(logger, msg)
		}
	}
	panic(msg)
}

func (c *cls) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Print(msg)
	c.Info(msg)
}

func (c *cls) Println(args ...interface{}) {
	fmt.Println(args...)
	format := ""
	for i := 0; i < len(args); i++ {
		format += "%v "
	}
	format = strings.TrimSuffix(format, " ")
	c.Info(fmt.Sprintf(format, args...))
}

func (c *cls) Report(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msgCol := color.GreenString(msg)
	fmt.Print(msgCol)
	c.Info(msg)
}

func (c *cls) Errorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	msgCol := color.RedString(msg)
	c.errorPrinter.Output(2, msgCol)
	c.Error(msg)
	return fmt.Errorf(format, args...)

}
