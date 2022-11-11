package emdien

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve/v2"
	"github.com/go-git/go-git/v5"
)

func updateLocalCache(index bleve.Index, cacheFolderPath string) {
	repoPath := getRepoPath(cacheFolderPath)
	repoWasCloned := cloneRepoIfDoesntExist(repoPath)

	// Pull the latest changes
	repoWasUpdated := false
	if !repoWasCloned {
		repoWasUpdated = updateRepo(repoPath)
	}

	// Re-index the data if there have been any changes
	if repoWasCloned || repoWasUpdated {
		reIndex(index, cacheFolderPath)
	}
}

func cloneRepoIfDoesntExist(repoPath string) bool {
	fmt.Println("Cloning MDN repository.")
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:        "http://github.com/mdn/content.git",
		Progress:   os.Stdout,
		RemoteName: "origin",
		// TODO: Decrease amount of data downloaded
		// ReferenceName: "main",
		// SingleBranch:  true,
		// Depth:         1,
	})
	if err == nil {
		fmt.Println("Cloned MDN repository successfully.")
		return true
	} else if err == git.ErrRepositoryAlreadyExists {
		fmt.Println("MDN repository already cloned.")
		return false
	} else {
		panic(err)
	}
}

func updateRepo(repoPath string) bool {
	fmt.Println("Updating MDN repository.")
	repository, errOpenRepo := git.PlainOpen(repoPath)
	if errOpenRepo != nil {
		panic(errOpenRepo)
	}
	workTree, errWorkTree := repository.Worktree()
	if errWorkTree != nil {
		panic(errWorkTree)
	}
	errPull := workTree.Pull(&git.PullOptions{
		Progress: os.Stdout,
		// TODO: Decrease amount of data downloaded
		// RemoteName:    "origin",
		// ReferenceName: "main",
		// SingleBranch:  true,
		// Depth:         1,
		// Force:         true,
	})
	if errPull == nil {
		fmt.Println("Updated MDN data.")
		return true
	} else if errPull == git.NoErrAlreadyUpToDate {
		fmt.Println("MDN data already up to date.")
		return false
	} else {
		panic(errPull)
	}
}
