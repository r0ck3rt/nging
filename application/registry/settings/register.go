/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package settings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/admpub/log"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
)

var settings = []*SettingForm{
	&SettingForm{
		Short: `系统设置`,
		Label: `系统设置`,
		Group: `base`,
		Tmpl:  []string{`manager/settings/base`},
	},
	&SettingForm{
		Short: `SMTP设置`,
		Label: `SMTP服务器设置`,
		Group: `smtp`,
		Tmpl:  []string{`manager/settings/smtp`},
	},
	&SettingForm{
		Short: `日志设置`,
		Label: `日志设置`,
		Group: `log`,
		Tmpl:  []string{`manager/settings/log`},
	},
}

func Settings() []*SettingForm {
	return settings
}

func Register(sf ...*SettingForm) {
	settings = append(settings, sf...)
	for _, s := range sf {
		if s.items != nil {
			AddDefaultConfig(s.Group, s.items)
		}
		if s.dataInitors != nil {
			s.dataInitors.Register(s.Group)
		}
		if s.dataFroms != nil {
			s.dataFroms.Register(s.Group)
		}
	}
}

func Get(group string) (int, *SettingForm) {
	for index, setting := range settings {
		if setting.Group == group {
			return index, setting
		}
	}
	return -1, nil
}

func RunHookPost(ctx echo.Context, groups ...string) error {
	n := len(groups)
	var i int
	var errs []string
	for _, setting := range settings {
		if n < 1 || com.InSlice(setting.Group, groups) {
			err := setting.RunHookPost(ctx)
			if err != nil {
				err = fmt.Errorf(`[config][group:%s] %w`, setting.Group, err)
				log.Error(err)
				errs = append(errs, err.Error())
			}
			if n > 0 {
				i++
				if i >= n {
					break
				}
			}
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}

func RunHookGet(ctx echo.Context, groups ...string) error {
	n := len(groups)
	var i int
	var errs []string
	for _, setting := range settings {
		if n < 1 || com.InSlice(setting.Group, groups) {
			err := setting.RunHookGet(ctx)
			if err != nil {
				err = fmt.Errorf(`[config][group:%s] %w`, setting.Group, err)
				log.Error(err)
				errs = append(errs, err.Error())
			}
			if n > 0 {
				i++
				if i >= n {
					break
				}
			}
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}
