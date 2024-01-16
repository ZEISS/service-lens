import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import type { WorkloadLensAnswer } from '@/db/models/workload-lenses-answers'

export type RiskCondition = string
export enum Risk {
  HIGH,
  MEDIUM,
  LOW,
  NO,
  UNKNOWN
}

export type Statistics = {
  totalQuestions?: number
  totalHighRisk?: number
  totalMediumRisk?: number
}

export const generateStats = (
  question: LensPillarQuestion,
  answer: WorkloadLensAnswer
): Statistics => {
  const stats = {}

  return stats
}

export type QuestionRef = string
export type QuestionId = bigint
export const groupAnswers = (answers: WorkloadLensAnswer[] = []) =>
  answers.reduce(
    (group, answer) =>
      group.set(
        answer.lensPillarQuestionId,
        answer.lensChoices?.flatMap(choice => choice.ref)
      ),
    new Map<QuestionId | undefined, QuestionRef[] | undefined>()
  )
