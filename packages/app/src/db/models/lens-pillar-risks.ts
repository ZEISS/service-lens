import {
  Column,
  CreatedAt,
  DataType,
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

export enum QuestionRisk {
  Unanswered = 'UNANSWERED',
  High = 'HIGH_RISK',
  Medium = 'MEDIUM_RISK',
  Low = 'LOW_RISK',
  None = 'NO_RISK'
}

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
  id!: bigint

  @ForeignKey(() => LensPillarQuestion)
  @Column
  questionId!: bigint

  @NotEmpty
  @Default(QuestionRisk.Unanswered)
  @Column(DataType.ENUM(...Object.values(QuestionRisk)))
  risk!: QuestionRisk

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  condition?: string

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
