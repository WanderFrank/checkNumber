package io

type Transport interface {
	Print(message string) error
	GetNumber() (int, error)
}