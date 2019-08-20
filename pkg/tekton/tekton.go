package tekton

import (
	"time"

	"github.com/golang/glog"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	clientsetversioned "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"

	"github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/cache"
	ctlruntimeconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

const (
	ApiVersion = "tekton.dev/v1alpha1"
)

type Tekton struct {
	TektonClient *versioned.Clientset
}

func NewTekton() (*Tekton, error) {
	cfg := ctlruntimeconfig.GetConfigOrDie()

	tektonClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building Serving clientset: %v", err)
	}

	t := &Tekton{
		TektonClient: tektonClient,
	}
	go func() {
		t.listWath()
	}()

	return t, nil
}

func (t *Tekton) listWath() error {
	restConfig := ctlruntimeconfig.GetConfigOrDie()

	versionedClientset, err := clientsetversioned.NewForConfig(restConfig)
	if err != nil {
		glog.Fatalf("failed to new versioned clientset: %s", err)
	}

	factory := externalversions.NewSharedInformerFactory(versionedClientset, time.Second*1800)
	informer := factory.Tekton().V1alpha1().PipelineRuns().Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    t.AddFunc,
		UpdateFunc: t.UpdateFun,
		DeleteFunc: t.DeleteFunc,
	})

	go func(ifm cache.SharedInformer) {
		stopper := make(chan struct{})
		defer close(stopper)
		ifm.Run(stopper)
	}(informer)

	return nil
}

func (t *Tekton) AddPipelineRun(pr *v1alpha1.PipelineRun) (p *v1alpha1.PipelineRun, err error) {
	pipelineRun, err := t.TektonClient.TektonV1alpha1().PipelineRuns(pr.Namespace).Get(pr.Name, metav1.GetOptions{})
	if err == nil {
		return pipelineRun, nil
	}

	if !errors.IsNotFound(err) {
		glog.Errorf("get pipelineRun error: %s ", err.Error())
		return nil, err
	}

	pipelineRun, err = t.TektonClient.TektonV1alpha1().PipelineRuns(pr.Namespace).Create(pr)
	if err == nil {
		return pipelineRun, nil
	}
	glog.Errorf("create PipelineRun error: %s", err.Error())
	return nil, err
}

func (t *Tekton) RemovePipelineRun(pr *v1alpha1.PipelineRun) error {
	return nil
}

func (t *Tekton) ListFinishedPipelineRun(ns string) (list []v1alpha1.PipelineRun, err error) {
	l, err := t.TektonClient.TektonV1alpha1().PipelineRuns(ns).List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("ListFinishedPipelineRun error:%s", err.Error())
		return nil, err
	}

	return l.Items, nil
}
