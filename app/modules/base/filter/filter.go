package filter

const (
	IgnoreLimit  = -1
	IgnoreOffset = -1
)

type Where map[string][]interface{}
type Joins map[string][]interface{}
type Keys map[string]bool
type Groups map[string]bool
type Filter interface {
	GetWhere() Where
	GetJoins() Joins
	GetLimit() int
	GetOffset() int
	GetOrderBy() []string
	GetKeys() Keys
	GetGroups() string
}
