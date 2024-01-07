import { Environment } from '@/db/models/environment'
import { v4 as uuidv4 } from 'uuid'
import type { EnvironmentCreationAttributes } from '../models/environment'
import { FindAndCountEnvironmentsSchema } from '../schemas/environments'
import { z } from 'zod'

export type Pagination = {
  offset?: number
  limit?: number
}

export async function createEnvironment({
  name,
  description
}: EnvironmentCreationAttributes) {
  const w = new Environment({ label: '', name, description })

  await w.validate()

  const workload = await w.save()

  return workload.dataValues
}

export const deleteEnvironment = async (id: string) =>
  await Environment.destroy({ where: { id } })

export const findAndCountEnvironments = async (
  opts: z.infer<typeof FindAndCountEnvironmentsSchema>
) =>
  await Environment.findAndCountAll({
    order: [['name', 'DESC']],
    ...opts
  })
