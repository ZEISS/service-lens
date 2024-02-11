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
  name: Roles
  description?: string
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type RoleCreationAttributes = Omit<
  RoleAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

export type Roles =
  | 'superadmin'
  | 'owner'
  | 'member'
  | 'developer'
  | 'viewer'
  | 'contributor'

@Table({
  tableName: 'roles',
  modelName: 'Role'
})
export class Role extends Model<RoleAttributes, RoleCreationAttributes> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  declare id: bigint

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  declare name: string

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
