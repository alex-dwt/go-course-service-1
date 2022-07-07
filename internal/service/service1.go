package service

import "fmt"

type Svc1 struct {
	svc2 *svc2
}

func NewService1() Svc1 {
	return Svc1{
		svc2: &svc2{
			svc3: &svc3{},
		},
	}
}

func (s *Svc1) Exec() error {
	if err := s.svc2.exec(); err != nil {
		return fmt.Errorf("exec svc2: %w", err)
	}

	return nil
}