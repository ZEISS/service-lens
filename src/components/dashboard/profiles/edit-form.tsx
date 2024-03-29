'use client'

import { PropsWithChildren, useMemo, useState } from 'react'
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
import { Button } from '@/components/ui/button'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { ProfileQuestion } from '@/db/models/profile-question'
import { Checkbox } from '@/components/ui/checkbox'
import { Profile } from '@/db/models/profile'
import {
  EditProfileFormValues,
  rhfEditProfileActionSchema,
  defaultValues
} from './edit-form.schema'

export type EditProfileFormProps = {
  editable?: boolean
  questions?: ProfileQuestion[]
  profile?: Profile
}

export function EditProfileForm({
  questions,
  profile,
  editable = false
}: PropsWithChildren<EditProfileFormProps>) {
  const [isEditable, setIsEditable] = useState(editable)
  const selectedChoices = useMemo(() => {
    const question = questions?.reduce(
      (questions, question) => ({
        ...questions,
        [question.ref]: []
      }),
      {} as Record<string, string[]>
    )

    profile?.answers?.forEach(
      answer =>
        question &&
        answer?.question &&
        question[answer.question.ref].push(answer.id)
    )

    return question
  }, [profile?.answers, questions])

  const form = useForm<EditProfileFormValues>({
    resolver: zodResolver(rhfEditProfileActionSchema),
    defaultValues: selectedChoices,
    mode: 'onChange'
  })

  async function onSubmit(data: EditProfileFormValues) {
    console.log(data)
  }

  return (
    <>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8"
          autoComplete="off"
        >
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
                        disabled={!isEditable}
                        control={form.control}
                        name={question.ref}
                        render={({ field, ...rest }) => {
                          return (
                            <FormItem
                              key={choice.id}
                              className="flex flex-row items-start space-y-0 my-4"
                            >
                              <FormControl>
                                <Checkbox
                                  className="mr-2"
                                  checked={field.value?.includes(choice.id)}
                                  disabled={!isEditable}
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
                              <FormMessage />
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
                    name={question.ref}
                    render={({ field }) => (
                      <div className="grid w-full">
                        <FormControl>
                          <Select
                            disabled={!isEditable}
                            value={field.value?.[0]}
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

          {isEditable && (
            <Button type="submit" disabled={form.formState.isSubmitting}>
              Update Profile
            </Button>
          )}
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
