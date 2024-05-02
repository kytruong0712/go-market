package model

// CategoriesHierarchy represents categories hierarchy
type CategoriesHierarchy struct {
	NestedCategories []Category
	NestedLevel      int
}
