package handler

import "fmt"

type CreateOpeningRequest struct {
	Id       uint   `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (cor *CreateOpeningRequest) Validate() error {
	if cor.Role == "" {
		return requiredParam("role", "string")
	}

	if cor.Company == "" {
		return requiredParam("company", "string")
	}

	if cor.Location == "" {
		return requiredParam("location", "string")
	}

	if cor.Link == "" {
		return requiredParam("link", "string")
	}

	if cor.Remote == nil {
		return requiredParam("remote", "boolean")
	}

	if cor.Salary <= 0 {
		return requiredParam("salary", "integer")
	}

	return nil
}

func requiredParam(paramName, paramType string) error {
	return fmt.Errorf("missing param: %s (type: %s) is required!", paramName, paramType)
}

type UpdateOpeningRequest struct {
	Id       uint   `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (cor *UpdateOpeningRequest) Validate() error {
	if cor.Role != "" || cor.Company != "" ||
		cor.Location != "" || cor.Link != "" ||
		cor.Remote != nil || cor.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one filed must be provided")
}
