import {
  Table,
  Model,
  CreatedAt,
  UpdatedAt,
  DeletedAt,
  Column,
  PrimaryKey,
  ForeignKey,
  AutoIncrement,
  DataType,
  Unique
} from 'sequelize-typescript'
import { WorkloadLensAnswer } from '@/db/models/workload-lenses-answers'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'

export interface WorkloadLensesAnswerChoiceAttributes {
  id: bigint
  answerId: bigint
  choiceId: bigint
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type WorkloadLensesAnswerChoiceCreationAttributes = Omit<
  WorkloadLensesAnswerChoiceAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'workloads-lenses-answers-choices'
})
export class WorkloadLensesAnswerChoice extends Model<
  WorkloadLensesAnswerChoiceAttributes,
  WorkloadLensesAnswerChoiceCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column
  declare id: bigint

  @ForeignKey(() => WorkloadLensAnswer)
  @Unique('workload-lens-pillar-answer')
  @Column
  declare answerId: bigint

  @ForeignKey(() => LensPillarChoice)
  @Unique('workload-lens-pillar-answer')
  @Column(DataType.BIGINT)
  declare choiceId: bigint

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
