package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arimatakao/mdx/mangadexapi"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	findCmd = &cobra.Command{
		Use:     "find",
		Aliases: []string{"f", "search"},
		Short:   "Find manga",
		Long:    "Search and print manga info. Sort by revelance asceding. Best results will be down",
		Run:     find,
	}
	title string
)

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().StringVarP(&title,
		"title", "t", "", "specifies the title of the manga to search for")

	findCmd.MarkFlagRequired("title")
}

func find(cmd *cobra.Command, args []string) {
	c := mangadexapi.NewClient("")

	spinner, _ := pterm.DefaultSpinner.Start("Searching manga...")

	response, err := c.Find(title, "25", "0")
	if err != nil {
		spinner.Fail("Failed to search manga")
		fmt.Printf("error while search manga: %v\n", err)
		os.Exit(1)
	}

	if response.Total == 0 {
		spinner.Warning("Nothing found...")
		os.Exit(0)
	}

	spinner.Success("Manga found!")
	fmt.Printf("\nTotal found: %d\n", response.Total)

	for _, m := range response.Data {
		fmt.Println("------------------------------")
		printMangaInfo(m)
	}

	printedCount, _ := strconv.Atoi("25")
	if response.Total > printedCount {
		fmt.Println("==============================")
		fmt.Printf("\nFull results (%d): https://mangadex.org/search?q=%s\n",
			response.Total, title)
	}
}
