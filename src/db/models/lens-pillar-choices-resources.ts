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
  AllowNull
} from 'sequelize-typescript'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'

export interface LensPillarChoiceResourceAttributes {
  id: string
  questionId: bigint
  description: string
  url?: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensPillarChoiceResourceCreationAttributes = Omit<
  LensPillarChoiceResourceAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses-pillars-choices-resources'
})
export class LensPillarChoiceResource extends Model<
  LensPillarChoiceResourceAttributes,
  LensPillarChoiceResourceCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column(DataType.INTEGER)
  id!: string

  @ForeignKey(() => LensPillarChoice)
  @Column
  choiceId?: bigint

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
