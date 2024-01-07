import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/components/ui/table'
import type { Lens } from '@/db/models/lens'
import Link from 'next/link'
import React from 'react'

export type LensesCardProps = {
  workloadId?: string
  lenses?: Lens[]
}

export function LensesCard({
  workloadId,
  lenses,
  ...props
}: LensesCardProps & React.HTMLAttributes<HTMLDivElement>) {
  return (
    <>
      <Card {...props}>
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl">Lenses</CardTitle>
        </CardHeader>
        <CardContent>
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Description</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {lenses?.map((lens, idx) => (
                <TableRow key={idx}>
                  <TableCell className="font-medium">
                    <Link
                      href={`/dashboard/workloads/${workloadId}/lenses/${lens.id}`}
                    >
                      {lens.name}
                    </Link>
                  </TableCell>
                  <TableCell>{lens.description}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </>
  )
}
