package printer

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Galdoba/devtools/printer/lvl"
	"github.com/fatih/color"
)

type printManager struct {
	file string
	//dest           *os.File
	fileLevel      int
	consoleLevel   int
	appName        string
	lastMessage    *time.Time
	consolecolor   bool
	durationReport bool
	logConsole     *log.Logger
	//logFile        *log.Logger
	// loggerALL      *log.Logger
	// loggerTRACE    *log.Logger
	// loggerDEBUG    *log.Logger
	// loggerNOTICE   *log.Logger
	// loggerINFO     *log.Logger
	// loggerREPORT   *log.Logger
	// loggerWARN     *log.Logger
	// loggerERROR    *log.Logger
	// loggerFATAL    *log.Logger
}

const (
	levelALL    = iota
	levelTRACE  //TRC
	levelDEBUG  //DBG
	levelNOTICE //NOT
	levelINFO   //INF
	levelREPORT //REP
	levelWARN   //WRN
	levelERROR  //ERR
	levelFATAL  //FAT
	levelOFF    //OFF
	timestamp   = "2006-01-02 15:04:05.999"
)

func New() *printManager {
	pm := printManager{
		consoleLevel: lvl.INFO,
		fileLevel:    lvl.OFF,
	}
	pm.logConsole = log.New(os.Stderr, "", 0)
	// if logFilePath != "" {
	// 	file, err := os.OpenFile(pm.file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// 	if err != nil {
	// 		return &pm
	// 	}
	// 	defer file.Close()
	// 	pm.file = logFilePath
	// 	pm.fileLevel = lvl.INFO
	// 	//pm.logFile = log.New(pm.dest, "", 0)

	// }

	return &pm
}

type Printer interface {
	Print(int, ...interface{})
	Println(int, ...interface{})
	Printf(int, string, ...interface{})
	Errorf(int, string, ...interface{}) error
}

func (pm *printManager) WithFile(logFilePath string) *printManager {
	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err.Error())
			return pm
		}
		defer file.Close()
		pm.file = logFilePath
		pm.fileLevel = lvl.INFO
		fmt.Println(pm.file, "set")
	}
	return pm
}

func (pm *printManager) WithFileLevel(fl int) *printManager {
	switch fl {
	default:
		return pm
	case lvl.ALL, lvl.TRACE, lvl.DEBUG, lvl.NOTICE, lvl.INFO, lvl.WARN, lvl.ERROR, lvl.FATAL, lvl.OFF:
		pm.fileLevel = fl
		return pm
	}
}

func (pm *printManager) WithConsoleLevel(fl int) *printManager {
	switch fl {
	default:
		return pm
	case lvl.ALL, lvl.TRACE, lvl.DEBUG, lvl.NOTICE, lvl.INFO, lvl.WARN, lvl.ERROR, lvl.FATAL, lvl.OFF:
		pm.consoleLevel = fl
		return pm
	}
}

func (pm *printManager) WithConsoleColors(ccol bool) *printManager {
	pm.consolecolor = ccol
	return pm
}

func (pm *printManager) WithAppName(appName string) *printManager {
	pm.appName = appName
	return pm
}

/*
log.Info("answer is %v",42)
log.Fprintf(lv.INFO,"answer is %v",42)
*/

func (pm *printManager) Printf(level int, format string, args ...interface{}) {
	if pm.consoleLevel > level && pm.fileLevel > level {
		return
	}
	if pm.consoleLevel <= level {
		text := formatForConsole(pm.consolecolor, level, format, args...)
		fmt.Print(text)
	}
	if pm.file != "" && pm.fileLevel <= level {
		t := time.Now()
		text := formatForFile(t, pm.appName, level, format, args...)
		text = strings.TrimSuffix(text, "\n")
		if level <= lvl.TRACE {
			if pm.lastMessage == nil {
				pm.lastMessage = &t
			}
			dur := time.Since(*pm.lastMessage)
			text += " ( " + dur.String() + " )"
			pm.lastMessage = &t
		}
		file, _ := os.OpenFile(pm.file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer file.Close()

		file.WriteString(text + "\n")
	}
}

func (pm *printManager) Errorf(format string, args ...interface{}) error {
	if pm.consoleLevel > lvl.ERROR && pm.fileLevel > lvl.ERROR {
		return fmt.Errorf(format, args...)
	}
	if pm.consoleLevel <= lvl.ERROR {
		text := formatForConsole(pm.consolecolor, lvl.ERROR, format, args...)
		fmt.Print(text)
	}
	if pm.file != "" && pm.fileLevel <= lvl.ERROR {
		t := time.Now()
		text := formatForFile(t, pm.appName, lvl.ERROR, format, args...)
		text = strings.TrimSuffix(text, "\n")
		file, _ := os.OpenFile(pm.file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer file.Close()

		file.WriteString(text + "\n")
	}
	return fmt.Errorf(format, args...)
}

func (pm *printManager) Print(level int, args ...interface{}) {
	if pm.consoleLevel > level && pm.fileLevel > level {
		return
	}
	if pm.consoleLevel <= level {
		text := formatForConsole(pm.consolecolor, level, "", args...)
		fmt.Print(text)
	}
	if pm.file != "" && pm.fileLevel <= level {
		t := time.Now()
		text := formatForFile(t, pm.appName, level, "", args...)
		text = strings.TrimSuffix(text, "\n")
		if level <= lvl.TRACE {
			if pm.lastMessage == nil {
				pm.lastMessage = &t
			}
			dur := time.Since(*pm.lastMessage)
			text += " ( " + dur.String() + " )"
			pm.lastMessage = &t
		}
		file, _ := os.OpenFile(pm.file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer file.Close()

		file.WriteString(text + "\n")
	}
}

func (pm *printManager) Println(level int, args ...interface{}) {
	if pm.consoleLevel > level && pm.fileLevel > level {
		return
	}
	if pm.consoleLevel <= level {
		text := formatForConsole(pm.consolecolor, level, "", args...)
		fmt.Println(text)
	}
	if pm.file != "" && pm.fileLevel <= level {
		t := time.Now()
		text := formatForFile(t, pm.appName, level, "", args...)
		text = strings.TrimSuffix(text, "\n")
		if level <= lvl.TRACE {
			if pm.lastMessage == nil {
				pm.lastMessage = &t
			}
			dur := time.Since(*pm.lastMessage)
			text += " ( " + dur.String() + " )"
			pm.lastMessage = &t
		}
		file, _ := os.OpenFile(pm.file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer file.Close()

		file.WriteString(text + "\n")
	}
}

// func (pm *printManager) Info(format string, args ...interface{}) {
// 	pm.lastMessage = time.Now()
// 	if pm.fileLevel >= levelINFO {
// 		fileText := formatForFile(pm.lastMessage, pm.appName, levelINFO, format, args...)
// 		writeToFile(pm.file, fileText)
// 	}
// 	if pm.consoleLevel >= levelINFO {
// 		text := formatForConsole(pm.consolecolor, levelINFO, format, args...)
// 		fmt.Println(text)
// 	}

// }

func formatForConsole(colorOutput bool, level int, format string, args ...interface{}) string {
	if format == "" {
		for i := 0; i < len(args); i++ {
			format += "%v"
		}
	}
	s := fmt.Sprintf(format, args...)

	// if level <= levelTRACE {
	// 	s += " (" + duration.String() + ")"
	// }
	if !colorOutput {
		return s
	}
	switch level {
	default:
	case levelALL, levelTRACE, levelDEBUG, levelINFO:
	case levelREPORT:
		s = color.HiGreenString(s)
	case levelWARN:
		s = color.YellowString(s)
	case levelERROR, levelFATAL:
		s = color.RedString(s)
	case levelNOTICE:
		s = color.CyanString(s)
	}
	return s
}

func formatForFile(t time.Time, caller string, level int, format string, args ...interface{}) string {
	s := t.Format(timestamp)
	for len(s) < 23 {
		s += "0"
	}
	s += " [" + levelStr(level) + "]"
	if caller != "" {
		s += " " + caller
	}
	s += ": "
	if format == "" {
		for i := 0; i < len(args); i++ {
			format += "%v"
		}
	}
	s += fmt.Sprintf(format, args...)

	return s
}

func writeToFile(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0770)
	if err != nil {
		return fmt.Errorf("open logfile fail: %v", err.Error())
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return fmt.Errorf("write to logfile fail: %v", err.Error())
	}
	return nil
}

func levelStr(level int) string {

	switch level {
	default:
		return "       "
	case levelALL:
		return " ALL  "
	case levelTRACE:
		return "TRACE "
	case levelNOTICE:
		return "NOTICE"
	case levelDEBUG:
		return "DEBUG "
	case levelINFO:
		return " INFO "
	case levelREPORT:
		return "REPORT"
	case levelWARN:
		return " WARN "
	case levelERROR:
		return "ERROR "
	case levelFATAL:
		return "FATAL "
	case levelOFF:
		return " OFF  "
	}
}

/*
pm := printer.New()
pm.WithWriters(os.StdErr)
pm.ShoutIf(INFO)
pm.PrintfAs(printer.INFO,"%v\n", arg)

print.Info("text :%v ", 42)

in file:
2024-05-17 21:30:01.000 INFO text: 42
glogerr

*/
