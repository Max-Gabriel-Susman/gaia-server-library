package gaia

/**************************************************

gaiaState : an instance of gaiaState can be inferred
upon by gaiaCore

**************************************************/
type gaiaState struct {
	stateID     string
	content     string
	dateCreated string
	annotation  stateAnnotation
}

func newGaiaState() (*gaiaState, error) {
	return &gaiaState{"", "", "", stateAnnotation{false}}, nil
}
