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
  ForeignKey
} from 'sequelize-typescript'
import { Permission } from '@/db/models/permissions'
import { Role } from '@/db/models/roles'

export interface RolePermissionAttributes {
  id: bigint
  slug: string
  description?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type RolePermissionCreationAttributes = Omit<
  RolePermissionAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
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
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: bigint

  @ForeignKey(() => Role)
  @Column(DataType.BIGINT)
  roleId?: bigint

  @ForeignKey(() => Permission)
  @Column(DataType.BIGINT)
  userId?: string

  @CreatedAt
  @Column
  createdAt?: Date

  @UpdatedAt
  @Column
  updatedAt?: Date

  @DeletedAt
  @Column
  deletedAt?: Date
}
