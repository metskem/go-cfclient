package resource

type RoutePolicy struct {
	Source        string                   `json:"source"`
	Metadata      *Metadata                `json:"metadata"`
	Relationships RoutePolicyRelationships `json:"relationships"`
	Resource      `json:",inline"`
}

type RoutePolicyCreate struct {
	Relationships RoutePolicyRelationships `json:"relationships"`
	Source        string                   `json:"source"`
}

type RoutePolicyList struct {
	Pagination Pagination     `json:"pagination"`
	Resources  []*RoutePolicy `json:"resources"`
}

type RoutePolicyRelationships struct {
	Route        ToOneRelationship `json:"route,omitempty"`
	App          ToOneRelationship `json:"app,omitempty"`
	Space        ToOneRelationship `json:"space,omitempty"`
	Organization ToOneRelationship `json:"organization,omitempty"`
}
