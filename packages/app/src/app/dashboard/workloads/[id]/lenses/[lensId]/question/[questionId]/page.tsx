import { api } from '@/trpc/server-http'
import { QuestionFormFactory } from '@/components/question/question-form'

export type PageProps = {
  params: { questionId: string; lensId: string; id: string }
}

export default async function Page({ params }: PageProps) {
  const question = await api.getLensQuestion.query(params.questionId)
  const answer = await api.getWorkloadAnswer.query({
    workloadId: params.id,
    lensPillarQuestionId: params.questionId
  })

  return (
    <>
      {question && (
        <QuestionFormFactory
          workloadId={params.id}
          lensPillarQuestionId={params.questionId}
          question={question}
          answer={answer}
        />
      )}
    </>
  )
}
