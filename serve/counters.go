package serve

import "expvar"

type counters struct {
	reqs        *expvar.Int
	resps       *expvar.Int
	pages       *expvar.Map
	statusCodes *expvar.Map
}

func newCounters() counters {
	return counters{
		reqs:        expvar.NewInt("requests"),
		resps:       expvar.NewInt("responses"),
		pages:       expvar.NewMap("pages"),
		statusCodes: expvar.NewMap("statusCodes"),
	}
}
