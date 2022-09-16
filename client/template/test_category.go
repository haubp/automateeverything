package template

type TestCategory struct {
	TestCategoryName string `json:"test_category_name"`
	TestCategoryGroups []TestGroup `json:"test_category_groups"`
}