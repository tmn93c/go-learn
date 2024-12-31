package router

import (
	"encoding/json"
	"my-project/internal/database"
	"net/http"

	"time"

	"github.com/go-chi/chi/v5"
)

var xmlBytes = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
"https://raw.githubusercontent.com/zhuxiujia/GoMybatis/master/mybatis-3-mapper.dtd">
<mapper>
	<select id="SelectAll">
		select * from biz_activity
	</select>
</mapper>`)

type Activity struct {
	Id         string    `json:"id"`
	Uuid       string    `json:"uuid"`
	Name       string    `json:"name"`
	PcLink     string    `json:"pcLink"`
	H5Link     string    `json:"h5Link"`
	Remark     string    `json:"remark"`
	Version    int       `json:"version"`
	CreateTime time.Time `json:"createTime"`
	DeleteFlag int       `json:"deleteFlag"`
}
type ActivityMapper struct {
	SelectAll func() ([]Activity, error)
}

type activityResource struct {
	service database.Service // Hold a reference to ServiceDb
}

var activityMapper ActivityMapper // Assuming you have defined this mapper
func NewActivityResource(service database.Service) *activityResource {
	return &activityResource{
		service: service,
	}
}

func (s *activityResource) RegisterActivityRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/page", s.getActivities)

	return r
}

// @Summary Get a Page Activity
// @Description Get a Page Activity
// @Produce json
// @Success 200 {object} map[string]string
// @Router /activity/page [get]
func (s *activityResource) getActivities(w http.ResponseWriter, r *http.Request) {
	engine := s.service.Mybatis()
	engine.WriteMapperPtr(&activityMapper, xmlBytes)

	activities, err := activityMapper.SelectAll()
	if err != nil {
		http.Error(w, "Failed to retrieve activities", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(activities)
}
