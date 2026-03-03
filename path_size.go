package code

import (
"fmt"
"os"
)

func GetSize(path string) (int64, error) {

info, err := os.Lstat(path)
if err !=nil {
return 0, err
}
if !info.IsDir() {
return info.Size(), nil
}

var totalSize int64
entries, err := os.ReadDir(path)
if err != nil {
return 0, err
}
for _, entry := range entries {
if entry.IsDir() {
continue
}

fileInfo, err := entry.Info()
if err != nil {
return 0, err
}
totalSize += fileInfo.Size()
}
return totalSize, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
size, err := GetSize(path)
if err != nil {
return "", err
}
if human {
return fmt.Sprintf("%s\t%s", FormatSize(size, true), path), nil
}
return fmt.Sprintf("%d\t%s", size, path), nil
}

func FormatSize(size int64, human bool) string {

if !human {
return fmt.Sprintf("%dB", size)
}
units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
if size < 1024 {
return fmt.Sprintf("%d%s", size, units[0])
}
value := float64(size)
unitIndex := 0
for value >= 1024 && unitIndex < len(units)-1 {
value /= 1024
unitIndex++
}
return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}
