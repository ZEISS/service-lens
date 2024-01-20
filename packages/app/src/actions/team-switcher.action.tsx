'use server'

import 'server-only'
import { cookies } from 'next/headers'
import { z } from 'zod'
import { rhfActionSetScopeSchema } from './teams-switcher.schema'
import { redirect } from 'next/navigation'

export async function rhfActionSetScope(
  opts: z.infer<typeof rhfActionSetScopeSchema>
) {
  cookies().set('scope', opts)
  redirect(`/teams/${opts}`)
}
