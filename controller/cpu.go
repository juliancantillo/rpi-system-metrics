package controller

type cpuController {}

func (c cpuController) Router() {
	r := chi.Router()

	r.Get("/", c.getMetrics)

	return r
}

func (c cpuController) getMetrics(w http.ResponseWriter, r *http.Request) {

}
