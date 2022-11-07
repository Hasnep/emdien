package emdien

import (
	"fmt"
	"os"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdn",
	Short: "MDN in the terminal",
	Long:  "Search and read MDN locally in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if update {
			return cobra.ArbitraryArgs(cmd, args)
		}
		return cobra.MinimumNArgs(1)(cmd, args)
	},
}

// CLI flags
var (
	update    bool
	rawOutput bool
)

func init() {
	// Add CLI flags
	rootCmd.Flags().BoolVarP(&update, "update", "u", false, "Update the local MDN data.")
	rootCmd.Flags().BoolVarP(&rawOutput, "raw-output", "r", false, "Output raw markdown instead of pretty rendered output.")
}

func Execute() error {
	return rootCmd.Execute()
}

func run(args []string) {
	cacheFolderPath := getCachePath()
	fmt.Println(cacheFolderPath)

	index := createIndexIfDoesntExist(getIndexPath(cacheFolderPath))

	if update {
		updateLocalCache(index, cacheFolderPath)
		os.Exit(0)
	}

	nDocumentsInIndex, errDocCount := index.DocCount()
	if errDocCount != nil {
		panic(errDocCount)
	}
	if nDocumentsInIndex == 0 {
		fmt.Println("No MDN documents could be found in the index, consider running `mdn --update` to update the index.")
		os.Exit(1)
	}

	searchTerm := strings.Join(args, " ")
	query := bleve.NewMatchQuery(searchTerm)
	search := bleve.NewSearchRequest(query)
	search.SortBy([]string{"-_score"})
	searchResult, err := index.Search(search)
	if err != nil {
		panic(err)
	}
	firstResult := searchResult.Hits[0]
	content, errReadFile := os.ReadFile(firstResult.ID)
	if errReadFile != nil {
		panic(errReadFile)
	}
	contentStr := string(content)
	if rawOutput {
		fmt.Println(contentStr)
	} else {
		renderedContent, errRender := glamour.Render(contentStr, "dark")
		if errRender != nil {
			panic(errRender)
		}
		fmt.Println(renderedContent)
	}
}
