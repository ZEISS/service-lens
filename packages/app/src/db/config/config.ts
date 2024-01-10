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
import { WorkloadLensesAnswer } from '../models/workload-lenses-answers'
import { WorkloadLensesAnswerChoice } from '../models/workload-lenses-answers-choices'

const env = process.env.NODE_ENV || 'development'
const isProduction = env === 'production'

const models = [
  Environment,
  Lens,
  LensPillar,
  LensPillarChoice,
  LensPillarQuestion,
  LensPillarQuestionRisk,
  Profile,
  ProfileQuestion,
  ProfileQuestionChoice,
  Solution,
  Workload,
  WorkloadEnvironment,
  WorkloadLens,
  WorkloadLensesAnswer,
  WorkloadLensesAnswerChoice,
  SolutionComment,
  SolutionTemplate,
  LensPillarResource,
  LensPillarChoiceResource,
  LensPillarQuestionResource,
  ProfileQuestionAnswer,
  User
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
