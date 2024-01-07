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
  ForeignKey,
  AutoIncrement,
  BelongsTo
} from 'sequelize-typescript'
import { User } from './users'
import { Solution } from '@/db/models/solution'

export interface SolutionCommentAttributes {
  id: bigint
  body: string
  userId: string
  solutionId: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type SolutionCommentCreationAttributes = Omit<
  SolutionCommentAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'solutions-comments'
})
export class SolutionComment extends Model<
  SolutionCommentAttributes,
  SolutionCommentCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: bigint

  @NotEmpty
  @Column(DataType.TEXT)
  body?: string

  @ForeignKey(() => Solution)
  @Column(DataType.UUIDV4)
  solutionId?: string

  @ForeignKey(() => User)
  @Column(DataType.INTEGER)
  userId?: string

  @BelongsTo(() => User)
  user?: User

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
