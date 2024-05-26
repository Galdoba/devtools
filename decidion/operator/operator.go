package operator

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

/*
prompts:
Input
Select
Confirm ?


*/

type selector struct {
}

func noValidation(s string) error {
	return nil
}

func Input(label string, validator ...func(string) error) (string, error) {
	input := ""
	inputComponent := huh.NewInput()
	inputComponent.Title(label)
	inputComponent.Value(&input)
	if len(validator) == 0 {
		inputComponent.Validate(noValidation)
	} else {
		inputComponent.Validate(validator[0])
	}
	form := huh.NewForm(huh.NewGroup(inputComponent))
	err := form.Run()
	return strings.TrimSuffix(input, " "), err

}

func Confirm(label string) bool {
	conf := true
	inputComponent := huh.NewConfirm()
	inputComponent.Title(label)
	inputComponent.Value(&conf)

	form := huh.NewForm(huh.NewGroup(inputComponent))
	form.Run()
	return conf

}

func Select[T comparable](label string, options ...T) (T, error) {
	noval := new(T)
	if len(options) == 1 {
		return options[0], nil
	}
	if len(options) < 1 {
		return *noval, fmt.Errorf("no options to select from")
	}
	picked := -1
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
	return options[picked], err

}

func SelectMulti[T comparable](label string, limit int, options ...T) ([]T, error) {
	if len(options) < 1 {
		return nil, fmt.Errorf("nothing to select from")
	}
	if limit < 1 {
		return nil, fmt.Errorf("limit must be > 0")
	}
	if len(options) == 1 {
		return []T{options[0]}, nil
	}
	picked := []int{}

	selectComponent := huh.NewMultiSelect[int]()
	selectComponent = selectComponent.Title(label).Limit(limit)
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

	selected := []T{}
	for _, s := range picked {
		selected = append(selected, options[s])
	}
	return selected, err

}
