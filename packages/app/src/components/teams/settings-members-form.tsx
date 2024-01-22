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
import { PropsWithChildren, use } from 'react'
import {
  SettingGeneralFormValues,
  settingsGeneralFormSchema
} from './settings-general-form.schema'
import { api } from '@/trpc/client'

export type SettingsMembersFormProps = {
  teamId: string
}

export function SettingsMembersForm({
  teamId = ''
}: PropsWithChildren<SettingsMembersFormProps>) {
  const team = use(api.teams.getUsersByName.query({ slug: teamId }))

  const form = useForm<SettingGeneralFormValues>({
    resolver: zodResolver(settingsGeneralFormSchema),
    defaultValues: {
      name: team?.name,
      slug: team?.slug,
      description: team?.description
    },
    mode: 'onChange'
  })

  return (
    <>{team?.users.map((user, index) => <div key={index}>{user.name}</div>)}</>
  )

  // return (

  //   <Form {...form}>
  //     <form className="space-y-8">
  //       <FormField
  //         control={form.control}
  //         name="name"
  //         render={({ field }) => (
  //           <FormItem>
  //             <FormLabel className="sr-only">Name</FormLabel>
  //             <FormControl>
  //               <Input disabled={true} placeholder="Name ..." {...field} />
  //             </FormControl>
  //             <FormDescription>This is the name of the team.</FormDescription>
  //             <FormMessage />
  //           </FormItem>
  //         )}
  //       />
  //       <FormField
  //         control={form.control}
  //         name="slug"
  //         render={({ field }) => (
  //           <FormItem>
  //             <FormLabel className="sr-only">Slug</FormLabel>
  //             <FormControl>
  //               <Input disabled={true} placeholder="Slug ..." {...field} />
  //             </FormControl>
  //             <FormDescription>
  //               {`This is the short name used for URLs (e.g.
  //               'solution-architects', 'order-service')`}
  //             </FormDescription>
  //             <FormMessage />
  //           </FormItem>
  //         )}
  //       />
  //       <FormField
  //         control={form.control}
  //         name="description"
  //         render={({ field }) => (
  //           <FormItem>
  //             <FormLabel className="sr-only">Description</FormLabel>
  //             <FormControl>
  //               <Input
  //                 disabled={true}
  //                 placeholder="Description ..."
  //                 {...field}
  //               />
  //             </FormControl>
  //             <FormDescription>
  //               This a brief description of the application instance.
  //             </FormDescription>
  //             <FormMessage />
  //           </FormItem>
  //         )}
  //       />
  //       <Button disabled={true} type="submit">
  //         Update settings
  //       </Button>
  //     </form>
  //   </Form>
  // )
}
