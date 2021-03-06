package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Reporter interface {
	Print(format string, args ...interface{})
	Println(format string, args ...interface{})
}

func Stdout() Reporter {
	return &Printer{
		w: os.Stdout,
	}
}

func Stderr() Reporter {
	return &Printer{
		w: os.Stderr,
	}
}

type Printer struct {
	w io.Writer
}

func (p *Printer) Print(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(p.w, format, args...)
}

func (p *Printer) Println(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	_, _ = fmt.Fprintf(p.w, "[%s] %s\n", time.Now().Format(time.RFC3339), msg)
}

const fallbackNextStep = "Press any key to continue..."

func PromptPressAnyKeyToContinue(nextStep string) error {
	if nextStep == "" {
		return nil
	}
	Stdout().Println("Next: %s. %s", nextStep, fallbackNextStep)
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadRune()
	return err
}
