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
  NotEmpty,
  Min,
  Max,
  AutoIncrement,
  HasMany
} from 'sequelize-typescript'
import { LensPillar } from './lens-pillars'
import { LensPillarChoice } from './lens-pillar-choices'
import { LensPillarQuestionRisk } from './lens-pillar-risks'
import { LensPillarQuestionResource } from '@/db/models/lens-pillar-questions-resources'

export interface LensPillarQuestionAttributes {
  id: bigint
  ref: string
  name: string
  pillarId: bigint
  description?: string
  risks?: LensPillarQuestionRisk[]
  questionAnswers?: LensPillarChoice[]
  resources?: LensPillarQuestionResource[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensPillarQuestionCreationAttributes = Omit<
  LensPillarQuestionAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses-pillars-questions'
})
export class LensPillarQuestion extends Model<
  LensPillarQuestionAttributes,
  LensPillarQuestionCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column
  id!: bigint

  @ForeignKey(() => LensPillar)
  @Column
  pillarId!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  ref!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name?: string

  @HasMany(() => LensPillarQuestionRisk, 'questionId')
  risks?: LensPillarQuestionRisk[]

  @HasMany(() => LensPillarChoice, 'questionId')
  questionAnswers?: LensPillarChoice[]

  @HasMany(() => LensPillarQuestionResource, 'questionId')
  resources?: LensPillarQuestionResource[]

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description?: string

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
