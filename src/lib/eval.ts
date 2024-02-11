import type { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import type { QuestionRef, QuestionId, RiskCondition } from './stats'
import type { WorkloadLensAnswer } from '@/db/models/workload-lenses-answers'

export type EvalContext = { [key: QuestionRef]: boolean }
export function evalInScope(
  js: RiskCondition = 'true',
  contextAsScope: EvalContext
): boolean {
  return new Function(`with (this) { return (${js}); }`).call(contextAsScope)
}
export const createContext = (choices: LensPillarChoice[] = []): EvalContext =>
  choices.reduce((ctx, choice) => ({ ...ctx, [choice.ref]: false }), {})
export const answerContext = (
  questionId: QuestionId,
  answers: WorkloadLensAnswer[] = []
) =>
  answers
    ?.filter(answer => answer.lensPillarQuestionId === questionId)
    ?.flatMap(answer => [...(answer.lensChoices ?? [])])
    ?.reduce((choices, choice) => ({ ...choices, [choice.ref]: true }), {})
