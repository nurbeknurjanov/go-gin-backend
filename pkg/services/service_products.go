package services

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
	"strconv"
)

type ProductsService struct {
	rootServices *Services
	repo         repositories.Products
	filesRepo    repositories.Files
}

func newProductsService(repo repositories.Products, filesRepo repositories.Files) *ProductsService {
	return &ProductsService{repo: repo, filesRepo: filesRepo}
}

func (s *ProductsService) Create(m *models.Product) error {
	if m.ImageID == nil {
		return s.repo.Create(m)
	}

	//db := s.productsRepo.(*repository.Repositories).IProductsRepository.(*repository.ProductsRepository).Db
	db := s.repo.(*repositories.ProductsSqlRepository).db

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := s.repo.CreateWithImage(m, tx); err != nil {
		return err
		//return tx.Rollback()
	}

	fm, err := s.filesRepo.Find(strconv.Itoa(*m.ImageID))
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

	if err = s.filesRepo.Update(fm, fp, tx); err != nil {
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

func (s *ProductsService) Update(m *models.Product, data *models.ProductPartial) error {
	return s.repo.Update(m, data)
}

func (s *ProductsService) Delete(m *models.Product) error {
	if m.Image != nil {
		f, err := s.filesRepo.Find(strconv.Itoa(*m.Image.ID))
		if err != nil {
			return err
		}

		if err := s.rootServices.Files.Delete(f); err != nil {
			return err
		}
	}

	return s.repo.Delete(m)
}

func (s *ProductsService) Find(ID string) (*models.Product, error) {
	return s.repo.Find(ID)
}

func (s *ProductsService) List(p *repositories.PaginationRequest, sort *repositories.Sort, f *models.ProductFilter) ([]*models.Product, error) {
	return s.repo.List(p, sort, f)
}

func (s *ProductsService) Count(f *models.ProductFilter) (*int, error) {
	return s.repo.Count(f)
}
