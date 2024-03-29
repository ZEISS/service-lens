import {
  BelongsTo,
  Column,
  CreatedAt,
  DataType,
  DeletedAt,
  ForeignKey,
  BelongsToMany,
  Max,
  Min,
  Model,
  NotEmpty,
  PrimaryKey,
  Table,
  UpdatedAt,
  HasMany,
  AllowNull,
  Default
} from 'sequelize-typescript'
import { Profile } from './profile'
import { Lens } from './lens'
import { WorkloadLens } from './workload-lens'
import { WorkloadEnvironment } from './workload-environment'
import { WorkloadLensAnswer } from './workload-lenses-answers'
import { Environment } from './environment'
import { Tag } from '@/db/models/tags'
import { TagTaggable } from './tags-taggable'
import { Ownership } from './ownership'
import { Team } from './teams'

export interface WorkloadAttributes {
  id: string
  name: string
  description?: string
  environments?: Environment[]
  answers?: WorkloadLens[]
  lenses?: Lens[]
  profilesId?: string
  profile?: Profile
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type WorkloadCreationAttributes = Omit<
  WorkloadAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'workloads',
  modelName: 'Workload'
})
export class Workload extends Model<
  WorkloadAttributes,
  WorkloadCreationAttributes
> {
  @PrimaryKey
  @AllowNull(false)
  @Default(DataType.UUIDV4)
  @Column(DataType.UUIDV4)
  id!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column(DataType.STRING)
  declare name: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  declare description?: string

  @ForeignKey(() => Profile)
  @Column
  declare profilesId?: number

  @BelongsTo(() => Profile)
  declare profile: Profile

  @BelongsToMany(() => Lens, () => WorkloadLens, 'workloadId', 'lensId')
  declare lenses: Lens[]

  @BelongsToMany(() => Tag, {
    through: {
      model: () => TagTaggable,
      unique: false,
      scope: {
        taggableType: 'workload'
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
        resourceType: 'workload'
      }
    },
    foreignKey: 'resourceId',
    otherKey: 'ownerId',
    constraints: false
  })
  declare teams: Team[]

  @BelongsToMany(
    () => Environment,
    () => WorkloadEnvironment,
    'workloadId',
    'environmentId'
  )
  environments?: Environment[]

  @HasMany(() => WorkloadLensAnswer, 'workloadId')
  answers?: WorkloadLensAnswer[]

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
