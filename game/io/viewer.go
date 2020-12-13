package io

type Viewer interface {
	PrintInputPrompt(currentCheckCount int)
	PrintTitle()
	PrintCheckResult(result CheckResult)
}

type CheckResult int

const (
	_ CheckResult = iota
	Equal
	More
	Less
)


