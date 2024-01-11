import {
  AutoIncrement,
  Column,
  CreatedAt,
  DataType,
  DeletedAt,
  ForeignKey,
  Model,
  PrimaryKey,
  Table,
  UpdatedAt
} from 'sequelize-typescript'
import { Team } from './team'
import { User } from './users'

export interface TeamMembersAttributes {
  id: bigint
  teamId: string
  userId: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type TeamMembersCreationAttributes = Omit<
  TeamMembersAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'teams-members'
})
export class TeamMembers extends Model<
  TeamMembersAttributes,
  TeamMembersCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: bigint

  @ForeignKey(() => Team)
  @Column(DataType.UUIDV4)
  teamId?: string

  @ForeignKey(() => User)
  @Column(DataType.UUIDV4)
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
