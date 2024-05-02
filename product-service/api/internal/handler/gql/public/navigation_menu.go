package public

import (
	"context"
	"fmt"

	"github.com/kytruong0712/go-market/product-service/api/internal/handler/gql/mod"
	"github.com/kytruong0712/go-market/product-service/api/internal/model"
)

// GetNavigationMenu gets navigation menu items
func (r *queryResolver) GetNavigationMenu(ctx context.Context) (*mod.NavigationMenu, error) {
	rs, err := r.categoryCtrl.GetCategoriesHierarchy(ctx)
	if err != nil {
		return nil, err
	}

	return newNavigationMenu(rs), nil
}

func newNavigationMenu(m model.CategoriesHierarchy) *mod.NavigationMenu {
	return &mod.NavigationMenu{
		NestedItems: toNestedItems(m.NestedCategories),
		NestedLevel: m.NestedLevel,
	}
}

func toNestedItems(cis []model.Category) []*mod.NavigationMenuItem {
	slice := make([]*mod.NavigationMenuItem, 0)
	for _, ci := range cis {
		slice = append(slice, toMenuItem(ci))
	}

	return slice
}

func toMenuItem(m model.Category) *mod.NavigationMenuItem {
	var imgUrl string
	if len(m.Images) > 0 {
		imgUrl = fmt.Sprintf("/%s/%s", m.Images[0].ImagePath, m.Images[0].ImageName)
	}

	return &mod.NavigationMenuItem{
		ID:          m.ID,
		ParentID:    &m.ParentID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   &m.CreatedAt,
		UpdatedAt:   &m.UpdatedAt,
		Status:      m.Status,
		ImageURL:    &imgUrl,
		SubItems:    appendSubItems(m.SubCategories),
	}
}

func appendSubItems(arr []model.Category) []*mod.NavigationMenuItem {
	var subItems []*mod.NavigationMenuItem
	for _, item := range arr {
		subItems = append(subItems, &mod.NavigationMenuItem{
			ID:          item.ID,
			ParentID:    &item.ParentID,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   &item.CreatedAt,
			UpdatedAt:   &item.UpdatedAt,
			Status:      item.Status,
			SubItems:    appendSubItems(item.SubCategories),
		})
	}

	return subItems
}
