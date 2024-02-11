import {
  Table,
  Model,
  CreatedAt,
  UpdatedAt,
  DeletedAt,
  Column,
  PrimaryKey,
  DataType,
  NotEmpty,
  Min,
  Max,
  AutoIncrement,
  ForeignKey,
  AllowNull
} from 'sequelize-typescript'
import { Optional } from 'sequelize'
import { Permission } from '@/db/models/permissions'
import { Role } from '@/db/models/roles'

export interface RolePermissionAttributes {
  id: bigint
  slug: string
  description?: string
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type RolePermissionCreationAttributes = Optional<
  RolePermissionAttributes,
  'id'
>

@Table({
  tableName: 'roles-permissions',
  modelName: 'RolePermission'
})
export class RolePermission extends Model<
  RolePermissionAttributes,
  RolePermissionCreationAttributes
> {
  @PrimaryKey
  @AllowNull(false)
  @AutoIncrement
  @Column(DataType.BIGINT)
  declare id: bigint

  @ForeignKey(() => Role)
  @Column(DataType.BIGINT)
  declare roleId: bigint

  @ForeignKey(() => Permission)
  @Column(DataType.BIGINT)
  declare userId: string

  @CreatedAt
  @Column
  declare createdAt?: Date

  @UpdatedAt
  @Column
  declare updatedAt?: Date

  @DeletedAt
  @Column
  declare deletedAt?: Date
}
