import {
  Table,
  Model,
  CreatedAt,
  UpdatedAt,
  DeletedAt,
  Column,
  PrimaryKey,
  DataType,
  NotEmpty,
  Min,
  Max,
  BelongsToMany
} from 'sequelize-typescript'
import { ProfileQuestionAnswer } from '@/db/models/profile-question-answers'
import { ProfileQuestionChoice } from '@/db/models/profile-question-choice'

export interface ProfileAttributes {
  id: string
  name: string
  description?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type ProfileCreationAttributes = Omit<
  ProfileAttributes,
  'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'profiles',
  modelName: 'Profile'
})
export class Profile extends Model<
  ProfileAttributes,
  ProfileCreationAttributes
> {
  @PrimaryKey
  @Column(DataType.UUIDV4)
  id!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name!: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description?: string

  @BelongsToMany(
    () => ProfileQuestionChoice,
    () => ProfileQuestionAnswer,
    'profileId',
    'choiceId'
  )
  answers?: ProfileQuestionChoice[]

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
