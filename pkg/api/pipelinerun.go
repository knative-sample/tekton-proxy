package api

import (
	"fmt"
	"net/http"

	"encoding/json"

	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/knative-sample/tekton-proxy/pkg/tekton"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *Api) SetPipelineRunRoutesV1(r *mux.Router) {
	router := r.PathPrefix("/v1").Subrouter()

	router.HandleFunc("/ping/{namespace}/pipelineruns/", a.AddPingPipelineRun).Methods("POST")
	router.HandleFunc("/ping/{namespace}/pipelineruns/", a.GetPingPipelineRun).Methods("GET")

	router.HandleFunc("/pipelinerun", a.AddPipelineRun).Methods("POST")
	router.HandleFunc("/pipelineruns", a.GetPipelineRun).Methods("GET")
}

func (a *Api) GetPingPipelineRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	st := time.Now().Nanosecond()
	ns := ""
	if params["namespace"] != "_all_" {
		ns = params["namespace"]
	}
	prs, err := a.Tekton.ListFinishedPipelineRun(ns)
	if err != nil {
		msg := fmt.Sprintf("ListFinishedPipelineRun Error : %v", err)
		glog.Error(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	resp := GetPingPipelineRunResponse{
		Len:          len(prs),
		PipeLineRuns: prs,
	}

	et := time.Now().Nanosecond()
	w.Header().Set("Response-Time", GetResponseTime(st, et))
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		glog.Errorf("response request error:%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST-3")
}
func GetResponseTime(st, et int) string {
	return fmt.Sprintf("%.3f", (float64(et)-float64(st))/1000000)

}

func (a *Api) AddPingPipelineRun(w http.ResponseWriter, r *http.Request) {
	st := time.Now().Nanosecond()
	args := &AddPipelineRunArgs{}
	if err := json.NewDecoder(r.Body).Decode(args); err != nil {
		glog.Errorf("AddPipelineRun parse args error:%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check args
	if args.Namespace == "" {
		msg := "AddPingPipelineRun args error, namespace is empty"
		glog.Errorf(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if args.Name == "" {
		msg := "AddPingPipelineRun args error, name is empty"
		glog.Errorf(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	pr := &v1alpha1.PipelineRun{
		ObjectMeta: v1.ObjectMeta{
			Name:      args.Name,
			Namespace: args.Namespace,
		},
		Spec: v1alpha1.PipelineRunSpec{
			PipelineRef: v1alpha1.PipelineRef{
				Name:       "tutorial-pipeline",
				APIVersion: tekton.ApiVersion,
			},
		},
	}
	p, err := a.Tekton.AddPipelineRun(pr)
	if err != nil {
		glog.Errorf("Run AddPipelineRun parse args error:%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := AddPipelineRunResponse{
		Name: p.Name,
		UID:  string(p.UID),
	}

	et := time.Now().Nanosecond()
	w.Header().Set("Response-Time", GetResponseTime(st, et))
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		glog.Errorf("response request error:%s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *Api) GetPipelineRun(w http.ResponseWriter, r *http.Request) {
}

func (a *Api) AddPipelineRun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello tekton-proxy!")
}
