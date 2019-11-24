package sorting


type Sortable interface {
	Greater(a, b Sortable) bool
}