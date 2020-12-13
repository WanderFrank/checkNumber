package io

type Transport interface {
	Print(message string)
	GetNumber() int
}