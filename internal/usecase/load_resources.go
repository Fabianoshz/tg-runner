package usecase

import (
	"log"
	"os"

	"github.com/fabianoshz/tg-runner/internal/entity"
	"gopkg.in/yaml.v3"
)

func (r LoadResourcesService) LoadResources(changelist string) []entity.Resource {
	var resources []entity.Resource

	yamlFile, err := os.ReadFile(changelist)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &resources)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return resources
}
