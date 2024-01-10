import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  Unique,
  Default
} from 'sequelize-typescript'

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
}
