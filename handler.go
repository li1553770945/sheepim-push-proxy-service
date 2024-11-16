package main

import (
	"context"
	"github.com/li1553770945/sheepim-push-proxy-service/biz/infra/container"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/project"
)

// ProjectsServiceImpl implements the last service interface defined in the IDL.
type ProjectServiceImpl struct{}

// GetProjects implements the ProjectsServiceImpl interface.
func (s *ProjectServiceImpl) GetProjects(ctx context.Context, req *project.ProjectsReq) (resp *project.ProjectsResp, err error) {
	// TODO: Your code here...
	App := container.GetGlobalContainer()
	resp, err = App.ProjectService.GetProjects(ctx, req)
	return
}

// GetProjectNum implements the ProjectsServiceImpl interface.
func (s *ProjectServiceImpl) GetProjectNum(ctx context.Context) (resp *project.ProjectNumResp, err error) {
	// TODO: Your code here...
	App := container.GetGlobalContainer()
	resp, err = App.ProjectService.GetProjectNum(ctx)
	return
}
