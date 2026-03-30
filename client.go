package hubspot

import (
	"net/http"
	"time"

	"github.com/scopiousdigital/hubspot-go/automation"
	"github.com/scopiousdigital/hubspot-go/cms"
	"github.com/scopiousdigital/hubspot-go/communicationpreferences"
	"github.com/scopiousdigital/hubspot-go/conversations"
	"github.com/scopiousdigital/hubspot-go/crm"
	"github.com/scopiousdigital/hubspot-go/events"
	"github.com/scopiousdigital/hubspot-go/files"
	"github.com/scopiousdigital/hubspot-go/marketing"
	"github.com/scopiousdigital/hubspot-go/oauth"
	"github.com/scopiousdigital/hubspot-go/settings"
	"github.com/scopiousdigital/hubspot-go/webhooks"
)

const (
	defaultBaseURL   = "https://api.hubapi.com"
	defaultUserAgent = "hubspot-go/0.1.0"
	defaultTimeout   = 30 * time.Second
)

// Client is the HubSpot API client.
type Client struct {
	httpClient *httpClient

	CRM                      *crm.Service
	CMS                      *cms.Service
	Marketing                *marketing.Service
	Automation               *automation.Service
	Conversations            *conversations.Service
	Events                   *events.Service
	Files                    *files.Service
	Settings                 *settings.Service
	OAuth                    *oauth.Service
	Webhooks                 *webhooks.Service
	CommunicationPreferences *communicationpreferences.Service
}

// NewClient creates a new HubSpot API client with the given options.
func NewClient(opts ...Option) *Client {
	hc := &httpClient{
		client:    &http.Client{Timeout: defaultTimeout},
		baseURL:   defaultBaseURL,
		userAgent: defaultUserAgent,
		headers:   make(map[string]string),
	}

	c := &Client{httpClient: hc}

	for _, opt := range opts {
		opt(c)
	}

	c.initialize()
	return c
}

func (c *Client) initialize() {
	c.CRM = crm.NewService(c.httpClient)
	c.CMS = cms.NewService(c.httpClient)
	c.Marketing = marketing.NewService(c.httpClient)
	c.Automation = automation.NewService(c.httpClient)
	c.Conversations = conversations.NewService(c.httpClient)
	c.Events = events.NewService(c.httpClient)
	c.Files = files.NewService(c.httpClient)
	c.Settings = settings.NewService(c.httpClient)
	c.OAuth = oauth.NewService(c.httpClient)
	c.Webhooks = webhooks.NewService(c.httpClient)
	c.CommunicationPreferences = communicationpreferences.NewService(c.httpClient)
}

// SetAccessToken updates the access token.
func (c *Client) SetAccessToken(token string) {
	c.httpClient.accessToken = token
}

// SetAPIKey updates the API key.
func (c *Client) SetAPIKey(apiKey string) {
	c.httpClient.apiKey = apiKey
}

// SetDeveloperAPIKey updates the developer API key.
func (c *Client) SetDeveloperAPIKey(key string) {
	c.httpClient.developerAPIKey = key
}
