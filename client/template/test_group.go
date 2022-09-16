package template

type TestGroup struct {
	TestGroupName string `json:"test_group_name"`
	TestGroupTests []Test `json:"test_group_tests"`
}