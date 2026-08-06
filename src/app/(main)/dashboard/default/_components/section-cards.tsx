import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { getTotalNumberOfDesigns } from "@/db/queries/designs"
import { getTotalNumberOfEnvironments } from "@/db/queries/environments"
import { getTotalNumberOfLenses } from "@/db/queries/lenses"
import { getTotalNumberOfWorkloads } from "@/db/queries/workloads"

export async function SectionCards() {
  const [totalWorkloads, totalDesigns, totalLenses, totalEnvironments] = await Promise.all([
    getTotalNumberOfWorkloads(),
    getTotalNumberOfDesigns(),
    getTotalNumberOfLenses(),
    getTotalNumberOfEnvironments(),
  ])

  return (
    <div className="grid @5xl/main:grid-cols-4 @xl/main:grid-cols-2 grid-cols-1 gap-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:from-primary/5 *:data-[slot=card]:to-card *:data-[slot=card]:shadow-xs dark:*:data-[slot=card]:bg-card">
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Total Workloads</CardDescription>
          <CardTitle className="font-semibold @[250px]/card:text-3xl text-2xl tabular-nums">{totalWorkloads}</CardTitle>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="text-muted-foreground">Active Workloads</div>
        </CardFooter>
      </Card>
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Total Designs</CardDescription>
          <CardTitle className="font-semibold @[250px]/card:text-3xl text-2xl tabular-nums">{totalDesigns}</CardTitle>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="text-muted-foreground">Active Designs</div>
        </CardFooter>
      </Card>
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Total Lenses</CardDescription>
          <CardTitle className="font-semibold @[250px]/card:text-3xl text-2xl tabular-nums">{totalLenses}</CardTitle>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="text-muted-foreground">Active Lenses</div>
        </CardFooter>
      </Card>
      <Card className="@container/card">
        <CardHeader>
          <CardDescription>Total Environments</CardDescription>
          <CardTitle className="font-semibold @[250px]/card:text-3xl text-2xl tabular-nums">
            {totalEnvironments}
          </CardTitle>
        </CardHeader>
        <CardFooter className="flex-col items-start gap-1.5 text-sm">
          <div className="text-muted-foreground">Active Environments</div>
        </CardFooter>
      </Card>
    </div>
  )
}
