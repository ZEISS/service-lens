"use client"

import type * as React from "react"

import { zodResolver } from "@hookform/resolvers/zod"
import { Controller, useForm, useWatch } from "react-hook-form"
import { toast } from "sonner"
import * as z from "zod"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Checkbox } from "@/components/ui/checkbox"
import { Field, FieldContent, FieldDescription, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Form } from "@/components/ui/form"
import { InputGroup, InputGroupAddon, InputGroupText, InputGroupTextarea } from "@/components/ui/input-group"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"
import type { TLensPillarQuestionWithChoices } from "@/db/schema"

const formSchema = z.object({
  notes: z.string().max(1024, "Notes must be at most 1024 characters."),
  choices: z.array(z.string()),
  doesNotApply: z.boolean().optional(),
})

export interface QuestionFormProps {
  question?: TLensPillarQuestionWithChoices
}

export function QuestionForm({ question }: QuestionFormProps) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      notes: "",
      choices: [],
      doesNotApply: false,
    },
  })

  const isDoesNotApply = useWatch({
    control: form.control,
    name: "doesNotApply",
  })

  function onSubmit(data: z.infer<typeof formSchema>) {
    toast("You submitted the following values:", {
      description: (
        <pre className="mt-2 w-[320px] overflow-x-auto rounded-md bg-code p-4 text-code-foreground">
          <code>{JSON.stringify(data, null, 2)}</code>
        </pre>
      ),
      position: "bottom-right",
    })
  }

  return (
    <form
      id="questions-form"
      onSubmit={form.handleSubmit(onSubmit)}
      className="@container/main flex flex-col gap-4 md:gap-6"
    >
      <Form {...form}>
        {/* Title */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">{question?.title}</CardTitle>
            <CardDescription>{question?.description || "No description provided."}</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex items-center space-x-2">
              <Controller
                name="doesNotApply"
                control={form.control}
                render={({ field }) => (
                  <FieldGroup className="max-w-sm">
                    <Field orientation="horizontal">
                      <Switch checked={field.value} onCheckedChange={field.onChange} />
                      <Label htmlFor="does-not-apply">
                        This question does not apply to the workload. {field.value}
                      </Label>
                    </Field>
                  </FieldGroup>
                )}
              />
            </div>
          </CardContent>
        </Card>

        {/* Choices */}
        <Card>
          <CardContent>
            <Controller
              name="choices"
              control={form.control}
              render={({ field }) => (
                <FieldGroup className="max-w-sm">
                  {question?.choices?.map((choice) => {
                    const isChecked = field.value.includes(choice.id)
                    return (
                      <Field orientation="horizontal" key={choice.id}>
                        <Checkbox
                          id={`question-checkbox-${choice.ref}`}
                          checked={isChecked}
                          onCheckedChange={() => {
                            const newValue = isChecked
                              ? field.value.filter((v) => v !== choice.id)
                              : [...field.value, choice.id]
                            field.onChange(newValue)
                          }}
                          disabled={isDoesNotApply}
                        />
                        <FieldContent>
                          <FieldLabel htmlFor={`question-checkbox-${choice.ref}`}>{choice.title}</FieldLabel>
                          <FieldDescription>{choice.description}</FieldDescription>
                        </FieldContent>
                      </Field>
                    )
                  })}
                </FieldGroup>
              )}
            />
          </CardContent>
        </Card>

        {/* Notes */}
        <Card>
          <CardContent>
            <FieldGroup>
              <Controller
                name="notes"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="form-rhf-demo-description">Notes</FieldLabel>
                    <InputGroup>
                      <InputGroupTextarea
                        {...field}
                        id="question-form-notes"
                        placeholder="Add your notes here ..."
                        rows={6}
                        className="min-h-24 resize-none"
                        aria-invalid={fieldState.invalid}
                      />
                      <InputGroupAddon align="block-end">
                        <InputGroupText className="tabular-nums">{field.value.length}/1024 characters</InputGroupText>
                      </InputGroupAddon>
                    </InputGroup>
                    <FieldDescription>Optional, add any additional notes or context here.</FieldDescription>
                    {fieldState.invalid && <FieldError errors={[fieldState.error]} />}
                  </Field>
                )}
              />
            </FieldGroup>
          </CardContent>
          <CardFooter>
            <Field orientation="horizontal">
              <Button type="button" variant="outline" onClick={() => form.reset()}>
                Reset
              </Button>
              <Button type="submit" form="questions-form">
                Save & Next
              </Button>
            </Field>
          </CardFooter>
        </Card>
      </Form>
    </form>
  )
}
