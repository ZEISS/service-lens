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
  AutoIncrement
} from 'sequelize-typescript'

export interface PermissionAttributes {
  id: bigint
  slug: string
  description?: string
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type PermissionCreationAttributes = Omit<
  PermissionAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

export type Permissions = 'admin' | 'write' | 'read'

@Table({
  tableName: 'permissions',
  modelName: 'Permission'
})
export class Permission extends Model<
  PermissionAttributes,
  PermissionCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  declare id: bigint

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  declare slug: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  declare description: string

  @CreatedAt
  @Column(DataType.DATE)
  declare createdAt: Date

  @UpdatedAt
  @Column(DataType.DATE)
  declare updatedAt: Date

  @DeletedAt
  @Column(DataType.DATE)
  declare deletedAt: Date
}
