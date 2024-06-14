model
  schema 1.1

type user

type team
  relations
    define admin: [user]
    define can_add_admin: can_add_owner
    define can_add_editor: can_add_admin or admin
    define can_add_owner: owner
    define can_add_viewer: can_add_editor or editor
    define can_create_environment: editor
    define can_create_lens: editor
    define can_create_profile: editor
    define can_create_workload: editor
    define can_delete: owner
    define can_delete_owner: can_add_owner
    define editor: [user] or admin
    define owner: [user]
    define viewer: [user] or editor or admin

type workload
  relations
    define admin: admin from team
    define can_delete: editor or admin
    define can_read: viewer
    define can_share: admin
    define can_write: editor or admin
    define editor: editor from team or admin
    define team: [team]
    define viewer: viewer from team or editor

type profile
  relations
    define admin: admin from team
    define can_delete: editor or admin
    define can_read: viewer
    define can_share: admin
    define can_write: editor or admin
    define editor: editor from team or admin
    define team: [team]
    define viewer: viewer from team or editor

type lens
  relations
    define admin: admin from team
    define can_delete: editor or admin
    define can_read: viewer
    define can_share: admin
    define can_write: editor or admin
    define editor: editor from team or admin
    define team: [team]
    define viewer: viewer from team or editor

type environment
  relations
    define admin: admin from team
    define can_delete: editor or admin
    define can_read: viewer
    define can_share: admin
    define can_write: editor or admin
    define editor: editor from team or admin
    define team: [team]
    define viewer: viewer from team or editor
