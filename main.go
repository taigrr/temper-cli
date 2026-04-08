package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"github.com/taigrr/temper"
)

var (
	version    = "dev"
	fahrenheit bool
	kelvin     bool
	jsonOutput bool
)

func main() {
	cmd := &cobra.Command{
		Use:     "temper-cli",
		Short:   "Read temperature from TEMPer USB sensors",
		Long:    "temper-cli discovers TEMPer USB HID temperature sensors and prints the current reading.",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		SilenceUsage: true,
	}

	cmd.Flags().BoolVarP(&fahrenheit, "fahrenheit", "f", false, "output temperature in Fahrenheit")
	cmd.Flags().BoolVarP(&kelvin, "kelvin", "k", false, "output temperature in Kelvin")
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "output readings as JSON")
	cmd.MarkFlagsMutuallyExclusive("fahrenheit", "kelvin")

	if err := fang.Execute(context.Background(), cmd); err != nil {
		os.Exit(1)
	}
}

func run() error {
	tempers, err := temper.FindTempers()
	if err != nil {
		return fmt.Errorf("finding temper devices: %w", err)
	}

	if len(tempers) == 0 {
		return fmt.Errorf("no temper devices found")
	}

	defer func() {
		for _, t := range tempers {
			t.Close()
		}
	}()

	unit := "celsius"
	if fahrenheit {
		unit = "fahrenheit"
	} else if kelvin {
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

	if jsonOutput {
		return FormatReadingsJSON(os.Stdout, readings)
	}

	for _, reading := range readings {
		FormatReading(os.Stdout, reading, unit)
	}

	return nil
}
