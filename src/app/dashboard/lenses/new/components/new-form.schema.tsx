import { z } from 'zod'

const readJSONFile = async (file: File) =>
  new Promise<string>((resolve, reject) => {
    const reader = new FileReader()
    reader.onload = event => resolve(event.target?.result as string)
    reader.onerror = error => reject(error)
    reader.readAsText(file)
  })

export const rhfActionSchema = z.object({
  name: z.string().min(3).max(256).default(''),
  spec: z
    .union([
      z
        .custom<FileList>()
        .refine(file =>
          typeof window === 'undefined'
            ? false
            : FileList && file.length > 0 && file.item(0)
        )
        .transform(async file => await readJSONFile(file.item(0) as File))
        .refine(file => file),
      z.string()
    ])
    .pipe(z.coerce.string()),
  description: z
    .string()
    .min(10, {
      message: 'Description must be at least 10 characters.'
    })
    .max(2024, {
      message: 'Description must be less than 2024 characters.'
    })
    .optional()
    .default('')
})
