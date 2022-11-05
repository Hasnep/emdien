package emdien

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdn",
	Short: "MDN in the terminal",
	Long:  `Now you can search MDN and see the results in the terminal!`,
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
	Args: cobra.MaximumNArgs(1),
}

var do_update bool

func init() {
	rootCmd.Flags().BoolVarP(&do_update, "update", "u", false, "Update the mdn docs")
}

func Execute() error {
	return rootCmd.Execute()
}
