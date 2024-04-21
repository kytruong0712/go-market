package public

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kytruong0712/go-market/product-service/api/internal/controller/categories"
	"github.com/kytruong0712/go-market/product-service/api/internal/handler/gql/scalar"
	"github.com/kytruong0712/go-market/product-service/api/internal/model"

	"github.com/kytruong0712/go-market/product-service/api/internal/handler/gql/mod"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type GetNavigationMenuInput struct {
	HierarchyLevel int
}

func (r *queryResolver) GetNavigationMenu(ctx context.Context, inp GetNavigationMenuInput) ([]*mod.NavigationMenu, error) {
	if inp.HierarchyLevel <= 0 {
		return nil, WebErrHierarchyLevelInvalid
	}

	items, err := r.categoryCtrl.GetNavigationItems(ctx, categories.GetNavigationItemsInput{
		HierarchyLevel: inp.HierarchyLevel,
	})
	if err != nil {
		return nil, err
	}

	return newNavigationMenuSlice(items), nil
}

func newNavigationMenuSlice(cis []model.Category) []*mod.NavigationMenu {
	slice := make([]*mod.NavigationMenu, 0, 0)
	for _, ci := range cis {
		slice = append(slice, newCategoryMenu(ci))
	}

	return slice
}

func newCategoryMenu(m model.Category) *mod.NavigationMenu {
	var imgUrl string
	if len(m.Images) > 0 {
		imgUrl = fmt.Sprintf("/%s/%s", m.Images[0].ImagePath, m.Images[0].ImageName)
	}

	nav := &mod.NavigationMenu{
		ID:          m.ID,
		ParentID:    &m.ParentID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   &m.CreatedAt,
		UpdatedAt:   &m.UpdatedAt,
		Status:      m.Status,
		ImageURL:    &imgUrl,
	}

	subItems, err := getSubItems(m)
	if err != nil {
		panic(err)
	}

	nav.SubItems = subItems

	return nav
}

func getSubItems(m model.Category) (scalar.JSONSlice, error) {
	var jsonSlice scalar.JSONSlice

	for _, v := range m.SubCategories {
		byteData, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		tmp := make(map[string]types.JSON)
		err = json.Unmarshal(byteData, &tmp)
		if err != nil {
			return nil, err
		}

		jsonSlice = append(jsonSlice, tmp)
	}

	return jsonSlice, nil
}
