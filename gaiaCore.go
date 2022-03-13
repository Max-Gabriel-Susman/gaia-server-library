package gaia

/**************************************************

gaiaCore : gaiaCore is an inference engine, it
encapsulates the specified ML solution of a
particular run and the operations and criteria
required for that solution to produce inferences
about the state of the gaia run.

* gaiaCores full scope of functionality is not yet
implemented.

**************************************************/
type gaiaCore struct {
	config gaiaCoreConfig
}

func (gc *gaiaCore) executeInferenceCycle(inference gaiaInference) gaiaInference {

	return inference
}

func (gc *gaiaCore) loadInference(inference gaiaInference) {

}
