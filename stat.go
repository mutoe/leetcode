package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var solutionUrlPrefix = "https://leetcode.com/problems/"

type Solution struct {
	id              int
	titleSlug       string
	timeComplexity  string
	timeSpent       string
	ranking         string
	spaceComplexity string
	spaceSpent      string
}

func main() {
	solutions := readSolutions(".")
	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i].id < solutions[j].id
	})
	updateSolutionsInReadme(solutions)
}

func readSolutions(dir string) (solutions []Solution) {
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	for {
		fileInfos, err := file.Readdir(10000)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("readdir error:", err)
			return
		}

		if len(fileInfos) == 0 {
			break
		}

		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() {
				reg := regexp.MustCompile(`^(\d+)_([a-z0-9_]+)$`)
				submatch := reg.FindStringSubmatch(fileInfo.Name())
				if submatch == nil {
					continue
				}

				id, _ := strconv.Atoi(submatch[1])
				solution, err := readFile(id, submatch[2])
				if err != nil {
					continue
				}
				solutions = append(solutions, solution)
			}
		}
		break
	}

	return
}

func readFile(id int, fileName string) (solution Solution, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(id, err)
		}
	}()

	filePath := fmt.Sprintf("./%d_%s/%s.go", id, fileName, fileName)
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("warning readfile: ", err)
		return solution, err
	}
	solution.id = id
	solution.titleSlug = fileName
	content := string(b)
	timeReg, _ := regexp.Compile(`// time: (O\(.+\)) ?(\d+)ms ?([\d.]+)%`)
	spaceReg, _ := regexp.Compile(`// space: (O\(.+\)) ?([\d.]+)M`)
	timeMatches := timeReg.FindStringSubmatch(content)
	spaceMatches := spaceReg.FindStringSubmatch(content)
	if len(timeMatches) == 0 || len(spaceMatches) == 0 {
		fmt.Println("Unresolved problem: ", id, "\t", fileName)
		return solution, errors.New("unresolved problem")
	}
	solution.timeComplexity = timeMatches[1]
	solution.timeSpent = timeMatches[2]
	solution.ranking = timeMatches[3]
	solution.spaceComplexity = spaceMatches[1]
	solution.spaceSpent = spaceMatches[2]
	return solution, nil
}

func updateSolutionsInReadme(solutions []Solution) {
	readmeFilePath := "./README.md"
	b, _ := ioutil.ReadFile(readmeFilePath)
	content := string(b)
	reg := regexp.MustCompile(`<!-- golang inject solutions start -->[\S\s]*<!-- golang inject solutions end -->`)
	replacer := `<!-- golang inject solutions start -->

|ID|Title|Time| |Space|Ranking|
|---:|--|--:|:--|:--|--:|
`

	for _, s := range solutions {
		title := regexp.MustCompile(`_`).ReplaceAllString(s.titleSlug, " ")
		titleSlug := regexp.MustCompile(`_`).ReplaceAllString(s.titleSlug, "-")
		link := solutionUrlPrefix + titleSlug
		rank, _ := strconv.ParseFloat(s.ranking, 0)
		var rankString string
		if rank >= 90.0 {
			rankString = fmt.Sprintf(`%g%% 🟢`, rank)
		} else if rank >= 50.0 {
			rankString = fmt.Sprintf(`%g%% 🟠`, rank)
		} else {
			rankString = fmt.Sprintf(`%g%% 🔴`, rank)
		}
		replacer += fmt.Sprintf("| %d\t| [%s](%s)\t| %sms\t| %s\t| %s\t| %s\t|\n", s.id, title, link, s.timeSpent, s.timeComplexity, s.spaceComplexity, rankString)
	}
	replacer += "\n<!-- golang inject solutions end -->"
	content = reg.ReplaceAllString(content, replacer)
	_ = ioutil.WriteFile(readmeFilePath, []byte(content), os.ModeAppend)
}
