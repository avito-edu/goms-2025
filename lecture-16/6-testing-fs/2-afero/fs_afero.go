package fs

import (
	"bufio"

	"github.com/spf13/afero"
)

func CountLines(fs afero.Fs, filename string) (int, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
