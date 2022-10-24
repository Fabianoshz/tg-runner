package usecase

import (
	"errors"
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

	// Look for all the .hcl files and get the dependencies.
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

	var orderExecution [][]entity.Resource

	for _, resource := range resources {
		var visited []string
		var groupOrder []entity.Resource

		groupOrder, _ = groupFileDependencies(resource.Path, dags, visited, groupOrder)
		orderExecution = append(orderExecution, groupOrder)
	}

	return orderExecution
}

// TODO use hclparser, doing this because I'm lazy.
// terragrunt checks for circular dependency, which
// is not fine for this, we need to have our own check.
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

func groupFileDependencies(path string, dags map[string][]string, visited []string, groupOrder []entity.Resource) ([]entity.Resource, error) {
	// If we visited this node before in this group
	// it means we've found a circular dependency.
	for _, v := range visited {
		if v == path {
			// TODO return error
			fmt.Println("Circular dependency detected.")
			return groupOrder, errors.New("empty name")
		}
	}

	visited = append(visited, path)

	// Recursivly check the children until we reach the top.
	// Children are processed before the parent because of this.
	for _, child := range dags[path] {
		groupOrder, _ = groupFileDependencies(child, dags, visited, groupOrder)
	}

	// We should only add a dependency that's not been
	// added before.
	add := true
	for _, v := range groupOrder {
		if v.Path == path {
			add = false
		}
	}

	if add {
		var a entity.Resource
		a.Action = entity.Undefined
		a.Path = path

		groupOrder = append(groupOrder, a)
	}

	return groupOrder, nil
}
