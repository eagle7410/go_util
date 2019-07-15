package lib

type (
	env struct {
		AllowedMethods *[]string
		IsDev *bool
	}
	InsertEnv interface {
		GetAllowedMethods () *[]string
		GetIsDev () * bool
	}
)
func (i *env) SetEnv (data InsertEnv) {
	i.AllowedMethods = data.GetAllowedMethods()
	i.IsDev = data.GetIsDev()
}

var Env env

