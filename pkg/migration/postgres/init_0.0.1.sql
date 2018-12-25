
-- --------------------
-- 更新时间触发器
-- --------------------
create or replace function upd_timestamp() returns trigger as
$$
begin
    new.updated_at = current_timestamp;
    return new;
end
$$
language plpgsql;


-- ----------------------------
-- 文件上传记录
--  ----------------------------
create  table file_upload (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  address VARCHAR(200),
  user_id integer ,
  file_name VARCHAR(200),
  type VARCHAR(20) NOT NULL default 'local',
  UNIQUE(address)
);

COMMENT ON TABLE file_upload IS '文件上传记录';
comment on column file_upload.id is '主键';
comment on column file_upload.created_at is '创建时间';
comment on column file_upload.updated_at is '修改时间';
comment on column file_upload.address is '目标文件';
comment on column file_upload.user_id is '上传用户';
comment on column file_upload.file_name is '文件名';
comment on column file_upload.type is '存储类型';


CREATE SEQUENCE file_upload_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table file_upload alter column id set default nextval('file_upload_id_seq');

CREATE TRIGGER file_upload_updated_at BEFORE UPDATE ON file_upload FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();


-- ----------------------------
-- 系统消息
--  ----------------------------
create  table notice (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  title varchar(50),
  description varchar(300) not null default '',
  type varchar(20),
  status varchar(20),
  extra varchar(50)
);

COMMENT ON TABLE notice IS '文件上传记录';
comment on column notice.id is '主键';
comment on column notice.created_at is '创建时间';
comment on column notice.updated_at is '修改时间';
comment on column notice.title is '标题';
comment on column notice.description is '内容';
comment on column notice.type is '类型';
comment on column notice.status is '状态';
comment on column notice.extra is '备注';


CREATE SEQUENCE notice_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table notice alter column id set default nextval('notice_id_seq');

CREATE TRIGGER notice_updated_at BEFORE UPDATE ON notice FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();

-- ----------------------------
-- 系统角色
--  ----------------------------
create  table role (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  code varchar(20),
  inherit_strings varchar(300),
  unique(code)
);

COMMENT ON TABLE role IS '系统角色';
comment on column role.id is '主键';
comment on column role.created_at is '创建时间';
comment on column role.updated_at is '修改时间';
comment on column role.code is '标题';
comment on column role.inherit_strings is '所继承角色ID逗号分隔';

CREATE SEQUENCE role_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table role alter column id set default nextval('role_id_seq');

CREATE TRIGGER role_updated_at BEFORE UPDATE ON role FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();

-- ----------------------------
-- 系统角色拥有的接口
--  ----------------------------
create  table role_api (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  role_id integer,
  system_api_id integer
);

COMMENT ON TABLE role_api IS '系统角色拥有的接口';
comment on column role_api.id is '主键';
comment on column role_api.created_at is '创建时间';
comment on column role_api.updated_at is '修改时间';
comment on column role_api.role_id is '角色ID';
comment on column role_api.system_api_id is '接口ID';

CREATE SEQUENCE role_api_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table role_api alter column id set default nextval('role_api_id_seq');

CREATE TRIGGER role_api_updated_at BEFORE UPDATE ON role_api FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();



-- ----------------------------
-- 系统接口
--  ----------------------------
create  table system_api (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  name varchar(50),
  address varchar(100),
  method varchar(10),
  display varchar(200),
  is_active bool default true,
  unique(display)
);

COMMENT ON TABLE system_api IS '系统接口';
comment on column system_api.id is '主键';
comment on column system_api.created_at is '创建时间';
comment on column system_api.updated_at is '修改时间';
comment on column system_api.name is '接口名称';
comment on column system_api.address is '接口地址';
comment on column system_api.method is '接口方法';
comment on column system_api.display is '显示名称';
comment on column system_api.is_active is '有效';

CREATE SEQUENCE system_api_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table system_api alter column id set default nextval('system_api_id_seq');

CREATE TRIGGER system_api_updated_at BEFORE UPDATE ON system_api FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();

-- ----------------------------
-- 系统用户
-- ----------------------------

create  table users (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  username VARCHAR(50) NOT NULL,
  alias  VARCHAR(50),
  avatar VARCHAR(200),
  email VARCHAR(50),
  pwd VARCHAR(200),
  mobile VARCHAR(50),
  is_active boolean DEFAULT true,
  is_admin boolean DEFAULT false,
  UNIQUE(username)
);

COMMENT ON TABLE users IS '系统用户';
comment on column users.id is '主键';
comment on column users.created_at is '创建时间';
comment on column users.updated_at is '修改时间';
comment on column users.username is '用户名';
comment on column users.email is '邮箱';
comment on column users.pwd is '密码';
comment on column users.is_active is '有效';
comment on column users.is_admin is '管理员';


CREATE SEQUENCE users_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table users alter column id set default nextval('users_id_seq');

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();

-- ----------------------------
-- 用户角色
--  ----------------------------
create  table user_role (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  role_id integer,
  user_id integer,
  is_active bool default true
);

COMMENT ON TABLE user_role IS '用户角色';
comment on column user_role.id is '主键';
comment on column user_role.created_at is '创建时间';
comment on column user_role.updated_at is '修改时间';
comment on column user_role.role_id is '角色ID';
comment on column user_role.user_id is '用户ID';
comment on column user_role.is_active is '有效';


CREATE SEQUENCE user_role_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table user_role alter column id set default nextval('user_role_id_seq');

CREATE TRIGGER user_role_updated_at BEFORE UPDATE ON role_api FOR EACH ROW EXECUTE PROCEDURE upd_timestamp();
