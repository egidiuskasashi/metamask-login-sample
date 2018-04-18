package auth

import (
	"github.com/m0t0k1ch1/metamask-login-sample/application"
	"github.com/m0t0k1ch1/metamask-login-sample/domain/common"
)

type ChallengeInput struct {
	*application.AddressHexInput
}

func NewChallengeInput(addressHex string) *ChallengeInput {
	return &ChallengeInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
	}
}

type AuthorizeInput struct {
	*application.AddressHexInput
	SigHex string
}

func NewAuthorizeInput(addressHex, sigHex string) *AuthorizeInput {
	return &AuthorizeInput{
		AddressHexInput: application.NewAddressHexInput(addressHex),
		SigHex:          sigHex,
	}
}

func (in *AuthorizeInput) Validate() error {
	if err := in.AddressHexInput.Validate(); err != nil {
		return err
	}
	if err := common.ValidateSignatureHex(in.SigHex); err != nil {
		return err
	}
	return nil
}

func (in *AuthorizeInput) Signature() common.Signature {
	return common.NewSignatureFromHex(in.SigHex)
}
