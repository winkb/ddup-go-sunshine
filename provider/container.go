package provider

type ProviderRegisterFuns []interface{}

type ProviderInterface interface {
	Register() ProviderRegisterFuns
}
