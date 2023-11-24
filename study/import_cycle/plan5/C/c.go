package C

type StringFunc func(s string) string

var funcMap = make(map[string]StringFunc)

func RegisterFunc(key string, f StringFunc) {
	funcMap[key] = f
}

func GetFunc(key string) StringFunc {
	return funcMap[key]
}
