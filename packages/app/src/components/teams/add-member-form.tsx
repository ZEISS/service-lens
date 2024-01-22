'use client'

import {
  Form,
  FormItem,
  FormLabel,
  FormMessage,
  FormControl,
  FormField
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { type PropsWithChildren } from 'react'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select'
import { zodResolver } from '@hookform/resolvers/zod'
import {
  rhfActionAddMemberSchema,
  defaultValues,
  type AddMembersFormValues
} from './add-member-form.schema'
import { rhfActionAddMember } from './add-member-form.action'
import { useFieldArray, useForm } from 'react-hook-form'
import { useAction } from '@/trpc/client'

export type AddMemberFormProps = {
  className?: string
}

export function AddMemberForm({
  ...props
}: PropsWithChildren<AddMemberFormProps>) {
  const form = useForm<AddMembersFormValues>({
    resolver: zodResolver(rhfActionAddMemberSchema),
    defaultValues
  })

  const { fields, append, prepend, remove, swap, move, insert } = useFieldArray(
    { control: form.control, name: 'members' }
  )

  const mutation = useAction(rhfActionAddMember)
  async function onSubmit(data: AddMembersFormValues) {
    await mutation.mutateAsync({ ...data })
  }

  return (
    <>
      <Form {...form}>
        <form
          action={rhfActionAddMember}
          onSubmit={form.handleSubmit(onSubmit)}
          className="py-4"
        >
          {fields.map((field, index) => (
            <div key={field.id} className="flex flex-row gap-x-2 py-2">
              <FormField
                control={form.control}
                name={`members.${index}.email`}
                render={({ field }) => (
                  <FormItem className="space-y-0 w-full">
                    <FormLabel className="sr-only">Email</FormLabel>
                    <FormControl>
                      <Input
                        placeholder="indy.jones@lucasfilm.com"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name={`members.${index}.type`}
                render={({ field }) => (
                  <FormItem className="min-w-[160px]">
                    <Select onValueChange={value => field.onChange(value)}>
                      <SelectTrigger>
                        <SelectValue placeholder="Not selected" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="member">Member</SelectItem>
                        <SelectItem value="owner">Owner</SelectItem>
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          ))}
          <div className="flex justify-between py-4">
            <Button
              variant={'outline'}
              className="items-stretch"
              onClick={() => append({ email: '', type: 'member' }, {})}
            >
              Add Member
            </Button>
            <Button type="submit" disabled={form.formState.isSubmitting}>
              Add
            </Button>
          </div>
        </form>
      </Form>
    </>
  )
}
