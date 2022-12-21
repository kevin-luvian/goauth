CREATE TABLE IF NOT EXISTS users_workspaces (
  user_id       INT NOT NULL,
  workspace_tag VARCHAR(50) NOT NULL,
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, workspace_tag),
  CONSTRAINT fk_users_workspaces_user
    FOREIGN KEY(user_id) REFERENCES users(id) 
      ON DELETE CASCADE,
  CONSTRAINT fk_users_workspaces_workspace
    FOREIGN KEY(workspace_tag) REFERENCES workspaces(tag) 
      ON DELETE CASCADE
);