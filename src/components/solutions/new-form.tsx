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
import { useEffect } from 'react'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import {
  rhfActionNewSolutionSchema,
  type NewSolutionFormValues
} from './new-form.schema'
import { rhfActionCreateNewSolution } from './new-form.action'
import { useForm } from 'react-hook-form'
import { useAction } from '@/trpc/client'
import { useRouter } from 'next/navigation'
import { SolutionTemplate } from '@/db/models/solution-templates'
import Markdown from 'react-markdown'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'

export type NewSolutionFormProps = {
  className?: string
  template?: SolutionTemplate
}

export function NewSolutionForm({ ...props }: NewSolutionFormProps) {
  const form = useForm<NewSolutionFormValues>({
    resolver: zodResolver(rhfActionNewSolutionSchema),
    defaultValues: {
      title: '',
      body: ''
    }
  })
  const router = useRouter()

  const mutation = useAction(rhfActionCreateNewSolution)
  async function onSubmit(data: NewSolutionFormValues) {
    await mutation.mutateAsync({ ...data })
  }

  useEffect(() => {
    if (mutation.status === 'success') {
      router.push(`/dashboard/solutions/${mutation.data?.id}`)
    }
  })

  return (
    <>
      <Form {...form}>
        <form
          action={rhfActionCreateNewSolution}
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8"
        >
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>
                  <h1>Title</h1>
                </FormLabel>
                <FormControl>
                  <Input placeholder="Title ..." {...field} />
                </FormControl>
                <FormDescription>Give it a great name.</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Tabs defaultValue="edit" className="w-full">
            <TabsList>
              <TabsTrigger value="edit">Edit</TabsTrigger>
              <TabsTrigger value="preview">Preview</TabsTrigger>
            </TabsList>
            <TabsContent value="edit">
              <FormField
                control={form.control}
                name="body"
                render={({ field }) => (
                  <FormItem>
                    <FormControl>
                      <Textarea
                        className="block w-full"
                        placeholder="Describe your solution ..."
                        rows={25}
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      This describes your solution in more detail.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </TabsContent>
            <TabsContent value="preview">
              <div className="border rounded p-4">
                <Markdown
                  components={{
                    h1(props) {
                      const { node, ...rest } = props
                      return (
                        <h1
                          className="scroll-m-20 text-4xl font-extrabold tracking-tight mt-6 lg:text-5x"
                          {...rest}
                        />
                      )
                    },
                    h2(props) {
                      const { node, ...rest } = props
                      return (
                        <h1
                          className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight mt-6 first:mt-0"
                          {...rest}
                        />
                      )
                    },
                    p(props) {
                      const { node, ...rest } = props
                      return (
                        <p
                          className="leading-7 [&:not(:first-child)]:mt-6"
                          {...rest}
                        />
                      )
                    }
                  }}
                >
                  {form.watch('body')}
                </Markdown>
              </div>
            </TabsContent>
          </Tabs>

          <Button
            type="submit"
            disabled={form.formState.isSubmitting || !form.formState.isValid}
          >
            New Solution
          </Button>
        </form>
      </Form>
    </>
  )
}
