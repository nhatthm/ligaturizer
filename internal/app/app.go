// Package app provides functionalities to run the application.
package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"go.nhat.io/ligaturizer/internal/python3"
)

// Run runs the application.
func Run() {
	defer python3.Finalize()

	if err := rootCommand().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, color.HiRedString("%s", err))
	}
}
