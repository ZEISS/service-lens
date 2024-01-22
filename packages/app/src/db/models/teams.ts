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
  Unique,
  Default,
  BelongsToMany
} from 'sequelize-typescript'
import { Optional } from 'sequelize'
import { Workload } from '@/db/models/workload'
import { Ownership } from '@/db/models/ownership'
import { User } from '@/db/models/users'
import { UserTeam } from '@/db/models/users-teams'

export interface TeamAttributes {
  id: string
  name: string
  slug: string
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
  @Max(128)
  @Column
  declare name: string

  @NotEmpty
  @Unique
  @Min(3)
  @Max(128)
  @Column
  declare slug: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  declare description?: string

  @BelongsToMany(() => Workload, {
    through: {
      model: () => Ownership,
      unique: false,
      scope: {
        resourceType: 'workload'
      }
    },
    foreignKey: 'ownerId',
    otherKey: 'resourceId',
    constraints: false
  })
  declare workloads: Workload[]

  @BelongsToMany(() => User, {
    through: {
      model: () => UserTeam,
      unique: false
    },
    foreignKey: 'teamId',
    otherKey: 'userId',
    constraints: false
  })
  declare users: User[]

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
