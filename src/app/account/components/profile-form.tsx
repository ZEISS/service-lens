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
import {
  ProfileFormValues,
  rhfActionProfileSchema,
  defaultValues
} from '../actions/profile.schema'
import { Session } from 'next-auth'

export interface ProfileFormProps {
  session?: Session
}

export function ProfileForm({ session }: ProfileFormProps) {
  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(rhfActionProfileSchema),
    defaultValues: {
      ...defaultValues,
      email: session?.user.email ?? '',
      name: session?.user.name ?? ''
    },
    mode: 'onChange'
  })

  function onSubmit(data: ProfileFormValues) {
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
          disabled={true}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormDescription>
                You cannot change the name right now.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          disabled={true}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input disabled={true} {...field} />
              </FormControl>
              <FormDescription>
                You cannot change this email right now.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit" disabled={true}>
          Update profile
        </Button>
      </form>
    </Form>
  )
}
