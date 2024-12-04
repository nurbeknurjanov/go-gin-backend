package service

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repository"
	"strconv"
)

type ProductsService struct {
	RootServices *Services
	productsRepo repository.IProductsRepository
	filesRepo    repository.IFilesRepository
}

func (s *ProductsService) CreateProduct(m *models.Product) error {
	if m.ImageID == nil {
		return s.productsRepo.CreateProduct(m)
	}

	//db := s.productsRepo.(*repository.Repositories).IProductsRepository.(*repository.ProductsRepository).Db
	db := s.productsRepo.(*repository.Repositories).Db

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := s.productsRepo.CreateProductWithImage(m, tx); err != nil {
		return err
		//return tx.Rollback()
	}

	fm, err := s.filesRepo.FindFile(strconv.Itoa(*m.ImageID))
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	fp := &models.FilePartial{}
	modelName := "Product"
	fp.ModelName = &modelName
	fp.ModelId = m.ID
	data := map[string]string{"type": "image"}
	fp.Data = &data

	if err = s.filesRepo.UpdateFile(fm, fp, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	m.Image = &models.File{ID: fm.ID, Ext: fm.Ext, UUID: fm.UUID, Data: fp.Data}

	return tx.Commit()
}

/*func (s *ProductsService) CreateProduct(m *models.Product) error {
	if err := s.productsRepo.CreateProduct(m); err != nil {
		return err
	}

	if m.ImageID != nil {
		fm, err := s.filesRepo.FindFile(strconv.Itoa(*m.ImageID))
		if err != nil {
			return err
		}

		fp := &models.FilePartial{}
		modelName := "Product"
		fp.ModelName = &modelName
		fp.ModelId = &m.ID
		data := map[string]string{"type": "image"}
		fp.Data = &data

		if err := s.filesRepo.UpdateFile(fm, fp); err != nil {
			return err
		}

		m.Image = &models.File{ID: &fm.ID, Ext: &fm.Ext, UUID: &fm.UUID, Data: fp.Data}
	}

	return nil
}*/

func (s *ProductsService) UpdateProduct(m *models.Product, data *models.ProductPartial) error {
	return s.productsRepo.UpdateProduct(m, data)
}

func (s *ProductsService) DeleteProduct(m *models.Product) error {
	if m.Image != nil {
		f, err := s.filesRepo.FindFile(strconv.Itoa(*m.Image.ID))
		if err != nil {
			return err
		}

		if err := s.RootServices.DeleteFile(f); err != nil {
			return err
		}
	}

	return s.productsRepo.DeleteProduct(m)
}

func (s *ProductsService) FindProduct(ID string) (*models.Product, error) {
	return s.productsRepo.FindProduct(ID)
}

func (s *ProductsService) ListProducts(p *repository.PaginationRequest, sort *repository.Sort, f *models.ProductFilter) ([]*models.Product, error) {
	return s.productsRepo.ListProducts(p, sort, f)
}

func (s *ProductsService) CountProducts(f *models.ProductFilter) (*int, error) {
	return s.productsRepo.CountProducts(f)
}

func newProductsService(repositories *repository.Repositories) *ProductsService {
	return &ProductsService{productsRepo: repositories, filesRepo: repositories}
}
