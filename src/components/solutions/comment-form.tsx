'use client'

import { Form, FormControl, FormField } from '@/components/ui/form'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter
} from '@/components/ui/card'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import { rhfActionSchema } from './comment-form.schema'
import { rhfAction } from './comment-form.action'
import { useForm } from 'react-hook-form'
import * as z from 'zod'
import { useAction } from '@/trpc/client'

export type CommentFormProps = {
  className?: string
  solutionId?: string
}

export function CommentForm({ solutionId, ...props }: CommentFormProps) {
  const form = useForm<z.infer<typeof rhfActionSchema>>({
    resolver: zodResolver(rhfActionSchema),
    defaultValues: {
      solutionId
    }
  })

  const mutation = useAction(rhfAction)
  async function onSubmit(data: z.infer<typeof rhfActionSchema>) {
    await mutation.mutateAsync({ ...data })
    form.reset({ body: '', solutionId })
  }

  return (
    <>
      <Form {...form}>
        <form
          action={rhfAction}
          onSubmit={form.handleSubmit(onSubmit)}
          className="py-4"
        >
          <FormField
            control={form.control}
            name="body"
            render={({ field }) => (
              <Card>
                <CardContent className="pt-6">
                  <FormControl>
                    <Textarea
                      {...field}
                      className="w-full"
                      placeholder="Add your comment..."
                    />
                  </FormControl>
                </CardContent>
                <CardFooter className="flex justify-between">
                  <CardDescription>Markdown is supported.</CardDescription>
                  <Button
                    variant="outline"
                    type="submit"
                    disabled={
                      form.formState.isSubmitting || !form.formState.isValid
                    }
                  >
                    Comment
                  </Button>
                </CardFooter>
              </Card>
            )}
          />
          <div className="flex justify-end"></div>
        </form>
      </Form>
    </>
  )
}
