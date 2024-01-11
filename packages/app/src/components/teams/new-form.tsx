'use client'

import {
  Form,
  FormControl,
  FormItem,
  FormLabel,
  FormDescription,
  FormMessage,
  FormField
} from '@/components/ui/form'
import { Textarea } from '@/components/ui/textarea'
import { useEffect } from 'react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import { rhfActionSchema } from './new-form.schema'
import { rhfAction } from './new-form.action'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { useAction } from '@/trpc/client'
import { useRouter } from 'next/navigation'
import type { PropsWithChildren } from 'react'
import type { NewTeamFormValues } from './new-form.schema'
import { defaultValues } from './new-form.schema'

export type NewTeamFormProps = {}

export function NewTeamForm({ ...props }: PropsWithChildren<NewTeamFormProps>) {
  const form = useForm<NewTeamFormValues>({
    resolver: zodResolver(rhfActionSchema),
    defaultValues,
    mode: 'onChange'
  })
  const router = useRouter()

  const mutation = useAction(rhfAction)
  const handleSubmit = async (data: z.infer<typeof rhfActionSchema>) =>
    await mutation.mutateAsync({ ...data })

  useEffect(() => {
    if (mutation.status === 'success') {
      router.push(`/teams/${mutation.data?.id}`)
    }
  })

  return (
    <>
      <Form {...form}>
        <form
          action={rhfAction}
          onSubmit={form.handleSubmit(handleSubmit)}
          className="space-y-8"
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  <h1>Name</h1>
                </FormLabel>
                <FormControl>
                  <Input placeholder="Team name ..." {...field} />
                </FormControl>
                <FormDescription>Give it a great name.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="contactEmail"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  <h1>Contact email</h1>
                </FormLabel>
                <FormControl>
                  <Input placeholder="team@acme.com" {...field} />
                </FormControl>
                <FormDescription>
                  Add a shared inbox for you team (optional).
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <div className="grid w-full">
                <FormItem>
                  <FormLabel>
                    <h1>Description</h1>
                  </FormLabel>
                  <FormControl>
                    <Textarea
                      {...field}
                      className="w-full"
                      placeholder="Add a description ..."
                    />
                  </FormControl>
                  <FormDescription>A desciption of your team</FormDescription>
                  <FormMessage />
                </FormItem>
              </div>
            )}
          />

          <Button
            type="submit"
            disabled={form.formState.isSubmitting || !form.formState.isValid}
          >
            Create team
          </Button>
        </form>
      </Form>
    </>
  )
}
