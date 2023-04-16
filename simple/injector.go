package simple

//go:build wireinject
// +build wireinject

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(
		NewSimpleRepository, NewSimpleService
	)
	return nil
}
