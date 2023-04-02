package toggl

import "time"

type Me struct {
	ID                 int       `json:"id"`
	APIToken           string    `json:"api_token"`
	Email              string    `json:"email"`
	Fullname           string    `json:"fullname"`
	Timezone           string    `json:"timezone"`
	DefaultWorkspaceID int       `json:"default_workspace_id"`
	BeginningOfWeek    int       `json:"beginning_of_week"`
	ImageURL           string    `json:"image_url"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	OpenidEmail        string    `json:"openid_email"`
	OpenidEnabled      bool      `json:"openid_enabled"`
	CountryID          int       `json:"country_id"`
	At                 time.Time `json:"at"`
	IntercomHash       string    `json:"intercom_hash"`
	OauthProviders     []string  `json:"oauth_providers"`
	HasPassword        bool      `json:"has_password"`
}

type Workspace struct {
	ID                          int         `json:"id"`
	OrganizationID              int         `json:"organization_id"`
	Name                        string      `json:"name"`
	Profile                     int         `json:"profile"`
	Premium                     bool        `json:"premium"`
	BusinessWs                  bool        `json:"business_ws"`
	Admin                       bool        `json:"admin"`
	SuspendedAt                 interface{} `json:"suspended_at"`
	ServerDeletedAt             interface{} `json:"server_deleted_at"`
	DefaultHourlyRate           interface{} `json:"default_hourly_rate"`
	RateLastUpdated             interface{} `json:"rate_last_updated"`
	DefaultCurrency             string      `json:"default_currency"`
	OnlyAdminsMayCreateProjects bool        `json:"only_admins_may_create_projects"`
	OnlyAdminsMayCreateTags     bool        `json:"only_admins_may_create_tags"`
	OnlyAdminsSeeBillableRates  bool        `json:"only_admins_see_billable_rates"`
	OnlyAdminsSeeTeamDashboard  bool        `json:"only_admins_see_team_dashboard"`
	ProjectsBillableByDefault   bool        `json:"projects_billable_by_default"`
	ReportsCollapse             bool        `json:"reports_collapse"`
	Rounding                    int         `json:"rounding"`
	RoundingMinutes             int         `json:"rounding_minutes"`
	APIToken                    string      `json:"api_token"`
	At                          time.Time   `json:"at"`
	LogoURL                     string      `json:"logo_url"`
	IcalURL                     string      `json:"ical_url"`
	IcalEnabled                 bool        `json:"ical_enabled"`
	CsvUpload                   interface{} `json:"csv_upload"`
	Subscription                interface{} `json:"subscription"`
}

type TimeEntry struct {
	ID              int           `json:"id"`
	WorkspaceID     int           `json:"workspace_id"`
	ProjectID       interface{}   `json:"project_id"`
	TaskID          interface{}   `json:"task_id"`
	Billable        bool          `json:"billable"`
	Start           time.Time     `json:"start"`
	Stop            interface{}   `json:"stop"`
	Duration        int           `json:"duration"`
	Description     string        `json:"description"`
	Tags            []interface{} `json:"tags"`
	TagIds          []interface{} `json:"tag_ids"`
	Duronly         bool          `json:"duronly"`
	At              time.Time     `json:"at"`
	ServerDeletedAt interface{}   `json:"server_deleted_at"`
	UserID          int           `json:"user_id"`
	UID             int           `json:"uid"`
	Wid             int           `json:"wid"`
}

type NewTimeEntry struct {
	CreatedWith string      `json:"created_with"`
	Description string      `json:"description"`
	Tags        []string    `json:"tags"`
	Billable    bool        `json:"billable"`
	WorkspaceID int         `json:"workspace_id"`
	ProjectID   int         `json:"project_id,omitempty"`
	Duration    int         `json:"duration"`
	Start       time.Time   `json:"start"`
	Stop        interface{} `json:"stop"`
}

type Project struct {
	ID                  int         `json:"id"`
	WorkspaceID         int         `json:"workspace_id"`
	ClientID            interface{} `json:"client_id"`
	Name                string      `json:"name"`
	IsPrivate           bool        `json:"is_private"`
	Active              bool        `json:"active"`
	At                  time.Time   `json:"at"`
	CreatedAt           time.Time   `json:"created_at"`
	ServerDeletedAt     interface{} `json:"server_deleted_at"`
	Color               string      `json:"color"`
	Billable            interface{} `json:"billable"`
	Template            interface{} `json:"template"`
	AutoEstimates       interface{} `json:"auto_estimates"`
	EstimatedHours      interface{} `json:"estimated_hours"`
	Rate                interface{} `json:"rate"`
	RateLastUpdated     interface{} `json:"rate_last_updated"`
	Currency            interface{} `json:"currency"`
	Recurring           bool        `json:"recurring"`
	RecurringParameters interface{} `json:"recurring_parameters"`
	CurrentPeriod       interface{} `json:"current_period"`
	FixedFee            interface{} `json:"fixed_fee"`
	ActualHours         int         `json:"actual_hours"`
	Wid                 int         `json:"wid"`
	Cid                 interface{} `json:"cid"`
}

type Tag struct {
	ID          int       `json:"id"`
	WorkspaceID int       `json:"workspace_id"`
	Name        string    `json:"name"`
	At          time.Time `json:"at"`
}
