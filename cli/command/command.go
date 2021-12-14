package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	TERMINAL_ON = iota
	TERMINAL_OFF
	BUFFER_ON
	BUFFER_OFF
	FILE
	STD_INPUT
)

type terminalCommand struct {
	programPath string
	args        []string
	term        bool
	buffer      bool
	filePaths   []string
	writersOUT  []io.Writer
	writersERR  []io.Writer
	stOut       string
	stErr       string
}

type commandInstruction struct {
	instType int
	arg      string
}

//New - создает и наполняет конструкт запускающийся в стандартном терминале
func New(inst ...commandInstruction) (*terminalCommand, error) {
	tc := terminalCommand{}
	for _, val := range inst {
		switch val.instType {
		default:
			continue
		case STD_INPUT:
			args := strings.Split(val.arg, " ")
			tc.programPath = args[0]
			tc.args = args[1:]
		case TERMINAL_ON:
			tc.term = true
		case TERMINAL_OFF:
			tc.term = false
		case BUFFER_ON:
			tc.buffer = true
		case BUFFER_OFF:
			tc.buffer = false
		case FILE:
			tc.filePaths = append(tc.filePaths, val.arg)
		}
	}
	if tc.programPath == "" {
		return nil, fmt.Errorf("command line undefined")
	}
	return &tc, nil
}

//Run - запускает объект обращаясь к стандартному терминалу
//ВНИМАНИЕ: Дефолтное состояние НЕ выводить информацию по ходу выполнения программы
//в консоль и буфер
func (tc *terminalCommand) Run() error {
	var o bytes.Buffer
	var e bytes.Buffer
	time.Sleep(time.Millisecond * 2)
	cmd := exec.Command(tc.programPath, tc.args...)
	//Control output for Console
	if tc.term {
		tc.writersOUT = append(tc.writersOUT, os.Stdout)
		tc.writersERR = append(tc.writersERR, os.Stderr)
	}
	//Control output for Buffer
	if tc.buffer {
		tc.writersOUT = append(tc.writersOUT, &o)
		tc.writersERR = append(tc.writersERR, &e)
	}
	//Control output for Files
	for _, fl := range tc.filePaths {
		f, err := os.OpenFile(fl, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		tc.writersOUT = append(tc.writersOUT, f)
		tc.writersERR = append(tc.writersERR, f)
	}
	//Setup writer(s)
	cmd.Stdout = io.MultiWriter(tc.writersOUT...)
	cmd.Stderr = io.MultiWriter(tc.writersERR...)
	err := cmd.Run()
	tc.stOut = o.String()
	tc.stErr = e.String()
	return err
}

//Line - задает командную строку (команда + аргументы)
func CommandLineArguments(prog string, args ...string) commandInstruction {
	comLine := prog + " "
	for _, arg := range args {
		comLine += arg
	}
	comLine = strings.TrimSuffix(comLine, " ")
	return commandInstruction{STD_INPUT, comLine}
}

//Set - контролирует вывод на out и error
//Доступные опции: TERMINAL_OFF, TERMINAL_ON, BUFFER_OFF, BUFFER_ON:
func Set(i int) commandInstruction {
	switch i {
	case TERMINAL_OFF, TERMINAL_ON, BUFFER_OFF, BUFFER_ON:
		return commandInstruction{i, ""}
	}
	return commandInstruction{}
}

//WriteToFile - добавляется файл в который будет писаться output и error
func WriteToFile(path string) commandInstruction {
	return commandInstruction{FILE, path}
}

func (tc *terminalCommand) StdOut() string {
	return tc.stOut
}

func (tc *terminalCommand) StdErr() string {
	return tc.stErr
}
