'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { toast } from '@/components/ui/use-toast'
import { use } from 'react'
import {
  SettingGeneralFormValues,
  settingsGeneralFormSchema
} from './settings-general-form.schema'
import { api } from '@/trpc/client'

export function SettingsGeneralForm({ teamId }: { teamId: string }) {
  const team = use(api.teams.get.query(teamId))

  const form = useForm<SettingGeneralFormValues>({
    resolver: zodResolver(settingsGeneralFormSchema),
    defaultValues: {
      name: team?.name,
      description: team?.description
    },
    mode: 'onChange'
  })

  function onSubmit(data: SettingGeneralFormValues) {
    toast({
      title: 'You submitted the following values:',
      description: (
        <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
          <code className="text-white">{JSON.stringify(data, null, 2)}</code>
        </pre>
      )
    })
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="sr-only">Name</FormLabel>
              <FormControl>
                <Input placeholder="Name ..." {...field} />
              </FormControl>
              <FormDescription>This is the name of the team.</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="description"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="sr-only">Description</FormLabel>
              <FormControl>
                <Input placeholder="Description ..." {...field} />
              </FormControl>
              <FormDescription>
                This a brief description of the application instance.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Update settings</Button>
      </form>
    </Form>
  )
}
