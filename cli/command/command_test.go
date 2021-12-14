package command

import (
	"fmt"
	"testing"
)

func Test_Command(t *testing.T) {
	tcmd, err := New(
		CommandLineArguments("ffmpeg -i d:\\MUX\\tests\\s05e01_Rostelecom_FLASH_YR05_18_19_NORA_16x9_STEREO_5_1_2_0_LTRT_EPISODE_E2291774_RUSSIAN_ENGLISH_10750107.mpg"),
		Set(BUFFER_ON),
		WriteToFile("d:\\MUX\\tests\\log2.txt"),
	)
	if err != nil {
		t.Errorf("func New() error: %v", err.Error())
	}
	if err := tcmd.Run(); err != nil {
		t.Errorf("func Run() error: %v", err.Error())
	}

	fmt.Println("///")
	fmt.Println("OUT")
	fmt.Println(tcmd.StdOut())
	fmt.Println("///")
	fmt.Println("ERR")
	fmt.Println(tcmd.stErr)

}
