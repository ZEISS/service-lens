import { Optional } from 'sequelize'
import {
  Model,
  Table,
  DataType,
  CreatedAt,
  UpdatedAt,
  DeletedAt,
  Column,
  Unique
} from 'sequelize-typescript'

export type ResourceType = 'workload' | 'solution' | 'lens' | 'profile'

export type OwnershipAttributes = {
  id: bigint
  resourceId: string
  resourceType: ResourceType
  ownerId: string
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
}

export type OwnershipCreationAttributes = Optional<OwnershipAttributes, 'id'>

@Table({ modelName: 'Ownership', tableName: 'ownership' })
export class Ownership extends Model<
  OwnershipAttributes,
  OwnershipCreationAttributes
> {
  @Unique('ownership_unique_constraint')
  @Column(DataType.UUID)
  declare resourceId: string

  @Unique('ownership_unique_constraint')
  @Column(DataType.UUID)
  declare ownerId: string

  @Unique('ownership_unique_constraint')
  @Column(DataType.STRING)
  declare resourceType: ResourceType

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
