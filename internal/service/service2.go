package service

import "fmt"

type svc2 struct {
	svc3 *svc3
}

func (s *svc2) exec() error {
	if err := s.svc3.exec(); err != nil {
		return fmt.Errorf("exec svc3: %w", err)
	}

	return nil
}
