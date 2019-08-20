package tekton

import (
	"github.com/golang/glog"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func (t *Tekton) AddFunc(obj interface{}) {
	pr := obj.(*v1alpha1.PipelineRun)
	glog.Infof("Add pipelinerun :%s", pr)
}

func (t *Tekton) UpdateFun(oldObj, newObj interface{}) {
	npr := newObj.(*v1alpha1.PipelineRun)
	opr := newObj.(*v1alpha1.PipelineRun)
	glog.Infof("UpdatePipelineRun :%s \n %s", opr, npr)
}
func (t *Tekton) DeleteFunc(obj interface{}) {
	pr := obj.(*v1alpha1.PipelineRun)
	glog.Infof("Delete PipelineRun:%s ", pr)
}
