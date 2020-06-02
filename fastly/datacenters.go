package fastly

// Datacenter represents a Fastly datacenter and its location
type Datacenter struct {
	Code        string `mapstructure:"code"`
	Coordinates struct {
		Latitude  float64 `mapstructure:"latitude"`
		Longitude float64 `mapstructure:"longitude"`
	} `mapstructure:"coordinates"`
	Group  string `mapstructure:"group"`
	Name   string `mapstructure:"name"`
	Shield string `mapstructure:"shield"`
}

func (c *Client) ListDatacenters() ([]*Datacenter, error) {
	resp, err := c.Get("/datacenters", nil)
	if err != nil {
		return nil, err
	}

	var dcs []*Datacenter
	if err := decodeBodyMap(resp.Body, &dcs); err != nil {
		return nil, err
	}

	return dcs, nil
}
