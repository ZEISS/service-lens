import { Team } from '@/db/models/teams'
import { User } from '@/db/models/users'
import type { FindOneUserSchema } from '../schemas/users'
import { Workload } from '@/db/models/workload'

export const findOneUser = async (opts: FindOneUserSchema) =>
  await User.findOne({
    where: { id: opts },
    include: [Team, Workload]
  })
