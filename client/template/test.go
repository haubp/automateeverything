package template

type Test struct {
	TestName string `json:"name"`
	TestSteps []Step `json:"test_steps"`
}