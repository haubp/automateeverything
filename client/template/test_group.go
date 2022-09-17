package template

type TestGroup struct {
	TestGroupName string `json:"test_group_name"`
	TestGroupTestCases []TestCase `json:"test_group_tests"`
}