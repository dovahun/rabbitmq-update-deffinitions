package src

type Config struct {
	Bindings     []Binding    `json:"bindings"`
	Exchanges    []Exchange   `json:"exchanges"`
	GlobalParams []any        `json:"global_parameters"`
	Parameters   []any        `json:"parameters"`
	Permissions  []Permission `json:"permissions"`
	Policies     []Policy     `json:"policies"`
	Queues       []Queue      `json:"queues"`
	Users        []User       `json:"users"`
	Vhosts       []Vhost      `json:"vhosts"`
}
type Binding struct {
	Arguments       map[string]any `json:"arguments"`
	Destination     string         `json:"destination"`
	DestinationType string         `json:"destination_type"`
	RoutingKey      string         `json:"routing_key"`
	Source          string         `json:"source"`
	Vhost           string         `json:"vhost"`
}
type Exchange struct {
	Arguments  map[string]any `json:"arguments"`
	AutoDelete bool           `json:"auto_delete"`
	Durable    bool           `json:"durable"`
	Internal   bool           `json:"internal"`
	Name       string         `json:"name"`
	Type       string         `json:"type"`
	Vhost      string         `json:"vhost"`
}

type Permission struct {
	Configure string `json:"configure"`
	Read      string `json:"read"`
	User      string `json:"user"`
	Vhost     string `json:"vhost"`
	Write     string `json:"write"`
}

type Policy struct {
	Definition map[string]any `json:"definition"`
	Name       string         `json:"name"`
	Pattern    string         `json:"pattern"`
	Vhost      string         `json:"vhost"`
}

type Queue struct {
	Arguments  map[string]any `json:"arguments"`
	AutoDelete bool           `json:"auto_delete"`
	Durable    bool           `json:"durable"`
	Name       string         `json:"name"`
	Vhost      string         `json:"vhost"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Tags     string `json:"tags"`
}

type Vhost struct {
	Name string `json:"name"`
}
