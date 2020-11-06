package code

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Galdoba/devtools/cli/user"
)

type dataFeed struct {
	structName string
	structMap  map[string]string
}

type testStruct struct {
	dataField1 int
	dataField2 []string
}

func (t *testStruct) DataField1() int {
	return t.dataField1
}

func (t *testStruct) SetDataField1(data int) {
	t.dataField1 = data
}

type methodConstructor interface {
	constructSetters()
	constructGetters()
}

func newDataFeed() (dataFeed, error) {
	df := dataFeed{}
	df.structMap = make(map[string]string)
	err := errors.New("No data")
	for err != nil {
		fmt.Print("Enter Struct Name: ")
		df.structName, err = user.InputStr()
	}
	err = errors.New("No data on fields. ")
	fmt.Print("Enter Struct fields data (blank slice to finish/[TAB] as separator): \n")
	for err != nil {
		fmt.Print("field ", len(df.structMap)+1, " --> ")
		sl, err := user.InputSliceStr("	")
		if err == nil {
			err = errors.New("There will be more")
		}
		if len(sl) > 2 {
			err = errors.New("Data unparseble (len != 2)")
			fmt.Print(err.Error() + "\n")
			continue
		}
		if len(sl) == 1 {
			err = errors.New("[TAB] must be used as separator")
			if sl[0] == "" {
				err = nil
				fmt.Print("[END INPUT]\n")
				fmt.Print("Constructing methods...\n")
				return df, err
			}
			fmt.Print(err.Error() + "\n")
			continue
		}
		if sl[1] == "" {
			err = errors.New("field type undetected")
			fmt.Print(err.Error() + "\n")
			continue
		}
		df.structMap[sl[0]] = sl[1]
	}
	return df, err
}

func newFile(fileName string) *os.File {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func addLineToFile(file string, newContent string) {
	lines := linesFromFile(file)
	lines = append(lines, newContent)
	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func linesFromFile(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func (df *dataFeed) constructDescription() {
	srtNme := string(byte(df.structName[0]))
	srtNme = strings.ToLower(srtNme)
	addLineToFile("output.txt", " ")
	addLineToFile("output.txt", "//"+df.structName+" - [SET DESCRIPTION OF OBJECT HERE]")
	addLineToFile("output.txt", "type "+df.structName+" struct {")
	for k, v := range df.structMap {
		addLineToFile("output.txt", "	"+k+"	"+v)
	}
	addLineToFile("output.txt", "}")
}

func (df *dataFeed) constructGetters() {
	srtNme := string(byte(df.structName[0]))
	srtNme = strings.ToLower(srtNme)
	for k, v := range df.structMap {
		field := k
		field = strings.Title(field)
		addLineToFile("output.txt", " ")
		addLineToFile("output.txt", "//"+field+" - returns "+srtNme+"."+k+" as a "+v)
		addLineToFile("output.txt", "func ("+srtNme+" *"+df.structName+") "+field+"() "+v+" {")
		addLineToFile("output.txt", "	return "+srtNme+"."+k)
		addLineToFile("output.txt", "}")
	}
}

func (df *dataFeed) constructSetters() {
	srtNme := string(byte(df.structName[0]))
	srtNme = strings.ToLower(srtNme)
	for k, v := range df.structMap {
		field := k
		field = strings.Title(field)
		addLineToFile("output.txt", " ")
		addLineToFile("output.txt", "//Set"+field+" - sets "+v+" value for "+srtNme+"."+k)
		addLineToFile("output.txt", "func ("+srtNme+" *"+df.structName+") Set"+field+"(data "+v+") {")
		addLineToFile("output.txt", "	"+srtNme+"."+k+" = data")
		addLineToFile("output.txt", "}")
	}
}

//ConstructStandardMethods - Создает output.txt в котором формирует структуру
// и все стандартные Геттеры и Сеттры с автоматическими коментариями
func ConstructStandardMethods() {
	df, err := newDataFeed()
	if err != nil {
		panic(err)
	}
	newFile("output.txt")
	df.constructDescription()
	df.constructGetters()
	df.constructSetters()
	fmt.Print("Construction done. Check 'output.txt' file in working directory.")
}
