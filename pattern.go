package main

type Cache interface {
	Name() string
	Set(key string, data interface{}) error
	Get(key string) (interface{}, error)
}
