import {
  Table,
  Model,
  Column,
  PrimaryKey,
  DataType,
  AutoIncrement,
  Default,
  CreatedAt,
  DeletedAt,
  UpdatedAt
} from 'sequelize-typescript'
import { Optional } from 'sequelize'

export interface SettingAttributes {
  id: string
  attributeName: string
  attributeValue: string
}

export type SettingCreationAttributes = Optional<SettingAttributes, 'id'>

@Table({
  tableName: 'settings',
  modelName: 'Settings'
})
export class Setting extends Model<
  SettingAttributes,
  SettingCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Default(DataType.UUIDV4)
  @Column(DataType.UUIDV4)
  declare id: string

  @Column(DataType.STRING)
  declare attributeName: string

  @Column(DataType.STRING)
  declare attributeValue: string

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
