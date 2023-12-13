package download

import (
	"fmt"
	"github.com/spf13/cobra"
)

var url string

var Root = &cobra.Command{
	Use:   "download",
	Short: "start download a url",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" starting download")
	},
}

func init() {
	Root.Flags().StringVarP(&url, "url", "u", "", "url downloading")
}
