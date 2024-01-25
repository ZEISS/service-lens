import { z } from 'zod'
import { publicProcedure, router } from '../trpc'
import { listLenses } from './actions/lenses'
import {
  getWorkload,
  getWorkloadAnswer,
  listWorkloads,
  findWorkloadLensQuestion,
  totalWorkloads,
  workloadsRouter
} from './actions/workloads'
import { getProfile, listProfilesQuestions } from './actions/profiles'
import { getLens, getLensQuestion } from './actions/lenses'
import { listProfiles } from './actions/profiles'
import { listEnvironments } from './actions/environments'
import {
  listSolutions,
  getSolution,
  findSolutionTemplates,
  getSolutionTemplate,
  totalSolutions,
  deleteSolutionTemplate
} from './actions/solutions'
import { checkPermission } from './actions/permissions'
import { lensRouter } from '@/server/routers/actions/lenses'
import { solutionsRouter } from './actions/solutions'
import { profilesRouter } from './actions/profiles'
import { teamsRouter } from './actions/teams'
import { usersRouter } from './actions/users'

export const appRouter = router({
  greeting: publicProcedure
    .input(
      z.object({
        text: z.string()
      })
    )
    .query(async opts => {
      console.log('request from', opts.ctx.headers?.['x-trpc-source'])
      return `hello ${opts.input.text} - ${Math.random()}`
    }),

  checkPermission: checkPermission,

  secret: publicProcedure.query(async opts => {
    if (!opts.ctx.session) {
      return 'You are not authenticated'
    }

    return "Cool, you're authenticated!"
  }),

  me: publicProcedure.query(opts => opts.ctx.session),

  deleteSolutionTemplate,
  findSolutionTemplates,
  findWorkloadLensQuestion,
  getLens,
  getLensQuestion,
  getProfile,
  getSolution,
  getSolutionTemplate,
  getWorkload,
  getWorkloadAnswer,
  listEnvironments,
  listLenses,
  listProfiles,
  listSolutions,
  listWorkloads,
  totalSolutions,
  totalWorkloads,
  listProfilesQuestions,

  // lenses
  lenses: lensRouter,

  // solutions
  solutions: solutionsRouter,

  // profiles
  profiles: profilesRouter,

  // workloads
  workloads: workloadsRouter,

  // teams
  teams: teamsRouter,

  // users
  users: usersRouter
})

export type AppRouter = typeof appRouter
