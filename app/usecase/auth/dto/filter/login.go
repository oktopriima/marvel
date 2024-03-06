package filter

import (
	"github.com/oktopriima/marvel/app/modules/base/filter"
	"github.com/oktopriima/marvel/app/modules/base/filter/mysqlfilter"
	"github.com/oktopriima/marvel/app/usecase/auth/dto"
)

type Login struct {
	credentials string
	password    string
}

type LoginFilter interface {
	filter.BaseFilter
}

type loginFilter struct {
	*mysqlfilter.BaseFilter
	Login
}

func (a *loginFilter) GetWhere() filter.Where {
	where := map[string][]interface{}{
		"email": {
			a.credentials,
		},
	}

	return where
}

func (a *loginFilter) GetLimit() int {
	return filter.IgnoreLimit
}

func (a *loginFilter) GetOffset() int {
	return filter.IgnoreOffset
}

func (a *loginFilter) GetJoins() filter.Joins {
	var join filter.Joins

	return join
}

func (a *loginFilter) GetOrderBy() []string {
	return []string{""}
}

func (a *loginFilter) GetGroups() string {
	return ""
}

func NewLoginFilter(request dto.EmailLoginRequest) LoginFilter {
	return &loginFilter{
		BaseFilter: mysqlfilter.NewBaseFilter(),
		Login: Login{
			credentials: request.Email,
			password:    request.Password,
		},
	}
}
