'use client'

import { useMemo, use } from 'react'
import { api } from '@/trpc/client'

export function EnvironmentsForm() {
  // const pagination = useDataTableContext()
  const environments = use(
    api.listEnvironments.query({
      limit: 10,
      offset: 0
    })
  )

  return (
    <>
      {environments.rows?.map(environment => (
        <div key={environment.id}>{environment.name}</div>
      ))}
    </>
  )
}
