import { isAllowed, protectedProcedure } from '../../trpc'
import {
  findAndCountSolutions,
  getSolution as gs,
  deleteSolutionComment as scd,
  findAndCountSolutionTemplates,
  findOneSolutionTemplate,
  countSolutions,
  destroySolutionTemplate,
  makeCopySolution,
  listSolutionByTeamSlug,
  destroySolution
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
  SolutionMakeCopySchema,
  ListSolutionByTeamSlug,
  DestroySolutionSchema
} from '../schemas/solution'
import { router } from '@/server/trpc'
import { revalidatePath } from 'next/cache'

export const listSolutions = protectedProcedure
  .input(SolutionListSchema)
  .query(async opts => findAndCountSolutions(opts.input))

export const getSolution = protectedProcedure
  .input(SolutionGetSchema)
  .query(async opts => await gs(opts.input))

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
  .query(async opts => {
    revalidatePath('/dashboard/solutions')
    return await makeCopySolution(opts.input)
  })

export const listByTeam = protectedProcedure
  .input(ListSolutionByTeamSlug)
  .query(async opts => await listSolutionByTeamSlug({ ...opts.input }))

export const deleteSolution = protectedProcedure
  .use(isAllowed('write'))
  .input(DestroySolutionSchema)
  .query(async opts => destroySolution(opts.input))

export const solutionsRouter = router({
  makeCopy: makeCopySolutionTemplate,
  delete: deleteSolution,
  deleteComment: deleteSolutionComment,
  deleteSolution: protectedProcedure
    .input(SolutionDeleteSchema)
    .query(async opts => await destroySolutionTemplate(opts.input)),
  get: getSolution,
  getSolutionTemplate,
  list: listSolutions,
  listSolutionTemplates: findSolutionTemplates,
  listByTeam: listByTeam,
  total: totalSolutions
})
