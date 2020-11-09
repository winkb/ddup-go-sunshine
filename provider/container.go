package container

type ProviderRegisterFuns []interface{}

type ProviderInterface interface {
	Register() ProviderRegisterFuns
}
