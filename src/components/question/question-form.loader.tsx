import { api } from '@/trpc/server-invoker'
import { QuestionFormFactory } from './question-form'

export type QuestionFormLoaderProps = {
  questionId: string
  id: string
}

export default async function QuestionFormLoader({
  questionId,
  id
}: QuestionFormLoaderProps) {
  const question = await api.getLensQuestion.query(questionId)
  const answer = await api.getWorkloadAnswer.query({
    workloadId: id,
    lensPillarQuestionId: questionId
  })

  return (
    <QuestionFormFactory
      workloadId={id}
      lensPillarQuestionId={questionId}
      answer={answer}
      question={question}
    />
  )
}
