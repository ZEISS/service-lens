import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  ForeignKey,
  Unique,
  Default,
  BelongsToMany,
  AllowNull
} from 'sequelize-typescript'

export interface UserPermissionAttributes {
  id: bigint
  userId: string
  teamId: string
  permission: string
}

export type UserPermissionCreationAttributes = Omit<
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
  userId?: string

  @AllowNull(false)
  @Column(DataType.UUIDV4)
  teamId?: bigint

  @Column
  permission?: string
}
