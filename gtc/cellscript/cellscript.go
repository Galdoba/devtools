package cellscript

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Galdoba/devtools/cli/command"
)

const (
	Status_Ready = iota
	Status_InProgress
	Status_Completed
	Status_Not_Initiated
	Status_Not_Ready
	Status_Conflict
)

type script struct {
	path   string
	args   []string
	stdout string
	stderr string
	proc   *process
}

func (sc *script) asses() error {
	absPath, _ := filepath.Abs(sc.path)
	f, err := os.OpenFile(absPath, os.O_RDONLY, 0777)
	if err != nil {
		return fmt.Errorf("script assesment failed: %v", err.Error())
	}
	defer f.Close()
	if !sc.proc.initiated && !sc.proc.ready {
		sc.proc.initiated = true
		sc.proc.ready = true
	} else {
		return fmt.Errorf("must not be initiated or ready at this point")
	}

	return nil
}

func (sc *script) Run() (string, string, error) {
	scArgs := append([]string{}, sc.path)
	scArgs = append(scArgs, sc.args...)
	sOut, sErr, err := command.Execute(strings.Join(scArgs, " "), command.Set(command.BUFFER_ON))
	switch runtime.GOOS {
	case "windows":
		if strings.HasSuffix(sc.path, ".bat") {
			sOut = strings.TrimSuffix(sOut, "\n")
		}
	}

	return sOut, sErr, err
}

func NewScript(path string, args ...string) *script {
	sc := script{}
	sc.path = path
	sc.args = args
	sc.proc = &process{}
	sc.stdout = "stdout"
	sc.stderr = "stderr"
	return &sc
}

type process struct {
	ready      bool
	initiated  bool
	inprogress bool
	completed  bool
}

type Process interface {
	Status() int //Тестируем состояние процесса, возвращаем ошибку с сообщением
	//Run(string, []string) (string, string, error) //запускаем процесс и читаем состояние (нужны райтеры)
}

func (pr *process) Status() int {
	stat := 0
	if pr.initiated {
		stat += 8
	}
	if pr.ready {
		stat += 4
	}
	if pr.inprogress {
		stat += 2
	}
	if pr.completed {
		stat += 1
	}
	switch stat {
	default:
		return Status_Not_Initiated
	case 8:
		return Status_Not_Ready
	case 11, 13, 14, 15:
		return Status_Conflict
	case 12:
		return Status_Ready
	case 10:
		return Status_InProgress
	case 9:
		return Status_Completed
	}
}
