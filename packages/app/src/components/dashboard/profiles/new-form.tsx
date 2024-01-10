'use client'

import { PropsWithChildren } from 'react'
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
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useEffect } from 'react'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import { rhfActionSchema } from './new-form.schema'
import { rhfAction } from './new-form.action'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { useAction } from '@/trpc/client'
import { useRouter } from 'next/navigation'
import { ProfileQuestion } from '@/db/models/profile-question'
import { Checkbox } from '@/components/ui/checkbox'
import { defaultValues } from './new-form.schema'

export type NewProfileFormProps = {
  questions?: ProfileQuestion[]
  selectedChoices?: Record<string, string[]>
}

export function NewProfileForm({
  questions,
  selectedChoices
}: PropsWithChildren<NewProfileFormProps>) {
  const form = useForm<z.infer<typeof rhfActionSchema>>({
    resolver: zodResolver(rhfActionSchema),
    defaultValues: {
      ...defaultValues,
      selectedChoices
    },
    mode: 'onChange'
  })
  const router = useRouter()

  const mutation = useAction(rhfAction)
  async function onSubmit(data: z.infer<typeof rhfActionSchema>) {
    await mutation.mutateAsync({ ...data })
  }

  useEffect(() => {
    if (mutation.status === 'success') {
      router.push(`/dashboard/profiles/${mutation.data?.id}`)
    }
  })

  return (
    <>
      <Form {...form}>
        <form
          action={rhfAction}
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8"
          autoComplete="off"
        >
          <Card>
            <CardHeader>
              <CardTitle>Overview</CardTitle>
            </CardHeader>
            <CardContent>
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem className="pb-6">
                    <FormControl>
                      <Input
                        {...field}
                        autoComplete="off"
                        placeholder="Add a name ..."
                      />
                    </FormControl>
                    <FormDescription>Give it a great name.</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="description"
                render={({ field }) => (
                  <FormItem>
                    <FormControl>
                      <Textarea
                        {...field}
                        className="w-full"
                        placeholder="Add a description ..."
                      />
                    </FormControl>
                    <FormDescription>
                      Provide a description for this profile.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </CardContent>
          </Card>

          {questions?.map((question, idx) => (
            <Card key={idx}>
              <CardHeader>
                <CardTitle>{question?.name}</CardTitle>
              </CardHeader>
              <CardContent>
                {question.isMultiple ? (
                  <div key={idx}>
                    {question?.choices?.map(choice => (
                      <FormField
                        key={choice.id}
                        control={form.control}
                        name={`selectedChoices.${question.ref}`}
                        render={({ field, ...rest }) => {
                          return (
                            <FormItem
                              key={choice.id}
                              className="flex flex-row items-start space-y-0 my-4"
                            >
                              <FormControl>
                                <Checkbox
                                  className="mr-2"
                                  checked={field.value.includes(choice.id)}
                                  onCheckedChange={checked => {
                                    return checked
                                      ? field.onChange([
                                          ...field.value,
                                          choice.id
                                        ])
                                      : field.onChange(
                                          field.value?.filter(
                                            value => value !== choice.id
                                          )
                                        )
                                  }}
                                />
                              </FormControl>
                              <FormLabel className="font-normal">
                                {choice.name}
                              </FormLabel>
                            </FormItem>
                          )
                        }}
                      />
                    ))}
                  </div>
                ) : (
                  <FormField
                    key={idx}
                    control={form.control}
                    name={`selectedChoices.${question.ref}`}
                    render={({ field }) => (
                      <div className="grid w-full">
                        <FormControl>
                          <Select
                            onValueChange={value => {
                              field.onChange([value])
                            }}
                          >
                            <FormControl>
                              <SelectTrigger>
                                <SelectValue placeholder="Not selected" />
                              </SelectTrigger>
                            </FormControl>
                            <SelectContent>
                              {question.choices?.map((choice, c) => (
                                <SelectItem key={c} value={choice.id}>
                                  {choice.name}
                                </SelectItem>
                              ))}
                            </SelectContent>
                          </Select>
                        </FormControl>
                        <FormDescription>
                          {question.description}
                        </FormDescription>
                        <FormMessage />
                      </div>
                    )}
                  />
                )}
              </CardContent>
            </Card>
          ))}

          <Button
            type="submit"
            disabled={form.formState.isSubmitting || !form.formState.isValid}
          >
            Add Profile
          </Button>
          <input
            autoComplete="false"
            name="hidden"
            type="text"
            style={{ display: 'none' }}
          ></input>
        </form>
      </Form>
    </>
  )
}
