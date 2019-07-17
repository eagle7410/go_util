package lib

type (
	DtoEnv interface {
		GetLinkIsDev () *bool
		GetLinkIsCorsAllowCredentials () *bool
		GetLinkAllowedMethods () *[]string
	}
	env struct {
		AllowedMethods *[]string
		IsCorsAllowCredentials, IsDev *bool
	}
)
func (i *env) SetEnv (iCnf DtoEnv) {
	i.AllowedMethods = iCnf.GetLinkAllowedMethods()
	i.IsDev = iCnf.GetLinkIsDev()
	i.IsCorsAllowCredentials = iCnf.GetLinkIsCorsAllowCredentials()
}

var Env env

