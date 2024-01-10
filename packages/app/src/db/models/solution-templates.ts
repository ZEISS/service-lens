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
  AutoIncrement,
  Min,
  Max
} from 'sequelize-typescript'

export interface SolutionTemplateAttributes {
  id: bigint
  title: string
  body?: string
  description?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type SolutionTemplateCreationAttributes = Omit<
  SolutionTemplateAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'solutions-templates'
})
export class SolutionTemplate extends Model<
  SolutionTemplateAttributes,
  SolutionTemplateCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.BIGINT)
  id!: bigint

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  title?: string

  @NotEmpty
  @Column(DataType.TEXT)
  body?: string

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
