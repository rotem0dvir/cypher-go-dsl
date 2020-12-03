package cypher_go_dsl

import "fmt"

type AliasedExpression struct {
	delegate Expression
	alias    string
	key      string
	notNil   bool
}

func (aliased AliasedExpression) isNotNil() bool {
	return aliased.notNil
}

func (aliased AliasedExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (aliased AliasedExpression) As(newAlias string) AliasedExpression {
	return AliasedExpression{delegate: aliased.delegate,
		alias: newAlias}
}

func (aliased AliasedExpression) accept(visitor *CypherRenderer) {
	aliased.key = fmt.Sprint(&aliased)
	(*visitor).enter(aliased)
	NameOrExpression(aliased.delegate).accept(visitor)
	(*visitor).leave(aliased)
}

func (aliased AliasedExpression) getKey() string {
	return aliased.key
}

func (aliased AliasedExpression) enter(renderer *CypherRenderer) {
	panic("implement me")
}

func (aliased AliasedExpression) leave(renderer *CypherRenderer) {
	panic("implement me")
}
