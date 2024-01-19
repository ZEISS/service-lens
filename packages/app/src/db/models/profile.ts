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
  AllowNull,
  Default,
  BelongsToMany
} from 'sequelize-typescript'
import { ProfileQuestionAnswer } from '@/db/models/profile-question-answers'
import { ProfileQuestionChoice } from '@/db/models/profile-question-choice'
import { Tag } from './tags'
import { TagTaggable } from './tags-taggable'
import { Team } from './teams'
import { Ownership } from './ownership'

export interface ProfileAttributes {
  id: string
  name: string
  description?: string
  tags: Tag[]
  teams: Team[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type ProfileCreationAttributes = Omit<
  ProfileAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
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
  @Default(DataType.UUIDV4)
  @AllowNull(false)
  @Column(DataType.UUIDV4)
  declare id: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column(DataType.STRING)
  declare name: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column(DataType.STRING)
  declare description: string

  @BelongsToMany(
    () => ProfileQuestionChoice,
    () => ProfileQuestionAnswer,
    'profileId',
    'choiceId'
  )
  answers?: ProfileQuestionChoice[]

  @BelongsToMany(() => Tag, {
    through: {
      model: () => TagTaggable,
      unique: false,
      scope: {
        taggableType: 'profile'
      }
    },
    otherKey: 'tagId',
    foreignKey: 'taggableId',
    constraints: false
  })
  declare tags: Tag[]

  @BelongsToMany(() => Team, {
    through: {
      model: () => Ownership,
      unique: false,
      scope: {
        resourceType: 'profile'
      }
    },
    foreignKey: 'resourceId',
    otherKey: 'ownerId',
    constraints: false
  })
  declare teams: Team[]

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
