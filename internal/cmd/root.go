package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "txtpix",
	Short: "A simple image tool utility",
	Long: `
txtpix turns images into text art and files into base64, from the terminal.

Algorithms:
  base64         encode any file to base64 text (default)
  base64_decode  decode a base64 file back to text
  ascii          image to ASCII art, brightness mapped to a character ramp
  unicode        image to Unicode art using half-block characters
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
