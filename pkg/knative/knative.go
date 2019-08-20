package knative

import (
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

type KnativeService struct {

}

func (k *KnativeService) AddPipelineRun(s *v1beta1.Service) error {
	return nil
}

func (k *KnativeService) RemovePipelineRun(s *v1beta1.Service) error {
	return nil
}

func (k *KnativeService) ListFinishedPipelineRun() ([]v1beta1.Service, error) {
	return nil, nil
}
