package template

type TestCase struct {
	TestName string `json:"name"`
	TestSteps []Step `json:"test_steps"`
}