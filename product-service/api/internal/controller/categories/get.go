package categories

import (
	"context"

	"github.com/kytruong0712/go-market/product-service/api/internal/model"
	"github.com/kytruong0712/go-market/product-service/api/internal/repository/category"
	"github.com/kytruong0712/go-market/product-service/api/internal/utils/ptr"
)

// GetCategoriesHierarchy gets categories hierarchy
func (i impl) GetCategoriesHierarchy(ctx context.Context) (model.CategoriesHierarchy, error) {
	inp := category.GetCategoriesInput{
		LoadImages:   true,
		IsNavigation: ptr.ToBoolPtr(true),
	}
	slice, err := i.repo.Category().GetCategories(ctx, inp)
	if err != nil {
		return model.CategoriesHierarchy{}, err
	}

	const (
		parentID     = 0
		currentLevel = 0
	)

	ch, lvl := buildCategoriesHierarchy(slice, parentID, currentLevel)

	return model.CategoriesHierarchy{
		NestedCategories: ch,
		NestedLevel:      lvl,
	}, nil
}

func buildCategoriesHierarchy(flattenCategories []model.Category, parentID int64, currentLevel int) ([]model.Category, int) {
	var categories []model.Category

	maxLevel := currentLevel
	for _, item := range flattenCategories {
		if item.ParentID == parentID {
			c := model.Category{
				ID:          item.ID,
				ParentID:    item.ParentID,
				Name:        item.Name,
				Code:        item.Code,
				Description: item.Description,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
				Status:      item.Status,
				Images:      item.Images,
			}

			var level int
			c.SubCategories, level = buildCategoriesHierarchy(flattenCategories, item.ID, currentLevel+1)
			categories = append(categories, c)

			if level > maxLevel {
				maxLevel = level
			}
		}
	}

	return categories, maxLevel
}
