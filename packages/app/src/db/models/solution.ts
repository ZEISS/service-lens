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

export interface SolutionAttributes {
  id: string
  title: string
  body: string
  user?: User
  userId?: string
  description?: string
  comments?: SolutionComment[]
  tags: Tag[]
  teams: Team[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type SolutionCreationAttributes = Omit<
  SolutionAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
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
  id?: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  title?: string

  @NotEmpty
  @Column(DataType.TEXT)
  body?: string

  @HasMany(() => SolutionComment, 'solutionId')
  comments?: SolutionComment[]

  @ForeignKey(() => User)
  @Column
  userId?: string

  @BelongsTo(() => User)
  user?: User

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
