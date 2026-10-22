import Form from "next/form"
import Link from "next/link"

import { Ellipsis } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardAction, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu"
import type { TLens } from "@/db/schema"

import { removeLensAction } from "./lens-card.action"

// LensCardProps interface for the LensCard component
export interface LensCardProps {
  workloadId: string
  lens: TLens
}

// LensCard component that displays a single lens card
export function LensCard({ lens, workloadId }: LensCardProps) {
  return (
    <Card key={lens.id}>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">{lens.name}</CardTitle>
        <CardAction>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost">
                <Ellipsis />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem variant="destructive">
                <Form action={removeLensAction}>
                  <input type="hidden" id="lensId" name="lensId" value={lens.id} />
                  <input type="hidden" id="workloadId" name="workloadId" value={workloadId} />
                  <button type="submit">Remove</button>
                </Form>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </CardAction>
      </CardHeader>
      <CardContent>{lens.description}</CardContent>
      <CardFooter>
        <Button variant="outline" asChild>
          <Link href={`/workloads/${workloadId}/lenses/${lens.id}`}>Review</Link>
        </Button>
      </CardFooter>
    </Card>
  )
}
