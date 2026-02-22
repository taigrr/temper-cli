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

	for _, t := range tempers {
		c, err := t.ReadC()
		if err != nil {
			return fmt.Errorf("reading temperature from %s: %w", t, err)
		}

		if fahrenheit {
			fmt.Printf("%.2f\n", c*9.0/5.0+32.0)
		} else {
			fmt.Printf("%.2f\n", c)
		}
	}

	return nil
}
