package lib

type (
	env struct {
		AllowedMethods *[]string
		IsDev *bool
	}
)
func (i *env) SetEnv (isDev *bool, allowedMethods *[]string) {
	i.AllowedMethods = allowedMethods
	i.IsDev = isDev
}

var Env env

