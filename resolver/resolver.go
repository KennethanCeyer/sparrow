package resolver

type Resolver interface {
	Tag(model interface{}) (err error)
	Resolve(data []byte) (resolvedData map[interface{}]interface{}, err error)
}
