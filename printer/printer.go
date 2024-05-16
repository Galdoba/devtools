package printer

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type printManager struct {
	file         string
	fileLevel    int
	consoleLevel int
	appName      string
	duration     time.Duration
	//consolecolor bool
	//durationReport bool
}

const (
	levelALL = iota
	levelTRACE
	levelDEBUG
	levelNOTICE
	levelINFO
	levelREPORT
	levelWARN
	levelERROR
	levelFATAL
	levelOFF
	timestamp = "2006-02-01 15:04:05"
)

func New(path string) *printManager {
	return &printManager{
		file: path,
	}
}

func (pm *printManager) WithFileLevel(fl int) *printManager {
	switch fl {
	default:
		return pm
	case levelALL, levelTRACE, levelDEBUG, levelNOTICE, levelINFO, levelWARN, levelERROR, levelFATAL, levelOFF:
		pm.fileLevel = fl
		return pm
	}
}

func (pm *printManager) WithConsoleLevel(fl int) *printManager {
	switch fl {
	default:
		return pm
	case levelALL, levelTRACE, levelDEBUG, levelNOTICE, levelINFO, levelWARN, levelERROR, levelFATAL, levelOFF:
		pm.fileLevel = fl
		return pm
	}
}

func (pm *printManager) WithAppName(appName string) *printManager {
	pm.appName = appName
	return pm
}

func (pm *printManager) Info(format string, args ...interface{}) []error {
	if pm.fileLevel >= levelINFO {
		fileText := formatForFile(pm.appName, levelINFO, pm.duration, format, args...)
		writeToFile(pm.file, fileText)
		fmt.Println()
	}

}

func formatForConsole(colorOutput bool, level int, duration time.Duration, format string, args ...interface{}) string {
	s := fmt.Sprintf(format, args...)
	if level <= levelTRACE {
		s += " (" + duration.String() + ")"
	}
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

func formatForFile(caller string, level int, duration time.Duration, format string, args ...interface{}) string {
	t := time.Now()
	s := t.Format(timestamp)
	if caller != "" {
		s += " " + caller
	}
	s += "[" + levelStr(level) + "]: "
	s += fmt.Sprintf(format, args...)
	if level <= levelTRACE {
		s += " (" + duration.String() + ")"
	}
	return s
}

func writeToFile(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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
		return "ALL    "
	case levelTRACE:
		return "TRACE  "
	case levelNOTICE:
		return "NOTICE "
	case levelDEBUG:
		return "DEBUG  "
	case levelINFO:
		return "INFO   "
	case levelREPORT:
		return "REPORT "
	case levelWARN:
		return "WARNING"
	case levelERROR:
		return "ERROR  "
	case levelFATAL:
		return "FATAL  "
	case levelOFF:
		return "OFF    "
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
