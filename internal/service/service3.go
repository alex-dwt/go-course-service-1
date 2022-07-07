package service

type svc3 struct {
}

func (s *svc3) exec() error {
	//return errors.New("could not read file")
	return ErrMyErrorBlue
}
