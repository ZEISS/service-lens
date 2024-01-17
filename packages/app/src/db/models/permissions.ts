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
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type PermissionCreationAttributes = Omit<
  PermissionAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

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
  id!: bigint

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  slug!: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description?: string

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
