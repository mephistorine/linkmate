package links

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mileusna/useragent"
	"github.com/oschwald/geoip2-golang"
	"github.com/samber/lo"
	"linkmate/internal/analytics"
	"linkmate/internal/config"
	authUtils "linkmate/internal/shared/auth-utils"
	"linkmate/internal/tags"
)

type Controller struct {
	linksRepository     *Repository
	tagsRepository      *tags.Repository
	analyticsRepository *analytics.Repository
	config              *config.Config
	geoip2Reader        *geoip2.Reader
}

type CreateLinkReqDto struct {
	Key string `json:"key"`
	Url string `json:"url"`
}

type CreateLinkResDto struct {
	Id         int       `json:"id"`
	Key        string    `json:"key"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"createTime"`
}

func NewController(repository *Repository, tagsRepository *tags.Repository, analyticsRepository *analytics.Repository, config2 *config.Config) *Controller {
	reader, err := geoip2.Open("GeoLite2-City.mmdb")

	if err != nil {
		panic("Not found geolite db")
	}

	return &Controller{linksRepository: repository, tagsRepository: tagsRepository, config: config2, analyticsRepository: analyticsRepository, geoip2Reader: reader}
}

// CreateLinkHandler godoc
//
//	@Summary	Create link
//	@Router		/links [post]
//	@Param		link	body		links.CreateLinkReqDto	true	"Create link"
//	@Success	200		{object}	links.CreateLinkResDto
//	@Tags		links
//	@Accept		json
//	@Produce	json
//
//	@Security	ApiKeyAuth
func (c Controller) CreateLinkHandler(ec echo.Context) error {
	dto := new(CreateLinkReqDto)

	if err := ec.Bind(dto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var key string

	if len(dto.Key) > 0 {
		key = dto.Key
	} else {
		newKey, err := gonanoid.New(c.config.LinkKeyLength)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		key = newKey
	}

	link, err := c.linksRepository.Create(CreateLinkDto{
		Key:    key,
		Url:    dto.Url,
		UserId: tokenData.UserId,
	})

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.JSON(http.StatusCreated, CreateLinkResDto{
		Id:         link.Id,
		Key:        link.Key,
		Url:        link.Url,
		CreateTime: link.CreateTime,
	})
}

// LinkListHandler godoc
//
//	@Summary	Link list
//	@Router		/links [get]
//	@Success	200	{object}	[]links.SingleLinkDto
//	@Tags		links
//	@Accept		json
//	@Produce	json
//
//	@Security	ApiKeyAuth
func (c Controller) LinkListHandler(ec echo.Context) error {
	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	links, err := c.linksRepository.FindManyByUserId(tokenData.UserId)

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tagsByLinkId, err := c.tagsRepository.FindManyByLinkIds(lo.Map(links, func(item *Link, _ int) int {
		return item.Id
	}))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var singleLinkDtos []SingleLinkDto

	for _, link := range links {
		tags := lo.Map(tagsByLinkId[link.Id], func(item *tags.Tag, _ int) int {
			return item.Id
		})
		singleLinkDtos = append(singleLinkDtos, SingleLinkDto{
			Id:         link.Id,
			Key:        link.Key,
			Url:        link.Url,
			UserId:     link.UserId,
			CreateTime: link.CreateTime,
			TagIds:     tags,
		})
	}

	return ec.JSON(http.StatusOK, singleLinkDtos)
}

// DeleteLinkHandler godoc
//
//	@Summary	Delete link
//	@Router		/links [delete]
//	@Param		id	query	string	true	"id"
//	@Tags		links
//	@Accept		json
//	@Produce	json
//
//	@Security	ApiKeyAuth
func (c Controller) DeleteLinkHandler(ec echo.Context) error {
	linkId, err := strconv.Atoi(ec.QueryParam("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	link, err := c.linksRepository.FindOneById(linkId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if link.UserId != tokenData.UserId {
		return echo.NewHTTPError(http.StatusForbidden, "access error")
	}

	return ec.NoContent(http.StatusOK)
}

// RedirectLinkHandler godoc
//
//	@Summary	Find link and redirect to original url
//	@Param		key	path	string	true	"key"
//	@Router		/{key} [get]
//	@Tags		links
//	@Accept		json
//	@Produce	json
func (c Controller) RedirectLinkHandler(ctx echo.Context) error {
	link, err := c.linksRepository.FindOneByKey(ctx.Param("key"))

	if err != nil {
		ctx.Logger().Error(err)
		return ctx.NoContent(http.StatusNotFound)
	}

	userAgent := ctx.Request().Header.Get("User-Agent")

	geolocation, err := getGeolocation(ctx.RealIP(), c.geoip2Reader)

	if len(userAgent) > 0 {
		result := useragent.Parse(userAgent)

		var deviceType string

		if result.Desktop {
			deviceType = "DESKTOP"
		} else if result.Tablet {
			deviceType = "TABLET"
		} else if result.Mobile {
			deviceType = "MOBILE"
		}

		var geolocationData analytics.AnalyticsGeolocationData

		if err == nil {
			geolocationData = analytics.AnalyticsGeolocationData{
				CountryName:      geolocation.CountryName,
				CountryCode:      geolocation.CountryCode,
				CountryGeoNameId: geolocation.CountryGeoNameId,
				CityName:         geolocation.CityName,
				CityGeoNameId:    geolocation.CityGeoNameId,
			}
		} else {
			ctx.Logger().Error("Geocoding data", err)
		}

		err := c.analyticsRepository.PushEvent(analytics.LinkAnalyticsEvent{
			LinkId:      link.Id,
			UserAgent:   userAgent,
			BrowserName: result.Name,
			DeviceType:  deviceType,
			OsName:      result.OS,
			IpAddress:   ctx.RealIP(),
			Source:      ctx.Request().Header.Get("Referer"),
			Geolocation: geolocationData,
		})

		if err != nil {
			ctx.Logger().Warn("Error with analytics")
		}
	} else {
		err := c.analyticsRepository.PushEvent(analytics.LinkAnalyticsEvent{
			LinkId: link.Id,
		})

		if err != nil {
			ctx.Logger().Warn("Error with analytics")
		}
	}

	return ctx.Redirect(http.StatusPermanentRedirect, link.Url)
}

type GeoData struct {
	CountryName      string
	CountryCode      string
	CountryGeoNameId uint
	CityName         string
	CityGeoNameId    uint
}

func getGeolocation(ip string, geoip2Reader *geoip2.Reader) (*GeoData, error) {
	ipParsed := net.ParseIP(ip)
	result, err := geoip2Reader.City(ipParsed)

	if err != nil {
		return nil, err
	}

	return &GeoData{
		CountryName:      result.Country.Names["en"],
		CountryCode:      result.Country.IsoCode,
		CountryGeoNameId: result.Country.GeoNameID,
		CityName:         result.City.Names["en"],
		CityGeoNameId:    result.City.GeoNameID,
	}, nil
}
