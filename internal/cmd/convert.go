package cmd

import (
	"fmt"

	"github.com/exec-cmd/txtpix/internal/engine"
	"github.com/exec-cmd/txtpix/internal/scibe"
	"github.com/spf13/cobra"
)

var algorithm string
var outputPath string

var cnvrtCmd = &cobra.Command{
	Short: "Convert a file using the selected algorithm",
	Long:  "Reads the input file, converts its contents using the selected algorithm, and prints the result",

	Use:          "convert <file>",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),

	RunE: runE,
}

func runE(cmd *cobra.Command, args []string) error {
	if _, ok := engine.Algs[algorithm]; !ok {
		return fmt.Errorf("unknown algorithm %q", algorithm)
	}

	res, err := engine.ConvertFile(args[0], algorithm)
	if err != nil {
		return err
	}

	fmt.Println(res)

	if outputPath == "" {
		return nil
	}

	if err := scibe.WriteStrInImage(outputPath, res); err != nil {
		return err
	}
	return nil
}

func init() {
	cnvrtCmd.Flags().StringVarP(&algorithm, "alg", "a", "base64", "Convert algorithm")
	cnvrtCmd.Flags().StringVarP(&outputPath, "out", "o", "", "Output file")

	rootCmd.AddCommand(cnvrtCmd)
}
