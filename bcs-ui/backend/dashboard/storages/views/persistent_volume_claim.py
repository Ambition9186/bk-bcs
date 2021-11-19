# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""
from backend.dashboard.permissions import DisableCommonClusterRequest
from backend.dashboard.viewsets import NamespaceScopeViewSet
from backend.resources.storages.persistent_volume_claim import PersistentVolumeClaim


class PersistentVolumeClaimViewSet(NamespaceScopeViewSet):
    """ PersistentVolumeClaim 相关接口 """

    resource_client = PersistentVolumeClaim

    def get_permissions(self):
        # 目前 公共集群 不对用户开放资源视图 PVC 功能
        return [DisableCommonClusterRequest(), *super().get_permissions()]
