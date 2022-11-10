package emdien

import (
	"fmt"
	"os"
	"path/filepath"

	// TODO: Parallelise indexing
	// "sync"

	"github.com/blevesearch/bleve/v2"
)

func createIndexIfDoesntExist(indexPath string) bleve.Index {
	if getDoesIndexExist(indexPath) {
		index, err := bleve.Open(indexPath)
		if err != nil {
			panic(err)
		}
		return index
	}
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexPath, mapping)
	if err != nil {
		panic(err)
	}
	return index
}

func indexWorker(
	// TODO: Parallelise indexing
	// waitGroup *sync.WaitGroup,
	index bleve.Index, filePath string,
) {
	if filepath.Ext(filePath) == ".md" {
		content, errReadFile := os.ReadFile(filePath)
		if errReadFile != nil {
			panic(errReadFile)
		}
		fmt.Println("Indexing", filePath)
		errIndex := index.Index(filePath, string(content))
		if errIndex != nil {
			panic(errIndex)
		}
	}
	// TODO: Parallelise indexing
	// waitGroup.Done()
}

func reIndex(index bleve.Index, cacheFolderPath string) bleve.Index {
	// Location to read data from
	dataFolderPath := filepath.Join(getRepoPath(cacheFolderPath), "files", "en-us")

	// Index the data in parallel
	fmt.Println("Indexing data.")
	// TODO: Parallelise indexing
	// var wg sync.WaitGroup
	errWalkDir := filepath.WalkDir(
		dataFolderPath,
		func(path string, dir_entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// TODO: Parallelise indexing
			// wg.Add(1)
			// go indexWorker(&wg, index, path)
			indexWorker(index, path)
			return nil
		},
	)
	// TODO: Parallelise indexing
	// wg.Wait() // Wait until all the workers have finished
	if errWalkDir != nil {
		panic(errWalkDir)
	}
	fmt.Println("Finished indexing data.")
	return index
}
