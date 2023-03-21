package types

import (
	"time"
)

// https://mholt.github.io/json-to-go/ to save time here
type Launch struct {
	Count    int    `json:"count,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous any    `json:"previous,omitempty"`
	Results  []struct {
		ID     string `json:"id,omitempty"`
		URL    string `json:"url,omitempty"`
		Slug   string `json:"slug,omitempty"`
		Name   string `json:"name,omitempty"`
		Status struct {
			ID          int    `json:"id,omitempty"`
			Name        string `json:"name,omitempty"`
			Abbrev      string `json:"abbrev,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"status,omitempty"`
		LastUpdated           time.Time `json:"last_updated,omitempty"`
		Net                   time.Time `json:"net,omitempty"`
		WindowEnd             time.Time `json:"window_end,omitempty"`
		WindowStart           time.Time `json:"window_start,omitempty"`
		Probability           int       `json:"probability,omitempty"`
		Holdreason            string    `json:"holdreason,omitempty"`
		Failreason            string    `json:"failreason,omitempty"`
		Hashtag               any       `json:"hashtag,omitempty"`
		LaunchServiceProvider struct {
			ID   int    `json:"id,omitempty"`
			URL  string `json:"url,omitempty"`
			Name string `json:"name,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"launch_service_provider,omitempty"`
		Rocket struct {
			ID            int `json:"id,omitempty"`
			Configuration struct {
				ID       int    `json:"id,omitempty"`
				URL      string `json:"url,omitempty"`
				Name     string `json:"name,omitempty"`
				Family   string `json:"family,omitempty"`
				FullName string `json:"full_name,omitempty"`
				Variant  string `json:"variant,omitempty"`
			} `json:"configuration,omitempty"`
		} `json:"rocket,omitempty"`
		Mission struct {
			ID               int    `json:"id,omitempty"`
			Name             string `json:"name,omitempty"`
			Description      string `json:"description,omitempty"`
			LaunchDesignator any    `json:"launch_designator,omitempty"`
			Type             string `json:"type,omitempty"`
			Orbit            struct {
				ID     int    `json:"id,omitempty"`
				Name   string `json:"name,omitempty"`
				Abbrev string `json:"abbrev,omitempty"`
			} `json:"orbit,omitempty"`
		} `json:"mission,omitempty"`
		Pad struct {
			ID        int    `json:"id,omitempty"`
			URL       string `json:"url,omitempty"`
			AgencyID  any    `json:"agency_id,omitempty"`
			Name      string `json:"name,omitempty"`
			InfoURL   any    `json:"info_url,omitempty"`
			WikiURL   string `json:"wiki_url,omitempty"`
			MapURL    string `json:"map_url,omitempty"`
			Latitude  string `json:"latitude,omitempty"`
			Longitude string `json:"longitude,omitempty"`
			Location  struct {
				ID                int    `json:"id,omitempty"`
				URL               string `json:"url,omitempty"`
				Name              string `json:"name,omitempty"`
				CountryCode       string `json:"country_code,omitempty"`
				MapImage          string `json:"map_image,omitempty"`
				TotalLaunchCount  int    `json:"total_launch_count,omitempty"`
				TotalLandingCount int    `json:"total_landing_count,omitempty"`
			} `json:"location,omitempty"`
			MapImage                  string `json:"map_image,omitempty"`
			TotalLaunchCount          int    `json:"total_launch_count,omitempty"`
			OrbitalLaunchAttemptCount int    `json:"orbital_launch_attempt_count,omitempty"`
		} `json:"pad,omitempty"`
		WebcastLive bool   `json:"webcast_live,omitempty"`
		Image       string `json:"image,omitempty"`
		Infographic any    `json:"infographic,omitempty"`
		Program     []struct {
			ID          int    `json:"id,omitempty"`
			URL         string `json:"url,omitempty"`
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
			Agencies    []struct {
				ID   int    `json:"id,omitempty"`
				URL  string `json:"url,omitempty"`
				Name string `json:"name,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"agencies,omitempty"`
			ImageURL       string    `json:"image_url,omitempty"`
			StartDate      time.Time `json:"start_date,omitempty"`
			EndDate        any       `json:"end_date,omitempty"`
			InfoURL        string    `json:"info_url,omitempty"`
			WikiURL        string    `json:"wiki_url,omitempty"`
			MissionPatches []any     `json:"mission_patches,omitempty"`
		} `json:"program,omitempty"`
		OrbitalLaunchAttemptCount      int `json:"orbital_launch_attempt_count,omitempty"`
		LocationLaunchAttemptCount     int `json:"location_launch_attempt_count,omitempty"`
		PadLaunchAttemptCount          int `json:"pad_launch_attempt_count,omitempty"`
		AgencyLaunchAttemptCount       int `json:"agency_launch_attempt_count,omitempty"`
		OrbitalLaunchAttemptCountYear  int `json:"orbital_launch_attempt_count_year,omitempty"`
		LocationLaunchAttemptCountYear int `json:"location_launch_attempt_count_year,omitempty"`
		PadLaunchAttemptCountYear      int `json:"pad_launch_attempt_count_year,omitempty"`
		AgencyLaunchAttemptCountYear   int `json:"agency_launch_attempt_count_year,omitempty"`
	} `json:"results,omitempty"`
}

// type Launch struct {
// 	Count    int    `json:"count"`
// 	Next     string `json:"next"`
// 	Previous any    `json:"previous"`
// 	Results  []struct {
// 		ID     string `json:"id"`
// 		URL    string `json:"url"`
// 		Slug   string `json:"slug"`
// 		Name   string `json:"name"`
// 		Status struct {
// 			ID          int    `json:"id"`
// 			Name        string `json:"name"`
// 			Abbrev      string `json:"abbrev"`
// 			Description string `json:"description"`
// 		} `json:"status"`
// 		LastUpdated           time.Time `json:"last_updated"`
// 		Net                   time.Time `json:"net"`
// 		WindowEnd             time.Time `json:"window_end"`
// 		WindowStart           time.Time `json:"window_start"`
// 		Probability           any       `json:"probability"`
// 		Holdreason            string    `json:"holdreason"`
// 		Failreason            string    `json:"failreason"`
// 		Hashtag               any       `json:"hashtag"`
// 		LaunchServiceProvider struct {
// 			ID   int    `json:"id"`
// 			URL  string `json:"url"`
// 			Name string `json:"name"`
// 			Type string `json:"type"`
// 		} `json:"launch_service_provider"`
// 		Rocket struct {
// 			ID            int `json:"id"`
// 			Configuration struct {
// 				ID       int    `json:"id"`
// 				URL      string `json:"url"`
// 				Name     string `json:"name"`
// 				Family   string `json:"family"`
// 				FullName string `json:"full_name"`
// 				Variant  string `json:"variant"`
// 			} `json:"configuration"`
// 		} `json:"rocket"`
// 		Mission struct {
// 			ID               int    `json:"id"`
// 			Name             string `json:"name"`
// 			Description      string `json:"description"`
// 			LaunchDesignator any    `json:"launch_designator"`
// 			Type             string `json:"type"`
// 			Orbit            struct {
// 				ID     int    `json:"id"`
// 				Name   string `json:"name"`
// 				Abbrev string `json:"abbrev"`
// 			} `json:"orbit"`
// 		} `json:"mission"`
// 		Pad struct {
// 			ID        int    `json:"id"`
// 			URL       string `json:"url"`
// 			AgencyID  int    `json:"agency_id"`
// 			Name      string `json:"name"`
// 			InfoURL   any    `json:"info_url"`
// 			WikiURL   any    `json:"wiki_url"`
// 			MapURL    any    `json:"map_url"`
// 			Latitude  string `json:"latitude"`
// 			Longitude string `json:"longitude"`
// 			Location  struct {
// 				ID                int    `json:"id"`
// 				URL               string `json:"url"`
// 				Name              string `json:"name"`
// 				CountryCode       string `json:"country_code"`
// 				MapImage          string `json:"map_image"`
// 				TotalLaunchCount  int    `json:"total_launch_count"`
// 				TotalLandingCount int    `json:"total_landing_count"`
// 			} `json:"location"`
// 			MapImage                  string `json:"map_image"`
// 			TotalLaunchCount          int    `json:"total_launch_count"`
// 			OrbitalLaunchAttemptCount int    `json:"orbital_launch_attempt_count"`
// 		} `json:"pad"`
// 		WebcastLive bool   `json:"webcast_live"`
// 		Image       string `json:"image"`
// 		Infographic any    `json:"infographic"`
// 		Program     []struct {

// 		} `json:"program"`
// 		OrbitalLaunchAttemptCount      int `json:"orbital_launch_attempt_count"`
// 		LocationLaunchAttemptCount     int `json:"location_launch_attempt_count"`
// 		PadLaunchAttemptCount          int `json:"pad_launch_attempt_count"`
// 		AgencyLaunchAttemptCount       int `json:"agency_launch_attempt_count"`
// 		OrbitalLaunchAttemptCountYear  int `json:"orbital_launch_attempt_count_year"`
// 		LocationLaunchAttemptCountYear int `json:"location_launch_attempt_count_year"`
// 		PadLaunchAttemptCountYear      int `json:"pad_launch_attempt_count_year"`
// 		AgencyLaunchAttemptCountYear   int `json:"agency_launch_attempt_count_year"`
// 	} `json:"results"`
// }
