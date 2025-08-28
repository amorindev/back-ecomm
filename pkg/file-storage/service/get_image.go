package service

import "context"

func (s *Service) GetImage(ctx context.Context, imgPath string) (string, error) {
	url, err := s.FileStgAdt.GetImage(ctx, imgPath)
	if err != nil {
		return "", err
	}
	return url, err
}
