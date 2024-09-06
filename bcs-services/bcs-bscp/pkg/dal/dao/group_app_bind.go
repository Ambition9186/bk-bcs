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

package dao

import (
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/criteria/errf"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/gen"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/i18n"
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/kit"
)

// GroupAppBind supplies all the group related operations.
type GroupAppBind interface {
	// BatchCreateWithTx batch create group app with transaction.
	BatchCreateWithTx(kit *kit.Kit, tx *gen.QueryTx, items []*table.GroupAppBind) error
	// BatchDeleteByGroupIDWithTx batch delete group app by group id with transaction.
	BatchDeleteByGroupIDWithTx(kit *kit.Kit, tx *gen.QueryTx, groupID, bizID uint32) error
	// BatchDeleteByAppIDWithTx batch delete group app by app id with transaction.
	BatchDeleteByAppIDWithTx(kit *kit.Kit, tx *gen.QueryTx, appID, bizID uint32) error
	// BatchListByGroupIDs batch list group app by group ids.
	BatchListByGroupIDs(kit *kit.Kit, bizID uint32, groupIDs []uint32) ([]*table.GroupAppBind, error)
	// Get get GroupAppBind by group id and app id.
	Get(kit *kit.Kit, groupID, appID, bizID uint32) (*table.GroupAppBind, error)
}

var _ GroupAppBind = new(groupAppDao)

type groupAppDao struct {
	genQ     *gen.Query
	idGen    IDGenInterface
	auditDao AuditDao
	lock     LockDao
}

// BatchCreateWithTx batch create group app with transaction.
func (dao *groupAppDao) BatchCreateWithTx(kit *kit.Kit, tx *gen.QueryTx, items []*table.GroupAppBind) error {
	if len(items) == 0 {
		return nil
	}

	// generate released config items id.
	ids, err := dao.idGen.Batch(kit, table.GroupAppBindTable, len(items))
	if err != nil {
		return errf.Errorf(errf.DBOpFailed, i18n.T(kit, "generate released config items id failed, err: %v", err))
	}

	for i, item := range items {
		// validate released config item field.
		if err := item.ValidateCreate(kit); err != nil {
			return err
		}
		item.ID = ids[i]
	}

	err = tx.Query.GroupAppBind.WithContext(kit.Ctx).Save(items...)
	if err != nil {
		return errf.Errorf(errf.DBOpFailed, i18n.T(kit, "batch create group app failed, err: %v", err))
	}

	return nil
}

// BatchDeleteByGroupIDWithTx batch delete group app by group id with transaction.
func (dao *groupAppDao) BatchDeleteByGroupIDWithTx(kit *kit.Kit, tx *gen.QueryTx, groupID, bizID uint32) error {

	if groupID == 0 {
		return errf.ErrInvalidIDF(kit)
	}

	if bizID == 0 {
		return errf.ErrInvalidBizIDF(kit)
	}

	m := tx.Query.GroupAppBind
	if _, err := tx.Query.GroupAppBind.WithContext(kit.Ctx).Where(
		m.GroupID.Eq(groupID), m.BizID.Eq(bizID)).Delete(); err != nil {
		return errf.Errorf(errf.DBOpFailed, i18n.T(kit, "batch delete group app failed, err: %v", err))
	}

	return nil
}

// BatchDeleteByAppIDWithTx batch delete group app by app id with transaction.
func (dao *groupAppDao) BatchDeleteByAppIDWithTx(kit *kit.Kit, tx *gen.QueryTx, appID, bizID uint32) error {

	if appID == 0 {
		return errf.ErrInvalidAppIDF(kit)
	}

	if bizID == 0 {
		return errf.ErrInvalidBizIDF(kit)
	}

	m := tx.GroupAppBind
	_, err := m.WithContext(kit.Ctx).Where(m.AppID.Eq(appID), m.BizID.Eq(bizID)).Delete()
	if err != nil {
		return errf.Errorf(errf.DBOpFailed, i18n.T(kit, "batch delete group app failed, err: %v", err))
	}

	return nil
}

// BatchListByGroupIDs batch list group app by group ids.
func (dao *groupAppDao) BatchListByGroupIDs(kit *kit.Kit,
	bizID uint32, groupIDs []uint32) ([]*table.GroupAppBind, error) {

	if bizID == 0 {
		return nil, errf.ErrInvalidBizIDF(kit)
	}

	if len(groupIDs) == 0 {
		return nil, nil
	}

	m := dao.genQ.GroupAppBind
	items, err := m.WithContext(kit.Ctx).Where(m.BizID.Eq(bizID), m.GroupID.In(groupIDs...)).Find()
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(kit, "list group app failed, err: %v", err))
	}
	return items, nil

}

// Get get GroupAppBind by group id and app id.
func (dao *groupAppDao) Get(kit *kit.Kit, groupID, appID, bizID uint32) (*table.GroupAppBind, error) {
	if bizID == 0 {
		return nil, errf.ErrInvalidBizIDF(kit)
	}
	if groupID == 0 {
		return nil, errf.ErrInvalidIDF(kit)
	}
	if appID == 0 {
		return nil, errf.ErrInvalidAppIDF(kit)
	}

	m := dao.genQ.GroupAppBind
	item, err := m.WithContext(kit.Ctx).Where(m.GroupID.Eq(groupID), m.AppID.Eq(appID), m.BizID.Eq(bizID)).Take()
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, "get group app bind failed, err: %v", err)
	}
	return item, nil
}
