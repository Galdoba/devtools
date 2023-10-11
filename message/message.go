package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	from_STATE = iota
	to_STATE
	title_STATE
	args_STATE

	prefix = ""
	indent = "  "
)

var unmarshalOrder = []string{
	"From",
	"To",
	"Title",
	"Args",
}

type message struct {
	From  string            `json:"From"`
	To    string            `json:"To"`
	Title string            `json:"Title,omitempty"`
	Args  map[string]string `json:"Args,omitempty"`
	List  []string          `json:"List"`
}

func (msg *message) UnmarshalJSON(data []byte) error {
	jsonStr := string(data)
	jsonSl := strings.Split(jsonStr, "\n")

	lastState := -1
	for lineNum, line := range jsonSl {
		catched := catchState(line)
		if catched != -1 {
			lastState = catched
			switch lastState {
			case 3:
				msg.Args = make(map[string]string)
				continue
			}
		}
		fmt.Printf("line %v: state %v: %v\n", lineNum, lastState, line)
		switch lastState {
		case from_STATE:
			fmt.Println("FROMLINE", line)
			msg.From = stringValueFrom(line)
		case to_STATE:
			msg.To = stringValueFrom(line)
			fmt.Println(msg.To)
		case title_STATE:
			msg.Title = stringValueFrom(line)
		case args_STATE:
			key, val := stringKVfrom(line)
			msg.Args[key] = val
		default:
		}

	}
	delete(msg.Args, "")
	fmt.Println(msg)
	return nil
}

func stringValueFrom(s string) string {
	fmt.Println("----", s)
	data := strings.Split(s, `": "`)
	if len(data) > 1 {
		return strings.TrimSuffix(data[1], `",`)
	}

	return ""
}

func stringKVfrom(s string) (string, string) {
	fmt.Println("----", s)
	for strings.HasPrefix(s, indent) {
		s = strings.TrimPrefix(s, indent)
	}
	fmt.Println("KV", s)
	s = strings.TrimPrefix(s, `"`)
	s = strings.TrimSuffix(s, `"`)
	data := strings.Split(s, `": "`)
	if len(data) == 2 {

		return data[0], data[1]
	}
	return "", ""
}

func catchState(s string) int {
	stateNum := -1
	for i, state := range unmarshalOrder {
		if countIndent(s, indent) != 1 {
			continue
		}
		if strings.Contains(s, state) {
			stateNum = i
		}
	}
	return stateNum
}

// indent - returns indentNum
func countIndent(str, ind string) int {
	count := 0
	for strings.HasPrefix(str, ind) {
		count++
		str = strings.TrimPrefix(str, ind)
	}
	return count
}

func (msg *message) Marshal() ([]byte, error) {
	return json.MarshalIndent(msg, prefix, indent)
}

func main() {
	msg := message{
		From:  "tracker",
		To:    "director",
		Title: "file detected",
		Args: map[string]string{
			"file": "amedia_s01e01.mp4",
		},
	}
	jsonData, err := msg.Marshal()
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
	filepath := "/home/galdoba/workbench/test.json"
	os.Create(filepath)
	os.WriteFile(filepath, jsonData, 0777)
	jsonData2, _ := os.ReadFile(filepath)
	msg2 := &message{}
	msg2.UnmarshalJSON(jsonData2)
	msg2.Title = "aaa"
	fmt.Println("msg2", msg2)
	os.Create(filepath + "2")
	jsondata3, _ := msg2.Marshal()
	os.WriteFile(filepath+"2", jsondata3, 0777)

}
