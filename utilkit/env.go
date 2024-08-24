package utilkit

import "strings"

type Env uint

const (
	EnvDev Env = iota + 1
	EnvTest
	EnvPre
	EnvProd
)

var EnvMap = map[Env]string{
	EnvDev:  "dev",
	EnvTest: "test",
	EnvPre:  "pre",
	EnvProd: "prod",
}

func FromEnv(env Env) string {
	if result, ok := EnvMap[env]; ok {
		return result
	}

	return EnvMap[EnvDev]
}

func ToEnv(sEnv string) Env {
	for k, v := range EnvMap {
		if v == strings.ToLower(sEnv) {
			return k
		}
	}

	return EnvDev
}

func IsLiveEnv(sEnv string) bool {
	return sEnv == EnvMap[EnvProd]
}

func GetEnvList() (envList []string) {
	for _, v := range EnvMap {
		envList = append(envList, v)
	}

	return
}
