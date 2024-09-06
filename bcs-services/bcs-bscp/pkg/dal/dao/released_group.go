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
	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/types"
)

// ReleasedGroup supplies all the group related operations.
type ReleasedGroup interface {
	// ListAll list all released groups in biz
	ListAll(kit *kit.Kit, bizID uint32) ([]*table.ReleasedGroup, error)
	// ListAllByGroupID list all released groups by groupID
	ListAllByGroupID(kit *kit.Kit, groupID, bizID uint32) ([]*table.ReleasedGroup, error)
	// ListAllByAppID list all released groups by appID
	ListAllByAppID(kit *kit.Kit, appID, bizID uint32) ([]*table.ReleasedGroup, error)
	// ListAllByReleaseID list all released groups by releaseID
	ListAllByReleaseID(kit *kit.Kit, releaseID, bizID uint32) ([]*table.ReleasedGroup, error)
	// CountGroupsReleasedApps counts each group's published apps.
	CountGroupsReleasedApps(kit *kit.Kit, opts *types.CountGroupsReleasedAppsOption) (
		[]*types.GroupPublishedAppsCount, error)
	// UpdateEditedStatusWithTx update edited status with transaction
	UpdateEditedStatusWithTx(kit *kit.Kit, tx *gen.QueryTx, edited bool, groupID, bizID uint32) error
	// BatchDeleteByAppIDWithTx batch delete by app id with transaction.
	BatchDeleteByAppIDWithTx(kit *kit.Kit, tx *gen.QueryTx, appID, bizID uint32) error
}

var _ ReleasedGroup = new(releasedGroupDao)

type releasedGroupDao struct {
	genQ     *gen.Query
	idGen    IDGenInterface
	auditDao AuditDao
	lock     LockDao
}

// ListAll list all released groups in biz
func (dao *releasedGroupDao) ListAll(kit *kit.Kit, bizID uint32) ([]*table.ReleasedGroup, error) {
	if bizID == 0 {
		return nil, errf.ErrInvalidBizIDF(kit)
	}

	m := dao.genQ.ReleasedGroup
	items, err := m.WithContext(kit.Ctx).Where(m.BizID.Eq(bizID)).Find()
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(kit, "list released groups failed, err: %v", err))
	}
	return items, nil
}

// ListByGroupID list released groups by groupID
func (dao *releasedGroupDao) ListAllByGroupID(kit *kit.Kit, groupID, bizID uint32) ([]*table.ReleasedGroup, error) {
	if bizID == 0 {
		return nil, errf.ErrInvalidBizIDF(kit)
	}

	m := dao.genQ.ReleasedGroup
	items, err := m.WithContext(kit.Ctx).Where(m.GroupID.Eq(groupID), m.BizID.Eq(bizID)).Find()
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(kit, "list released groups failed, err: %v", err))
	}
	return items, nil
}

// ListByGroupID list released groups by appID
func (dao *releasedGroupDao) ListAllByAppID(kit *kit.Kit, appID, bizID uint32) ([]*table.ReleasedGroup, error) {
	if bizID == 0 {
		return nil, errf.ErrInvalidBizIDF(kit)
	}

	if appID == 0 {
		return nil, errf.ErrInvalidAppIDF(kit)
	}

	m := dao.genQ.ReleasedGroup
	items, err := m.WithContext(kit.Ctx).Where(m.AppID.Eq(appID), m.BizID.Eq(bizID)).Find()
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(kit, "list released group failed, err: %v", err))
	}
	return items, nil
}

// ListAllByReleaseID list all released groups by releaseID
func (dao *releasedGroupDao) ListAllByReleaseID(kit *kit.Kit, releaseID, bizID uint32) ([]*table.ReleasedGroup, error) {
	if bizID == 0 {
		return nil, errf.ErrInvalidBizIDF(kit)
	}

	if releaseID == 0 {
		return nil, errf.ErrInvalidIDF(kit)
	}

	m := dao.genQ.ReleasedGroup
	items, err := m.WithContext(kit.Ctx).Where(m.ReleaseID.Eq(releaseID), m.BizID.Eq(bizID)).Find()
	if err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(kit, "list released groups failed, err: %v", err))
	}
	return items, nil
}

// CountGroupsReleasedApps counts each group's published apps.
func (dao *releasedGroupDao) CountGroupsReleasedApps(kit *kit.Kit, opts *types.CountGroupsReleasedAppsOption) (
	[]*types.GroupPublishedAppsCount, error) {

	if opts == nil {
		return nil, errf.New(errf.InvalidParameter, i18n.T(kit, "count groups released apps option is nil"))
	}

	if err := opts.Validate(kit, nil); err != nil {
		return nil, err
	}

	counts := make([]*types.GroupPublishedAppsCount, 0)

	m := dao.genQ.ReleasedGroup
	if err := m.WithContext(kit.Ctx).
		Select(m.GroupID, m.AppID.Distinct().Count().As("counts"), m.Edited.Max().As("edited")).
		Where(m.BizID.Eq(opts.BizID), m.GroupID.In(opts.Groups...)).
		Group(m.GroupID).
		Scan(&counts); err != nil {
		return nil, errf.Errorf(errf.DBOpFailed, i18n.T(kit, "counts each group's published apps failed, err: %v", err))
	}
	return counts, nil
}

// UpdateEditedStatusWithTx update edited status with transaction
func (dao *releasedGroupDao) UpdateEditedStatusWithTx(kit *kit.Kit,
	tx *gen.QueryTx, edited bool, groupID, bizID uint32) error {
	if bizID == 0 {
		return errf.ErrInvalidBizIDF(kit)
	}
	if groupID == 0 {
		// group id is 0, means it is a default group,can not be edited
		return errf.ErrInvalidIDF(kit)
	}

	m := tx.ReleasedGroup

	if _, err := m.WithContext(kit.Ctx).
		Where(m.GroupID.Eq(groupID), m.BizID.Eq(bizID)).
		Update(m.Edited, edited); err != nil {
		return errf.Errorf(errf.DBOpFailed, i18n.T(kit, "update edited status failed, err: %v", err))
	}
	return nil
}

// BatchDeleteByAppIDWithTx batch delete by app id with transaction.
func (dao *releasedGroupDao) BatchDeleteByAppIDWithTx(kit *kit.Kit, tx *gen.QueryTx, appID, bizID uint32) error {
	if bizID == 0 {
		return errf.ErrInvalidBizIDF(kit)
	}
	if appID == 0 {
		return errf.ErrInvalidAppIDF(kit)
	}

	m := tx.ReleasedGroup
	_, err := m.WithContext(kit.Ctx).Where(m.AppID.Eq(appID), m.BizID.Eq(bizID)).Delete()
	if err != nil {
		return errf.Errorf(errf.DBOpFailed, i18n.T(kit, "batch delete release groups failed, err: %v", err))
	}
	return err
}
