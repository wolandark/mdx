package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// general flags
	mangaurl string
	mangaId  string
)

var (
	versionApp bool
	versionAPI bool

	rootCmd = &cobra.Command{
		Use:   "mdx",
		Short: "manga downloader from MangaDex website",
		Long: `mdx is a command-line interface program for downloading manga from the MangaDex - https://mangadex.org .
The program uses MangaDex API (https://api.mangadex.org/docs) to fetch manga content.`,
		Run: func(cmd *cobra.Command, args []string) {
			if versionApp {
				fmt.Println(MDX_APP_VERSION)
				os.Exit(0)
			}

			if versionAPI {
				fmt.Println(MANGADEX_API_VERSION)
				os.Exit(0)
			}

			cmd.Help()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	rootCmd.Flags().BoolP("help", "h", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&versionApp, "version", "v", false, "version of application")
	rootCmd.Flags().BoolVarP(&versionAPI, "version-api", "a", false, "version of API")
}
