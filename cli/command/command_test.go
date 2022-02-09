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

func Test_Command_untested(t *testing.T) {
	path := "d:\\IN\\IN_2022-02-07\\Eiffel_AUDIORUS51.m4a"
	tcmd, err := New(
		CommandLineArguments(fmt.Sprintf("loudnorm %v -scan", path)),
		Set(BUFFER_ON),
		WriteToFile("d:\\IN\\IN_2022-02-07\\Eiffel_AUDIORUS51.txt"),
	)
	if err != nil {
		t.Errorf("func New() error: %v", err.Error())
	}
	if err := tcmd.Run_untested(); err != nil {
		t.Errorf("func Run() error: %v", err.Error())
	}

	fmt.Println("///")
	fmt.Println("OUT")
	fmt.Println(tcmd.StdOut())
	fmt.Println("///")
	fmt.Println("ERR")
	fmt.Println(tcmd.StdErr())

}
