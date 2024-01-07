import { protectedProcedure } from '../../trpc'
import { v4 as uuidv4 } from 'uuid'
import {
  findAndCountSolutions,
  addSolution as as,
  getSolution as gs,
  deleteSolutionComment as scd,
  findAndCountSolutionTemplates,
  findOneSolutionTemplate,
  countSolutions,
  destroySolutionTemplate,
  makeCopySolution
} from '@/db/services/solutions'
import {
  SolutionListSchema,
  SolutionAddSchema,
  SolutionGetSchema,
  SolutionCommentDeleteSchema,
  SolutionTemplateListSchema,
  SolutionTemplateGetSchema,
  SolutionDeleteSchema,
  SolutionTemplateDeleteSchema,
  SolutionMakeCopySchema
} from '../schemas/solution'
import { router } from '@/server/trpc'

export const listSolutions = protectedProcedure
  .input(SolutionListSchema)
  .query(async opts => findAndCountSolutions(opts.input))

export const addSolution = protectedProcedure
  .input(SolutionAddSchema)
  .query(async opts => as({ id: uuidv4(), ...opts.input }))

export const getSolution = protectedProcedure
  .input(SolutionGetSchema)
  .query(async opts => {
    const s = await gs(opts.input)
    console.log(s?.comments)

    return s
  })

export const deleteSolutionComment = protectedProcedure
  .input(SolutionCommentDeleteSchema)
  .query(async opts => scd(opts.input))

export const findSolutionTemplates = protectedProcedure
  .input(SolutionTemplateListSchema)
  .query(async opts => findAndCountSolutionTemplates(opts.input))

export const getSolutionTemplate = protectedProcedure
  .input(SolutionTemplateGetSchema)
  .query(async opts => findOneSolutionTemplate(opts.input))

export const totalSolutions = protectedProcedure.query(
  async _ => await countSolutions()
)

export const deleteSolutionTemplate = protectedProcedure
  .input(SolutionTemplateDeleteSchema)
  .query(async opts => await destroySolutionTemplate(opts.input))

export const makeCopySolutionTemplate = protectedProcedure
  .input(SolutionMakeCopySchema)
  .query(async opts => await makeCopySolution(opts.input))

export const solutionsRouter = router({
  add: addSolution,
  makeCopy: makeCopySolutionTemplate,
  delete: deleteSolutionComment,
  deleteSolution: protectedProcedure
    .input(SolutionDeleteSchema)
    .query(async opts => await destroySolutionTemplate(opts.input)),
  get: getSolution,
  getSolutionTemplate,
  list: listSolutions,
  listSolutionTemplates: findSolutionTemplates,
  total: totalSolutions
})
