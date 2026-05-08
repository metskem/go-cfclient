package client

import (
	"context"
	"net/url"

	"github.com/cloudfoundry/go-cfclient/v3/internal/path"
	"github.com/cloudfoundry/go-cfclient/v3/resource"
)

type RoutePolicyClient commonClient

// RoutePolicyListOptions list filters
type RoutePolicyListOptions struct {
	*ListOptions

	RouteGUIDs        Filter `qs:"route_guids"`
	SpaceGUIDs        Filter `qs:"space_guids"`
	SourceGUIDs       Filter `qs:"source_guids"`
	Sources           Filter `qs:"sources"`
	OrganizationGUIDs Filter `qs:"organization_guids"`
}

// NewRoutePolicyListOptions creates new options to pass to list
func NewRoutePolicyListOptions() *RoutePolicyListOptions {
	return &RoutePolicyListOptions{
		ListOptions: NewListOptions(),
	}
}

func (o RoutePolicyListOptions) ToQueryString() (url.Values, error) {
	return o.ListOptions.ToQueryString(o)
}

// Create a new route policy
func (c *RoutePolicyClient) Create(ctx context.Context, r *resource.RoutePolicyCreate) (*resource.RoutePolicy, error) {
	var routePolicy resource.RoutePolicy
	_, err := c.client.post(ctx, "/v3/route_policies", r, &routePolicy)
	if err != nil {
		return nil, err
	}
	return &routePolicy, nil
}

// Delete the specified route policy asynchronously and return a jobGUID.
func (c *RoutePolicyClient) Delete(ctx context.Context, guid string) (string, error) {
	return c.client.delete(ctx, path.Format("/v3/route_policies/%s", guid))
}

// First returns the first route policy matching the options or an error when less than 1 match
func (c *RoutePolicyClient) First(ctx context.Context, opts *RoutePolicyListOptions) (*resource.RoutePolicy, error) {
	return First[*RoutePolicyListOptions, *resource.RoutePolicy](opts, func(opts *RoutePolicyListOptions) ([]*resource.RoutePolicy, *Pager, error) {
		return c.List(ctx, opts)
	})
}

// Get the specified route policy
func (c *RoutePolicyClient) Get(ctx context.Context, guid string) (*resource.RoutePolicy, error) {
	var routePolicy resource.RoutePolicy
	err := c.client.get(ctx, path.Format("/v3/route_policies/%s", guid), &routePolicy)
	if err != nil {
		return nil, err
	}
	return &routePolicy, nil
}

// List pages all the route policies the user has access to
func (c *RoutePolicyClient) List(ctx context.Context, opts *RoutePolicyListOptions) ([]*resource.RoutePolicy, *Pager, error) {
	if opts == nil {
		opts = NewRoutePolicyListOptions()
	}

	var res resource.RoutePolicyList
	err := c.client.list(ctx, "/v3/route_policies", opts.ToQueryString, &res)
	if err != nil {
		return nil, nil, err
	}
	pager := NewPager(res.Pagination)
	return res.Resources, pager, nil
}

// ListAll retrieves all route policies the user has access to
func (c *RoutePolicyClient) ListAll(ctx context.Context, opts *RoutePolicyListOptions) ([]*resource.RoutePolicy, error) {
	if opts == nil {
		opts = NewRoutePolicyListOptions()
	}
	return AutoPage[*RoutePolicyListOptions, *resource.RoutePolicy](opts, func(opts *RoutePolicyListOptions) ([]*resource.RoutePolicy, *Pager, error) {
		return c.List(ctx, opts)
	})
}

// Single returns a single route policy matching the options or an error if not exactly 1 match
func (c *RoutePolicyClient) Single(ctx context.Context, opts *RoutePolicyListOptions) (*resource.RoutePolicy, error) {
	return Single[*RoutePolicyListOptions, *resource.RoutePolicy](opts, func(opts *RoutePolicyListOptions) ([]*resource.RoutePolicy, *Pager, error) {
		return c.List(ctx, opts)
	})
}
