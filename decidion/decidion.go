package decidion

/*
prompts:
Input
Select
Confirm ?


*/

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
)

type Selector interface {
	Select(opts ...interface{}) (int, error)
}

//One - select one from slice
func One[T comparable](label string, dice Selector, manual bool, options ...T) T {
	if len(options) == 1 {
		return options[0]
	}
	if len(options) < 1 {
		panic("no options to select from")
	}
	picked := -1

	switch manual {
	case true:
		selectComponent := huh.NewSelect[int]()
		selectComponent = selectComponent.Title(label)
		opts := []huh.Option[int]{}
		for i, opt := range options {
			key := fmt.Sprintf("%v", opt)
			opt := huh.NewOption[int](key, i)
			opts = append(opts, opt)
		}
		selectComponent = selectComponent.Options(opts...)
		selectComponent = selectComponent.Value(&picked)
		form := huh.NewForm(huh.NewGroup(selectComponent))
		err := form.Run()
		if err != nil {
			panic(err.Error())
		}
	case false:
		err := errors.New("not selected")
		picked, err = autoPickGen[T](options, dice)
		if err != nil {
			panic(err.Error())
		}
	}
	if picked == -1 {
		panic("can't auto pick")
	}
	return options[picked]
}

//One_Exclude - select one from slice AND return list of leftover options
func One_Exclude[T comparable](label string, dice Selector, manual bool, options ...T) (T, []T) {
	selected := One[T](label, dice, manual, options...)
	leftOver := excludeType(options, selected)
	return selected, leftOver

}

//Few - select options n times from slice
func Few[T comparable](n int, label string, dice Selector, manual bool, options ...T) []T {
	if len(options) < 1 {
		panic("no options to select from")
	}
	selected := []T{}
	currentLabel := label
	for len(selected) < n {
		selected = append(selected, One[T](currentLabel, dice, manual, options...))
		currentLabel = ""
		if len(selected) > 0 {
			for _, a := range selected {
				currentLabel += fmt.Sprintf("%v\n", a)
			}
		}
		currentLabel = fmt.Sprintf("%v%v", currentLabel, label)
	}
	return selected
}

//Few_Exclude - select n options from slice AND return leftover options
func Few_Exclude[T comparable](n int, label string, dice Selector, manual bool, options ...T) ([]T, []T) {
	if n <= 0 {
		return nil, options
	}
	if len(options) < 1 {
		panic("no options to select from")
	}
	if len(options) < n {
		return options, nil
	}
	selected := []T{}
	currentLabel := label
	for len(selected) < n {
		selectedNow := One[T](currentLabel, dice, manual, options...)
		selected = append(selected, selectedNow)
		currentLabel = ""
		if len(selected) > 0 {
			for _, a := range selected {
				currentLabel += fmt.Sprintf("%v\n", a)
			}
		}
		currentLabel = fmt.Sprintf("%v%v", currentLabel, label)
		options = excludeType(options, selectedNow)
		//options = leftOver
	}
	return selected, options
}

///////////////////

func autoPickGen[T comparable](sl []T, dice Selector) (int, error) {
	return dice.Select(sl)
}

func excludeType[T comparable](sl []T, elem T) []T {
	leftover := []T{}
	for i, val := range sl {
		if val == elem {
			continue
		}
		leftover = append(leftover, sl[i])
	}
	return leftover
}
