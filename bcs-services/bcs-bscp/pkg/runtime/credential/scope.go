/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package credential provides credential scope related operations.
package credential

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gobwas/glob"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/i18n"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/kit"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/tools"
)

// Scope defines the credential scope expression.
type Scope string

// Validate validate a credential scope is valid or not.
func (cs Scope) Validate(kit *kit.Kit) error {
	strs := strings.Split(string(cs), "/")
	if len(strs) < 2 {
		return errors.New(i18n.T(kit, "invalid credential scope %s", string(cs)))
	}
	for _, str := range strs {
		if len(str) == 0 {
			return errors.New(i18n.T(kit, "invalid credential scope %s", string(cs)))
		}
	}
	return nil
}

// Split 拆分成 app / scope 格式
func (cs Scope) Split() (app string, scope string, err error) {
	index := strings.Index(string(cs), "/")
	if index == -1 {
		return "", "", fmt.Errorf("invalid credential scope %s", cs)
	}

	app = string(cs[:index])
	scope = string(cs[index:])
	return app, scope, nil
}

// MatchApp checks if the credential scope matches the app name.
func (cs Scope) MatchApp(kit *kit.Kit, name string) (bool, error) {
	if err := cs.Validate(kit); err != nil {
		return false, err
	}

	appPattern := strings.Split(string(cs), "/")[0]
	g, err := glob.Compile(appPattern)
	if err != nil {
		return false, errors.New(i18n.T(kit, "checks if the credential scope matches the app name failed, err: %v", err))
	}
	return g.Match(name), nil
}

// MatchConfigItem checks if the credential scope matches the config item.
func (cs Scope) MatchConfigItem(kit *kit.Kit, path, name string) (bool, error) {
	if err := cs.Validate(kit); err != nil {
		return false, err
	}
	configItemPattern := strings.SplitN(string(cs), "/", 2)[1]
	ok, err := tools.MatchConfigItem(configItemPattern, path, name)
	if err != nil {
		return false, errors.New(i18n.T(kit, "checks if the credential scope matches the config item failed, err: %v", err))
	}
	return ok, nil
}

// New 通过 app scope 组装
func New(kit *kit.Kit, app string, scope string) (Scope, error) {
	if len(app) == 0 {
		return "", errors.New(i18n.T(kit, "app is required"))
	}

	if len(scope) == 0 {
		return "", errors.New(i18n.T(kit, "scope is required"))
	}

	return Scope(app + scope), nil
}
