package cypher_go_dsl

import (
	"errors"
	"strings"
)

type Parameter struct {
	name   string
	key    string
	err    error
	notNil bool
}

func ParameterCreate(name string) Parameter {
	if name == "" {
		return ParameterError(errors.New("the name of the parameter is required"))
	}
	if strings.HasPrefix(name, "$") {
		return ParameterCreate(name[1:])
	}
	parameter := Parameter{
		name:   name,
		notNil: true,
	}
	parameter.key = getAddress(&parameter)
	return parameter
}

func ParameterError(err error) Parameter {
	return Parameter{
		err: err,
	}
}

func (p Parameter) getError() error {
	return p.err
}

func (p Parameter) accept(visitor *CypherRenderer) {
	visitor.enter(p)
	visitor.leave(p)
}

func (p Parameter) enter(renderer *CypherRenderer) {
	renderer.append("$").append(p.name)
}

func (p Parameter) leave(renderer *CypherRenderer) {
}

func (p Parameter) getKey() string {
	return p.key
}

func (p Parameter) isNotNil() bool {
	return p.notNil
}

func (p Parameter) GetExpressionType() ExpressionType {
	return "Parameter"
}
