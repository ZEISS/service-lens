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
import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import { WorkloadLensesAnswerChoice } from './workload-lenses-answers-choices'

export interface WorkloadLensAnswerAttributes {
  id: bigint
  workloadId: string
  lensPillarQuestionId: string
  notes?: string
  lensChoices?: LensPillarChoice[]
  doesNotApply?: boolean
  doesNotApplyReason?: string
  risk: QuestionRisk
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type WorkloadLensAnswerCreationAttributes = Omit<
  WorkloadLensAnswerAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

export const questionRisk = DataType.ENUM(
  'UNANSWERED',
  'HIGH_RISK',
  'MEDIUM_RISK',
  'LOW_RISK',
  'NO_RISK'
)
export type QuestionRisk = NonNullable<typeof questionRisk>

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
  declare id: bigint

  @ForeignKey(() => Workload)
  @Unique('workload-lens-pillar-question')
  @Column(DataType.UUID)
  declare workloadId: string

  @ForeignKey(() => LensPillarQuestion)
  @Unique('workload-lens-pillar-question')
  @Column(DataType.BIGINT)
  declare lensPillarQuestionId: bigint

  @AllowNull
  @Min(12)
  @Max(2048)
  @Column(DataType.STRING)
  declare notes?: string

  @Default(false)
  @Column(DataType.BOOLEAN)
  declare doesNotApply: boolean

  @Column(DataType.STRING)
  declare doesNotApplyReason: string

  @BelongsToMany(
    () => LensPillarChoice,
    () => WorkloadLensesAnswerChoice,
    'answerId',
    'choiceId'
  )
  declare lensChoices?: LensPillarChoice[]

  @NotEmpty
  @Default('UNANSWERED')
  @Column(
    DataType.ENUM(
      'UNANSWERED',
      'HIGH_RISK',
      'MEDIUM_RISK',
      'LOW_RISK',
      'NO_RISK'
    )
  )
  declare risk: QuestionRisk

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
