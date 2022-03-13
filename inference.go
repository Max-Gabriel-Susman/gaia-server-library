package gaia

/**************************************************

inference : an instance of inference that has been
produced by gaiaCore

**************************************************/
type inference struct {
	content     []float64
	name        string
	dateCreated string
	model       machineLearningModel
}
