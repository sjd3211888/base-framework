// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package migration

import (
	"github.com/itcloudy/base-framework/pkg/migration/postgres"
)

var migrations = []*migration{
	&migration{
		version: "0.0.1",
		data:    postgres.Init,
	},
}
var updateMigrations []*migration