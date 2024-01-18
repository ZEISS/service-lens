import { Optional } from 'sequelize'
import {
  Model,
  Table,
  DataType,
  CreatedAt,
  UpdatedAt,
  DeletedAt,
  Column,
  Unique
} from 'sequelize-typescript'

export type TaggableType = 'workload' | 'solution' | 'lens' | 'profile'

export type TagTaggableAttributes = {
  id: bigint
  taggableId: string
  taggableType: TaggableType
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type TagTaggableCreationAttributes = Optional<
  TagTaggableAttributes,
  'id'
>

@Table({ modelName: 'TagTaggable', tableName: 'tags-taggable' })
export class TagTaggable extends Model<
  TagTaggableAttributes,
  TagTaggableCreationAttributes
> {
  @Unique('tt_unique_constraint')
  @Column(DataType.BIGINT)
  declare tagId: bigint

  @Unique('tt_unique_constraint')
  @Column(DataType.UUID)
  declare taggableId: string

  @Unique('tt_unique_constraint')
  @Column(DataType.STRING)
  declare taggableType: TaggableType

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
