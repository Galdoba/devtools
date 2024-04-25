package configbuilder

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
)

func TestSourcePath(t *testing.T) {
	cb := New("yaml")

	err := cb.SetSourceDir(`C:\Users\Admin\go\src\github.com\Galdoba\ffstuff\app\mfrip\config\`)
	if err != nil {
		t.Errorf("met error: %v", err.Error())
	}
	fmt.Println(cb.sourcePath)
	fmt.Println(cb.app)
	fmt.Println(cb.pathToConfig)
	f := model.NewField(cb.language).WithSource("Source1").WithDataType("string").WithDesignation("Apppaaa").WithComment("This is a comment")
	f2 := model.NewField(cb.language).WithSource("Source2").WithDataType("string").WithDesignation("Apppaa  a  222").WithValue("key", "val")
	if err := cb.AddField(f); err != nil {
		t.Errorf("add field fail: %v", err.Error())

	}
	if err := cb.AddField(f2); err != nil {
		t.Errorf("add field fail: %v", err.Error())
	}
	text, err := cb.GenerateSource()
	fmt.Println(err)
	fmt.Println("===========================")
	fmt.Println(text)
	// f, _ := os.OpenFile(cb.sourcePath, os.O_CREATE|os.O_WRONLY, 0777)
	// defer f.Close()
	// f.WriteString(text)

}
