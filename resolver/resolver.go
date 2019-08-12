package resolver

type Resolver interface {
	Resolve(data []byte) (resolvedData map[interface{}]interface{}, err error)
}
