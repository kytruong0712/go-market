package categories

import (
	"context"

	"github.com/kytruong0712/go-market/product-service/api/internal/repository/category"

	"github.com/kytruong0712/go-market/product-service/api/internal/model"
)

type GetNavigationItemsInput struct {
	HierarchyLevel int
}

// GetNavigationItems gets navigation items
func (i impl) GetNavigationItems(ctx context.Context, inp GetNavigationItemsInput) ([]model.Category, error) {
	slice, err := i.repo.Category().GetCategories(ctx, category.GetCategoriesInput{
		LoadImages:          true,
		LoadNavigationItems: true,
	})
	if err != nil {
		return nil, err
	}

	var currentHierarchyLevel = 1

	return buildCatalogsHierarchy(slice, 0, inp.HierarchyLevel, currentHierarchyLevel), nil
}

func buildCatalogsHierarchy(flattenCategories []model.Category, parentID int64, maxHierarchyLevel, currentHierarchyLevel int) []model.Category {
	if currentHierarchyLevel > maxHierarchyLevel {
		return nil
	}

	var categories []model.Category

	i := 0
	for i < len(flattenCategories) {
		if flattenCategories[i].ParentID == parentID {
			c := model.Category{
				ID:            flattenCategories[i].ID,
				ParentID:      flattenCategories[i].ParentID,
				Name:          flattenCategories[i].Name,
				Code:          flattenCategories[i].Code,
				Description:   flattenCategories[i].Description,
				CreatedAt:     flattenCategories[i].CreatedAt,
				UpdatedAt:     flattenCategories[i].UpdatedAt,
				Status:        flattenCategories[i].Status,
				Images:        flattenCategories[i].Images,
				SubCategories: buildCatalogsHierarchy(flattenCategories[i+1:], flattenCategories[i].ID, maxHierarchyLevel, currentHierarchyLevel+1),
			}

			categories = append(categories, c)
			flattenCategories = append(flattenCategories[:i], flattenCategories[i+1:]...)
		} else {
			i++
		}
	}

	return categories
}
