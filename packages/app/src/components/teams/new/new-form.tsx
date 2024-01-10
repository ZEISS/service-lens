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
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select'
import { api } from '@/trpc/client'
import { use } from 'react'
import { FancyMultiSelect } from '@/components/fancy-multi-select'
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

export type NewProfilesFormProps = {
  className?: string
}

export function NewProfilesForm({ ...props }: NewProfilesFormProps) {
  const profiles = use(api.listProfiles.query({}))
  const environments = use(api.listEnvironments.query({}))
  const lenses = use(api.listLenses.query({}))

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
  const handleSubmit = async (data: z.infer<typeof rhfActionSchema>) =>
    await mutation.mutateAsync({ ...data })

  useEffect(() => {
    if (mutation.status === 'success') {
      router.push(`/dashboard/workloads/${mutation.data?.id}`)
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
                  <Input {...field} />
                </FormControl>
                <FormDescription>Give it a great name.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="environmentsIds"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Environments</FormLabel>
                <FormControl>
                  <FancyMultiSelect
                    placeholder="Select environments ..."
                    onValueChange={field.onChange}
                    dataValues={environments?.rows.map(env => ({
                      value: env.id,
                      label: env.name
                    }))}
                  />
                </FormControl>
                <FormDescription>Select matching environments.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="profilesId"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Profile</FormLabel>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Select a profile" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectGroup>
                      {profiles?.rows.map((profile, idx) => (
                        <SelectItem key={idx} value={profile.id}>
                          {profile.name}
                        </SelectItem>
                      ))}
                    </SelectGroup>
                  </SelectContent>
                </Select>
                <FormDescription>Select matching profile.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="lensesIds"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Lenses</FormLabel>
                <FormControl>
                  <FancyMultiSelect
                    placeholder="Select lenses ..."
                    onValueChange={field.onChange}
                    dataValues={lenses?.rows.map(lens => ({
                      value: lens.id,
                      label: lens.name
                    }))}
                  />
                </FormControl>
                <FormDescription>Select the matching lenses.</FormDescription>
                <FormMessage />
              </FormItem>
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
                    className="w-full"
                    placeholder="Add a description ..."
                  />
                </FormControl>
                <FormDescription>
                  A desciption of your workload.
                </FormDescription>
                <FormMessage />
              </div>
            )}
          />

          <Button
            type="submit"
            disabled={form.formState.isSubmitting || !form.formState.isValid}
          >
            Add Workload
          </Button>
        </form>
      </Form>
    </>
  )
}
