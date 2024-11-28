package analytics

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type AnalyticsGeolocationData struct {
	CountryName      string `json:"countryName,omitempty"`
	CountryCode      string `json:"countryCode,omitempty"`
	CountryGeoNameId uint   `json:"countryGeonameId,omitempty"`
	CityName         string `json:"cityName,omitempty"`
	CityGeoNameId    uint   `json:"cityGeonameId,omitempty"`
}

func (g AnalyticsGeolocationData) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *AnalyticsGeolocationData) Scan(value interface{}) error {
	b, ok := value.([]byte)

	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &g)
}

type LinkAnalyticsEvent struct {
	Id          string
	LinkId      int
	CreateTime  time.Time
	UserAgent   string
	BrowserName string
	DeviceType  string
	OsName      string
	Source      string
	IpAddress   string
	Geolocation AnalyticsGeolocationData
}

func (r *Repository) PushEvent(dto LinkAnalyticsEvent) error {
	_, err := r.db.Exec(
		`INSERT INTO link_analytics_events(
            link_id,
            user_agent,
            browser_name,
            device_type,
            os_name,
            source,
            ip_address,
            geolocation
         )
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		dto.LinkId, dto.UserAgent, dto.BrowserName, dto.DeviceType, dto.OsName, dto.Source, dto.IpAddress, dto.Geolocation)
	return err
}

func (r *Repository) FindManyBetweenInterval(start time.Time, end time.Time) ([]LinkAnalyticsEvent, error) {
	rows, err := r.db.Query(
		`SELECT
            id,
            link_id,
            create_time,
            user_agent,
            browser_name,
            device_type,
            os_name,
            source,
            ip_address,
            geolocation
        FROM
            link_analytics_events
        WHERE
            create_time BETWEEN $1 AND $2`,
		start,
		end,
	)

	if err != nil {
		return nil, err
	}

	var events []LinkAnalyticsEvent

	for rows.Next() {
		var event LinkAnalyticsEvent
		rows.Scan(&event.Id, &event.LinkId, &event.CreateTime, &event.UserAgent, &event.BrowserName, &event.DeviceType, &event.OsName, &event.Source, &event.IpAddress, &event.Geolocation)
		events = append(events, event)
	}

	return events, nil
}
