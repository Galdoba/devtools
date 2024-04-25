package cmd

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
)

func userInput(q string, defaults ...string) string {
	var out string
	for _, v := range defaults {
		out = v
	}
	form := huh.NewForm(
		// Gather some final details about the order.
		huh.NewGroup(
			huh.NewInput().
				Title(q).
				Value(&out).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					if str == "Frank" {
						return errors.New("Sorry, we don’t serve customers named Frank.")
					}
					return nil
				}),
		),
	)
	err := form.Run()
	if err != nil {
		panic("fatal err: " + err.Error())
	}
	return out
}

func userSelect(question string, options ...string) string {
	var picked int
	selectComponent := huh.NewSelect[int]()
	selectComponent = selectComponent.Title(question)
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
	return options[picked]
}

func userConfirm(text string) bool {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(text).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm)),
	)
	err := form.Run()
	if err != nil {
		panic("fatal err: " + err.Error())
	}
	return confirm
}
