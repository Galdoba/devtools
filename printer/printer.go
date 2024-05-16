package printer

import "os"

type printManager struct {
	writerOut       *os.File
	writerErr       *os.File
	writersOther    []*os.File
	alertLevel      int
	timeStampPrefix bool
	caller          string
}

/*
pm := printer.New()
pm.WithWriters(os.StdErr)
pm.ShoutIf(INFO)
pm.PrintfAs(printer.INFO,"%v\n", arg)

*/
