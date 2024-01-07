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
  AllowNull,
  Default
} from 'sequelize-typescript'
import { LensPillar } from '@/db/models/lens-pillars'

export interface LensPillarQuestionResourceAttributes {
  id: string
  questionId: bigint
  description: string
  url?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensPillarQuestionResourceCreationAttributes = Omit<
  LensPillarQuestionResourceAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses-pillars-questions-resources'
})
export class LensPillarQuestionResource extends Model<
  LensPillarQuestionResourceAttributes,
  LensPillarQuestionResourceCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: string

  @ForeignKey(() => LensPillar)
  @Column
  questionId!: bigint

  @NotEmpty
  @Column
  description!: string

  @AllowNull
  @Column
  url?: string

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
