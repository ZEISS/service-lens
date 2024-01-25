import { protectedProcedure } from '../../trpc'
import {
  LensGetSchema,
  LensGetQuestionSchema,
  LensListSchema,
  ListLensByTeamSlug
} from '../schemas/lens'
import { deleteLens as dl, listLensByTeamSlug } from '@/db/services/lenses'
import { getLens as gl } from '@/db/services/lenses'
import {
  findOneLensPillarQuestion,
  findAndCountLenses
} from '@/db/services/lenses'
import { router } from '@/server/trpc'

export const getLens = protectedProcedure
  .input(LensGetSchema)
  .query(async opts => await gl(opts.input))

export const getLensQuestion = protectedProcedure
  .input(LensGetQuestionSchema)
  .query(async opts => await findOneLensPillarQuestion(opts.input))

export const listLenses = protectedProcedure
  .input(LensListSchema)
  .query(async opts => await findAndCountLenses({ ...opts.input }))

export const listByTeam = protectedProcedure
  .input(ListLensByTeamSlug)
  .query(async opts => await listLensByTeamSlug({ ...opts.input }))

export const lensRouter = router({
  get: getLens,
  getQuestion: getLensQuestion,
  list: listLenses,
  listByTeam: listByTeam
})
