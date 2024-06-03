package handler

import (
	"fmt"
	"testing"
	"time"
)

func TestAllawance(t *testing.T) {
	j := pseudoJob()
	j.SetSchedule("* * * * *")
	if err := j.Save(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\cronex\handler\`); err != nil {
		fmt.Println(err.Error())
	}
	time.Sleep(time.Second)
	j = pseudoJob()
	j.SetSchedule("* * * * *")
	j.Save(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\cronex\handler\`)
	err := Start(`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\cronex\handler\`, Cycle(15))
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	// job, _ := job.Load(`C:\Users\Admin\go\src\github.com\Galdoba\devtools\cronex\cronex_job_20240529172050972.json`)
	// job.DoOnce = true
	// fmt.Println(job)
	// hnd := New(`C:\Users\Admin\go\src\github.com\Galdoba\devtools\cronex\`).With(
	// 	Cycle(10),
	// )
	// errsCount := 0
	// for {
	// 	if errsCount > 20 {
	// 		break
	// 	}
	// 	fi, err := os.ReadDir(hnd.storage)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		errsCount++
	// 	}
	// 	for _, f := range fi {
	// 		if f.IsDir() {
	// 			continue
	// 		}
	// 		jb, err := job.Load(hnd.storage + f.Name())
	// 		if err != nil {
	// 			fmt.Println(f.Name(), err.Error())
	// 			continue
	// 		}
	// 		fmt.Println("HANDLE:", jb.ID)
	// 		if err := hnd.Handle(jb); err != nil {
	// 			fmt.Println(err.Error())
	// 			errsCount++
	// 		}
	// 	}
	// 	fmt.Printf("sleep for %v seconds...\r", hnd.cycle)
	// 	time.Sleep(time.Second * time.Duration(hnd.cycle))
	// }
	// fmt.Println("HANDLE:")
	// err := hnd.Handle(job)
	// if err != nil {
	// 	fmt.Println("OUTERR:", err.Error())
	// }
}
