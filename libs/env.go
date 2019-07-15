package lib

type (
	env struct {
		AllowedMethods *[]string
		IsDev *bool
	}
	InsertEnv interface {
		GetAllowedMethods () []string
		GetIsDev () * bool
	}
)
func (i *env) SetEnv (data InsertEnv) {
	methods := data.GetAllowedMethods()
	i.AllowedMethods = &methods
	i.IsDev = data.GetIsDev()
}

var Env env

