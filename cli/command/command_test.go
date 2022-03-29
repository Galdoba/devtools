package command

import (
	"fmt"
	"os"
	"testing"
)

func Test_Command(t *testing.T) {
	tcmd, err := New(
		CommandLineArguments("ffprobe -v verbose -f lavfi -i amovie=d:\\\\IN\\\\IN_2022-03-28\\\\geroi_zakalennye_severnoy_shirotoy_2015__hd_rus20.m4a,asetnsamples=48000,astats=metadata=1:reset=1 -show_entries frame=pkt_pts_time:frame_tags=lavfi.astats.Overall.RMS_level,lavfi.astats.1.RMS_level,lavfi.astats.2.RMS_level,lavfi.astats.3.RMS_level,lavfi.astats.4.RMS_level,lavfi.astats.5.RMS_level,lavfi.astats.6.RMS_level -of csv=p=0"),
		Set(TERMINAL_ON),
		WriteToFile("d:\\IN\\IN_testInput\\log3.txt"),
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
	ex, err := os.Executable()
	fmt.Println("launch position:", ex, err)

	hn, err := os.Hostname()
	fmt.Println("Host Name:", hn, err)

}
