import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  ForeignKey,
  Unique,
  Default,
  BelongsToMany
} from 'sequelize-typescript'
import { User } from './users'
import { Role } from './roles'
import { Team } from './teams'

export interface UserRoleAttributes {
  id: string
  email: string
  name: string
  emailVerified: string
  image?: string
}

export type UserRoleCreationAttributes = Omit<UserRoleAttributes, 'id'>

@Table({
  tableName: 'users',
  timestamps: false,
  underscored: true
})
export class UserRole extends Model<
  UserRoleAttributes,
  UserRoleCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: bigint

  @ForeignKey(() => User)
  @Column(DataType.UUIDV4)
  userId?: string

  @ForeignKey(() => Role)
  @Column(DataType.BIGINT)
  roleId?: bigint

  @ForeignKey(() => Team)
  @Column(DataType.UUIDV4)
  teamId?: bigint
}
