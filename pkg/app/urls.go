package app

import "github.com/go-chi/chi"

const (
	// api entrypoints
	ApiV1      = "/api/v1/"
	ProcessCSV = "process_csv/"
	Result     = "result/{id}/"

	FullProcessCSV = ApiV1 + ProcessCSV
	FullResult     = ApiV1 + Result
)

func (a *App) MountHandlers() {
	a.router.Route(ApiV1, func(r chi.Router) {
		r.Post(ProcessCSV, a.HandleProcessCSV)
		r.Get(Result, a.HandleResult)
	})

}
