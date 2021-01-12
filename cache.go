package mc

type Cache interface {
	Get(interface{}) (interface{}, bool)
	Add(interface{}, interface{})
	Remove(interface{})
}
