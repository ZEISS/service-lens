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
  HasMany,
  ForeignKey,
  BelongsTo,
  Default,
  AllowNull,
  BelongsToMany
} from 'sequelize-typescript'
import { SolutionComment } from './solution-comments'
import { User } from './users'
import { Team } from './teams'
import { TagTaggable } from './tags-taggable'
import { Tag } from './tags'
import { Ownership } from './ownership'
import { Optional } from 'sequelize'

export interface SolutionAttributes {
  id: string
  title: string
  body: string
  user: User
  userId: string
  description?: string
  comments?: SolutionComment[]
  tags?: Tag[]
  teams?: Team[]
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type SolutionCreationAttributes = Optional<
  SolutionAttributes,
  'id' | 'user' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'solutions',
  modelName: 'Solution'
})
export class Solution extends Model<
  SolutionAttributes,
  SolutionCreationAttributes
> {
  @PrimaryKey
  @AllowNull(false)
  @Default(DataType.UUIDV4)
  @Column(DataType.UUIDV4)
  declare id: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  declare title: string

  @NotEmpty
  @Column(DataType.TEXT)
  declare body: string

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column(DataType.STRING)
  declare description?: string

  @HasMany(() => SolutionComment, 'solutionId')
  declare comments?: SolutionComment[]

  @ForeignKey(() => User)
  @Column
  declare userId: string

  @BelongsTo(() => User)
  declare user: User

  @BelongsToMany(() => Tag, {
    through: {
      model: () => TagTaggable,
      unique: false,
      scope: {
        taggableType: 'solution'
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
        resourceType: 'solution'
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
