package git

import "os"

func folderNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func createFolder(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func createFolderIfNotExists(path string) (bool, error) {
	if folderNotExists(path) {
		if err := createFolder(path); err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, nil
}
