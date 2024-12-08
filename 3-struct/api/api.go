package api

import "bins/config"

type Api struct{
	Cfg config.Config
}

func NewApi(cfg *config.Config) *Api{
	return &Api{Cfg: *cfg}
}