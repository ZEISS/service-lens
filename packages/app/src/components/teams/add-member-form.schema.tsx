import { z } from 'zod'

export const rhfActionAddMemberSchema = z.object({
  members: z
    .array(
      z.object({
        email: z.string().email(),
        type: z.union([z.literal('member'), z.literal('owner')])
      })
    )
    .min(1)
    .max(100)
})

export type AddMembersFormValues = z.infer<typeof rhfActionAddMemberSchema>
export const defaultValues: Partial<AddMembersFormValues> = {
  members: [{ email: '', type: 'member' }]
}
