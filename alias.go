package devtools

import (
	"bufio"
	"os"
	"strings"
)

//Сокращает имя по гибкому списку в файле. Если не находит знакомый префикс,
//то ничего не делает
func AliasByPrefixFromFile(filepath string, name string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return name
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		alias := scanner.Text()
		if strings.HasPrefix(name, alias) {
			return alias
		}
	}

	if err := scanner.Err(); err != nil {
		return name
	}
	return name
}
