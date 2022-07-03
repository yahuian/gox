package filex

import (
	"errors"
	"io/fs"
	"os"
)

// Exist check if a file or dir exist. when error isn't nil, bool value is unreliable don't use it!
// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func Exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	// maybe file exist but permission denied, disk damage
	return false, err
}
