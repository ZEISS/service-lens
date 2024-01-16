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
  NotEmpty,
  Unique
} from 'sequelize-typescript'
import { Workload } from './workload'
import { QuestionRisk } from './lens-pillar-risks'
import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import { WorkloadLensesAnswerChoice } from './workload-lenses-answers-choices'

export interface WorkloadLensAnswerAttributes {
  id: string
  workloadId: string
  lensPillarQuestionId: string
  notes?: string
  lensChoices?: LensPillarChoice[]
  doesNotApply?: boolean
  doesNotApplyReason?: string
  risk: QuestionRisk
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type WorkloadLensAnswerCreationAttributes = Omit<
  WorkloadLensAnswerAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  modelName: 'WorkloadLensAnswer',
  tableName: 'workloads-lenses-answers'
})
export class WorkloadLensAnswer extends Model<
  WorkloadLensAnswerAttributes,
  WorkloadLensAnswerCreationAttributes
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

  @NotEmpty
  @Default(QuestionRisk.Unanswered)
  @Column(DataType.ENUM(...Object.values(QuestionRisk)))
  risk!: QuestionRisk

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
