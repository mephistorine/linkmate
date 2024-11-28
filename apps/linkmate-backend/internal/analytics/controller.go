package analytics

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"linkmate/internal/config"
)

type Controller struct {
	config *config.Config
	repo   *Repository
}

func NewController(repository *Repository, conf *config.Config) *Controller {
	return &Controller{
		config: conf,
		repo:   repository,
	}
}

type DashboardParamsReqDto struct {
	IntervalStart time.Time `json:"intervalStart"`
	IntervalEnd   time.Time `json:"intervalEnd"`
}

type LinkVisit struct {
	LinkId     int `json:"linkId"`
	VisitCount int `json:"visitCount"`
}

type BrowserDeviceVisit struct {
	BrowserName string `json:"browserName"`
	VisitCount  int    `json:"visitCount"`
}

type OsDeviceVisit struct {
	OsName     string `json:"osName"`
	VisitCount int    `json:"visitCount"`
}

type SizeDeviceVisit struct {
	SizeName   string `json:"sizeName"`
	VisitCount int    `json:"visitCount"`
}

type CountryVisit struct {
	CountryCode string `json:"countryCode"`
	VisitCount  int    `json:"visitCount"`
}

type SourceVisit struct {
	SourceUrl  string `json:"sourceUrl"`
	VisitCount int    `json:"visitCount"`
}

type DashboardDataResDto struct {
	TotalVisits      int                  `json:"totalVisits"`
	LinkVisits       []LinkVisit          `json:"linkVisits"`
	Browsers         []BrowserDeviceVisit `json:"browsers"`
	OperationSystems []OsDeviceVisit      `json:"operationSystems"`
	Sizes            []SizeDeviceVisit    `json:"sizes"`
	Countries        []CountryVisit       `json:"countries"`
}

func (c *Controller) DashboardHandler(ctx echo.Context) error {
	dto := new(DashboardParamsReqDto)

	if err := ctx.Bind(dto); err != nil {
		ctx.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	events, err := c.repo.FindManyBetweenInterval(dto.IntervalStart, dto.IntervalEnd)

	if err != nil {
		ctx.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, DashboardDataResDto{
		TotalVisits:      len(events),
		LinkVisits:       calculateLinkVisits(events),
		Browsers:         calculateBrowsers(events),
		OperationSystems: calculateOperationSystems(events),
		Sizes:            calculateSizes(events),
		Countries:        calculateCountries(events),
	})
}

func calculateLinkVisits(events []LinkAnalyticsEvent) []LinkVisit {
	groups := lo.GroupBy(events, func(item LinkAnalyticsEvent) int {
		return item.LinkId
	})

	var links []LinkVisit

	for linkId, events := range groups {
		links = append(links, LinkVisit{
			LinkId:     linkId,
			VisitCount: len(events),
		})
	}

	return links
}

func calculateBrowsers(events []LinkAnalyticsEvent) []BrowserDeviceVisit {
	groups := lo.GroupBy(events, func(item LinkAnalyticsEvent) string {
		if len(item.BrowserName) > 0 {
			return item.BrowserName
		}

		return "Unknown"
	})

	var deviceVisits []BrowserDeviceVisit

	for browserName, events := range groups {
		deviceVisits = append(deviceVisits, BrowserDeviceVisit{
			BrowserName: browserName,
			VisitCount:  len(events),
		})
	}

	return deviceVisits
}

func calculateSizes(events []LinkAnalyticsEvent) []SizeDeviceVisit {
	groups := lo.GroupBy(events, func(item LinkAnalyticsEvent) string {
		if len(item.DeviceType) > 0 {
			return item.DeviceType
		}

		return "Unknown"
	})

	var sizeVisits []SizeDeviceVisit

	for sizeName, events := range groups {
		sizeVisits = append(sizeVisits, SizeDeviceVisit{
			SizeName:   sizeName,
			VisitCount: len(events),
		})
	}

	return sizeVisits
}

func calculateOperationSystems(events []LinkAnalyticsEvent) []OsDeviceVisit {
	groups := lo.GroupBy(events, func(item LinkAnalyticsEvent) string {
		if len(item.OsName) > 0 {
			return item.OsName
		}

		return "Unknown"
	})

	var osVisits []OsDeviceVisit

	for osName, events := range groups {
		osVisits = append(osVisits, OsDeviceVisit{
			OsName:     osName,
			VisitCount: len(events),
		})
	}

	return osVisits
}

func calculateCountries(events []LinkAnalyticsEvent) []CountryVisit {
	groups := lo.GroupBy(events, func(item LinkAnalyticsEvent) string {
		if len(item.Geolocation.CountryCode) > 0 {
			return item.Geolocation.CountryCode
		}

		return "Unknown"
	})

	var visits []CountryVisit

	for countryCode, events := range groups {
		visits = append(visits, CountryVisit{
			CountryCode: countryCode,
			VisitCount:  len(events),
		})
	}

	return visits
}
