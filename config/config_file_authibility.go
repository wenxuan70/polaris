/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or serverAuthibilityied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package config

import (
	"context"

	api "github.com/polarismesh/polaris-server/common/api/v1"
)

// CreateConfigFile 创建配置文件
func (cs *serverAuthibility) CreateConfigFile(ctx context.Context,
	configFile *api.ConfigFile) *api.ConfigResponse {

	authCtx := cs.collectBaseTokenInfo(ctx)
	if err := cs.authChecker.VerifyCredential(authCtx); err != nil {
		return api.NewConfigFileResponseWithMessage(convertToErrCode(err), err.Error())
	}

	ctx = authCtx.GetRequestContext()

	return cs.targetServer.CreateConfigFile(ctx, configFile)
}

// GetConfigFileBaseInfo 获取配置文件，只返回基础元信息
func (cs *serverAuthibility) GetConfigFileBaseInfo(ctx context.Context, namespace,
	group, name string) *api.ConfigResponse {

	return cs.targetServer.GetConfigFileBaseInfo(ctx, namespace, group, name)
}

// GetConfigFileRichInfo 获取单个配置文件基础信息，包含发布状态等信息
func (cs *serverAuthibility) GetConfigFileRichInfo(ctx context.Context, namespace,
	group, name string) *api.ConfigResponse {

	return cs.targetServer.GetConfigFileRichInfo(ctx, namespace, group, name)
}

// SearchConfigFile 查询配置文件
func (cs *serverAuthibility) SearchConfigFile(ctx context.Context, namespace, group, name,
	tags string, offset, limit uint32) *api.ConfigBatchQueryResponse {

	return cs.targetServer.SearchConfigFile(ctx, namespace, group, name, tags, offset, limit)
}

// UpdateConfigFile 更新配置文件
func (cs *serverAuthibility) UpdateConfigFile(ctx context.Context, configFile *api.ConfigFile) *api.ConfigResponse {

	authCtx := cs.collectBaseTokenInfo(ctx)
	if err := cs.authChecker.VerifyCredential(authCtx); err != nil {
		return api.NewConfigFileResponseWithMessage(convertToErrCode(err), err.Error())
	}

	ctx = authCtx.GetRequestContext()

	return cs.targetServer.UpdateConfigFile(ctx, configFile)
}

// DeleteConfigFile 删除配置文件，删除配置文件同时会通知客户端 Not_Found
func (cs *serverAuthibility) DeleteConfigFile(ctx context.Context, namespace, group,
	name, deleteBy string) *api.ConfigResponse {

	authCtx := cs.collectBaseTokenInfo(ctx)
	if err := cs.authChecker.VerifyCredential(authCtx); err != nil {
		return api.NewConfigFileResponseWithMessage(convertToErrCode(err), err.Error())
	}

	ctx = authCtx.GetRequestContext()

	return cs.targetServer.DeleteConfigFile(ctx, namespace, group, name, deleteBy)
}

// BatchDeleteConfigFile 批量删除配置文件
func (cs *serverAuthibility) BatchDeleteConfigFile(ctx context.Context, configFiles []*api.ConfigFile,
	operator string) *api.ConfigResponse {

	authCtx := cs.collectBaseTokenInfo(ctx)
	if err := cs.authChecker.VerifyCredential(authCtx); err != nil {
		return api.NewConfigFileResponseWithMessage(convertToErrCode(err), err.Error())
	}

	ctx = authCtx.GetRequestContext()

	return cs.targetServer.BatchDeleteConfigFile(ctx, configFiles, operator)
}
