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
  Default,
  HasOne
} from 'sequelize-typescript'
import { Profile } from './profile'
import { Lens } from './lens'
import { WorkloadLens } from './workload-lens'
import { WorkloadEnvironment } from './workload-environment'
import { WorkloadLensAnswer } from './workload-lenses-answers'
import { Environment } from './environment'
import { Tag } from '@/db/models/tags'
import { TagTaggable } from './tags-taggable'

export interface WorkloadAttributes {
  id: string
  name: string
  description?: string
  environments?: Environment[]
  answers?: WorkloadLens[]
  lenses?: Lens[]
  profilesId: string
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
  @Column
  name!: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description?: string

  @ForeignKey(() => Profile)
  @Column
  profilesId?: number

  @BelongsTo(() => Profile)
  profile?: Profile

  @BelongsToMany(() => Lens, () => WorkloadLens, 'workloadId', 'lensId')
  lenses?: Lens[]

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
  declare tags?: Tag[]

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
  @Column
  createdAt?: Date

  @UpdatedAt
  @Column
  updatedAt?: Date

  @DeletedAt
  @Column
  deletedAt?: Date
}
