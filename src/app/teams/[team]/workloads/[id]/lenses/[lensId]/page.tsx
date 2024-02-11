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
import { questionRisk } from '@/db/models/workload-lenses-answers'
import DateFormat from '@/components/date-format'
import { Separator } from '@/components/ui/separator'

export type PageProps = {
  params: { lensId: string; id: string }
}

export default async function Page({ params }: PageProps) {
  const lens = await api.getLens.query(params?.lensId)
  const workload = await api.workloads.get.query(params?.id)
  // const stats: { [key: string]: number } =
  //   workload?.answers?.reduce(
  //     (answers, answer) => ({
  //       ...answers
  //       // [answer.risk.toString()]: answers[answer.risk] + 1
  //     }),
  //     Object.values(QuestionRisk).reduce(
  //       (risks, risk) => ({ ...risks, [risk]: 0 }),
  //       {} as { [key: string]: number }
  //     )
  //   ) ?? {}

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
          <CardTitle>Risks</CardTitle>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Risk</TableHead>
                <TableHead>Number</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {/* {Object?.entries(stats ?? {}).map(([k, v], n) => (
                <TableRow key={n}>
                  <TableCell className="font-medium">{k}</TableCell>
                  <TableCell>{v}</TableCell>
                </TableRow>
              ))} */}
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
