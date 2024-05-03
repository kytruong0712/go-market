package category

import (
	"context"
	"fmt"
	"github.com/kytruong0712/go-market/product-service/api/internal/model"
	"github.com/kytruong0712/go-market/product-service/api/internal/repository/dbmodel"

	pkgerrors "github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetCategoriesInput represents input to load
type GetCategoriesInput struct {
	LoadImages   bool
	IsNavigation *bool
}

// GetCategories gets category items which available as navigation item
func (i impl) GetCategories(ctx context.Context, inp GetCategoriesInput) ([]model.Category, error) {
	qms := []qm.QueryMod{
		dbmodel.CategoryWhere.Status.EQ(model.CategoryStatusActive.String()),
	}

	if inp.IsNavigation != nil {
		qms = append(qms, dbmodel.CategoryWhere.IsNavgitation.EQ(*inp.IsNavigation))
	}

	if inp.LoadImages {
		qms = append(qms, qm.Load(dbmodel.CategoryRels.CategoryImages,
			dbmodel.CategoryImageWhere.Status.EQ(model.CategoryStatusActive.String()),
			dbmodel.CategoryImageWhere.ImageType.EQ(null.StringFrom(model.CategoryImageTypeMenuIcon.String()))))

		qms = append(qms, qm.Load(fmt.Sprintf("%v.%v", dbmodel.CategoryRels.CategoryImages, dbmodel.CategoryImageRels.UploadedFile),
			dbmodel.UploadedFileWhere.Status.EQ(model.CategoryStatusActive.String()),
			dbmodel.UploadedFileWhere.FileType.EQ(model.FileTypeImage.String())))
	}

	dbItems, err := dbmodel.Categories(qms...).All(ctx, i.dbConn)

	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var modelItems []model.Category
	for _, dbItem := range dbItems {
		modelItems = append(modelItems, toCategoryModel(dbItem))
	}

	return modelItems, nil
}

func toCategoryModel(dbItem *dbmodel.Category) model.Category {
	m := model.Category{
		ID:           dbItem.ID,
		ParentID:     dbItem.ParentID.Int64,
		Name:         dbItem.CategoryName,
		Code:         dbItem.CategoryCode,
		Description:  dbItem.Description,
		IsNavigation: dbItem.IsNavgitation,
		CreatedAt:    dbItem.CreatedAt,
		UpdatedAt:    dbItem.UpdatedAt,
		Status:       model.CategoryStatus(dbItem.Status),
	}

	if dbItem.R != nil && dbItem.R.CategoryImages != nil {
		cis := dbItem.R.CategoryImages
		for _, ci := range cis {
			m.Images = append(m.Images, toCategoryImageModel(ci))
		}
	}

	return m
}

func toCategoryImageModel(dbItem *dbmodel.CategoryImage) model.CategoryImage {
	m := model.CategoryImage{
		ID:         dbItem.ID,
		CategoryID: dbItem.CategoryID,
		CreatedAt:  dbItem.CreatedAt,
		UpdatedAt:  dbItem.UpdatedAt,
		ImageType:  model.CategoryImageType(dbItem.ImageType.String),
		Status:     model.CategoryStatus(dbItem.Status),
	}

	if dbItem.R != nil && dbItem.R.UploadedFile != nil {
		m.ImagePath = dbItem.R.UploadedFile.FilePath
		m.ImageName = dbItem.R.UploadedFile.FileName
	}

	return m
}
