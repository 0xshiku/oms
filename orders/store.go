package main

import "context"

type store struct {
	// TODO: add mongodb instance
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context) error {
	return nil
}
