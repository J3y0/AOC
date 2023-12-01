package main

import (
	"bufio"
    "fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
    TotalSize int
    FileSize int
    DirNames []string
}

// Problem with same dir name in 2 diff dirs -> need to save more than just one dir as key -> path with parent_dir ?

func ComputeTotalSize(key string, dir map[string]Directory) int {
    var sum int
    if len(dir[key].DirNames) == 0 {
        return dir[key].FileSize
    } else {
        for _, k := range dir[key].DirNames {
            sum += ComputeTotalSize(k, dir)
        }
    }
    return sum + dir[key].FileSize
}

func RemoveLast(slice []string) []string {
    return slice[:len(slice)-1]
}

func ParseInput(r io.Reader) (map[string]Directory, error) {
    directories := make(map[string]Directory)
    pwd := make([]string, 0)
    sc := bufio.NewScanner(r)

    var readLs bool
    var curDir string
    for sc.Scan() {
        line := sc.Text()
        if strings.HasPrefix(line, "$ cd") {
            dir := strings.Split(line, " ")[2]
            readLs = false
            if dir == ".." {
                pwd = RemoveLast(pwd)
            } else {
                curDir = ""
                for _, elt := range pwd {
                    curDir += elt + "_"
                }
                curDir += dir
                pwd = append(pwd, dir)
                directories[curDir] = Directory{}
            }
        }

        if strings.HasPrefix(line, "$ ls") {
            readLs = true
        }

        if readLs {
            splitted := strings.Split(line, " ")
            if splitted[0] == "dir" {
                if entry, ok := directories[curDir]; ok {
                    entry.DirNames = append(entry.DirNames, curDir + "_" + splitted[1])
                    directories[curDir] = entry
                }
            } else {
                if entry, ok := directories[curDir]; ok {
                    size, _ := strconv.Atoi(splitted[0])
                    entry.FileSize += size
                    directories[curDir] = entry
                }
            }
        }
    }
    if err := sc.Err(); err != nil {
        log.Fatalf("scan file error: %v", err)
        return nil, err
    }

    // Compute total size, as we now have every child directory
    for k := range directories {
        if entry, ok := directories[k]; ok {
            entry.TotalSize = ComputeTotalSize(k, directories)
            directories[k] = entry
        }
    }
    return directories, nil
}

func SolvePart1(directories map[string]Directory) int {
    var result int
    const maxSize int = 100000
    for _, d := range directories {
        if d.TotalSize <= maxSize {
            result += d.TotalSize
        }
    }
    return result
}

func SolvePart2(directories map[string]Directory) int {
    minToDelete := 30000000 -  (70000000 - directories["/"].TotalSize)
    minCandidates := 70000000
    for _, d := range directories {
        if d.TotalSize >= minToDelete && d.TotalSize < minCandidates {
            minCandidates = d.TotalSize
        }
    }
    return minCandidates
}

func main() {
    file, err := os.Open("./data/day7.txt")
    if err != nil {
        log.Fatalln("Error while opening the file")
        os.Exit(1)
    }
    defer file.Close()

    var directories map[string]Directory
    directories, err = ParseInput(file)
    if err != nil {
        log.Fatalln("Error while parsing the file")
        os.Exit(1)
    }

    // log.Printf("%v", directories)
    
    part1 := SolvePart1(directories)
    fmt.Printf("Part 1 result: %d\n", part1)

    part2 := SolvePart2(directories)
    fmt.Printf("Part 2 result: %d\n", part2)
    os.Exit(0)
}
