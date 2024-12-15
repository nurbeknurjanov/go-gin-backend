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

	fm, err := s.filesRepo.Find(strconv.Itoa(*m.ImageID))
	if err != nil {
		return err
	}
	return s.repo.CreateWithImage(m, fm, s.filesRepo)
}

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
