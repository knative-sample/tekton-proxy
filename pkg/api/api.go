package api

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/knative-sample/tekton-proxy/pkg/tekton"
)

type Api struct {
	Port   int
	Tekton *tekton.Tekton
}

func StartApi(args *StartApiArgs) error {
	t, err := tekton.NewTekton()
	if err != nil {
		glog.Errorf("StartApi error:%s", err.Error())
		return err
	}
	api := &Api{
		Port:   8080,
		Tekton: t,
	}
	if args.Port != 0 {
		api.Port = args.Port
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.Hello).Methods("GET")
	router.HandleFunc("/health/liveness", api.LivenessProbe).Methods("GET")
	router.HandleFunc("/health/readiness", api.ReadinessProbe).Methods("GET")

	// set routes
	api.SetPipelineRunRoutesV1(router)
	api.SetTenantRoutesV1(router)

	go func() {
		glog.Infof("api server is started")
		http.Handle("/", router)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", api.Port), router); err != nil {
			glog.Errorf("API Listen error:%s ", err.Error())
		}
	}()

	return nil
}

func (a *Api) LivenessProbe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "liveness")
}

func (a *Api) ReadinessProbe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "readiness")
}

func (a *Api) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello tekton-proxy!")
}
