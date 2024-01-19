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
  HasMany,
  Default,
  AllowNull,
  BelongsToMany
} from 'sequelize-typescript'
import { LensPillar } from './lens-pillars'
import type { Spec } from '../schemas/spec'
import { Tag } from './tags'
import { TagTaggable } from './tags-taggable'
import { Team } from './teams'
import { Ownership } from './ownership'

export interface LensAttributes {
  id: string
  version: number
  spec: object
  name: string
  isDraft: boolean
  description?: string
  pillars?: LensPillar[]
  tags: Tag[]
  teams: Team[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensCreationAttributes = Omit<
  LensAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses',
  modelName: 'Lens'
})
export class Lens extends Model<LensAttributes, LensCreationAttributes> {
  @PrimaryKey
  @AllowNull(false)
  @Default(DataType.UUIDV4)
  @Column(DataType.UUIDV4)
  id!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name!: string

  @NotEmpty
  @Column
  version!: string

  @NotEmpty
  @Column(DataType.JSONB)
  spec!: Spec

  @Default(true)
  @Column
  isDraft?: boolean

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description?: string

  @HasMany(() => LensPillar, 'lensId')
  pillars?: LensPillar[]

  @BelongsToMany(() => Tag, {
    through: {
      model: () => TagTaggable,
      unique: false,
      scope: {
        taggableType: 'lens'
      }
    },
    otherKey: 'tagId',
    foreignKey: 'taggableId',
    constraints: false
  })
  declare tags: Tag[]

  @BelongsToMany(() => Team, {
    through: {
      model: () => Ownership,
      unique: false,
      scope: {
        resourceType: 'lens'
      }
    },
    foreignKey: 'resourceId',
    otherKey: 'ownerId',
    constraints: false
  })
  declare teams: Team[]

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
