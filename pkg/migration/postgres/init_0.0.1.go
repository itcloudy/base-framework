// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package postgres

var Init = `
-- ----------------------------
-- Table structure for departments
-- ----------------------------
DROP TABLE IF EXISTS "public"."departments";
CREATE TABLE "public"."departments" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "name" text COLLATE "pg_catalog"."default",
  "department_no" text COLLATE "pg_catalog"."default",
  "parent_id" text COLLATE "pg_catalog"."default",
  "province" text COLLATE "pg_catalog"."default",
  "city" text COLLATE "pg_catalog"."default",
  "district" text COLLATE "pg_catalog"."default",
  "location" text COLLATE "pg_catalog"."default",
  "counter_tel" text COLLATE "pg_catalog"."default",
  "department_tel" text COLLATE "pg_catalog"."default",
  "remarks" text COLLATE "pg_catalog"."default",
  "status" int4,
  "pic" text COLLATE "pg_catalog"."default",
  "pic_tel" text COLLATE "pg_catalog"."default",
  "level" int4,
  "sort" int4,
  "chain_table_id" int4,
  "chain_table_name" text COLLATE "pg_catalog"."default",
  "block_id" int4,
  "hash_content" text COLLATE "pg_catalog"."default",
  "chain_err" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Primary Key structure for table departments
-- ----------------------------
ALTER TABLE "public"."departments" ADD CONSTRAINT "departments_pkey" PRIMARY KEY ("id");
`
