package public

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/kytruong0712/go-market/product-service/api/internal/controller/categories"
	"github.com/kytruong0712/go-market/product-service/api/internal/handler/gql/mod"
	"github.com/kytruong0712/go-market/product-service/api/internal/model"
	"github.com/kytruong0712/go-market/product-service/api/internal/utils/ptr"

	"github.com/stretchr/testify/require"
)

func TestQueryResolver_GetNavigationMenu(t *testing.T) {
	type mockCtrl struct {
		want    model.CategoriesHierarchy
		wantErr error
	}

	dt := time.Now()

	tcs := map[string]struct {
		mockCtrl mockCtrl
		want     *mod.NavigationMenu
		wantErr  error
	}{
		"err": {
			mockCtrl: mockCtrl{
				wantErr: errors.New("err"),
			},
			wantErr: errors.New("err"),
		},
		"success": {
			mockCtrl: mockCtrl{
				want: model.CategoriesHierarchy{
					NestedCategories: []model.Category{
						{
							ID:          1,
							ParentID:    0,
							Name:        "Cate 1",
							Code:        "cate-1",
							Description: "Cate 1",
							Status:      "ACTIVE",
							CreatedAt:   dt,
							UpdatedAt:   dt,
							Images: []model.CategoryImage{
								{
									ID:         1,
									CategoryID: 1,
									ImagePath:  "foo/bar",
									ImageName:  "cate1.png",
									ImageType:  "CATEGORY_MENU_ICON",
									Status:     "ACTIVE",
									CreatedAt:  dt,
									UpdatedAt:  dt,
								},
							},
							SubCategories: []model.Category{
								{
									ID:          3,
									ParentID:    1,
									Name:        "Cate 1.1",
									Code:        "cate-1.1",
									Description: "Cate 1.1",
									Status:      "ACTIVE",
									CreatedAt:   dt,
									UpdatedAt:   dt,
									SubCategories: []model.Category{
										{
											ID:          5,
											ParentID:    3,
											Name:        "Cate 1.1.1",
											Code:        "cate-1.1.1",
											Description: "Cate 1.1.1",
											Status:      "ACTIVE",
											CreatedAt:   dt,
											UpdatedAt:   dt,
										},
									},
								},
							},
						},
						{
							ID:          2,
							ParentID:    0,
							Name:        "Cate 2",
							Code:        "cate-2",
							Description: "Cate 2",
							Status:      "ACTIVE",
							CreatedAt:   dt,
							UpdatedAt:   dt,
							Images: []model.CategoryImage{
								{
									ID:         2,
									CategoryID: 2,
									ImagePath:  "foo/bar",
									ImageName:  "cate2.png",
									ImageType:  "CATEGORY_MENU_ICON",
									Status:     "ACTIVE",
									CreatedAt:  dt,
									UpdatedAt:  dt,
								},
							},
							SubCategories: []model.Category{
								{
									ID:          4,
									ParentID:    2,
									Name:        "Cate 2.1",
									Code:        "cate-2.1",
									Description: "Cate 2.1",
									Status:      "ACTIVE",
									CreatedAt:   dt,
									UpdatedAt:   dt,
								},
							},
						},
					},
					NestedLevel: 3,
				},
			},
			want: &mod.NavigationMenu{
				NestedItems: []*mod.NavigationMenuItem{
					{
						ID:          1,
						ParentID:    ptr.ToIntPtr(int64(0)),
						Name:        "Cate 1",
						Description: "Cate 1",
						Status:      "ACTIVE",
						ImageURL:    ptr.ToStringPtr("/foo/bar/cate1.png"),
						CreatedAt:   &dt,
						UpdatedAt:   &dt,
						SubItems: []*mod.NavigationMenuItem{
							{
								ID:          3,
								ParentID:    ptr.ToIntPtr(int64(1)),
								Name:        "Cate 1.1",
								Description: "Cate 1.1",
								Status:      "ACTIVE",
								CreatedAt:   &dt,
								UpdatedAt:   &dt,
								SubItems: []*mod.NavigationMenuItem{
									{
										ID:          5,
										ParentID:    ptr.ToIntPtr(int64(3)),
										Name:        "Cate 1.1.1",
										Description: "Cate 1.1.1",
										Status:      "ACTIVE",
										CreatedAt:   &dt,
										UpdatedAt:   &dt,
									},
								},
							},
						},
					},
					{
						ID:          2,
						ParentID:    ptr.ToIntPtr(int64(0)),
						Name:        "Cate 2",
						Description: "Cate 2",
						Status:      "ACTIVE",
						ImageURL:    ptr.ToStringPtr("/foo/bar/cate2.png"),
						CreatedAt:   &dt,
						UpdatedAt:   &dt,
						SubItems: []*mod.NavigationMenuItem{
							{
								ID:          4,
								ParentID:    ptr.ToIntPtr(int64(2)),
								Name:        "Cate 2.1",
								Description: "Cate 2.1",
								Status:      "ACTIVE",
								CreatedAt:   &dt,
								UpdatedAt:   &dt,
							},
						},
					},
				},
				NestedLevel: 3,
			},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			mockCategoriesCtrl := categories.NewMockController(t)
			mockCategoriesCtrl.On("GetCategoriesHierarchy", ctx).Return(tc.mockCtrl.want, tc.mockCtrl.wantErr)

			r := &queryResolver{
				resolver: &resolver{
					categoryCtrl: mockCategoriesCtrl,
				},
			}

			rs, err := r.GetNavigationMenu(ctx)
			if tc.wantErr != nil {
				require.Error(t, tc.wantErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, rs)
			}
		})
	}
}
