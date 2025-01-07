package usecase

import "testing"

type Testcase[T Validator, R any] struct {
	Input    T
	Expected R
}

type Tester[T Validator, R any, U Usecase[T, R]] struct {
	usecase U
	t       *testing.T

	validator func(expected R, actual R) (bool, error)
	testcases []Testcase[T, R]
}

func NewTester[T Validator, R any, U Usecase[T, R]](usecase U, t *testing.T, validator func(R, R) (bool, error)) *Tester[T, R, U] {
	return &Tester[T, R, U]{usecase: usecase, t: t, validator: validator}
}

func (t *Tester[T, R, U]) AddTestcase(testcase ...Testcase[T, R]) {
	t.testcases = append(t.testcases, testcase...)
}

func (t *Tester[T, R, U]) Run() {
	for _, testcase := range t.testcases {
		if testcase.Input.Validate() != nil {
			t.t.Errorf("invalid input: %v", testcase.Input.Validate())
		}

		actual, err := t.usecase.Execute(testcase.Input)
		if err != nil {
			t.t.Errorf("unexpected error: %v", err)
			return
		}

		if ok, err := t.validator(testcase.Expected, actual); err != nil || !ok {
			t.t.Errorf("unexpected result: %v", err)
			return
		}
	}
}
