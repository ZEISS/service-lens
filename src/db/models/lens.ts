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
  Default
} from 'sequelize-typescript'
import { LensPillar } from './lens-pillars'
import type { Spec } from '../schemas/spec'

export interface LensAttributes {
  id: string
  version: number
  spec: object
  name: string
  isDraft: boolean
  description?: string
  pillars?: LensPillar[]
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type LensCreationAttributes = Omit<
  LensAttributes,
  'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'lenses',
  modelName: 'Lens'
})
export class Lens extends Model<LensAttributes, LensCreationAttributes> {
  @PrimaryKey
  @Column(DataType.UUIDV4)
  id!: string

  @NotEmpty
  @Min(3)
  @Max(256)
  @Column
  name!: string

  @NotEmpty
  @Column
  version!: string

  @NotEmpty
  @Column(DataType.JSONB)
  spec!: Spec

  @Default(true)
  @Column
  isDraft!: boolean

  @NotEmpty
  @Min(12)
  @Max(2048)
  @Column
  description?: string

  @HasMany(() => LensPillar, 'lensId')
  pillars?: LensPillar[]

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
