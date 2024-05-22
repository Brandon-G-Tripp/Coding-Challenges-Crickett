package cut

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func CutSecondField(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, "\t")
        if len(fields) >= 2 {
            fmt.Println(fields[1])
        }
    }

    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %v", err)
    }

    return nil
}
