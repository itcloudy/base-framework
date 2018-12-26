// Copyright 2018 itcloudy@qa.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package postgres

var Update_0_0_2 = `
create  table update_test (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP 
);
`
