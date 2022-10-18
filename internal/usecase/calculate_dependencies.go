package usecase

import (
	"log"

	"github.com/fabianoshz/tg-runner/internal/entity"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Dependencies struct {
	Dependencies []Dependency `hcl:"dependency,block"`
}

type Dependency struct {
	ConfigPath string `hcl:"config_path"`
	Type       string `hcl:"type,label"`
}

var configSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "dependency",
			LabelNames: []string{"dependency"},
		},
	},
}

var dependencySchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name:     "config_path",
			Required: true,
		},
	},
}

func (d CalculateDependenciesService) CalculateDependencies(resources []entity.Resource) [][]entity.Resource {
	// TODO calculate dependencies of resources
	// TODO order dependencies of resources

	for _, v := range resources {
		parser := hclparse.NewParser()
		out, _ := parser.ParseHCLFile(v.Path + "/terragrunt.hcl")

		content, _, _ := out.Body.PartialContent(configSchema)

		for _, b := range content.Blocks {
			content, _ := b.Body.Content(dependencySchema)
			log.Printf("Configuration is %#v", content.Attributes["config_path"])
		}

		// TODO resolve dependencies
		// TODO resolve find_in_parent_folders()
	}

	var ordered [][]entity.Resource

	ordered = append(ordered, resources)

	return ordered
}
