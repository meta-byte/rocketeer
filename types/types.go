package types

import (
	"time"
)

// https://mholt.github.io/json-to-go/ to save time here
type Launch struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		Slug   string `json:"slug"`
		Name   string `json:"name"`
		Status struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Abbrev      string `json:"abbrev"`
			Description string `json:"description"`
		} `json:"status"`
		LastUpdated           time.Time `json:"last_updated"`
		Net                   time.Time `json:"net"`
		WindowEnd             time.Time `json:"window_end"`
		WindowStart           time.Time `json:"window_start"`
		Probability           any       `json:"probability"`
		Holdreason            string    `json:"holdreason"`
		Failreason            string    `json:"failreason"`
		Hashtag               any       `json:"hashtag"`
		LaunchServiceProvider struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"launch_service_provider"`
		Rocket struct {
			ID            int `json:"id"`
			Configuration struct {
				ID       int    `json:"id"`
				URL      string `json:"url"`
				Name     string `json:"name"`
				Family   string `json:"family"`
				FullName string `json:"full_name"`
				Variant  string `json:"variant"`
			} `json:"configuration"`
		} `json:"rocket"`
		Mission struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			LaunchDesignator any    `json:"launch_designator"`
			Type             string `json:"type"`
			Orbit            struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Abbrev string `json:"abbrev"`
			} `json:"orbit"`
		} `json:"mission"`
		Pad struct {
			ID        int    `json:"id"`
			URL       string `json:"url"`
			AgencyID  int    `json:"agency_id"`
			Name      string `json:"name"`
			InfoURL   any    `json:"info_url"`
			WikiURL   any    `json:"wiki_url"`
			MapURL    any    `json:"map_url"`
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
			Location  struct {
				ID                int    `json:"id"`
				URL               string `json:"url"`
				Name              string `json:"name"`
				CountryCode       string `json:"country_code"`
				MapImage          string `json:"map_image"`
				TotalLaunchCount  int    `json:"total_launch_count"`
				TotalLandingCount int    `json:"total_landing_count"`
			} `json:"location"`
			MapImage                  string `json:"map_image"`
			TotalLaunchCount          int    `json:"total_launch_count"`
			OrbitalLaunchAttemptCount int    `json:"orbital_launch_attempt_count"`
		} `json:"pad"`
		WebcastLive                    bool   `json:"webcast_live"`
		Image                          string `json:"image"`
		Infographic                    any    `json:"infographic"`
		Program                        []any  `json:"program"`
		OrbitalLaunchAttemptCount      int    `json:"orbital_launch_attempt_count"`
		LocationLaunchAttemptCount     int    `json:"location_launch_attempt_count"`
		PadLaunchAttemptCount          int    `json:"pad_launch_attempt_count"`
		AgencyLaunchAttemptCount       int    `json:"agency_launch_attempt_count"`
		OrbitalLaunchAttemptCountYear  int    `json:"orbital_launch_attempt_count_year"`
		LocationLaunchAttemptCountYear int    `json:"location_launch_attempt_count_year"`
		PadLaunchAttemptCountYear      int    `json:"pad_launch_attempt_count_year"`
		AgencyLaunchAttemptCountYear   int    `json:"agency_launch_attempt_count_year"`
	} `json:"results"`
}
