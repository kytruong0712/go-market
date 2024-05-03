package categories

import (
	"context"
	"errors"
	"testing"

	"github.com/kytruong0712/go-market/product-service/api/internal/model"
	"github.com/kytruong0712/go-market/product-service/api/internal/repository"
	"github.com/kytruong0712/go-market/product-service/api/internal/repository/category"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestImpl_GetCategoriesHierarchy(t *testing.T) {
	type mockGetCategoriesRepo struct {
		want    []model.Category
		wantErr error
	}

	parentCate1 := model.Category{
		ID:           1,
		ParentID:     0,
		Name:         "Cate 1",
		Code:         "cate-1",
		Description:  "Cte 1",
		IsNavigation: true,
		IsFiltering:  true,
		Status:       "ACTIVE",
		Images: []model.CategoryImage{
			{
				ID:         1,
				CategoryID: 1,
				ImagePath:  "/foo/bar",
				ImageName:  "cate1.png",
				ImageType:  "CATEGORY_MENU_ICON",
				Status:     "ACTIVE",
			},
		},
	}
	parentCate2 := model.Category{
		ID:           2,
		ParentID:     0,
		Name:         "Cate 2",
		Code:         "cate-2",
		Description:  "Cate 2",
		IsNavigation: true,
		IsFiltering:  true,
		Status:       "ACTIVE",
		Images: []model.CategoryImage{
			{
				ID:         2,
				CategoryID: 2,
				ImagePath:  "/foo/bar",
				ImageName:  "cate2.png",
				ImageType:  "CATEGORY_MENU_ICON",
				Status:     "ACTIVE",
			},
		},
	}
	subCate1 := model.Category{
		ID:           3,
		ParentID:     1,
		Name:         "Cate 1.1",
		Code:         "cate-1.1",
		Description:  "Cate 1.1",
		IsNavigation: true,
		IsFiltering:  true,
		Status:       "ACTIVE",
	}
	subCate2 := model.Category{
		ID:           4,
		ParentID:     2,
		Name:         "Cate 2.1",
		Code:         "cate-2.1",
		Description:  "Cate 2.1",
		IsNavigation: true,
		IsFiltering:  true,
		Status:       "ACTIVE",
	}
	subCate3 := model.Category{
		ID:           5,
		ParentID:     3,
		Name:         "Cate 1.1.1",
		Code:         "cate-1.1.1",
		Description:  "Cate 1.1.1",
		IsNavigation: true,
		IsFiltering:  true,
		Status:       "ACTIVE",
	}

	tcs := map[string]struct {
		mockGetCategoriesRepo mockGetCategoriesRepo
		want                  model.CategoriesHierarchy
		wantErr               error
	}{
		"err": {
			mockGetCategoriesRepo: mockGetCategoriesRepo{
				wantErr: errors.New("err"),
			},
			wantErr: errors.New("err"),
		},
		"success": {
			mockGetCategoriesRepo: mockGetCategoriesRepo{
				want: []model.Category{
					parentCate1,
					parentCate2,
					subCate1,
					subCate2,
					subCate3,
				},
			},
			want: model.CategoriesHierarchy{
				NestedCategories: []model.Category{
					{
						ID:          1,
						ParentID:    0,
						Name:        "Cate 1",
						Code:        "cate-1",
						Description: "Cte 1",
						Status:      "ACTIVE",
						Images: []model.CategoryImage{
							{
								ID:         1,
								CategoryID: 1,
								ImagePath:  "/foo/bar",
								ImageName:  "cate1.png",
								ImageType:  "CATEGORY_MENU_ICON",
								Status:     "ACTIVE",
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
								SubCategories: []model.Category{
									{
										ID:          5,
										ParentID:    3,
										Name:        "Cate 1.1.1",
										Code:        "cate-1.1.1",
										Description: "Cate 1.1.1",
										Status:      "ACTIVE",
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
						Images: []model.CategoryImage{
							{
								ID:         2,
								CategoryID: 2,
								ImagePath:  "/foo/bar",
								ImageName:  "cate2.png",
								ImageType:  "CATEGORY_MENU_ICON",
								Status:     "ACTIVE",
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
			mockCategoryRepo := category.NewMockRepository(t)
			mockCategoryRepo.On("GetCategories", ctx, mock.Anything).Return(tc.mockGetCategoriesRepo.want, tc.mockGetCategoriesRepo.wantErr)

			mockRepo := repository.NewMockRegistry(t)
			mockRepo.On("Category").Return(mockCategoryRepo)

			rs, err := New(mockRepo).GetCategoriesHierarchy(ctx)
			if tc.wantErr != nil {
				require.Error(t, tc.wantErr, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, rs)
			}
		})
	}
}
