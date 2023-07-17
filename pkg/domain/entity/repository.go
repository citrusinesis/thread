package entity

import (
	"github.com/citrusinesis/thread/pkg/ent"
	"sync"
)

var repositoryOnce sync.Once

type Repository struct {
	client *ent.Client
}

func NewItemRepository(client *ent.Client) (repository *Repository) {
	repositoryOnce.Do(func() {
		repository = &Repository{client}
	})
	return
}

/*
	TODO: Define repositories.

	Example:
		func (r *Repository) BEHAVE() (ent.Items, error) {
			return r.client.ENTITY.Query().All(context.Background())
		}
*/
