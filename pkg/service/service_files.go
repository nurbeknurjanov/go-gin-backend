package service

import (
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"os"
)

type FilesService struct {
	filesRepo repositories.IFilesRepository
}

func newFilesService(repositories *repositories.Repositories) *FilesService {
	return &FilesService{filesRepo: repositories}
}

func (s *FilesService) CreateFile(m *models.File) error {
	return s.filesRepo.CreateFile(m)
}

func (s *FilesService) ListFiles(p *repositories.PaginationRequest, sort *repositories.Sort, f *models.FileFilter) ([]*models.File, error) {
	return s.filesRepo.ListFiles(p, sort, f)
}

func (s *FilesService) CountFiles(f *models.FileFilter) (*int, error) {
	return s.filesRepo.CountFiles(f)
}

func (s *FilesService) DeleteFile(m *models.File) error {
	if err := os.Remove("public/upload/" + m.FileName()); err != nil {
		return err
	}
	return s.filesRepo.DeleteFile(m)
}

func (s *FilesService) FindFile(ID string) (*models.File, error) {
	return s.filesRepo.FindFile(ID)
}
