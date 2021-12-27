package users

type Service interface {
	Init() ([]User, error)
	GetAll() ([]User, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (User, error)
	GetUser(id int) (User, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaDeCreacion string) (User, error) {

	id, err := s.repository.GetId()

	if err != nil {
		return User{}, nil
	}

	user, err := s.repository.Store(id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *service) Init() ([]User, error) {
	users, err := s.repository.Init()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetUser(id int) (User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
