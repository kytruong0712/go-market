package category

import (
	"context"
	"database/sql"
	"testing"

	"github.com/kytruong0712/go-market/product-service/api/internal/model"
	"github.com/kytruong0712/go-market/product-service/api/internal/utils/ptr"
	"github.com/kytruong0712/go-market/product-service/api/internal/utils/test"

	"github.com/stretchr/testify/require"
)

func TestImpl_GetCategories(t *testing.T) {
	item1 := model.Category{
		ID:           1,
		ParentID:     0,
		Name:         "Điện thoại, Tablet",
		Code:         "dien-thoai-tablet",
		Description:  "Điện thoại, Tablet",
		IsNavigation: true,
		Status:       "ACTIVE",
		Images: []model.CategoryImage{
			{
				ID:         1,
				CategoryID: 1,
				ImagePath:  "foo/bar",
				ImageName:  "cate1.png",
				ImageType:  "CATEGORY_MENU_ICON",
				Status:     "ACTIVE",
			},
		},
		SubCategories: nil,
	}
	item2 := model.Category{
		ID:           2,
		ParentID:     0,
		Name:         "Laptop",
		Code:         "laptop",
		Description:  "Laptop",
		IsNavigation: true,
		Status:       "ACTIVE",
		Images: []model.CategoryImage{
			{
				ID:         2,
				CategoryID: 2,
				ImagePath:  "foo/bar",
				ImageName:  "cate2.png",
				ImageType:  "CATEGORY_MENU_ICON",
				Status:     "ACTIVE",
			},
		},
		SubCategories: nil,
	}
	item3 := model.Category{
		ID:           3,
		ParentID:     0,
		Name:         "Âm thanh",
		Code:         "am-thanh",
		Description:  "Âm thanh",
		IsNavigation: true,
		Status:       "ACTIVE",
		Images: []model.CategoryImage{
			{
				ID:         3,
				CategoryID: 3,
				ImagePath:  "foo/bar",
				ImageName:  "cate3.png",
				ImageType:  "CATEGORY_MENU_ICON",
				Status:     "ACTIVE",
			},
		},
		SubCategories: nil,
	}
	item4 := model.Category{
		ID:            4,
		ParentID:      1,
		Name:          "Thương hiệu điện thoại",
		Code:          "thuong-hieu-dien-thoai",
		Description:   "Thương hiệu điện thoại",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item5 := model.Category{
		ID:            5,
		ParentID:      1,
		Name:          "Thương hiệu Tablet",
		Code:          "thuong-hieu-tablet",
		Description:   "Thương hiệu Tablet",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item6 := model.Category{
		ID:            6,
		ParentID:      2,
		Name:          "Chọn theo hãng",
		Code:          "chon-theo-hang",
		Description:   "Chọn theo hãng",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item7 := model.Category{
		ID:            7,
		ParentID:      2,
		Name:          "Tầm giá",
		Code:          "tam-gia",
		Description:   "Tầm giá",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item8 := model.Category{
		ID:            8,
		ParentID:      3,
		Name:          "Chọn loại tai nghe",
		Code:          "chon-loai-tai-nghe",
		Description:   "Chọn loại tai nghe",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item9 := model.Category{
		ID:            9,
		ParentID:      3,
		Name:          "Hãng tai nghe",
		Code:          "hang-tai-nghe",
		Description:   "Hãng tai nghe",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item10 := model.Category{
		ID:            10,
		ParentID:      4,
		Name:          "Apple",
		Code:          "apple",
		Description:   "Apple",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item11 := model.Category{
		ID:            11,
		ParentID:      4,
		Name:          "Samsung",
		Code:          "samsung",
		Description:   "Samsung",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item12 := model.Category{
		ID:            12,
		ParentID:      5,
		Name:          "Apple",
		Code:          "apple",
		Description:   "Apple",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item13 := model.Category{
		ID:            13,
		ParentID:      5,
		Name:          "Samsung",
		Code:          "samsung",
		Description:   "Samsung",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item14 := model.Category{
		ID:            14,
		ParentID:      6,
		Name:          "Dell",
		Code:          "dell",
		Description:   "Dell",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item15 := model.Category{
		ID:            15,
		ParentID:      6,
		Name:          "Lenovo",
		Code:          "lenovo",
		Description:   "Lenovo",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item16 := model.Category{
		ID:            16,
		ParentID:      7,
		Name:          "Dưới 10 triệu",
		Code:          "duoi-10-trieu",
		Description:   "Dưới 10 triệu",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item17 := model.Category{
		ID:            17,
		ParentID:      7,
		Name:          "Từ 10 - 15 triệu",
		Code:          "10-15-trieu",
		Description:   "Từ 10 - 15 triệu",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item18 := model.Category{
		ID:            18,
		ParentID:      8,
		Name:          "Tai nghe true-wireless",
		Code:          "tai-nghe-true-wireless",
		Description:   "Tai nghe true-wireless",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item19 := model.Category{
		ID:            19,
		ParentID:      8,
		Name:          "Tai nghe có dây",
		Code:          "tai-nghe-co-day",
		Description:   "Tai nghe có dây",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item20 := model.Category{
		ID:            20,
		ParentID:      9,
		Name:          "Sennheiser",
		Code:          "sennheiser",
		Description:   "Sennheiser",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item21 := model.Category{
		ID:            21,
		ParentID:      9,
		Name:          "Sony",
		Code:          "sony",
		Description:   "Sony",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item22 := model.Category{
		ID:            22,
		ParentID:      10,
		Name:          "Iphone 15 series",
		Code:          "iphone-15-series",
		Description:   "Iphone 15 series",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item23 := model.Category{
		ID:            23,
		ParentID:      14,
		Name:          "Dell XPS series",
		Code:          "dell-xps-series",
		Description:   "Dell XPS series",
		IsNavigation:  true,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item24 := model.Category{
		ID:            26,
		ParentID:      0,
		Name:          "Promotion event",
		Code:          "promotion-event",
		Description:   "Promotion event",
		IsNavigation:  false,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}
	item25 := model.Category{
		ID:            27,
		ParentID:      4,
		Name:          "Nokia",
		Code:          "nokia",
		Description:   "Nokia",
		IsNavigation:  false,
		Status:        "ACTIVE",
		Images:        nil,
		SubCategories: nil,
	}

	tcs := map[string]struct {
		input   GetCategoriesInput
		want    []model.Category
		wantLen int
		wantErr error
	}{
		"success | load image: true | load nav items: true": {
			input: GetCategoriesInput{
				LoadImages:   true,
				IsNavigation: ptr.ToBoolPtr(true),
			},
			want: []model.Category{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10, item11, item12,
				item13, item14, item15, item16, item17, item18, item19, item20, item21, item22, item23},
			wantLen: 23,
		},
		"success | load image: true | load nav items: false": {
			input: GetCategoriesInput{
				LoadImages:   true,
				IsNavigation: ptr.ToBoolPtr(false),
			},
			want:    []model.Category{item24, item25},
			wantLen: 2,
		},
		"success | load image: true": {
			input: GetCategoriesInput{
				LoadImages: true,
			},
			want: []model.Category{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10, item11, item12,
				item13, item14, item15, item16, item17, item18, item19, item20, item21, item22, item23, item24, item25},
			wantLen: 25,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			test.WithTxDB(t, func(tx *sql.Tx) {
				ctx := context.Background()

				repo := New(tx)
				rs, err := repo.GetCategories(ctx, tc.input)
				if tc.wantErr != nil {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.wantLen, len(rs))
					for idx, exp := range tc.want {
						test.Compare(t, exp, rs[idx], model.Category{}, "CreatedAt", "UpdatedAt", "Images")
						test.Compare(t, exp.Images, rs[idx].Images, model.CategoryImage{}, "CreatedAt", "UpdatedAt")
					}
				}
			})
		})
	}
}
