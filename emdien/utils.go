package emdien

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

// Functions to construct data paths

func getCachePath() string {
	return filepath.Join(xdg.CacheHome, "emdien")
}

func getRepoPath(cacheFolderPath string) string {
	return filepath.Join(cacheFolderPath, "repo")
}

func getIndexPath(cacheFolderPath string) string {
	return filepath.Join(cacheFolderPath, "index")
}

// Checking folder existence

func getDoesFolderExist(folderPath string) bool {
	_, err := os.Stat(folderPath)
	return !os.IsNotExist(err)
}

func getDoesIndexExist(indexPath string) bool {
	return getDoesFolderExist(indexPath)
}
