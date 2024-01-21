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
import { Role } from './roles'
import { Team } from './teams'

export interface UserRoleAttributes {
  id: bigint
  userId: string
  roleId: bigint
  teamId: string
}

export type UserRoleCreationAttributes = Omit<UserRoleAttributes, 'id'>

@Table({
  tableName: 'users-roles'
})
export class UserRole extends Model<
  UserRoleAttributes,
  UserRoleCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  declare id: bigint

  @ForeignKey(() => User)
  @Column(DataType.UUIDV4)
  declare userId: string

  @ForeignKey(() => Role)
  @Column(DataType.BIGINT)
  declare roleId: bigint

  @ForeignKey(() => Team)
  @Column(DataType.UUIDV4)
  declare teamId: string
}
