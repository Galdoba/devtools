package schedule

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/devtools/datastruct/slice"
)

const (
	SECOND = iota
	MINUTE
	HOUR
	DayOfMonth
	MONTH
	DayOfWeek
)

type blockIndex struct {
	index     map[int]int
	allowance map[int][]int
}

type Schedule interface {
	Allowance(int) []int
}

func TimeValid(sched Schedule) bool {
	now := time.Now().Local()
	if !slice.HasElement(sched.Allowance(SECOND), now.Second()) {
		return false
	}
	if !slice.HasElement(sched.Allowance(MINUTE), now.Minute()) {
		return false
	}
	if !slice.HasElement(sched.Allowance(HOUR), now.Hour()) {
		return false
	}
	if !slice.HasElement(sched.Allowance(DayOfMonth), now.Day()) {
		return false
	}
	if !slice.HasElement(sched.Allowance(MONTH), int(now.Month())) {
		return false
	}
	if !slice.HasElement(sched.Allowance(DayOfWeek), int(now.Weekday())) {
		return false
	}
	return true

}

func newBlockIndex(i int) (*blockIndex, error) {
	bi := blockIndex{}
	bi.index = make(map[int]int)
	switch i {
	default:
		return nil, fmt.Errorf("invalid number of blocks")
	case 5:
		for j, block := range []int{MINUTE, HOUR, DayOfMonth, MONTH, DayOfWeek} {
			bi.index[block] = j
		}
	case 6:
		for j, block := range []int{SECOND, MINUTE, HOUR, DayOfMonth, MONTH, DayOfWeek} {
			bi.index[block] = j
		}
	}
	bi.allowance = make(map[int][]int)
	return &bi, nil
}

func index(bi *blockIndex, block int) int {
	if i, ok := bi.index[block]; ok {
		return i
	}
	return -1
}

func (bi *blockIndex) Allowance(block int) []int {
	index := index(bi, block)
	if index == -1 {
		if block == SECOND {
			return allNumbers(0, 59)
		}
		return []int{}
	}
	return bi.allowance[index]
}

func New(sch string) (Schedule, error) {
	min, max := 0, 0
	err := fmt.Errorf("not parsed")
	parts := strings.Split(sch, " ")
	bi, err := newBlockIndex(len(parts))
	if err != nil {
		return nil, err
	}
	for i, k := range bi.index {
		switch i {
		default:
			return nil, fmt.Errorf("unknown block type met")
		case SECOND, MINUTE:
			min, max = 0, 59
		case HOUR:
			min, max = 0, 23
		case DayOfMonth:
			min, max = 1, 31
		case MONTH:
			min, max = 1, 12
		case DayOfWeek:
			min, max = 0, 7
		}
		allowance, err := parseBlock(parts[k], min, max)
		if err != nil {
			return nil, fmt.Errorf("block %v parsing failed: %v", k, err)
		}
		bi.allowance[k] = allowance
	}

	return bi, nil
}

func parseBlock(s string, min, max int) ([]int, error) {
	if s == "*" {
		return allNumbers(min, max), nil
	}
	if s == "" {
		return []int{}, fmt.Errorf("no data in block")
	}
	mm := []int{}
	segm := strings.Split(s, ",")
	for _, seg := range segm {
		if strings.Contains(seg, "/") {
			sp, err := splitDevision(seg)
			if err != nil {
				return mm, fmt.Errorf("parsing failed: %v", err.Error())
			}
			mm = append(mm, sp...)
			continue
		}
		if strings.Contains(seg, "-") {
			sp, err := splitOrderly(seg)
			if err != nil {
				return mm, fmt.Errorf("parsing failed: %v", err.Error())
			}
			mm = append(mm, sp...)
			continue
		}
		if val, err := strconv.Atoi(seg); err == nil {
			mm = append(mm, val)
			continue
		}
		return mm, fmt.Errorf("parsing failed: %v", s)
	}
	minutes, err := checkVals(mm, min, max)
	if err != nil {
		return mm, fmt.Errorf("parsing failed: %v", err.Error())
	}
	return minutes, nil
}

func splitOrderly(s string) ([]int, error) {
	sp := []int{}
	data := strings.Split(s, "-")
	if len(data) != 2 {
		return sp, fmt.Errorf("'%v': expect to have 2 numbers", s)
	}
	min, err := strconv.Atoi(data[0])
	if err != nil {
		return sp, fmt.Errorf("%v: first number: %v", s, err.Error())
	}
	max, err2 := strconv.Atoi(data[1])
	if err2 != nil {
		return sp, fmt.Errorf("%v: second number: %v", s, err2.Error())
	}
	for i := min; i <= max; i++ {
		sp = append(sp, i)
	}
	return sp, nil
}

func splitDevision(s string) ([]int, error) {
	sp := []int{}
	data := strings.Split(s, "/")
	if len(data) != 2 {
		return sp, fmt.Errorf("'%v': expect to have 2 numbers", s)
	}
	dev, err := strconv.Atoi(data[0])
	if err != nil {
		return sp, fmt.Errorf("'%v': first number: %v", s, err.Error())
	}
	dvtr, err2 := strconv.Atoi(data[1])
	if err2 != nil || dvtr == 0 {
		return sp, fmt.Errorf("'%v': can't devide by zero", s)
	}
	for i := 0; i < dev; i++ {
		if i%dvtr == 0 {
			sp = append(sp, i)
		}
	}
	return sp, nil
}

func allNumbers(min, max int) []int {
	sl := []int{}
	for i := min; i <= max; i++ {
		sl = append(sl, i)
	}
	return sl
}

func checkVals(sl []int, min, max int) ([]int, error) {
	formatted := []int{}
	if len(sl) < 1 {
		return formatted, fmt.Errorf("no values parsed")
	}
	sort.Ints(sl)
	if sl[0] < min {
		return formatted, fmt.Errorf("bad values parsed: %v", sl[0])
	}
	if sl[len(sl)-1] > max {
		return formatted, fmt.Errorf("bad values parsed: %v", sl[len(sl)-1])
	}
	for _, i := range sl {
		formatted = appendUnique(formatted, i)
	}
	return formatted, nil
}

func appendUnique(sl []int, elem int) []int {
	for _, i := range sl {
		if i == elem {
			return sl
		}
	}
	return append(sl, elem)
}

func AllBlocks() []int {
	return []int{
		SECOND,
		MINUTE,
		HOUR,
		DayOfMonth,
		MONTH,
		DayOfWeek,
	}
}
