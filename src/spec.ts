import Ajv, { JSONSchemaType } from 'ajv'
const ajv = new Ajv()

interface Spec {
  name: string
  description?: string
  pillars: [{ id: string }]
}

export const schema: JSONSchemaType<Spec> = {
  type: 'object',
  properties: {
    name: { type: 'string' },
    description: { type: 'string', nullable: true },
    pillars: {
      type: 'array',
      minItems: 1,
      additionalItems: false,
      items: [
        {
          type: 'object',
          properties: {
            id: {
              type: 'string'
            }
          },
          required: ['id']
        }
      ]
    }
  },
  required: ['name', 'pillars'],
  additionalProperties: false
}
