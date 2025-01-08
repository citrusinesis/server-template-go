package validator_test

import (
	"errors"
	"strings"
	"testing"

	"example/pkg/validator"
)

type SimpleDTO struct {
	Name string
}

func (s *SimpleDTO) Validate() error {
	if s.Name == "" {
		return errors.New("Name is empty")
	}
	return nil
}

type NestedDTO struct {
	Title        string
	SubItem      *SimpleDTO
	ItemChildren []SimpleDTO
}

func (n *NestedDTO) Validate() error {
	if n.Title == "" {
		return errors.New("Title is empty")
	}
	return nil
}

type MixedDTO struct {
	Number  int
	SubMap  map[string]*SimpleDTO
	SubList []*SimpleDTO
}

func TestValidate_NilInput(t *testing.T) {
	if err := validator.Validate(nil); err != nil {
		t.Errorf("expected no error for nil input, got: %v", err)
	}
}

func TestValidate_NonStruct(t *testing.T) {
	num := 42
	if err := validator.Validate(num); err != nil {
		t.Errorf("expected no error for non-struct, got: %v", err)
	}
}

func TestValidate_SimpleDTO_Valid(t *testing.T) {
	dto := &SimpleDTO{Name: "ValidName"}
	if err := validator.Validate(dto); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidate_SimpleDTO_Invalid(t *testing.T) {
	dto := &SimpleDTO{Name: ""}
	err := validator.Validate(dto)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if !strings.Contains(err.Error(), "Name is empty") {
		t.Errorf("expected error to contain 'Name is empty', got: %v", err)
	}
}

func TestValidate_NestedDTO_Valid(t *testing.T) {
	dto := &NestedDTO{
		Title: "Some Title",
		SubItem: &SimpleDTO{
			Name: "Valid SubItem",
		},
		ItemChildren: []SimpleDTO{
			{Name: "Child1"},
			{Name: "Child2"},
		},
	}

	if err := validator.Validate(dto); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidate_NestedDTO_Invalid(t *testing.T) {
	dto := &NestedDTO{
		Title: "",
		SubItem: &SimpleDTO{
			Name: "",
		},
		ItemChildren: []SimpleDTO{
			{Name: ""},
			{Name: "Child2"},
		},
	}

	err := validator.Validate(dto)
	if err == nil {
		t.Fatal("expected multiple validation errors, got nil")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "Title is empty") {
		t.Errorf("expected error for empty Title, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "SubItem: Name is empty") {
		t.Errorf("expected error for empty SubItem.Name, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "ItemChildren[0]: Name is empty") {
		t.Errorf("expected error for empty Child1.Name, got: %s", errMsg)
	}
}

func TestValidate_MixedDTO(t *testing.T) {
	dto := &MixedDTO{
		Number: 100,
		SubMap: map[string]*SimpleDTO{
			"key1": {Name: ""},
			"key2": {Name: "Ok"},
		},
		SubList: []*SimpleDTO{
			{Name: "Ok"},
			{Name: ""},
		},
	}

	err := validator.Validate(dto)
	if err == nil {
		t.Fatal("expected validation errors in SubMap and SubList, got nil")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "SubList[1]: Name is empty") {
		t.Errorf("expected SubMap[key1].Name is empty error, got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "SubList[1]: Name is empty") {
		t.Errorf("expected SubList[1].Name is empty error, got: %s", errMsg)
	}
}
