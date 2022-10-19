package usecase

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/fabianoshz/tg-runner/internal/entity"
)

func (d CalculateDependenciesService) CalculateDependencies(resources []entity.Resource, rootdir string) [][]entity.Resource {
	wg := new(sync.WaitGroup)
	var dags = make(map[string][]string)

	filepath.Walk(rootdir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			obj, err := regexp.Match(`(.*).hcl`, []byte(info.Name()))
			if err != nil {
				return err
			}

			if obj && string(info.Name()[0]) != "." {
				fmt.Println("file path:", filepath.Dir(path))

				wg.Add(1)
				go getFileDependencies(filepath.Dir(path), dags, wg)
			}
		}

		return nil
	})

	wg.Wait()

	var orderExecution [][]string

	for _, resource := range resources {
		var visited []string
		var groupOrder []string

		groupOrder = groupFileDependencies(resource.Path, dags, visited, groupOrder)
		orderExecution = append(orderExecution, groupOrder)
	}

	var ordered [][]entity.Resource
	return ordered
}

// TODO use hclparser, doing this because I'm lazy.
func getFileDependencies(path string, dags map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()

	var command = "terragrunt graph-dependencies --terragrunt-non-interactive | grep \"" + path + "\" | awk '{print $3}' | cut -d '\"' -f2"

	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = path
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	for _, s := range strings.Split(string(out), "\n") {
		if len(s) > 0 {
			dags[path] = append(dags[path], s)
		}
	}
}

func groupFileDependencies(path string, dags map[string][]string, visited []string, groupOrder []string) []string {
	// TODO return error if this node has been visited already
	visited = append(visited, path)

	for _, child := range dags[path] {
		groupOrder = append(groupOrder, child)

		groupFileDependencies(child, dags, visited, groupOrder)
	}

	groupOrder = append(groupOrder, path)
	return groupOrder
}
