package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"github.com/taigrr/temper"
)

var version = "dev" // overridable via -ldflags

type commandOptions struct {
	fahrenheit bool
	kelvin     bool
	jsonOutput bool
}

func init() {
	if version != "dev" {
		return
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			version = v
		}
	}
}

func main() {
	cmd := newRootCommand()
	if err := fang.Execute(context.Background(), cmd); err != nil {
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	opts := commandOptions{}

	cmd := &cobra.Command{
		Use:     "temper-cli",
		Short:   "Read temperature from TEMPer USB sensors",
		Long:    "temper-cli discovers TEMPer USB HID temperature sensors and prints the current reading.",
		Version: version,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		SilenceUsage: true,
	}

	cmd.Flags().BoolVarP(&opts.fahrenheit, "fahrenheit", "f", false, "output temperature in Fahrenheit")
	cmd.Flags().BoolVarP(&opts.kelvin, "kelvin", "k", false, "output temperature in Kelvin")
	cmd.Flags().BoolVarP(&opts.jsonOutput, "json", "j", false, "output readings as JSON")
	cmd.MarkFlagsMutuallyExclusive("fahrenheit", "kelvin")

	return cmd
}

func run(opts commandOptions) error {
	tempers, err := temper.FindTempers()
	if err != nil {
		return fmt.Errorf("finding temper devices: %w", err)
	}

	if len(tempers) == 0 {
		return fmt.Errorf("no temper devices found")
	}

	defer func() {
		for _, t := range tempers {
			if closeErr := t.Close(); closeErr != nil {
				fmt.Fprintf(os.Stderr, "closing %s: %v\n", t, closeErr)
			}
		}
	}()

	unit := "celsius"
	if opts.fahrenheit {
		unit = "fahrenheit"
	} else if opts.kelvin {
		unit = "kelvin"
	}

	var readings []Reading
	for _, t := range tempers {
		celsius, readErr := t.ReadC()
		if readErr != nil {
			return fmt.Errorf("reading temperature from %s: %w", t, readErr)
		}
		readings = append(readings, NewReading(t.String(), float64(celsius)))
	}

	if opts.jsonOutput {
		return FormatReadingsJSON(os.Stdout, readings)
	}

	showLabels := len(readings) > 1
	for _, reading := range readings {
		if showLabels {
			if err := FormatLabeledReading(os.Stdout, reading, unit); err != nil {
				return fmt.Errorf("writing labeled reading for %s: %w", reading.Device, err)
			}
			continue
		}
		if err := FormatReading(os.Stdout, reading, unit); err != nil {
			return fmt.Errorf("writing reading for %s: %w", reading.Device, err)
		}
	}

	return nil
}
