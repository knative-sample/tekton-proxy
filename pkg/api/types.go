package api

import "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"

type StartTektonPipelineRunArgs struct {
}

type StartTektonPipelineRunResponse struct {
}

type CommonResponse struct {
}

type StartApiArgs struct {
	Port int
}

type AddPipelineRunArgs struct {
	GitRepo          string
	GitBranch        string
	GitUserName      string
	GitPassword      string
	NewImageFullName string
	NewImageTag      string
	CRUser           string
	CRPassword       string

	Name      string
	Namespace string
}

type AddPipelineRunResponse struct {
	Response
	UID  string
	Name string
}

type GetPingPipelineRunArgs struct {
	Namespace string
}

type GetPingPipelineRunResponse struct {
	Response
	Len          int
	PipeLineRuns []v1alpha1.PipelineRun
}

type Response struct {
	RequestId string
}
