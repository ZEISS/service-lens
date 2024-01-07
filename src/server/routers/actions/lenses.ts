import { protectedProcedure } from '../../trpc'
import {
  LensDeleteSchema,
  LensGetSchema,
  LensGetQuestionSchema,
  LensListSchema
} from '../schemas/lens'
import { deleteLens as dl } from '@/db/services/lenses'
import { getLens as gl } from '@/db/services/lenses'
import {
  findOneLensPillarQuestion,
  findAndCountLenses
} from '@/db/services/lenses'
import { router } from '@/server/trpc'

export const deleteLens = protectedProcedure
  .input(LensDeleteSchema)
  .query(async opts => await dl(opts.input))

export const getLens = protectedProcedure
  .input(LensGetSchema)
  .query(async opts => await gl(opts.input))

export const getLensQuestion = protectedProcedure
  .input(LensGetQuestionSchema)
  .query(async opts => await findOneLensPillarQuestion(opts.input))

export const listLenses = protectedProcedure
  .input(LensListSchema)
  .query(async opts => await findAndCountLenses({ ...opts.input }))

export const lensRouter = router({
  delete: deleteLens,
  get: getLens,
  getQuestion: getLensQuestion,
  list: listLenses
})
