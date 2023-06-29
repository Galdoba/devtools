package errmaker

import (
	"fmt"
	"runtime"
	"strings"
)

func Errorf(format string, args ...interface{}) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	return nil
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

// MyCaller returns the caller of the function that called it :)
func MyCaller() string {
	// Skip GetCallerFunctionName and the function to get the caller of

	return getFrame(2).Function
}

// bar is what we want to see in the output - it is our "caller"
func ErrorFrom(err error, args ...interface{}) error {
	str := "\n" + MyCaller() + "("
	for _, arg := range args {
		str += fmt.Sprintf("%v, ", arg)
	}
	str = strings.TrimSuffix(str, ", ") + "): "
	return fmt.Errorf("%v%v", str, err.Error())
}

func main() {
	err := someFunc(1, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func someFunc(a, b int) error {
	err := deepFunc(a, b, a+b)
	if err != nil {
		return ErrorFrom(err, a, b)
	}
	return nil
}

func deepFunc(args ...int) error {
	e := 0
	for _, a := range args {
		e += a
	}
	//e := c + d + 2
	err := Errorf("gave result %v", e)
	return ErrorFrom(err, args)
}
