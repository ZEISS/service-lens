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
import {
  SettingGeneralFormValues,
  settingsGeneralFormSchema,
  defaultValues
} from './settings-general-form.schema'

export function SettingsGeneralForm() {
  const form = useForm<SettingGeneralFormValues>({
    resolver: zodResolver(settingsGeneralFormSchema),
    defaultValues
  })

  return (
    <Form {...form}>
      <form className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="sr-only">Name</FormLabel>
              <FormControl>
                <Input placeholder="Name ..." {...field} />
              </FormControl>
              <FormDescription>This is the name of the site.</FormDescription>
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
                This is a brief description of the site.
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
