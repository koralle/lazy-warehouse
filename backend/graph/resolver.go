package graph

import "github.com/koralle/lazy-warehouse/backend/graph/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Srv service.Services
}
