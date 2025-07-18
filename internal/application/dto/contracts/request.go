package contracts

import (
	"strings"

	"github.com/Azzurriii/slythr/internal/domain/errors"
)

var (
	ErrInvalidAddress = errors.ErrInvalidAddress
)

type GetContractSourceCodeRequest struct {
	Address string `json:"address" binding:"required"`
	Network string `json:"network" binding:"required"`
}

func (r *GetContractSourceCodeRequest) Validate() error {
	if r.Address == "" {
		return errors.ErrInvalidAddress
	}

	if !strings.HasPrefix(r.Address, "0x") {
		return errors.ErrInvalidAddress
	}

	if len(r.Address) != 42 {
		return errors.ErrInvalidAddress
	}

	return nil
}

type CreateContractRequest struct {
	Address         string `json:"address" binding:"required"`
	Network         string `json:"network" binding:"required"`
	SourceCode      string `json:"source_code" binding:"required"`
	ContractName    string `json:"contract_name" binding:"required"`
	CompilerVersion string `json:"compiler_version" binding:"required"`
}

func (r *CreateContractRequest) Validate() error {
	if r.Address == "" {
		return errors.ErrInvalidAddress
	}

	if !strings.HasPrefix(r.Address, "0x") {
		return errors.ErrInvalidAddress
	}

	if len(r.Address) != 42 {
		return errors.ErrInvalidAddress
	}

	if strings.TrimSpace(r.Network) == "" {
		return errors.ErrInvalidNetwork
	}

	if strings.TrimSpace(r.SourceCode) == "" {
		return errors.ErrEmptySourceCode
	}

	if strings.TrimSpace(r.ContractName) == "" {
		return errors.ErrInvalidContractName
	}

	if strings.TrimSpace(r.CompilerVersion) == "" {
		return errors.ErrInvalidCompilerVersion
	}

	return nil
}
