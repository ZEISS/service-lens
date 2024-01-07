'use client'

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import DateFormat from '@/components/date-format'
import { Workload } from '@/db/models/workload'

export type PropertiesCard = {
  workload?: Workload
}

export function PropertiesCard({ workload }: PropertiesCard) {
  return (
    <Card className="my-4">
      <CardHeader className="space-y-1">
        <CardTitle className="text-2xl">Properties</CardTitle>
      </CardHeader>
      <CardContent className="grid gap-4">
        <div className="flex justify-between">
          <div className="w-2/4">
            <div className="space-y-1 py-2">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                Name
              </h2>
              <p>{workload?.name}</p>
            </div>
            <div className="space-y-1 py-2">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                UUID
              </h2>
              <p>{workload?.id}</p>
            </div>
            <div className="space-y-1 py-2">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                Created
              </h2>
              <DateFormat date={workload?.dataValues?.createdAt} />
            </div>
            <div className="space-y-1 py-2">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                Last updated
              </h2>
              <DateFormat date={workload?.dataValues?.updatedAt} />
            </div>
          </div>
          <div className="w-2/4">
            <div className="space-y-1">
              <h2 className="text-l font-semibold tracking-tight text-muted-foreground">
                Environment
              </h2>
              <p>
                {workload?.environments?.map(env => (
                  <Button
                    key={env.id}
                    variant="outline"
                    size="sm"
                    className="h-8 border-dashed mr-2"
                    disabled={true}
                  >
                    {env.name}
                  </Button>
                ))}
              </p>
            </div>
          </div>
        </div>
        <Separator />
        <p>{workload?.description || 'No description provided.'}</p>
      </CardContent>
      <CardFooter></CardFooter>
    </Card>
  )
}
