package usecase

type Validator interface {
	Validate() error
}

type Usecase[T Validator, R any] interface {
	Execute(T) (R, error)
}
