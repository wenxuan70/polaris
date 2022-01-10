/*
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
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package httpserver

import (
	api "github.com/polarismesh/polaris-server/common/api/v1"
	"sync"
)

type connection struct {
	finishChan chan interface{}
	handler    *Handler
}

//client -> connection
var conns = new(sync.Map)

func (h *HTTPServer) addConn(clientId string, watchConfigFiles []*api.ClientConfigFileInfo,
	handler *Handler, finishChan chan interface{}) {
	conns.Store(clientId, &connection{
		finishChan: finishChan,
		handler:    handler,
	})

	h.configServer.WatchCenter().AddWatcher(clientId, watchConfigFiles, func(clientId string, rsp *api.ConfigClientResponse) bool {
		conn, ok := conns.Load(clientId)
		if ok {
			c := conn.(*connection)
			c.handler.WriteHeaderAndProto(rsp)
			c.finishChan <- new(interface{})
		}
		return true
	})
}

func (h *HTTPServer) removeConn(clientId string, watchConfigFiles []*api.ClientConfigFileInfo) {
	conns.Delete(clientId)
	h.configServer.WatchCenter().RemoveWatcher(clientId, watchConfigFiles)
}
