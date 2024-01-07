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
  ForeignKey
} from 'sequelize-typescript'
import { Profile } from '@/db/models/profile'
import { ProfileQuestionChoice } from '@/db/models/profile-question-choice'

export interface ProfileQuestionAnswerAttributes {
  id: number
  choiceId: bigint
  profileId: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type ProfileQuestionAnswerCreationAttributes = Omit<
  ProfileQuestionAnswerAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'profiles-questions-answers'
})
export class ProfileQuestionAnswer extends Model<
  ProfileQuestionAnswerAttributes,
  ProfileQuestionAnswerCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: string

  @ForeignKey(() => Profile)
  @Column(DataType.UUID)
  profileId?: string

  @ForeignKey(() => ProfileQuestionChoice)
  @Column
  choiceId?: bigint

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
