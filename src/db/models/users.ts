import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  Unique,
  Default,
  CreatedAt,
  DeletedAt,
  UpdatedAt,
  BelongsToMany
} from 'sequelize-typescript'
import { UserTeam } from './users-teams'
import { Team } from './teams'
import { Workload } from './workload'
import { Ownership } from './ownership'

export interface UserAttributes {
  id: string
  email: string
  name: string
  emailVerified: string
  image?: string
}

export type UserCreationAttributes = Omit<UserAttributes, 'id'>

@Table({
  tableName: 'users',
  timestamps: false,
  underscored: true
})
export class User extends Model<UserAttributes, UserCreationAttributes> {
  @PrimaryKey
  @AutoIncrement
  @Default(DataType.UUIDV4)
  @Column(DataType.UUIDV4)
  id!: string

  @Column
  name?: string

  @Unique
  @Column
  email!: string

  @Column
  emailVerified?: Date

  @Column
  image?: string

  @BelongsToMany(() => Team, () => UserTeam, 'userId', 'teamId')
  teams?: Team[]

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
