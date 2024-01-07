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

export interface EnvironmentAttributes {
  id: bigint
  name: string
  description: string
  label: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type EnvironmentCreationAttributes = Omit<
  EnvironmentAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'environments',
  modelName: 'Environment'
})
export class Environment extends Model<
  EnvironmentAttributes,
  EnvironmentCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  label!: string

  @NotEmpty
  @Column
  labelColor!: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description!: string

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
