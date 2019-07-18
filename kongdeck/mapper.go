package kongdeck

import (
	"regexp"
	"strconv"

	"github.com/davidcv5/kongup/kongfig"
	deck "github.com/hbagdi/deck/file"
	"github.com/hbagdi/go-kong/kong"
)

func fromKongfig(kongfig *kongfig.Config) (*deck.Content, error) {
	var services []deck.Service
	var consumers []deck.Consumer
	for _, api := range kongfig.Apis {
		service, err := mapService(&api)
		services = append(services, *service)
		if err != nil {
			return nil, err
		}
	}
	for _, consumer := range kongfig.Consumers {
		c, err := mapConsumer(&consumer)
		consumers = append(consumers, *c)
		if err != nil {
			return nil, err
		}
	}
	var file deck.Content
	file.Services = services
	file.Consumers = consumers
	return &file, nil
}

func mapService(api *kongfig.API) (*deck.Service, error) {
	s := &deck.Service{}
	s.Name = kong.String(api.Name)
	s.ConnectTimeout = kong.Int(api.Attributes.UpstreamConnectTimeout)
	s.WriteTimeout = kong.Int(api.Attributes.UpstreamSendTimeout)
	s.ReadTimeout = kong.Int(api.Attributes.UpstreamReadTimeout)
	s.Retries = kong.Int(api.Attributes.Retries)

	exp := regexp.MustCompile(`(?P<scheme>https?)?://(?P<host>[\w\.\{\}]+)(?P<port>:\d{4,5})?(?P<path>/[^#]*)?`)
	match := exp.FindStringSubmatch(api.Attributes.UpstreamURL)
	for i, name := range exp.SubexpNames() {
		if i == 0 {
			continue
		}
		if name == "scheme" {
			s.Protocol = kong.String(match[i])
		} else if name == "host" {
			s.Host = kong.String(match[i])
		} else if name == "port" {
			if p, err := strconv.Atoi(match[i]); err == nil {
				s.Port = kong.Int(p)
			}
		} else if name == "path" {
			s.Path = kong.String(match[i])
		}

	}

	r, err := mapRoute(api)
	if err != nil {
		return nil, err
	}
	s.Routes = []*deck.Route{r}

	plugins, err := mapPlugin(&api.Plugins)
	if err != nil {
		return nil, err
	}
	s.Plugins = plugins

	return s, nil
}

func mapRoute(api *kongfig.API) (*deck.Route, error) {
	r := &deck.Route{}
	r.Name = kong.String(api.Name)
	r.Service = &kong.Service{
		Name: kong.String(api.Name),
	}
	r.PreserveHost = kong.Bool(api.Attributes.PreserveHost)
	r.StripPath = kong.Bool(api.Attributes.StripURI)
	r.Paths = kong.StringSlice(api.Attributes.Uris...)

	return r, nil
}

func mapPlugin(plugins *[]kongfig.Plugin) ([]*deck.Plugin, error) {
	result := []*deck.Plugin{}
	if plugins == nil || len(*plugins) == 0 {
		return result, nil
	}
	for _, p := range *plugins {
		plugin := &deck.Plugin{}
		plugin.Name = kong.String(p.Name)
		plugin.Enabled = kong.Bool(p.Attributes.Enabled)
		p.Attributes.Config.DeepCopyInto(&plugin.Config)
		result = append(result, plugin)
	}
	return result, nil
}

func mapConsumer(con *kongfig.Consumer) (*deck.Consumer, error) {
	c := &deck.Consumer{}
	c.Username = kong.String(con.Username)
	if con.CustomID != nil {
		c.CustomID = kong.String(*con.CustomID)
	}
	return c, nil
}
