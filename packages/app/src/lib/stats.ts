import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { WorkloadLensAnswer } from '@/db/models/workload-lenses-answers'

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
