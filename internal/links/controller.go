package links

import (
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"linkmate/internal/auth"
	"linkmate/internal/config"
	"net/http"
	"strconv"
	"time"
)

type Controller struct {
	linksRepository *Repository
	config          *config.Config
}

type CreateLinkReqDto struct {
	Key string `json:"key"`
	Url string `json:"url"`
}

type CreateLinkResDto struct {
	Id         string    `json:"id"`
	Key        string    `json:"key"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"createTime"`
}

func NewController(repository *Repository, config2 *config.Config) *Controller {
	return &Controller{linksRepository: repository, config: config2}
}

// CreateLinkHandler godoc
//
//	@Summary	Create link
//	@Router		/links [post]
//
//	@Success	200	{object}	links.CreateLinkResDto
//
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

	tokenData, err := auth.ParseJwt(ec.Get("user"))

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
	tokenData, err := auth.ParseJwt(ec.Get("user"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	links, err := c.linksRepository.FindManyByUserId(tokenData.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var singleLinkDtos []SingleLinkDto

	for _, link := range links {
		singleLinkDtos = append(singleLinkDtos, SingleLinkDto{
			Id:         link.Id,
			Key:        link.Key,
			Url:        link.Url,
			UserId:     link.UserId,
			CreateTime: link.CreateTime,
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

	tokenData, err := auth.ParseJwt(ec.Get("user"))

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
//	@Router		/:key [get]
//	@Tags		links
//	@Accept		json
//	@Produce	json
func (c Controller) RedirectLinkHandler(ctx echo.Context) error {
	link, err := c.linksRepository.FindOneByKey(ctx.Param("key"))

	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusMovedPermanently, link.Url)
}
