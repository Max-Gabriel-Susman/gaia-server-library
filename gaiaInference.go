package gaia

/**************************************************

gaiaInference : an instance of gaiaInference
represents an inference that may be performed by
gaiaCore and then passed on to --?

**************************************************/
type gaiaInference struct {
	inference inference
	evidence  evidence
}
