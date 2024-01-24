import { SolutionComment } from '@/db/models/solution-comments'
import { SolutionTemplate } from '@/db/models/solution-templates'
import { User } from '@/db/models/users'
import { Solution } from '@/db/models/solution'
import { type SolutionCreate, type DestroySolution } from '../schemas/solutions'
import sequelize from '../config/config'
import {
  SolutionCommentAddSchema,
  SolutionCommentDeleteSchema,
  SolutionsGetSchema,
  FindAndCountSolutionsSchema,
  FindAndCountSolutionTemplates,
  FindOneSolutionTemplate,
  DestroySolutionTemplateSchema,
  MakeCopySolutionTemplateSchema,
  MakeCopySolutionSchema,
  ListSolutionByTeamSlug
} from '../schemas/solutions'
import { z } from 'zod'
import { Team } from '../models/teams'
import { Ownership } from '../models/ownership'

export const countSolutions = async () => await Solution.count()

export const createSolution = async (opts: SolutionCreate) =>
  await sequelize.transaction(async transaction => {
    const workload = await Solution.create({ ...opts }, { transaction })

    const _ = await Ownership.create(
      {
        ownerId: opts.ownerId,
        resourceId: workload.id,
        resourceType: 'solution'
      },
      { transaction }
    )

    return workload
  })

export const findAndCountSolutions = async (
  opts: z.infer<typeof FindAndCountSolutionsSchema>
) => await Solution.findAndCountAll({ offset: opts.offset, limit: opts.limit })

export const destroySolution = async (opts: DestroySolution) =>
  sequelize.transaction(
    async transaction => await Solution.destroy({ where: { id: opts } })
  )

export const makeCopySolution = async (
  opts: z.infer<typeof MakeCopySolutionSchema>
) =>
  sequelize.transaction(async transaction => {
    const solution = await Solution.findOne({ where: { id: opts } })

    return Solution.create(
      {
        title: `${solution?.title} (Copy)`,
        body: solution?.body ?? '',
        description: solution?.description
      },
      { transaction }
    )
  })

export const getSolution = async (opts: z.infer<typeof SolutionsGetSchema>) =>
  await Solution.findOne({
    where: { id: opts },
    include: [
      { model: User },
      { model: SolutionComment, include: [{ model: User }] }
    ]
  })

export const addSolutionComment = async (
  opts: z.infer<typeof SolutionCommentAddSchema>
) => (await (await SolutionComment.create({ ...opts })).save()).dataValues

export const deleteSolutionComment = async (
  opts: z.infer<typeof SolutionCommentDeleteSchema>
) => SolutionComment.destroy({ where: { id: opts } })

export const findAndCountSolutionTemplates = async (
  opts: z.infer<typeof FindAndCountSolutionTemplates>
) =>
  await SolutionTemplate.findAndCountAll({
    offset: opts.offset,
    limit: opts.limit
  })

export const destroySolutionTemplate = async (
  opts: z.infer<typeof DestroySolutionTemplateSchema>
) => await SolutionTemplate.destroy({ where: { id: opts } })

export const findOneSolutionTemplate = async (
  opts: z.infer<typeof FindOneSolutionTemplate>
) =>
  await SolutionTemplate.findOne({
    where: { id: opts }
  })

export const makeCopySolutionTemplate = async (
  opts: z.infer<typeof MakeCopySolutionTemplateSchema>
) =>
  sequelize.transaction(async transaction => {
    const tmpl = await SolutionTemplate.findOne({ where: { id: opts } })

    return SolutionTemplate.create(
      {
        title: `${tmpl?.title} (Copy)`,
        body: tmpl?.body,
        description: tmpl?.description
      },
      { transaction }
    )
  })

export const listSolutionByTeamSlug = async (opts: ListSolutionByTeamSlug) =>
  await Solution.findAndCountAll({
    offset: opts.offset,
    limit: opts.limit,
    include: [{ model: Team, where: { slug: opts.slug } }]
  })
