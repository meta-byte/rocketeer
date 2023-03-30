package types

import (
	"time"
)

// https://mholt.github.io/json-to-go/ to save time here
type CachedLaunch struct {
	ID             string    `json:"id,omitempty"`
	NET            time.Time `json:"countdown,omitempty"`
	LaunchProvider string    `json:"launch_provider,omitempty"`
	LaunchData     Launch    `json:"launch,omitempty"`
	Users          []string  `json:"users,omitempty"`
}

type Agency struct {
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Program struct {
	ID             int       `json:"id,omitempty"`
	URL            string    `json:"url,omitempty"`
	Name           string    `json:"name,omitempty"`
	Description    string    `json:"description,omitempty"`
	Agencies       []Agency  `json:"agencies,omitempty"`
	ImageURL       string    `json:"image_url,omitempty"`
	StartDate      time.Time `json:"start_date,omitempty"`
	EndDate        any       `json:"end_date,omitempty"`
	InfoURL        string    `json:"info_url,omitempty"`
	WikiURL        string    `json:"wiki_url,omitempty"`
	MissionPatches []any     `json:"mission_patches,omitempty"`
}

type LaunchStatus struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Abbrev      string `json:"abbrev,omitempty"`
	Description string `json:"description,omitempty"`
}

type LaunchResults struct {
	Count    int      `json:"count,omitempty"`
	Next     string   `json:"next,omitempty"`
	Previous any      `json:"previous,omitempty"`
	Results  []Launch `json:"results,omitempty"`
}

type LaunchServiceProvider struct {
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type RocketConfiguration struct {
	ID       int    `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	Name     string `json:"name,omitempty"`
	Family   string `json:"family,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Variant  string `json:"variant,omitempty"`
}

type Rocket struct {
	ID            int                 `json:"id,omitempty"`
	Configuration RocketConfiguration `json:"configuration,omitempty"`
}

type Orbit struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Abbrev string `json:"abbrev,omitempty"`
}

type Mission struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	LaunchDesignator any    `json:"launch_designator,omitempty"`
	Type             string `json:"type,omitempty"`
	Orbit            Orbit  `json:"orbit,omitempty"`
}

type Location struct {
	ID                int    `json:"id,omitempty"`
	URL               string `json:"url,omitempty"`
	Name              string `json:"name,omitempty"`
	CountryCode       string `json:"country_code,omitempty"`
	MapImage          string `json:"map_image,omitempty"`
	TotalLaunchCount  int    `json:"total_launch_count,omitempty"`
	TotalLandingCount int    `json:"total_landing_count,omitempty"`
}

type Pad struct {
	ID                        int      `json:"id,omitempty"`
	URL                       string   `json:"url,omitempty"`
	AgencyID                  any      `json:"agency_id,omitempty"`
	Name                      string   `json:"name,omitempty"`
	InfoURL                   any      `json:"info_url,omitempty"`
	WikiURL                   string   `json:"wiki_url,omitempty"`
	MapURL                    string   `json:"map_url,omitempty"`
	Latitude                  string   `json:"latitude,omitempty"`
	Longitude                 string   `json:"longitude,omitempty"`
	Location                  Location `json:"location,omitempty"`
	MapImage                  string   `json:"map_image,omitempty"`
	TotalLaunchCount          int      `json:"total_launch_count,omitempty"`
	OrbitalLaunchAttemptCount int      `json:"orbital_launch_attempt_count,omitempty"`
}

type Launch struct {
	ID                             string                `json:"id,omitempty"`
	URL                            string                `json:"url,omitempty"`
	Slug                           string                `json:"slug,omitempty"`
	Name                           string                `json:"name,omitempty"`
	Status                         LaunchStatus          `json:"status,omitempty"`
	LastUpdated                    time.Time             `json:"last_updated,omitempty"`
	Net                            time.Time             `json:"net,omitempty"`
	WindowEnd                      time.Time             `json:"window_end,omitempty"`
	WindowStart                    time.Time             `json:"window_start,omitempty"`
	Probability                    int                   `json:"probability,omitempty"`
	Holdreason                     string                `json:"holdreason,omitempty"`
	Failreason                     string                `json:"failreason,omitempty"`
	Hashtag                        any                   `json:"hashtag,omitempty"`
	LaunchServiceProvider          LaunchServiceProvider `json:"launch_service_provider,omitempty"`
	Rocket                         Rocket                `json:"rocket,omitempty"`
	Mission                        Mission               `json:"mission,omitempty"`
	Pad                            Pad                   `json:"pad,omitempty"`
	WebcastLive                    bool                  `json:"webcast_live,omitempty"`
	Image                          string                `json:"image,omitempty"`
	Infographic                    any                   `json:"infographic,omitempty"`
	Program                        []Program             `json:"program,omitempty"`
	OrbitalLaunchAttemptCount      int                   `json:"orbital_launch_attempt_count,omitempty"`
	LocationLaunchAttemptCount     int                   `json:"location_launch_attempt_count,omitempty"`
	PadLaunchAttemptCount          int                   `json:"pad_launch_attempt_count,omitempty"`
	AgencyLaunchAttemptCount       int                   `json:"agency_launch_attempt_count,omitempty"`
	OrbitalLaunchAttemptCountYear  int                   `json:"orbital_launch_attempt_count_year,omitempty"`
	LocationLaunchAttemptCountYear int                   `json:"location_launch_attempt_count_year,omitempty"`
	PadLaunchAttemptCountYear      int                   `json:"pad_launch_attempt_count_year,omitempty"`
	AgencyLaunchAttemptCountYear   int                   `json:"agency_launch_attempt_count_year,omitempty"`
}
