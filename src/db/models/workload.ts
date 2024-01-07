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
  HasMany
} from 'sequelize-typescript'
import { Profile } from './profile'
import { Lens } from './lens'
import { WorkloadLens } from './workload-lens'
import { WorkloadEnvironment } from './workload-environment'
import { Environment } from './environment'

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
  'createdAt' | 'updatedAt' | 'deletedAt'
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

  @BelongsToMany(
    () => Environment,
    () => WorkloadEnvironment,
    'workloadId',
    'environmentId'
  )
  environments?: Environment[]

  @HasMany(() => WorkloadLens, 'workloadId')
  answers?: WorkloadLens[]

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
