package cls

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const (
	//levels
	LV_TRACE   = iota //пишем начало/конец циклов и значения переменных
	LV_DEBUG          //пишем начало/конец функций и значения переменных
	LV_NOTICE         //пишем начало экспортируемымых функций
	LV_INFO           //пишем общая информация
	LV_REPORT         //пишем завершение процессов
	LV_WARNING        //пишем не критичные ошибки
	LV_ERROR          //пишем ошибки приводящие к рестарту процесса
	LV_FATAL          //пишем ошибки приводящие концу программы
	lv_none           //не пишем
)

//cls - Combined Log Structure
type cls struct {
	levelControl map[int][]*log.Logger
	printerMode  bool
	files        []string
}

func New() *cls {
	ls := cls{}
	ls.levelControl = make(map[int][]*log.Logger)
	return &ls
}

func writeMsg(logger *log.Logger, sep string, a ...interface{}) {
	format := ""
	for i := 0; i < len(a); i++ {
		format += "%v" + sep
	}
	msg := fmt.Sprintf(format, a...)

	logger.Print(msg)
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

	for i := lv; i < lv_none; i++ {
		logger := log.New(f, leveltag(i)+prefix, flags)
		l.levelControl[i] = append(l.levelControl[i], logger)
	}
	l.files = append(l.files, path)
	return nil
}

func (l *cls) AddStdErr(lv int, prefix string, flags int) error {
	for _, loggers := range l.levelControl {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr {
				return fmt.Errorf("stderr was already added")
			}
		}
	}
	for i := lv; i < lv_none; i++ {
		prf := leveltag(i) + prefix
		if flags == 0 {
			prf = ""
		}
		logger := log.New(os.Stderr, prf, flags)
		l.levelControl[i] = append(l.levelControl[i], logger)
	}
	return nil
}

func leveltag(lv int) string {
	switch lv {
	default:
		return ""
	case LV_TRACE:
		return "TRACE: "
	case LV_DEBUG:
		return "DEBUG: "
	case LV_NOTICE:
		return " NOTE: "
	case LV_INFO:
		return " INFO: "
	case LV_REPORT:
		return "REPOR: "
	case LV_WARNING:
		return "WRNNG: "
	case LV_ERROR:
		return "ERROR: "
	case LV_FATAL:
		return "FATAL: "
	}
}

func (c *cls) Trace(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_TRACE]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
	c.printerMode = false
}

func (c *cls) Debug(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_DEBUG]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
	c.printerMode = false
}

func (c *cls) Notice(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_NOTICE]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
	c.printerMode = false
}

func (c *cls) Info(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_INFO]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
	c.printerMode = false
}

func (c *cls) Report(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_REPORT]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
}

func (c *cls) Warning(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_WARNING]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
	c.printerMode = false
}

func (c *cls) Error(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_ERROR]; ok {
		for _, logger := range loggers {
			if logger.Writer() == os.Stderr && c.printerMode {
				continue
			}
			writeMsg(logger, " ", a...)
		}
	}
	c.printerMode = false
}

func (c *cls) Fatal(a ...interface{}) {
	if loggers, ok := c.levelControl[LV_FATAL]; ok {
		for _, logger := range loggers {
			writeMsg(logger, " ", a...)
		}
	}
	panic(fmt.Sprintf("FATAL: %v", a))
}

func (c *cls) Fprintf(lv int, format string, args ...interface{}) {
	c.printerMode = true
	msg := fmt.Sprintf(format, args...)
	switch lv {
	default:
		return
	case LV_TRACE:
		fmt.Print(color.HiBlackString(msg))
		c.Trace(msg)
	case LV_DEBUG:
		fmt.Print(color.HiBlackString(msg))
		c.Debug(msg)
	case LV_NOTICE:
		fmt.Print(color.HiCyanString(msg))
		c.Notice(msg)
	case LV_INFO:
		fmt.Print(msg)
		c.Report(msg)
	case LV_REPORT:
		fmt.Print(color.HiGreenString(msg))
		c.Report(msg)
	case LV_WARNING:
		fmt.Print(color.HiYellowString(msg))
		c.Warning(msg)
	case LV_ERROR:
		fmt.Print(color.RedString(msg))
		c.Error(msg)
	case LV_FATAL:
		fmt.Print(color.RedString(msg))
		c.Fatal(msg)
	}
}

func (c *cls) Errorf(format string, args ...interface{}) error {
	c.printerMode = true
	msg := fmt.Sprintf(format, args...)
	fmt.Print(color.RedString(msg))
	c.Error(msg)
	return fmt.Errorf(format, args...)
}
