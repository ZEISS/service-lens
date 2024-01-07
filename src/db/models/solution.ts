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
  BelongsTo
} from 'sequelize-typescript'
import { SolutionComment } from './solution-comments'
import { User } from './users'

export interface SolutionAttributes {
  id: string
  title: string
  body: string
  user?: User
  userId?: string
  description?: string
  comments?: SolutionComment[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type SolutionCreationAttributes = Omit<
  SolutionAttributes,
  'createdAt' | 'updatedAt' | 'deletedAt'
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
