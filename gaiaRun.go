package gaia

type gaiaRun struct {
}

func NewGaiaRun() (*gaiaRun, error) {
	return &gaiaRun{}, nil
}

func (gr *gaiaRun) executeRun() (*gaiaState, int, error) {
	state, err := newGaiaState()
	return state, 0, err
}
