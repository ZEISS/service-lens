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

export interface RoleAttributes {
  id: bigint
  name: string
  description?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type RoleCreationAttributes = Omit<
  RoleAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'roles',
  modelName: 'Role'
})
export class Role extends Model<RoleAttributes, RoleCreationAttributes> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: bigint

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name!: string

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
