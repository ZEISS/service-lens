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
  Unique
} from 'sequelize-typescript'
import { WorkloadLensesAnswer } from '@/db/models/workload-lenses-answers'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'

export interface WorkloadLensesAnswerChoiceAttributes {
  id: string
  answerId: string
  choiceId: string
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
  id!: string

  @ForeignKey(() => WorkloadLensesAnswer)
  @Unique('workload-lens-pillar-answer')
  @Column
  answerId?: string

  @ForeignKey(() => LensPillarChoice)
  @Unique('workload-lens-pillar-answer')
  @Column
  choiceId?: string

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
