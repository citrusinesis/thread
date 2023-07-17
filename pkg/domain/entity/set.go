package entity

import "github.com/google/wire"

// Set Dependency injection target in entity
var Set = wire.NewSet(NewItemRepository, NewItemService, NewItemHandler, NewItemGroup)
