import { z } from "zod"

/**
 * Resource schema - links to external documentation or resources
 */
export const resourceSchema = z.object({
  url: z.string().url(),
  description: z.string(),
})

export type Resource = z.infer<typeof resourceSchema>

/**
 * Choice schema - represents a selectable option in a question
 */
export const choiceSchema = z.object({
  ref: z.string().min(1),
  title: z.string().min(1),
  description: z.string(),
})

export type Choice = z.infer<typeof choiceSchema>

/**
 * Risk level enum
 */
export const riskLevelSchema = z.enum(["NO_RISK", "MEDIUM_RISK", "HIGH_RISK"])

export type RiskLevel = z.infer<typeof riskLevelSchema>

/**
 * Risk schema - defines risk level based on condition
 * Condition can be a boolean expression with choice refs or "default"
 */
export const riskSchema = z.object({
  risk: riskLevelSchema,
  condition: z.string().min(1),
})

export type Risk = z.infer<typeof riskSchema>

/**
 * Question schema - represents a single evaluation question
 */
export const questionSchema = z.object({
  ref: z.string().min(1),
  title: z.string().min(1),
  description: z.string(),
  resources: z.array(resourceSchema).optional().default([]),
  choices: z.array(choiceSchema).min(1),
  risks: z.array(riskSchema).min(1),
})

export type Question = z.infer<typeof questionSchema>

/**
 * Pillar schema - represents a category/pillar of evaluation
 */
export const pillarSchema = z.object({
  ref: z.string().min(1),
  name: z.string().min(1),
  description: z.string(),
  questions: z.array(questionSchema).min(1),
  resources: z.array(resourceSchema).optional().default([]),
})

export type Pillar = z.infer<typeof pillarSchema>

/**
 * Lens specification schema - the root schema for a lens definition
 */
export const lensSpecSchema = z.object({
  version: z.number().int().positive(),
  name: z.string().min(1),
  description: z.string(),
  pillars: z.array(pillarSchema).min(1),
})

export type LensSpec = z.infer<typeof lensSpecSchema>

/**
 * JSON Schema representation of the Lens Spec
 */
export const lensSpecJsonSchema = {
  $schema: "http://json-schema.org/draft-07/schema#",
  title: "LensSpec",
  description: "Schema for defining a lens specification for evaluating workloads against best practices",
  type: "object",
  required: ["version", "name", "description", "pillars"],
  properties: {
    version: {
      type: "integer",
      minimum: 1,
      description: "Version number of the lens specification",
    },
    name: {
      type: "string",
      minLength: 1,
      description: "Name of the lens",
    },
    description: {
      type: "string",
      description: "Description of what the lens evaluates",
    },
    pillars: {
      type: "array",
      minItems: 1,
      description: "Array of pillars (categories) for evaluation",
      items: {
        $ref: "#/definitions/Pillar",
      },
    },
  },
  definitions: {
    Resource: {
      type: "object",
      required: ["url", "description"],
      properties: {
        url: {
          type: "string",
          format: "uri",
          description: "URL to the external resource",
        },
        description: {
          type: "string",
          description: "Description of the resource",
        },
      },
    },
    Choice: {
      type: "object",
      required: ["ref", "title", "description"],
      properties: {
        ref: {
          type: "string",
          minLength: 1,
          description: "Unique reference identifier for the choice",
        },
        title: {
          type: "string",
          minLength: 1,
          description: "Display title for the choice",
        },
        description: {
          type: "string",
          description: "Detailed description of the choice",
        },
      },
    },
    Risk: {
      type: "object",
      required: ["risk", "condition"],
      properties: {
        risk: {
          type: "string",
          enum: ["NO_RISK", "MEDIUM_RISK", "HIGH_RISK"],
          description: "Risk level",
        },
        condition: {
          type: "string",
          minLength: 1,
          description: "Boolean expression using choice refs (e.g., 'choice_1 && choice_2') or 'default'",
        },
      },
    },
    Question: {
      type: "object",
      required: ["ref", "title", "description", "choices", "risks"],
      properties: {
        ref: {
          type: "string",
          minLength: 1,
          description: "Unique reference identifier for the question",
        },
        title: {
          type: "string",
          minLength: 1,
          description: "Display title for the question",
        },
        description: {
          type: "string",
          description: "Detailed description of the question",
        },
        resources: {
          type: "array",
          items: {
            $ref: "#/definitions/Resource",
          },
          default: [],
          description: "Optional resources related to the question",
        },
        choices: {
          type: "array",
          minItems: 1,
          items: {
            $ref: "#/definitions/Choice",
          },
          description: "Available choices for the question",
        },
        risks: {
          type: "array",
          minItems: 1,
          items: {
            $ref: "#/definitions/Risk",
          },
          description: "Risk definitions based on choice selections",
        },
      },
    },
    Pillar: {
      type: "object",
      required: ["ref", "name", "description", "questions"],
      properties: {
        ref: {
          type: "string",
          minLength: 1,
          description: "Unique reference identifier for the pillar",
        },
        name: {
          type: "string",
          minLength: 1,
          description: "Display name for the pillar",
        },
        description: {
          type: "string",
          description: "Description of the pillar",
        },
        questions: {
          type: "array",
          minItems: 1,
          items: {
            $ref: "#/definitions/Question",
          },
          description: "Questions within this pillar",
        },
        resources: {
          type: "array",
          items: {
            $ref: "#/definitions/Resource",
          },
          default: [],
          description: "Optional resources related to the pillar",
        },
      },
    },
  },
} as const

/**
 * Parse and validate a lens specification from JSON
 * @param data - The JSON data to parse
 * @returns Validated LensSpec object
 * @throws ZodError if validation fails
 */
export function parseLensSpec(data: unknown): LensSpec {
  return lensSpecSchema.parse(data)
}

/**
 * Safely parse a lens specification, returning success/error result
 * @param data - The JSON data to parse
 * @returns SafeParseReturnType with success/error information
 */
export function safeParseLensSpec(data: unknown) {
  return lensSpecSchema.safeParse(data)
}
