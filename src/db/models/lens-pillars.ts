import {
  Table,
  Model,
  CreatedAt,
  UpdatedAt,
  DeletedAt,
  Column,
  PrimaryKey,
  DataType,
  ForeignKey,
  AutoIncrement,
  NotEmpty,
  Min,
  Max,
  HasMany
} from 'sequelize-typescript'
import { Lens } from './lens'
import { LensPillarQuestion } from './lens-pillar-questions'
import { LensPillarResource } from './lens-pillar-resources'

export interface LensPillarAttributes {
  id: bigint
  name: string
  ref: string
  description: string
  questions: LensPillarQuestion[]
  resources: LensPillarResource[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensPillarCreationAttributes = Omit<
  LensPillarAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses-pillars'
})
export class LensPillar extends Model<
  LensPillarAttributes,
  LensPillarCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column
  id!: bigint

  @ForeignKey(() => Lens)
  @Column(DataType.UUIDV4)
  lensId!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  ref!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name!: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description!: string

  @HasMany(() => LensPillarQuestion, 'pillarId')
  questions?: LensPillarQuestion[]

  @HasMany(() => LensPillarResource, 'pillarId')
  resources?: LensPillarResource[]

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
