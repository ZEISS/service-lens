'use server'

import 'server-only'
import { cookies } from 'next/headers'
import { z } from 'zod'
import { rhfActionSetScopeSchema } from './teams-switcher.schema'
import { redirect } from 'next/navigation'
import { revalidatePath } from 'next/cache'

export async function rhfActionSetScope(
  opts: z.infer<typeof rhfActionSetScopeSchema>
) {
  cookies().set('scope', opts)
  revalidatePath(opts !== 'personal' ? `/teams/${opts}` : `/home`)
  redirect(opts !== 'personal' ? `/teams/${opts}` : `/home`)
}
