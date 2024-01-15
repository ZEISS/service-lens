import { api } from '@/trpc/server-http'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
  CardFooter
} from '@/components/ui/card'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/components/ui/table'
import DateFormat from '@/components/date-format'
import { Separator } from '@/components/ui/separator'

export type PageProps = {
  params: { lensId: string; id: string }
}

function evalInScope(js: string, contextAsScope: Object) {
  return new Function(`with (this) { return (${js}); }`).call(contextAsScope)
}

export default async function Page({ params }: PageProps) {
  const lens = await api.getLens.query(params?.lensId)
  const workload = await api.workloads.get.query(params?.id)

  const answers = workload?.answers?.reduce((prev, curr) => {
    prev.set(
      curr.lensPillarQuestionId,
      curr.lensChoices?.map(choice => choice.ref)
    )

    return prev
  }, new Map<bigint | undefined, string[] | undefined>())

  lens?.pillars?.forEach(
    pillars =>
      pillars.questions?.forEach(question => {
        const ctx = Object.fromEntries(
          question?.questionAnswers?.map(answer => [answer.ref, false]) ?? []
        )
        answers?.get(question.id)?.forEach(answer => (ctx[answer] = true))

        question.risks?.forEach(risk => {
          try {
            const truth = evalInScope(risk.condition ?? '', ctx)
            if (truth) {
              console.log(risk.risk)
            }
          } catch {
            return
          }
        })
      })
  )

  return (
    <>
      <Card>
        <CardHeader>
          <CardTitle>Overview</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex items-center justify-between">
            <div className="space-y-1">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                Version
              </h2>
              <p>{lens?.version}</p>
            </div>
            <div className="space-y-1">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                Last updated
              </h2>
              <DateFormat date={lens?.dataValues?.updatedAt} />
            </div>
          </div>
          <Separator />
          <p>{lens?.description || 'No description provided.'}</p>
        </CardContent>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Pillars</CardTitle>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Questions</TableHead>
                <TableHead>High Risk</TableHead>
                <TableHead>Medium Risk</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {lens?.pillars?.map(pillar => (
                <TableRow key={pillar.id}>
                  <TableCell className="font-medium">{pillar.name}</TableCell>
                  <TableCell>{pillar.questions?.length}</TableCell>
                  <TableCell></TableCell>
                  <TableCell className="text-right"></TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
        <CardFooter>
          <CardDescription>This is test</CardDescription>
        </CardFooter>
      </Card>
    </>
  )
}
