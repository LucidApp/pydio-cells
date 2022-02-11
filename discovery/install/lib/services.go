/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package lib

import (
	"context"
	"fmt"
	"github.com/pydio/cells/v4/common/plugins"
	registry2 "github.com/pydio/cells/v4/common/proto/registry"
	"github.com/pydio/cells/v4/common/registry"
	"github.com/pydio/cells/v4/common/service"
	servicecontext "github.com/pydio/cells/v4/common/service/context"
	"path/filepath"
	"sync"

	"github.com/ory/hydra/x"

	"github.com/pydio/cells/v4/common"
	"github.com/pydio/cells/v4/common/config"
	"github.com/pydio/cells/v4/common/proto/install"
)

type DexClient struct {
	Id                     string
	Name                   string
	Secret                 string
	RedirectURIs           []string
	IdTokensExpiry         string
	RefreshTokensExpiry    string
	OfflineSessionsSliding bool
}

func addBoltDbEntry(sName string, db ...string) error {
	bDir, e := config.ServiceDataDir(common.ServiceGrpcNamespace_ + sName)
	if e != nil {
		return e
	}
	dbName := sName + ".db"
	if len(db) > 0 {
		dbName = db[0]
	}
	return config.SetDatabase(common.ServiceGrpcNamespace_+sName, "boltdb", filepath.Join(bDir, dbName))
}

func addBleveDbEntry(sName string, db ...string) error {
	bDir, e := config.ServiceDataDir(common.ServiceGrpcNamespace_ + sName)
	if e != nil {
		return e
	}
	dbName := sName + ".bleve"
	if len(db) > 0 {
		dbName = db[0]
	}
	configKey := common.ServiceGrpcNamespace_ + sName
	if len(db) > 1 {
		configKey = db[1]
	}
	return config.SetDatabase(configKey, "bleve", filepath.Join(bDir, dbName))
}

var (
	listRegistry registry.Registry
	loadRegistry sync.Once
)

func ListServicesWithStorage() (ss []service.Service, e error) {
	loadRegistry.Do(func() {
		ctx := context.Background()
		reg, err := registry.OpenRegistry(ctx, "memory:///?cache=shared")
		if err != nil {
			e = err
		}
		plugins.Init(servicecontext.WithRegistry(ctx, reg), "main")
		listRegistry = reg
	})
	if e != nil {
		return nil, e
	}
	items, er := listRegistry.List(registry.WithType(registry2.ItemType_SERVICE))
	if er != nil {
		return nil, er
	}
	for _, i := range items {
		var srv service.Service
		if i.As(&srv) {
			if len(srv.Options().Storages) > 0 {
				ss = append(ss, srv)
			}
		} else {
			fmt.Println("cannot convert", i)
		}
	}
	return ss, nil
}

func actionConfigsSet(c *install.InstallConfig) error {
	save := false

	// OAuth web
	oauthWeb := common.ServiceWebNamespace_ + common.ServiceOAuth
	secret, err := x.GenerateSecret(32)
	if err != nil {
		return err
	}

	/*
		if e := addBoltDbEntry(common.ServiceActivity, "activities.db"); e != nil {
			return e
		}
		if e := addBoltDbEntry(common.ServiceChat); e != nil {
			return e
		}
		if e := addBoltDbEntry(common.ServiceDocStore); e != nil {
			return e
		}
		if e := addBoltDbEntry(common.ServiceJobs); e != nil {
			return e
		}
		if e := addBleveDbEntry(common.ServiceJobs, "tasklogs.bleve?mapping=logs&rotationSize=-1", common.ServiceGrpcNamespace_+common.ServiceJobs+".index"); e != nil {
			return e
		}
		if e := addBleveDbEntry(common.ServiceLog, "syslog.bleve?mapping=logs"); e != nil {
			return e
		}
		if e := addBoltDbEntry(common.ServiceVersions); e != nil {
			return e
		}
		if e := addBleveDbEntry(common.ServiceSearch, "searchengine.bleve?rotationSize=-1"); e != nil {
			return e
		}
	*/

	// Easy finding usage of srvUrl
	configKeys := map[string]interface{}{
		"services/" + oauthWeb + "/insecureRedirects": []string{"#insecure_binds...#/auth/callback"},
		"services/" + oauthWeb + "/secret":            string(secret),
	}

	for path, def := range configKeys {
		val := config.Get(path)
		var data interface{}
		if val.Scan(&data); data != def {
			val.Set(def)
			save = true
		}
	}

	if save {
		config.Save("cli", "Install / Setting Dex default settings")
	}

	return nil
}
