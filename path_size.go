package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func sizeToString(i64 int64, human bool) string {
	hSize := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	var unit float64 = 1024

	if human && i64 > int64(unit) {
		e := math.Floor(math.Log(float64(i64)) / math.Log(unit))
		val := float64(i64) / math.Pow(unit, e)

		return fmt.Sprintf("%.1f%s", val, hSize[int(e)])
	}

	return fmt.Sprintf("%d%s", i64, hSize[0])
}

func isHiddenPath(fileName string) bool {
	return strings.HasPrefix(fileName, ".")
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	path, err := filepath.Abs(path)

	if err != nil {
		return "", err
	}

	baseInfo, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !all && isHiddenPath(baseInfo.Name()) {
		return "", fmt.Errorf("file not found")
	}

	if baseInfo.IsDir() {
		total, err := calcRecursiveDir(path, recursive, all)

		if err != nil {
			return "", err
		}

		return sizeToString(total, human), nil
	}

	return sizeToString(baseInfo.Size(), human), nil
}

// calcRecursiveDir - calculate size of files in derictory and, if recursive flag setted, in subdirectroies
func calcRecursiveDir(path string, recursive bool, all bool) (int64, error) {
	dir, err := os.ReadDir(path)

	if err != nil {
		return 0, err
	}

	var total int64

	for _, f := range dir {
		newPath := filepath.Join(path, f.Name())
		fileInfo, err := os.Lstat(newPath)

		if !all && isHiddenPath(fileInfo.Name()) {
			continue
		}

		if err != nil {
			return 0, err
		}

		if recursive && fileInfo.IsDir() {
			cTotal, err := calcRecursiveDir(newPath, recursive, all)
			if err != nil {
				return 0, err
			}
			total += cTotal
		} else {
			total += fileInfo.Size()

		}
	}

	return total, nil
}
