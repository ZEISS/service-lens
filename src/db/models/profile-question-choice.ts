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
  ForeignKey,
  BelongsTo
} from 'sequelize-typescript'
import { ProfileQuestion } from '@/db/models/profile-question'

export interface ProfileQuestionChoiceAttributes {
  id: number
  name: string
  ref: string
  description: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type ProfileQuestionChoiceCreationAttributes = Omit<
  ProfileQuestionChoiceAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'profiles-questions-choices'
})
export class ProfileQuestionChoice extends Model<
  ProfileQuestionChoiceAttributes,
  ProfileQuestionChoiceCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: string

  @ForeignKey(() => ProfileQuestion)
  @Column(DataType.BIGINT)
  questionId?: bigint

  @BelongsTo(() => ProfileQuestion)
  question?: ProfileQuestion

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  ref?: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name?: string

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
