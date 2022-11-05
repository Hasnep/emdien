package emdien

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/adrg/xdg"
	"github.com/blevesearch/bleve/v2"

	// index "github.com/blevesearch/bleve_index_api"
	"github.com/go-git/go-git/v5"
)

func run(args []string) {
	cache_dir := get_cache_dir()

	if do_update {
		do_the_update(cache_dir)
	}

	if len(args) == 0 {
		if do_update {
			// Not supplying a search term is fine if you're doing an update
			os.Exit(0)
		} else {
			fmt.Println("No search term or update. :(")
			os.Exit(1)
		}
	}
	if len(args) > 1 {
		fmt.Println("Too many search terms")
		os.Exit(1)
	}
	search_term := args[0]
	fmt.Println("You searched for ", search_term)

	index := create_index_if_not_exists(cache_dir)

	// search for some text
	query := bleve.NewMatchQuery("html")
	search := bleve.NewSearchRequest(query)
	search_result, err := index.Search(search)
	if err != nil {
		panic(err)
	}
	fmt.Println(search_result)
	// for _,search_result := range *search_results {

	// }

	// 	if render_markdown {

	// 		in := `# Hello World

	// 		This is a simple example of Markdown rendering with Glamour!
	// 		Check out the [other examples](https://github.com/charmbracelet/glamour/tree/master/examples) too.

	// 		Bye!
	// 		`

	// 		out, err := glamour.Render(in, "dark")
	// 		if err == nil {
	// 			fmt.Println(out)
	// 		}
	// 	}

}

func clone_repo_if_doesnt_exist(repo_path string) {
	_, err := git.PlainClone(repo_path, false, &git.CloneOptions{
		URL:      "http://github.com/mdn/content.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err == nil {
		fmt.Println("Cloned successfully.")
	} else if err == git.ErrRepositoryAlreadyExists {
	} else {
		fmt.Println("Failed to clone. :(")
		os.Exit(1)
	}
}

func get_repo_path(cache_dir string) string {
	return filepath.Join(cache_dir, "repo")
}

func get_index_path(cache_dir string) string {
	return filepath.Join(cache_dir, "index")
}

func do_the_update(cache_dir string) {
	repo_path := get_repo_path(cache_dir)
	clone_repo_if_doesnt_exist(repo_path)
	// repository, open_err := git.PlainOpen(cache_dir)
	// if open_err != nil {
	// 	fmt.Println("failed to open repo once cloned.")
	// }
	// work_tree, work_tree_err := repository.Worktree()
	// if work_tree_err != nil {
	// 	fmt.Println("Could not access files :(")
	// 	os.Exit(1)
	// }
	// pull_err := work_tree.Pull(&git.PullOptions{RemoteName: "origin"})
	// if pull_err == nil {
	// 	fmt.Println("Updated :)")
	// } else if pull_err == git.NoErrAlreadyUpToDate {
	// 	fmt.Println("Already up to date :)")
	// } else {
	// 	fmt.Println("Could not update :(")
	// 	fmt.Println(pull_err)
	// }
	reindex(cache_dir)
}

func create_index_if_not_exists(index_path string) bleve.Index {
	if get_does_index_exist(index_path) {
		index, err := bleve.Open("example.bleve")
		if err != nil {
			panic(err)
		}
		return index
	}
	mapping := bleve.NewIndexMapping()
	index, bleve_err := bleve.New(index_path, mapping)
	if bleve_err != nil {
		fmt.Println("error creating bleve: ", bleve_err)
		os.Exit(1)
	}
	return index
}

func index_worker(wg *sync.WaitGroup, index bleve.Index, file_path string) {
	if filepath.Ext(file_path) == ".md" {
		// fmt.Println("Indexing", file_path)
		content, read_err := ioutil.ReadFile(file_path)
		if read_err != nil {
			fmt.Println("reading error:", read_err)
			os.Exit(1)
		}
		index_err := index.Index(file_path, string(content))
		if index_err != nil {
			panic(index_err)
		}
	}
	wg.Done()
}

func reindex(cache_dir string) bleve.Index {
	index_path := get_index_path(cache_dir)
	index := create_index_if_not_exists(index_path)
	data_folder_path := filepath.Join(get_repo_path(cache_dir), "files", "en-us")
	// n_files, count_err := count_files_recursively(data_folder_path)
	// if count_err != nil {
	// 	panic(count_err)
	// }
	var wg sync.WaitGroup
	// wg.Add(n_files)
	walk_err := filepath.WalkDir(data_folder_path,
		func(path string, dir_entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			wg.Add(1)
			go index_worker(&wg, index, path)
			return nil
		})
	wg.Wait()
	if walk_err != nil {
		fmt.Println("errored :(")
		os.Exit(1)
	}
	return index
}

// func get_does_local_repo_exist(repo_path string) bool {
// 	return does_folder_exist(filepath.Join(repo_path, ".git"))
// }

func get_does_index_exist(index_path string) bool {
	return does_folder_exist(index_path)
}

func does_folder_exist(folder_path string) bool {
	_, err := os.Stat(folder_path)
	return !os.IsNotExist(err)
}

func get_docs_folder(cache_dir string, language string) string {
	return filepath.Join(cache_dir, language)
}

func get_cache_dir() string {
	return filepath.Join(xdg.CacheHome, "emdien")
}
