package model

import "time"

// CategoryStatus represents the status of the category
type CategoryStatus string

const (
	// CategoryStatusActive means the category is active
	CategoryStatusActive CategoryStatus = "ACTIVE"
	// CategoryStatusInactive means the category is inactive
	CategoryStatusInactive CategoryStatus = "INACTIVE"
	// CategoryStatusDeleted means the category is deleted. This is for archival only
	CategoryStatusDeleted CategoryStatus = "DELETED"
)

// String returns string type of custom type
func (c CategoryStatus) String() string {
	return string(c)
}

// CategoryImageType represents the type of the category image
type CategoryImageType string

const (
	// CategoryImageTypeBanner means the category image type banner
	CategoryImageTypeBanner CategoryImageType = "CATEGORY_BANNER"
	// CategoryImageTypeMenuIcon means the category image type menu icon
	CategoryImageTypeMenuIcon CategoryImageType = "CATEGORY_MENU_ICON"
)

// String returns string type of custom type
func (c CategoryImageType) String() string {
	return string(c)
}

// Category represents the category to populate
type Category struct {
	ID            int64
	ParentID      int64
	Name          string
	Code          string
	Description   string
	IsNavigation  bool
	IsFiltering   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        CategoryStatus
	Images        []CategoryImage
	SubCategories []Category
}

// CategoryImage represents the category images
type CategoryImage struct {
	ID         int64
	CategoryID int64
	ImagePath  string
	ImageName  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ImageType  CategoryImageType
	Status     CategoryStatus
}
