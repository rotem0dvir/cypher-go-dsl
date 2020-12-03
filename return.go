package cypher_go_dsl

import "fmt"

type Return struct {
	distinct Distinct
	body     ReturnBody
	key      string
	notNil   bool
}

func ReturnCreate(distinctInstance Distinct, body ReturnBody) Return {
	return Return{
		distinct: distinctInstance,
		body:     body,
		notNil:   true,
	}
}

func (r Return) isNotNil() bool {
	return r.notNil
}

func (r Return) getKey() string {
	return r.key
}

func (r Return) accept(visitor *CypherRenderer) {
	r.key = fmt.Sprint(&r)
	(*visitor).enter(r)
	VisitIfNotNull(r.distinct, visitor)
	r.body.accept(visitor)
	(*visitor).leave(r)
}

func ReturnByMultiVariable(distinct bool, returnItems ExpressionList, order Order, skip Skip, limit Limit) Return {
	var distinctInstance Distinct
	if distinct {
		distinctInstance = DISTINCT_INSTANCE
	}
	body := ReturnBodyCreate(returnItems, order, skip, limit)
	return ReturnCreate(distinctInstance, body)
}

func (r Return) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString("RETURN  ")
}

func (r Return) leave(renderer *CypherRenderer) {
}
