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
import { Workload } from '@/db/models/workload'
import { WorkloadLens } from '@/db/models/workload-lens'

export interface LensAttributes {
  id: string
  version: number
  spec: object
  name: string
  isDraft: boolean
  description?: string
  pillars?: LensPillar[]
  workloads?: Workload[]
  tags?: Tag[]
  teams?: Team[]
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
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
  declare id: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column(DataType.STRING)
  declare name: string

  @NotEmpty
  @Column(DataType.STRING)
  declare version: string

  @NotEmpty
  @Column(DataType.JSONB)
  declare spec: Spec

  @Default(true)
  @Column(DataType.BOOLEAN)
  declare isDraft: boolean

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column(DataType.STRING)
  declare description?: string

  @HasMany(() => LensPillar, 'lensId')
  declare pillars?: LensPillar[]

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

  @BelongsToMany(() => Workload, {
    through: {
      model: () => WorkloadLens
    },
    otherKey: 'workloadId',
    foreignKey: 'lensId'
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
