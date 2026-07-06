"use client"

import { Button } from "@/components/ui/button"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useSession } from "@/lib/auth-client"
import { useForm } from "react-hook-form"
// import { showSubmittedData } from '@/lib/show-submitted-data'
import { zodResolver } from "@hookform/resolvers/zod"
import type { User } from "better-auth"
import { z } from "zod"

const FormSchema = z.object({
  name: z.string().min(2, { message: "Name must be at least 2 characters." }),
  email: z.string().email({ message: "Please enter a valid email address." }),
})

type ProfileFormValues = z.infer<typeof FormSchema>

// This can come from your database or API.
const defaultValues: Partial<User> = {
  name: "Indy Jones",
  email: "indy@jones.com",
}

export function ProfileForm() {
  const session = useSession()

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      name: session.data?.user.name || defaultValues.name,
      email: session.data?.user.email || defaultValues.email,
    },
    mode: "onChange",
  })

  // const { fields, append } = useFieldArray({
  //     name: 'urls',
  //     control: form.control,
  // })

  return (
    <Form {...form}>
      <form
        // onSubmit={form.handleSubmit((data) => showSubmittedData(data))}
        className="space-y-8"
      >
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormDescription>
                This is your public display name. It can be your real name or a pseudonym. You can only change this once
                every 30 days.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          disabled
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormDescription>Your email address cannot be changed.</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <div>
          {/* {fields.map((field, index) => (
                        <FormField
                            control={form.control}
                            key={field.id}
                            name={`urls.${index}.value`}
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel
                                        className={cn(index !== 0 && 'sr-only')}
                                    >
                                        URLs
                                    </FormLabel>
                                    <FormDescription
                                        className={cn(index !== 0 && 'sr-only')}
                                    >
                                        Add links to your website, blog, or
                                        social media profiles.
                                    </FormDescription>
                                    <FormControl
                                        className={cn(index !== 0 && 'mt-1.5')}
                                    >
                                        <Input {...field} />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    ))} */}
        </div>
        <Button type="submit">Update profile</Button>
      </form>
    </Form>
  )
}
