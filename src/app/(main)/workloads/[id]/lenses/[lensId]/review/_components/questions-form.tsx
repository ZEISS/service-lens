"use client"

import type * as React from "react"

import { zodResolver } from "@hookform/resolvers/zod"
import { Controller, useForm, useWatch } from "react-hook-form"
import { toast } from "sonner"
import * as z from "zod"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Checkbox } from "@/components/ui/checkbox"
import { Field, FieldContent, FieldDescription, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field"
import { Form } from "@/components/ui/form"
import { InputGroup, InputGroupAddon, InputGroupText, InputGroupTextarea } from "@/components/ui/input-group"
import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"
import type { TLensPillarQuestionWithChoices } from "@/db/schema"

const formSchema = z.object({
  title: z
    .string()
    .min(5, "Bug title must be at least 5 characters.")
    .max(32, "Bug title must be at most 32 characters."),
  description: z
    .string()
    .min(20, "Description must be at least 20 characters.")
    .max(1024, "Description must be at most 1024 characters."),
  notes: z.string().max(1024, "Notes must be at most 1024 characters."),
  choices: z.array(
    z.object({
      id: z.string(),
      ref: z.string(),
      title: z.string(),
      description: z.string(),
    }),
  ),
  doesNotApply: z.boolean().optional(),
})

export interface QuestionFormProps {
  question?: TLensPillarQuestionWithChoices
}

export function QuestionForm({ question }: QuestionFormProps) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      title: "",
      description: "",
      notes: "",
      choices: question?.choices ?? [],
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
      classNames: {
        content: "flex flex-col gap-2",
      },
      style: {
        "--border-radius": "calc(var(--radius)  + 4px)",
      } as React.CSSProperties,
    })
  }

  return (
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
                    <Label htmlFor="does-not-apply">This question does not apply to the workload. {field.value}</Label>
                  </Field>
                </FieldGroup>
              )}
            />
          </div>
        </CardContent>
      </Card>

      {/* Choices */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">Answers</CardTitle>
          <CardDescription>Select the answers that best fit the workload.</CardDescription>
        </CardHeader>
        <CardContent>
          <Controller
            name="choices"
            control={form.control}
            render={({ field, fieldState, formState }) => (
              <FieldGroup className="max-w-sm">
                {field.value.map((choice, index) => (
                  <Field orientation="horizontal" key={index}>
                    <Checkbox
                      id={`question-checkbox-${index}`}
                      name={`question-checkbox-${index}`}
                      disabled={isDoesNotApply}
                    />
                    <FieldContent>
                      <FieldLabel htmlFor={`question-checkbox-${index}`}>
                        {choice.ref}: {choice.title}
                      </FieldLabel>
                      <FieldDescription>{choice.description}</FieldDescription>
                    </FieldContent>
                  </Field>
                ))}
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
      </Card>
    </Form>
  )
}
