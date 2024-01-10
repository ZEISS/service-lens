import {
  Column,
  CreatedAt,
  DataType,
  DeletedAt,
  ForeignKey,
  Model,
  PrimaryKey,
  Table,
  AutoIncrement,
  UpdatedAt
} from 'sequelize-typescript'
import { Workload } from './workload'
import { Environment } from '@/db/models/environment'

export interface WorkloadEnvironmentAttributes {
  id: number
  environmentId: bigint
  workloadId: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type WorkloadEnvironmentCreationAttributes = Omit<
  WorkloadEnvironmentAttributes,
  'id' | 'createdAt' | 'updatedAt' | 'deletedAt'
>

@Table({
  tableName: 'workloads-environment',
  modelName: 'WorkloadEnvironment'
})
export class WorkloadEnvironment extends Model<
  WorkloadEnvironmentAttributes,
  WorkloadEnvironmentCreationAttributes
> {
  @PrimaryKey
  @AutoIncrement
  @Column
  id!: number

  @ForeignKey(() => Workload)
  @Column(DataType.UUIDV4)
  workloadId?: string

  @ForeignKey(() => Environment)
  @Column
  environmentId?: bigint

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
