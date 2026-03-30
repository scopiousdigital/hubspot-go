# hubspot-go

A Go client library for the [HubSpot API](https://developers.hubspot.com/docs/api/overview), ported from the official [@hubspot/api-client](https://www.npmjs.com/package/@hubspot/api-client) Node.js package.

## Installation

```bash
go get github.com/scopiousdigital/hubspot-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    hubspot "github.com/scopiousdigital/hubspot-go"
    "github.com/scopiousdigital/hubspot-go/crm"
)

func main() {
    client := hubspot.NewClient(
        hubspot.WithAccessToken("your-access-token"),
    )

    // Create a contact
    contact, err := client.CRM.Contacts.Create(context.Background(), &crm.SimplePublicObjectInputForCreate{
        Properties: crm.Properties{
            "email":     "jane@example.com",
            "firstname": "Jane",
            "lastname":  "Doe",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created contact: %s\n", contact.ID)
}
```

## Authentication

The client supports three authentication methods:

```go
// Private app access token (recommended)
client := hubspot.NewClient(hubspot.WithAccessToken("pat-xxx"))

// API key (legacy)
client := hubspot.NewClient(hubspot.WithAPIKey("xxx"))

// Developer API key
client := hubspot.NewClient(hubspot.WithDeveloperAPIKey("xxx"))

// You can also update credentials at runtime
client.SetAccessToken("new-token")
```

## Configuration

```go
client := hubspot.NewClient(
    hubspot.WithAccessToken("token"),

    // Automatic retries (0-6) on 5xx and 429 responses
    hubspot.WithRetries(3),

    // Client-side rate limiting (requests/sec, burst)
    hubspot.WithRateLimiter(10, 5),

    // Custom base URL
    hubspot.WithBaseURL("https://api.hubapi.com"),

    // Custom HTTP client
    hubspot.WithHTTPClient(&http.Client{Timeout: 60 * time.Second}),

    // Default headers for all requests
    hubspot.WithHeaders(map[string]string{
        "X-Custom-Header": "value",
    }),
)
```

## CRM

All standard CRM objects (contacts, companies, deals, tickets, etc.) share the same API surface:

### CRUD Operations

```go
import "github.com/scopiousdigital/hubspot-go/crm"

// Create
contact, err := client.CRM.Contacts.Create(ctx, &crm.SimplePublicObjectInputForCreate{
    Properties: crm.Properties{"email": "jane@example.com", "firstname": "Jane"},
})

// Read
contact, err := client.CRM.Contacts.GetByID(ctx, "501", &crm.GetByIDOptions{
    Properties:   []string{"email", "firstname", "lastname"},
    Associations: []string{"companies"},
})

// Update
contact, err := client.CRM.Contacts.Update(ctx, "501", &crm.SimplePublicObjectInput{
    Properties: crm.Properties{"firstname": "Updated"},
}, nil)

// Delete (archive)
err := client.CRM.Contacts.Archive(ctx, "501")

// GDPR delete
err := client.CRM.Contacts.GdprDelete(ctx, &crm.PublicGdprDeleteInput{
    ObjectID: "501",
})
```

### List & Pagination

```go
// Single page
page, err := client.CRM.Contacts.List(ctx, &crm.ListOptions{
    Limit:      100,
    Properties: []string{"email", "firstname"},
})

// Auto-paginate all results
all, err := client.CRM.Contacts.GetAll(ctx, &crm.GetAllOptions{
    Properties: []string{"email", "firstname"},
})
```

### Search

```go
results, err := client.CRM.Contacts.Search(ctx, &crm.PublicObjectSearchRequest{
    FilterGroups: []crm.FilterGroup{
        {
            Filters: []crm.Filter{
                {
                    PropertyName: "email",
                    Operator:     crm.FilterOperatorEQ,
                    Value:        "jane@example.com",
                },
            },
        },
    },
    Properties: []string{"email", "firstname"},
    Limit:      10,
})
```

### Batch Operations

```go
// Batch create
result, err := client.CRM.Contacts.BatchCreate(ctx, &crm.BatchCreateInput{
    Inputs: []crm.SimplePublicObjectInputForCreate{
        {Properties: crm.Properties{"email": "a@example.com"}},
        {Properties: crm.Properties{"email": "b@example.com"}},
    },
})

// Batch read
result, err := client.CRM.Contacts.BatchRead(ctx, &crm.BatchReadInput{
    Properties: []string{"email", "firstname"},
    Inputs:     []crm.ObjectID{{ID: "1"}, {ID: "2"}},
})

// Batch update
result, err := client.CRM.Contacts.BatchUpdate(ctx, &crm.BatchUpdateInput{
    Inputs: []crm.BatchObjectInput{
        {ID: "1", Properties: crm.Properties{"firstname": "Updated"}},
    },
})

// Batch archive
err := client.CRM.Contacts.BatchArchive(ctx, &crm.BatchArchiveInput{
    Inputs: []crm.ObjectID{{ID: "1"}, {ID: "2"}},
})

// Batch upsert
result, err := client.CRM.Contacts.BatchUpsert(ctx, &crm.BatchUpsertInput{
    Inputs: []crm.BatchObjectUpsertInput{
        {ID: "jane@example.com", IDProperty: "email", Properties: crm.Properties{"firstname": "Jane"}},
    },
})
```

### Merge

```go
result, err := client.CRM.Contacts.Merge(ctx, &crm.PublicMergeInput{
    PrimaryObjectID: "501",
    ObjectIDToMerge: "502",
})
```

### All CRM Object Types

The same methods are available on all standard CRM object services:

```go
client.CRM.Contacts
client.CRM.Companies
client.CRM.Deals
client.CRM.Tickets
client.CRM.Products
client.CRM.LineItems
client.CRM.Quotes
client.CRM.Calls
client.CRM.Emails
client.CRM.Meetings
client.CRM.Notes
client.CRM.Tasks
client.CRM.Communications
client.CRM.PostalMail
client.CRM.FeedbackSubmissions
client.CRM.Goals
client.CRM.Leads
client.CRM.DealSplits
client.CRM.Taxes
```

## Error Handling

```go
contact, err := client.CRM.Contacts.GetByID(ctx, "nonexistent", nil)
if err != nil {
    // Check for specific error types
    if hubspot.IsNotFound(err) {
        fmt.Println("Contact not found")
    } else if hubspot.IsRateLimited(err) {
        fmt.Println("Rate limited — retries are handled automatically with WithRetries()")
    }

    // Access full error details
    var apiErr *hubspot.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("Status: %d, Category: %s, Message: %s\n",
            apiErr.HTTPStatusCode, apiErr.Category, apiErr.Message)
    }
}
```

## API Coverage

| API Group | Package | Status |
|-----------|---------|--------|
| CRM (Contacts, Companies, Deals, Tickets, ...) | `crm` | Implemented |
| CRM (Owners, Properties, Pipelines, Schemas, Associations, Lists, ...) | `crm` | Implemented |
| CMS (Blogs, HubDB, Pages, Domains, ...) | `cms` | Implemented |
| Marketing (Forms, Emails, Events, Transactional) | `marketing` | Implemented |
| Automation (Actions) | `automation` | Implemented |
| Conversations (Visitor Identification) | `conversations` | Implemented |
| Events (Send, Batch) | `events` | Implemented |
| Files (Files, Folders) | `files` | Implemented |
| Settings (Users, Roles, Teams, Business Units) | `settings` | Implemented |
| OAuth (Tokens, Access Tokens, Refresh Tokens) | `oauth` | Implemented |
| Webhooks (Settings, Subscriptions) | `webhooks` | Implemented |
| Communication Preferences (Definitions, Status) | `communicationpreferences` | Implemented |

## Package Structure

```
hubspot-go/
├── client.go, config.go, errors.go        # Root: Client, options, error handling
├── http.go, retry.go                       # HTTP transport with retry & rate limiting
├── internal/api/                           # Shared Requester interface
├── crm/                                    # CRM objects, owners, properties, pipelines, ...
├── cms/                                    # Blogs, HubDB, pages, domains, ...
├── marketing/                              # Forms, emails, events, transactional
├── automation/                             # Workflow actions
├── conversations/                          # Visitor identification
├── events/                                 # Custom events
├── files/                                  # File & folder management
├── settings/                               # Users, roles, teams, business units
├── oauth/                                  # OAuth token management
├── webhooks/                               # Webhook settings & subscriptions
└── communicationpreferences/               # Subscription definitions & status
```

## License

MIT
