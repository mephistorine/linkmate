package tags

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"linkmate/internal/config"
	authUtils "linkmate/internal/shared/auth-utils"
	"net/http"
	"strconv"
	"time"
)

type Controller struct {
	tagsRepository *Repository
	config         *config.Config
}

type CreateTagReqDto struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type CreateTagResDto struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	UserId     int       `json:"userId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type UpdateTagReqDto struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type UpdateTagResDto struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	UserId     int       `json:"userId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type SingleTagResDto struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	UserId     int       `json:"userId"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type AddTagsToLinkReqDto struct {
	LinkId int   `json:"linkId"`
	TagIds []int `json:"tagIds"`
}

type RemoveTagsToLinkReqDto struct {
	LinkId int   `json:"linkId"`
	TagIds []int `json:"tagIds"`
}

func NewController(repository *Repository, config2 *config.Config) *Controller {
	return &Controller{tagsRepository: repository, config: config2}
}

// CreateTagHandler godoc
//
//	@summary	Create tag
//	@router		/tags [post]
//	@param		tag	body		tags.CreateTagReqDto	true	"Create tag"
//	@success	201	{object}	tags.CreateTagResDto
//	@tags		tags
//	@accept		json
//	@produce	json
//	@security	ApiKeyAuth
func (c *Controller) CreateTagHandler(ec echo.Context) error {
	dto := new(CreateTagReqDto)

	if err := ec.Bind(dto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tag, err := c.tagsRepository.Create(CreateTagDto{
		Name:   dto.Name,
		Color:  dto.Color,
		UserId: tokenData.UserId,
	})

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.JSON(http.StatusCreated, CreateTagResDto{
		Id:         tag.Id,
		Name:       tag.Name,
		Color:      tag.Color,
		UserId:     tag.UserId,
		CreateTime: tag.CreateTime,
		UpdateTime: tag.UpdateTime,
	})
}

// UpdateTagHandler godoc
//
//	@summary	Update tag
//	@router		/tags [put]
//	@param		tag	body		tags.UpdateTagReqDto	true	"Update tag"
//	@success	200	{object}	tags.UpdateTagResDto
//	@tags		tags
//	@accept		json
//	@produce	json
//	@security	ApiKeyAuth
func (c *Controller) UpdateTagHandler(ec echo.Context) error {
	dto := new(UpdateTagReqDto)

	if err := ec.Bind(dto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tagId, err := strconv.Atoi(ec.QueryParam("id"))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tag, err := c.tagsRepository.FindById(tagId)

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	if tag.UserId != tokenData.UserId {
		ec.Logger().Error(fmt.Sprintf("Access denied to update tag with id=%d by user with id=%d", tagId, tokenData.UserId))
		return echo.NewHTTPError(http.StatusForbidden)
	}

	updatedTag, err := c.tagsRepository.UpdateById(tagId, UpdateTagDto{
		Name:  dto.Name,
		Color: dto.Color,
	})

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.JSON(http.StatusOK, UpdateTagResDto{
		Id:         updatedTag.Id,
		Name:       updatedTag.Name,
		Color:      updatedTag.Color,
		UserId:     updatedTag.UserId,
		CreateTime: updatedTag.CreateTime,
		UpdateTime: updatedTag.UpdateTime,
	})
}

// GetTags godoc
//
//	@summary	Get tags
//	@router		/tags [get]
//	@success	200	{object}	[]tags.SingleTagResDto
//	@tags		tags
//	@accept		json
//	@produce	json
//	@security	ApiKeyAuth
func (c *Controller) GetTags(ec echo.Context) error {
	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tags, err := c.tagsRepository.FindManyByUserId(tokenData.UserId)

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	var dtos []SingleTagResDto

	for _, tag := range tags {
		dtos = append(dtos, SingleTagResDto{
			Id:         tag.Id,
			Name:       tag.Name,
			Color:      tag.Color,
			UserId:     tag.UserId,
			CreateTime: tag.CreateTime,
			UpdateTime: tag.UpdateTime,
		})
	}

	return ec.JSON(http.StatusOK, dtos)
}

// DeleteTagHandler godoc
//
//	@summary	Delete tag
//	@router		/tags [delete]
//	@param		id	query	int	true	"Tag id"
//	@success	200
//	@tags		tags
//	@accept		json
//	@produce	json
//	@security	ApiKeyAuth
func (c *Controller) DeleteTagHandler(ec echo.Context) error {
	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tagId, err := strconv.Atoi(ec.QueryParam("id"))

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	tag, err := c.tagsRepository.FindById(tagId)

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound)
	}

	if tag.UserId != tokenData.UserId {
		ec.Logger().Error(fmt.Sprintf("Access denied to delete tag with id=%d by user with id=%d", tagId, tokenData.UserId))
		return echo.NewHTTPError(http.StatusForbidden)
	}

	if err := c.tagsRepository.DeleteById(tagId); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.NoContent(http.StatusOK)
}

// AddTagsToLinkHandler godoc
//
//	@summary	Add tags to link
//	@router		/tags/settings [post]
//	@param		tag	body	tags.AddTagsToLinkReqDto	true	"Add tags to link"
//	@success	200
//	@tags		tags
//	@accept		json
//	@produce	json
//	@security	ApiKeyAuth
func (c *Controller) AddTagsToLinkHandler(ec echo.Context) error {
	dto := new(AddTagsToLinkReqDto)

	if err := ec.Bind(dto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := c.tagsRepository.CreateTagSettings(dto.LinkId, dto.TagIds); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.NoContent(http.StatusOK)
}

// RemoveTagsFromLinkHandler godoc
//
//	@summary	Remove tags from link
//	@router		/tags/settings [delete]
//	@param		tag	body	tags.RemoveTagsToLinkReqDto	true	"Remove tags from link"
//	@success	200
//	@tags		tags
//	@accept		json
//	@produce	json
//	@security	ApiKeyAuth
func (c *Controller) RemoveTagsFromLinkHandler(ec echo.Context) error {
	dto := new(RemoveTagsToLinkReqDto)

	if err := ec.Bind(dto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := c.tagsRepository.DeleteTagSettings(dto.LinkId, dto.TagIds); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.NoContent(http.StatusOK)
}
