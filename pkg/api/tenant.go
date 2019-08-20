package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *Api) SetTenantRoutesV1(r *mux.Router) {
	router := r.PathPrefix("/v1/tenant").Subrouter()

	router.HandleFunc("/init", a.NewTenant).Methods("POST")
}

func (a *Api) NewTenant(w http.ResponseWriter, r *http.Request) {
	// TODO 创建的所有 k8s 资源都要带上 label，方便后续的管理和扩展
	fmt.Fprint(w, "create new tenant")
}
