package usecase

import "github.com/fabianoshz/tg-runner/internal/entity"

func (d CalculateDependenciesService) CalculateDependencies(resources []entity.Resource) [][]entity.Resource {
	// TODO calculate dependencies of resources
	// TODO order dependencies of resources

	var ordered [][]entity.Resource

	ordered = append(ordered, resources)

	return ordered
}
