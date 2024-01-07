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
  Min,
  Max,
  Default,
  AllowNull,
  BelongsToMany,
  Unique
} from 'sequelize-typescript'
import { Workload } from './workload'
import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import { WorkloadLensesAnswerChoice } from './workload-lenses-answers-choices'

export interface WorkloadLensesAnswerAttributes {
  id: string
  workloadId: string
  lensPillarQuestionId: string
  notes?: string
  lensChoices?: LensPillarChoice[]
  doesNotApply?: boolean
  doesNotApplyReason?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type WorkloadLensesAnswerCreationAttributes = Omit<
  WorkloadLensesAnswerAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'workloads-lenses-answers'
})
export class WorkloadLensesAnswer extends Model<
  WorkloadLensesAnswerAttributes,
  WorkloadLensesAnswerCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: string

  @ForeignKey(() => Workload)
  @Unique('workload-lens-pillar-question')
  @Column(DataType.UUID)
  workloadId?: string

  @ForeignKey(() => LensPillarQuestion)
  @Unique('workload-lens-pillar-question')
  @Column
  lensPillarQuestionId?: bigint

  @AllowNull
  @Min(12)
  @Max(2048)
  @Column
  notes?: string

  @Default(false)
  @Column
  doesNotApply?: boolean

  @Column
  doesNotApplyReason?: string

  @BelongsToMany(
    () => LensPillarChoice,
    () => WorkloadLensesAnswerChoice,
    'answerId',
    'choiceId'
  )
  lensChoices?: LensPillarChoice[]

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
