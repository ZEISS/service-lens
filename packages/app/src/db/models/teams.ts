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
  AllowNull,
  Default
} from 'sequelize-typescript'
import { Optional } from 'sequelize'

export interface TeamAttributes {
  id: string
  name: string
  description?: string
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type TeamCreationAttributes = Optional<TeamAttributes, 'id'>

@Table({
  tableName: 'teams',
  modelName: 'Team'
})
export class Team extends Model<TeamAttributes, TeamCreationAttributes> {
  @PrimaryKey
  @Default(DataType.UUIDV4)
  @AllowNull(false)
  @Column(DataType.UUIDV4)
  declare id: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  declare name: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  declare description?: string

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
