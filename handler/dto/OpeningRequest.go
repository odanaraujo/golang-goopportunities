package dto

import (
	"fmt"
	"strings"
)

func ErrParamIsRequesd(name string, typ string) error {
	return fmt.Errorf("param %s is required and should be %s", name, typ)
}

type OpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Link     string  `json:"link"`
	Remote   *bool   `json:"remote"`
	Salary   float64 `json:"salary"`
}

func (request *OpeningRequest) Validate() error {

	if request == nil {
		return fmt.Errorf("request body is empty or malformed")
	}

	if strings.TrimSpace(request.Role) == "" {
		return ErrParamIsRequesd("role", "string")
	}

	if strings.TrimSpace(request.Company) == "" {
		return ErrParamIsRequesd("company", "string")
	}

	if strings.TrimSpace(request.Location) == "" {
		return ErrParamIsRequesd("location", "string")
	}

	if strings.TrimSpace(request.Link) == "" {
		return ErrParamIsRequesd("link", "string")
	}

	if request.Remote == nil {
		return ErrParamIsRequesd("remote", "bool")
	}

	if request.Salary <= 0 {
		return ErrParamIsRequesd("salary", "float64")
	}

	return nil
}
