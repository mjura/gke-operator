package services

import (
	"context"

	"golang.org/x/oauth2"
	gkeapi "google.golang.org/api/container/v1"
	"google.golang.org/api/option"
)

type GKEClusterServiceInerface interface {
	Create(parent string, cluster *gkeapi.CreateClusterRequest) (*gkeapi.Operation, error)
	List(parent string) (*gkeapi.ListClustersResponse, error)
	Get(name string) (*gkeapi.Cluster, error)
}

type gkeClusterService struct {
	svc *gkeapi.ProjectsLocationsClustersService
}

func NewGKEClusterService(ctx context.Context, ts oauth2.TokenSource) (GKEClusterServiceInerface, error) {
	svc, err := gkeapi.NewService(ctx, option.WithHTTPClient(oauth2.NewClient(ctx, ts)))
	if err != nil {
		return nil, err
	}
	return &gkeClusterService{
		svc: svc.Projects.Locations.Clusters,
	}, nil
}

func (g *gkeClusterService) Create(parent string, cluster *gkeapi.CreateClusterRequest) (*gkeapi.Operation, error) {
	return g.svc.Create(parent, cluster).Do()
}

func (g *gkeClusterService) List(parent string) (*gkeapi.ListClustersResponse, error) {
	return g.svc.List(parent).Do()
}

func (g *gkeClusterService) Get(name string) (*gkeapi.Cluster, error) {
	return g.svc.Get(name).Do()
}

type GKENodePoolServiceInerface interface {
	Create(parent string, cluster *gkeapi.CreateNodePoolRequest) (*gkeapi.Operation, error)
	List(parent string) (*gkeapi.ListNodePoolsResponse, error)
	Get(name string) (*gkeapi.NodePool, error)
}

type gkeNodePoolService struct {
	svc *gkeapi.ProjectsLocationsClustersNodePoolsService
}

func NewGKENodePoolService(ctx context.Context, ts oauth2.TokenSource) (GKENodePoolServiceInerface, error) {
	svc, err := gkeapi.NewService(ctx, option.WithHTTPClient(oauth2.NewClient(ctx, ts)))
	if err != nil {
		return nil, err
	}
	return &gkeNodePoolService{
		svc: svc.Projects.Locations.Clusters.NodePools,
	}, nil
}

func (g *gkeNodePoolService) Create(parent string, nodepool *gkeapi.CreateNodePoolRequest) (*gkeapi.Operation, error) {
	return g.svc.Create(parent, nodepool).Do()
}

func (g *gkeNodePoolService) List(parent string) (*gkeapi.ListNodePoolsResponse, error) {
	return g.svc.List(parent).Do()
}

func (g *gkeNodePoolService) Get(name string) (*gkeapi.NodePool, error) {
	return g.svc.Get(name).Do()
}
