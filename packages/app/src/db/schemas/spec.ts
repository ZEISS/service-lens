import { z } from 'zod'

export const Resource = z.object({
  description: z.string().min(10).max(2048),
  url: z.string().url()
})

export type Resource = z.infer<typeof Resource>

export const Risk = z.object({
  risk: z.union([
    z.literal('HIGH_RISK'),
    z.literal('MEDIUM_RISK'),
    z.literal('LOW_RISK')
  ]),
  condition: z.string()
})

export type Risk = z.infer<typeof Risk>

export const Improvement = z.object({
  description: z.string().min(10).max(2048),
  url: z.string().url()
})

export type Improvement = z.infer<typeof Improvement>

export const Choice = z.object({
  id: z.string(),
  title: z.string().min(3).max(256),
  resources: z.array(Resource).optional(),
  improvments: z.array(Improvement).optional()
})

export type Choice = z.infer<typeof Choice>

export const Question = z.object({
  id: z.string(),
  title: z.string().min(3).max(256),
  description: z.string().min(10).max(2048),
  risks: z.array(Risk),
  resources: z.array(Resource).optional(),
  choices: z.array(Choice)
})

export type Question = z.infer<typeof Question>

export const Pillar = z.object({
  id: z.string(),
  name: z.string().min(3).max(256),
  description: z.string().min(10).max(2048),
  questions: z.array(Question),
  resources: z.array(Resource).optional()
})

export type Pillar = z.infer<typeof Pillar>

export const Spec = z.object({
  version: z.number(),
  name: z.string().min(3).max(256),
  description: z.string().min(10).max(2048),
  pillars: z.array(Pillar)
})

export type Spec = z.infer<typeof Spec>
