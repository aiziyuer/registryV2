package registry

import (
	"github.com/aiziyuer/registryV2/impl/common"
	"github.com/aiziyuer/registryV2/impl/handler"
	"net/http"
)

type (
	Endpoint struct {
		Schema string
		Host   string
	}

	Registry struct {
		Auth          *common.Auth
		Endpoint      *Endpoint
		Client        *http.Client
		HandlerFacade *handler.Facade
	}
)

func NewClient(c *http.Client, endpoint *Endpoint, auth *common.Auth) *Registry {

	return &Registry{
		Auth: auth,
		Endpoint: &Endpoint{
			Schema: endpoint.Schema,
			Host:   endpoint.Host,
		},
		Client: c,
		HandlerFacade: &handler.Facade{
			Client: c,
			PatternHandlerMap: map[string]handler.Handler{
				".+": {
					RequestFns: map[string]handler.RequestHandlerFunc{
						"common": (&handler.AuthRequestHandler{
							Client: c,
							Auth:   auth,
						}).RequestHandlerFunc,
					},
					ResponseFns: map[string]handler.ResponseHandlerFunc{},
				},
			},
		},
	}
}

func (r *Registry) Ping() error {

	q, err := handler.NewApiRequest(`
	{
		"Method": "GET",
		"Path": "/v2/",
		"Schema": "{{ .Schema }}",
		"Host": "{{ .Host }}",
		"Header": {
			"Content-Type": "application/json; charset=utf-8",
			"Authorization": "{{ .Token }}",
		},
		"Body": "",
	}
`, handler.ApiRequestInput{
		"Schema": r.Endpoint.Schema,
		"Host":   r.Endpoint.Host,
		"Token":  "",
	})
	if err != nil {
		return err
	}

	req, _ := q.Wrapper()
	resp, _ := r.HandlerFacade.Do(req)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	return err
}

func (r *Registry) Login() error {

	if err := r.Ping(); err != nil {
		return err
	}

	return nil
}

type ResponseFunc func(resp *http.Response)

func (r *Registry) Do(template string, input *handler.ApiRequestInput, fn ResponseFunc) error {

	q, err := handler.NewApiRequest(template, *input)
	if err != nil {
		return err
	}

	req, _ := q.Wrapper()
	resp, err := r.HandlerFacade.Do(req)
	if err != nil {
		return err
	}

	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	fn(resp)

	return nil

}
