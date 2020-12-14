package cypher_go_dsl

func VisitIfNotNull(dest interface{}, visitor *CypherRenderer) {
	visitable, isVisitable := dest.(Visitable)
	if isVisitable && visitable.isNotNil() {
		visitable.accept(visitor)
	}
}

type CanHasError interface {
	getError() error
}

type Visitable interface {
	CanHasError
	accept(visitor *CypherRenderer)
	enter(renderer *CypherRenderer)
	leave(renderer *CypherRenderer)
	getKey() string
	isNotNil() bool
}

type SubVisitable interface {
	Visitable
	PrepareVisit(visitable Visitable) Visitable
}
