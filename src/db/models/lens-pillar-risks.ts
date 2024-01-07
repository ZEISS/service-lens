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
  ForeignKey
} from 'sequelize-typescript'
import { LensPillarQuestion } from './lens-pillar-questions'

export interface LensPillarQuestionRiskAttributes {
  id: bigint
  lensId: string
  risk: string
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
  tableName: 'lenses-pillars-risks'
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
  questionId?: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  risk?: string

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
