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
	for _, in := range inst {
		tc.AddInstruction(in)
	}
	if tc.programPath == "" {
		return nil, fmt.Errorf("command line undefined")
	}
	return &tc, nil
}

func (in *commandInstruction) String() string {
	return fmt.Sprintf("Instruction type: %v, arg: '%v'", in.instType, in.arg)
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

//AddInstruction - добавляет в инструкции информацию о том что и как делать
//там где инструкции противоречат друг другу приоритетной будет послеледняя
func (tc *terminalCommand) AddInstruction(ti commandInstruction) {
	switch ti.instType {
	default:
	case STD_INPUT:
		args := strings.Split(ti.arg, " ")
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
		if ti.arg != "" {
			tc.filePaths = append(tc.filePaths, ti.arg)
		}
	}
}

//StdOut - возвращает стандартный вывод
func (tc *terminalCommand) StdOut() string {
	return tc.stOut
}

//StdErr - возвращает стандартную ошибку
func (tc *terminalCommand) StdErr() string {
	return tc.stErr
}

func RunSilent(cmmnd string, args ...string) (string, error) {
	comm, err := New(
		CommandLineArguments(cmmnd, args...),
		Set(TERMINAL_OFF),
		Set(BUFFER_ON),
	)
	if err != nil {
		return "", err
	}
	runErr := comm.Run()
	return comm.StdOut(), runErr
}
