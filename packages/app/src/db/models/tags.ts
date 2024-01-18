import {
  Model,
  DataType,
  Table,
  Column,
  PrimaryKey,
  AutoIncrement,
  AllowNull,
  UpdatedAt,
  CreatedAt,
  DeletedAt,
  BelongsToMany,
  Scopes
} from 'sequelize-typescript'
import { Optional } from 'sequelize'

import { Workload } from '@/db/models/workload'
import { TagTaggable } from './tags-taggable'

export type TagAttributes = {
  id: bigint
  name: string
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type TagCreationAttributes = Optional<TagAttributes, 'id'>

@Table({
  modelName: 'Tag',
  tableName: 'tags'
})
export class Tag extends Model<TagAttributes, TagCreationAttributes> {
  @PrimaryKey
  @AutoIncrement
  @AllowNull(false)
  @Column(DataType.BIGINT)
  declare id: bigint

  @AllowNull(false)
  @Column(DataType.STRING)
  declare name: string

  @BelongsToMany(() => Workload, {
    through: {
      model: () => TagTaggable,
      unique: false
    },
    foreignKey: 'tagId',
    otherKey: 'taggableId',
    constraints: false
  })
  declare workloads: Workload[]

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
