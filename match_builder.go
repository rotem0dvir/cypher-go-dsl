package cypher_go_dsl

type MatchBuilder struct {
	patternList []PatternElement
	condition   ConditionBuilder
	optional    bool
	notNil      bool
}

func (builder MatchBuilder) buildMatch() Match {
	pattern := Pattern{patternElements: builder.patternList}
	conditionBuilder := builder.condition
	var optionalWhere Where = Where{}
	if conditionBuilder.condition != nil {
		builtCondition := conditionBuilder.buildCondition()
		optionalWhere = WhereCreate(builtCondition)
	}
	return Match{
		optional:      builder.optional,
		pattern:       pattern,
		optionalWhere: optionalWhere,
		notNil:        true,
	}
}
