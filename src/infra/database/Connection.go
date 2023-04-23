package database

type Connection[T any] interface {
	Info()
	Client() T
}
