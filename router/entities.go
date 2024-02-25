package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	// contImpl "github.com/tonychill/ifitu/services/content/service"
	corImpl "github.com/tonychill/ifitu/services/coordinator/service"
	finImpl "github.com/tonychill/ifitu/services/finance/service"
	// idImpl "github.com/tonychill/ifitu/services/identity/service"
)

type Router interface {
	Shutdown(ctx context.Context) error
}

type _Store struct {
	fs *session.Session
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}
type ServiceImplementations struct {
	CoordinatorService *corImpl.ServiceImpl
	// ContentService    *contImpl.ServiceImpl
	// IdentityService   *idImpl.ServiceImpl
	FinanceService *finImpl.ServiceImpl
}

// Provides a landing zone for incoming requests. The router will route requests to the
// concierge and then from there to the appropriate service(s).
type routerImpl struct {
	grpcSvrShutdown func()
	app             *fiber.App
	GrpcServerPort  string `envconfig:"SVC_GRPC_PORT" required:"true"`
	corImpl         *corImpl.ServiceImpl

	// The following services are here only so that they're registered by the router
	// and therefor can be called by the concierge. They should not be called directly.
	// This allows them to be called by the concierge. While it is possible for the
	// router to allow external callers to call the services independantly, it is not recommended.
	// The router's main purpose is to route requests to the concierge. The concierge will then
	// call the domain services as needed.
	// contImpl *contImpl.ServiceImpl
	// idImpl   *idImpl.ServiceImpl
	finImpl *finImpl.ServiceImpl
}

type Checkout struct {
	Rules           string      `json:"rules"`
	SessionLink     string      `json:"session_link"`
}
type Task struct {
	ID           string      `json:"id"`
	CustomID     interface{} `json:"custom_id"`
	CustomItemID int         `json:"custom_item_id"`
	Name         string      `json:"name"`
	TextContent  string      `json:"text_content"`
	Description  string      `json:"description"`
	Status       struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Color      string `json:"color"`
		Orderindex int    `json:"orderindex"`
		Type       string `json:"type"`
	} `json:"status"`
	Orderindex  string      `json:"orderindex"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
	DateClosed  interface{} `json:"date_closed"`
	DateDone    interface{} `json:"date_done"`
	Archived    bool        `json:"archived"`
	Creator     struct {
		ID             int         `json:"id"`
		Username       string      `json:"username"`
		Color          string      `json:"color"`
		Email          string      `json:"email"`
		ProfilePicture interface{} `json:"profilePicture"`
	} `json:"creator"`
	Assignees []interface{} `json:"assignees"`
	Watchers  []struct {
		ID             int         `json:"id"`
		Username       string      `json:"username"`
		Color          string      `json:"color"`
		Initials       string      `json:"initials"`
		Email          string      `json:"email"`
		ProfilePicture interface{} `json:"profilePicture"`
	} `json:"watchers"`
	Checklists   []interface{} `json:"checklists"`
	Tags         []interface{} `json:"tags"`
	Parent       interface{}   `json:"parent"`
	Priority     interface{}   `json:"priority"`
	DueDate      interface{}   `json:"due_date"`
	StartDate    interface{}   `json:"start_date"`
	Points       interface{}   `json:"points"`
	TimeEstimate interface{}   `json:"time_estimate"`
	TimeSpent    int           `json:"time_spent"`
	CustomFields []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		TypeConfig struct {
		} `json:"type_config"`
		DateCreated    string `json:"date_created"`
		HideFromGuests bool   `json:"hide_from_guests"`
		Value          string `json:"value"`
		Required       bool   `json:"required"`
	} `json:"custom_fields"`
	Dependencies []interface{} `json:"dependencies"`
	LinkedTasks  []interface{} `json:"linked_tasks"`
	Locations    []interface{} `json:"locations"`
	TeamID       string        `json:"team_id"`
	URL          string        `json:"url"`
	Sharing      struct {
		Public               bool        `json:"public"`
		PublicShareExpiresOn interface{} `json:"public_share_expires_on"`
		PublicFields         []string    `json:"public_fields"`
		Token                interface{} `json:"token"`
		SeoOptimized         bool        `json:"seo_optimized"`
	} `json:"sharing"`
	PermissionLevel string `json:"permission_level"`
	List            struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"list"`
	Project struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"project"`
	Folder struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
	Attachments []interface{} `json:"attachments"`
}