package viewer

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

type LanguageData struct {
	Title       string
	InputNumber string
	Less        string
	More        string
	Equal       string
}
