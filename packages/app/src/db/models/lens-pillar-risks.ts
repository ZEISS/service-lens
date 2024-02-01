import {
  Column,
  CreatedAt,
  DeletedAt,
  Model,
  PrimaryKey,
  Table,
  AutoIncrement,
  UpdatedAt,
  Min,
  Max,
  NotEmpty,
  Default,
  ForeignKey
} from 'sequelize-typescript'
import { LensPillarQuestion } from './lens-pillar-questions'
import { type QuestionRisk, questionRisk } from './workload-lenses-answers'

export interface LensPillarQuestionRiskAttributes {
  id: bigint
  questionId: bigint
  risk: QuestionRisk
  condition: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensPillarQuestionCreationAttributes = Omit<
  LensPillarQuestionRiskAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses-pillars-risks',
  modelName: 'LensPillarQuestionRisk'
})
export class LensPillarQuestionRisk extends Model<
  LensPillarQuestionRiskAttributes,
  LensPillarQuestionCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column
  declare id: bigint

  @ForeignKey(() => LensPillarQuestion)
  @Column
  declare questionId: bigint

  @NotEmpty
  @Default('UNANSWERED')
  @Column(questionRisk)
  declare risk: QuestionRisk

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  declare condition: string

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
