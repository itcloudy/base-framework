
-- --------------------
-- 更新时间触发器
-- --------------------
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

-- ----------------------------
-- 数据库升级日志表
-- ----------------------------
CREATE TABLE migration_history (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  version VARCHAR(50) NOT NULL ,
  data text NOT NULL,
  UNIQUE(version)
);
COMMENT ON TABLE migration_history IS '数据库升级';
comment on column migration_history.id is '主键';
comment on column migration_history.created_at is '创建时间';
comment on column migration_history.version is '升级版本';
comment on column migration_history.data is '升级内容';

CREATE SEQUENCE migration_history_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table migration_history alter column id set default nextval('migration_history_id_seq');


CREATE TRIGGER update_migration_history_updated_at BEFORE UPDATE ON migration_history FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

-- ----------------------------
-- 系统用户
-- ----------------------------

create  table users (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  username VARCHAR(50) NOT NULL,
  alias  VARCHAR(50),
  head_image VARCHAR(200),
  email VARCHAR(50),
  pwd VARCHAR(200),
  mobile VARCHAR(50),
  is_active boolean DEFAULT true,
  is_admin boolean DEFAULT false,
  UNIQUE(username),
);

COMMENT ON TABLE users IS '系统用户';
comment on column users.created_at is '创建时间';
comment on column users.updated_at is '更新时间';
comment on column users.username is '用户名';
comment on column users.email is '邮箱';
comment on column users.pwd is '密码';
comment on column users.username is '有效';
comment on column users.username is '管理员';


CREATE SEQUENCE users_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table users alter column id set default nextval('users_id_seq');

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
