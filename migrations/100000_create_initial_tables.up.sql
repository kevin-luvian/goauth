CREATE TABLE IF NOT EXISTS users (
  id      SERIAL NOT NULL PRIMARY KEY,
  tag     VARCHAR(50) NOT NULL UNIQUE,
  name    VARCHAR(225) NOT NULL,
  email   VARCHAR(225) NOT NULL UNIQUE,
  hpass   VARCHAR(225) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS workspaces (
  tag        VARCHAR(50) NOT NULL PRIMARY KEY,
  name       VARCHAR(225) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE IF NOT EXISTS apptokens (
  id            SERIAL NOT NULL PRIMARY KEY,
  token         VARCHAR(225) NOT NULL UNIQUE,
  workspace_tag VARCHAR(50) NOT NULL,
  created_by    INT,
  CONSTRAINT fk_apptokens_users
    FOREIGN KEY(created_by) REFERENCES users(id) 
      ON DELETE SET NULL,
  CONSTRAINT fk_apptokens_workspaces
    FOREIGN KEY(workspace_tag) REFERENCES workspaces(tag) 
      ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS roles (
  id            SERIAL NOT NULL PRIMARY KEY,
  workspace_tag VARCHAR(50) NOT NULL,
  name          VARCHAR(225) NOT NULL,
  created_by    INT,
  UNIQUE (workspace_tag, name),
  CONSTRAINT fk_roles_users
    FOREIGN KEY(created_by) REFERENCES users(id) 
      ON DELETE SET NULL,
  CONSTRAINT fk_roles_workspaces
    FOREIGN KEY(workspace_tag) REFERENCES workspaces(tag) 
      ON DELETE CASCADE
);