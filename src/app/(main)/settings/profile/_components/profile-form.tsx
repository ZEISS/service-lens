"use client"

import { Button } from "@/components/ui/button"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useSession } from "@/lib/auth-client"
import { useMemo, useEffect } from "react"
import { useForm } from "react-hook-form"
import { toast } from "sonner"
import { zodResolver } from "@hookform/resolvers/zod"
import type { User } from "better-auth"
import { z } from "zod"
import { client } from "@/lib/auth-client"

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

  useEffect(() => {
     form.reset({
       name: session.data?.user.name || defaultValues.name,
       email: session.data?.user.email || defaultValues.email,
     });
  }, [session]);


  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(FormSchema),
    defaultValues: useMemo(() => {
            return {
              name: session.data?.user.name || defaultValues.name,
              email: session.data?.user.email || defaultValues.email,
            }
        }, [session]),
    mode: "onChange"
  })

  const onSubmit = form.handleSubmit(async (data) => {
    try {
      await client.updateUser(data)
      toast.success("Successfully updated profile")
    } catch (error) {
      toast.error("Failed to update profile")
    }
  })

  return (
    <Form {...form}>
      <form className="space-y-8" onSubmit={onSubmit}>
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
        <Button type="submit">Update profile</Button>
      </form>
    </Form>
  )
}
