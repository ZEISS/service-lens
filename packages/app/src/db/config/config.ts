import { Environment } from '../models/environment'
import { Lens } from '../models/lens'
import { LensPillar } from '../models/lens-pillars'
import { LensPillarChoice } from '../models/lens-pillar-choices'
import { LensPillarChoiceResource } from '../models/lens-pillar-choices-resources'
import { LensPillarQuestion } from '../models/lens-pillar-questions'
import { LensPillarQuestionResource } from '../models/lens-pillar-questions-resources'
import { LensPillarQuestionRisk } from '../models/lens-pillar-risks'
import { LensPillarResource } from '../models/lens-pillar-resources'
import { Profile } from '../models/profile'
import { ProfileQuestion } from '../models/profile-question'
import { ProfileQuestionAnswer } from '../models/profile-question-answers'
import { ProfileQuestionChoice } from '../models/profile-question-choice'
import { Sequelize, SequelizeOptions } from 'sequelize-typescript'
import { Solution } from '../models/solution'
import { SolutionComment } from '../models/solution-comments'
import { SolutionTemplate } from '../models/solution-templates'
import { User } from '../models/users'
import { Workload } from '../models/workload'
import { WorkloadEnvironment } from '../models/workload-environment'
import { WorkloadLens } from '../models/workload-lens'
import { WorkloadLensAnswer } from '../models/workload-lenses-answers'
import { WorkloadLensesAnswerChoice } from '../models/workload-lenses-answers-choices'
import { Team } from '../models/teams'
import { Role } from '../models/roles'
import { Permission } from '../models/permissions'
import { RolePermission } from '../models/roles-permissions'
import { UserRole } from '../models/users-roles'
import { UserTeam } from '../models/users-teams'
import { UserPermission } from '../models/users-permissions'
import { Tag } from '@/db/models/tags'
import { TagTaggable } from '@/db/models/tags-taggable'
import { Ownership } from '@/db/models/ownership'

const env = process.env.NODE_ENV || 'development'
const isProduction = env === 'production'

const models = [
  Environment,
  Lens,
  LensPillar,
  LensPillarChoice,
  LensPillarChoiceResource,
  LensPillarQuestion,
  LensPillarQuestionResource,
  LensPillarQuestionRisk,
  LensPillarResource,
  Ownership,
  Permission,
  Profile,
  ProfileQuestion,
  ProfileQuestionAnswer,
  ProfileQuestionChoice,
  Role,
  RolePermission,
  Solution,
  SolutionComment,
  SolutionTemplate,
  Tag,
  TagTaggable,
  Team,
  User,
  UserPermission,
  UserRole,
  UserTeam,
  Workload,
  WorkloadEnvironment,
  WorkloadLens,
  WorkloadLensAnswer,
  WorkloadLensesAnswerChoice
]

export interface Config {
  [index: string]: SequelizeOptions
}

export const config: Config = {
  development: {
    username: process.env.DB_USER,
    password: process.env.DB_PASS,
    database: process.env.DB_NAME,
    host: process.env.DB_HOST,
    dialectModule: require('pg'),
    dialect: 'postgres',
    models
  },
  test: {
    username: process.env.DB_USER,
    password: process.env.DB_PASS,
    database: process.env.DB_NAME,
    host: process.env.DB_HOST,
    dialectModule: require('pg'),
    dialect: 'postgres',
    logging: false,
    models
  },
  production: {
    username: process.env.DB_USER,
    password: process.env.DB_PASS,
    database: process.env.DB_NAME,
    host: process.env.DB_HOST,
    dialect: 'postgres',
    dialectModule: require('pg'),
    dialectOptions: {
      ssl: false
    },
    models
  }
}

const connection = new Sequelize({ ...config[env] })

export const initDB = async () => {
  await connection.authenticate()
  !isProduction && (await connection.sync({ alter: true }))
}

export { Sequelize }
export default connection
