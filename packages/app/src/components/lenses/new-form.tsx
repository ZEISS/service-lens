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
import type { PropsWithChildren } from 'react'
import { Textarea } from '@/components/ui/textarea'
import { useEffect, useMemo } from 'react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import { rhfActionSchema } from './new-form.schema'
import { rhfAction } from './new-form.action'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { useAction } from '@/trpc/client'
import { useRouter } from 'next/navigation'

export type NewLensFormProps = {}

export function NewLensForm({ ...props }: PropsWithChildren<NewLensFormProps>) {
  const form = useForm<z.infer<typeof rhfActionSchema>>({
    resolver: zodResolver(rhfActionSchema),
    defaultValues: {
      name: '',
      description: ''
    },
    mode: 'onChange'
  })
  const router = useRouter()

  const mutation = useAction(rhfAction)
  async function onSubmit(data: z.infer<typeof rhfActionSchema>) {
    await mutation.mutateAsync({ ...data })
  }

  const isLoading = useMemo(
    () => mutation.status === 'loading',
    [mutation.status]
  )

  useEffect(() => {
    if (mutation.status === 'success') {
      router.push(`/dashboard/lenses/${mutation.data?.id}`)
    }
  }, [mutation.status, mutation.data?.id, router])

  const fileRef = form.register('spec', { required: true })

  return (
    <>
      <Form {...form}>
        <form
          action={rhfAction}
          onSubmit={form.handleSubmit(onSubmit)}
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
                  <Input disabled={isLoading} {...field} />
                </FormControl>
                <FormDescription>Give it a great name.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="spec"
            render={({ field }) => (
              <div className="grid w-full max-w-sm items-center gap-1.5">
                <FormLabel>Specification</FormLabel>
                <FormControl>
                  <Input
                    type="file"
                    disabled={isLoading}
                    placeholder=""
                    {...fileRef}
                  />
                </FormControl>
                <FormDescription>
                  This must follow the Lens Format Specification (max. 3MB).
                </FormDescription>
                <FormMessage />
              </div>
            )}
          />
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <div className="grid w-full">
                <FormControl>
                  <Textarea
                    {...field}
                    disabled={isLoading}
                    className="w-full"
                    placeholder="Add a description ..."
                  />
                </FormControl>
                <FormDescription>A desciption for your lens.</FormDescription>
                <FormMessage />
              </div>
            )}
          />
          <Button type="submit" disabled={isLoading || !form.formState.isValid}>
            New Lens
          </Button>
        </form>
      </Form>
    </>
  )
}
