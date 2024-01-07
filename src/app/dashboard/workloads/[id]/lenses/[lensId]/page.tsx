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

export default async function Page({ params }: PageProps) {
  const lens = await api.getLens.query(params?.lensId)

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
              <p>{lens?.dataValues?.version}</p>
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
