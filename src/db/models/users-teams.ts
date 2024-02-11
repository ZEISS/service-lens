import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  ForeignKey
} from 'sequelize-typescript'
import { User } from './users'
import { Team } from './teams'

export interface UserTeamAttributes {
  id: bigint
  userId: string
  teamId: string
}

export type UserRoleCreationAttributes = Omit<UserTeamAttributes, 'id'>

@Table({
  tableName: 'users-teams'
})
export class UserTeam extends Model<
  UserTeamAttributes,
  UserRoleCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  declare id: bigint

  @ForeignKey(() => User)
  @Column(DataType.UUIDV4)
  declare userId: string

  @ForeignKey(() => Team)
  @Column(DataType.UUIDV4)
  declare teamId: string
}
