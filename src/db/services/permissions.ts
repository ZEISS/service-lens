import type { FindOnePermissionSchema } from '../schemas/permissions'
import { UserPermission } from '../models/users-permissions'

export const findOnePermission = async (opts: FindOnePermissionSchema) =>
  await UserPermission.count({ where: { ...opts } })
