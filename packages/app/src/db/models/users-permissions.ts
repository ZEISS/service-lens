import { Table, Model, Column, DataType, AllowNull } from 'sequelize-typescript'
import { Optional } from 'sequelize'

export type UserPermissions = 'read' | 'write' | 'admin' | 'superadmin'

export interface UserPermissionAttributes {
  id: bigint
  userId: string
  teamId: string
  permission: string
}

export type UserPermissionCreationAttributes = Optional<
  UserPermissionAttributes,
  'id'
>

@Table({
  tableName: 'vw_user_teams_permissions'
})
export class UserPermission extends Model<
  UserPermissionAttributes,
  UserPermissionCreationAttributes
> {
  @AllowNull(false)
  @Column(DataType.UUIDV4)
  declare userId: string

  @AllowNull(false)
  @Column(DataType.UUIDV4)
  declare teamId: bigint

  @Column(DataType.STRING)
  declare permission: UserPermissions
}
