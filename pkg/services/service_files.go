package services

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/repositories"
	"os"
)

type FilesService struct {
	repo repositories.Files
}

func newFilesService(repo repositories.Files) *FilesService {
	return &FilesService{repo: repo}
}

func (s *FilesService) Create(m *models.File) error {
	return s.repo.Create(m)
}

func (s *FilesService) List(p *repositories.PaginationRequest, sort *repositories.Sort, f *models.FileFilter) ([]*models.File, error) {
	return s.repo.List(p, sort, f)
}

func (s *FilesService) Count(f *models.FileFilter) (*int, error) {
	return s.repo.Count(f)
}

func (s *FilesService) Delete(m *models.File) error {
	if err := os.Remove("public/upload/" + m.FileName()); err != nil {
		return err
	}
	return s.repo.Delete(m)
}

func (s *FilesService) Find(ID string) (*models.File, error) {
	return s.repo.Find(ID)
}
