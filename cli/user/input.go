package user

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//InputStr - Awaits User-feeded input ended by byte '0x13' (\n)
//короче надо нажать много клавишь а потом 'Enter'.
func InputStr() (string, error) {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	switch runtime.GOOS {
	default:
		return "", errors.New("WARNING: Unknown System")
	case "windows":
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
	case "linux":
		str = strings.TrimSuffix(str, "\n")
	}
	return str, nil
}

//InputInt - same as InputStr, but convert input to Int
func InputInt() (int, error) {
	s, err := InputStr()
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

//InputFloat64 - same as InputStr, but convert input to float64
func InputFloat64() (float64, error) {
	s, err := InputStr()
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

//InputSliceStr - same as InputStr, but convert input to slice of strings
//with sp as separotor
func InputSliceStr(sp string) ([]string, error) {
	s, err := InputStr()
	if err != nil {
		return nil, err
	}
	sl := strings.Split(s, sp)
	return sl, nil
}

//InputSliceInt - same as InputSliceStr, but convert input to slice of Int's
//with space (' ') as separotor
func InputSliceInt() ([]int, error) {
	sl, err := InputSliceStr(" ")
	if err != nil {
		return nil, err
	}
	var sli []int
	for i := range sl {
		sla, err := strconv.Atoi(sl[i])
		if err != nil {
			return nil, err
		}
		sli = append(sli, sla)
	}
	return sli, nil
}

//Confirm - ask 'yes' or 'no' input
func Confirm(q string) (bool, error) {
	fmt.Print(q)
	answer, err := InputStr()
	answer = strings.ToUpper(answer)
	if err != nil {
		return false, err
	}
	switch answer {
	default:
		return false, errors.New("answer not clear")
	case "N", "NO", "0", "Т":
		return false, nil
	case "Y", "YES", "1", "Н":
		return true, nil
	}
}

//ChooseOne -  Awaits User-feeded input int
//выводит вопрос и ждет ответа от пользователя
func ChooseOne(q string, answers []string) (int, error) {
	if len(answers) == 0 {
		return -1, errors.New("no answers provided")
	}
	fmt.Print(q + "\n")
	if len(answers) == 1 {
		return 0, nil
	}
	for i, an := range answers {
		fmt.Print("[", i, "] - ", an, "\n")
	}
	answer, err := InputInt()
	if err != nil {
		return -1, err
	}
	if answer < 0 || answer >= len(answers) {
		return -1, errors.New("answer [" + strconv.Itoa(answer) + "] is outside of options array (0-" + strconv.Itoa(len(answers)) + ")")
	}
	return answer, nil
}
