import { db } from "@/db/index"
import { environments, type TNewEnvironment } from "@/db/schema"

export async function seedEnvironment(input: { names: string[] }) {
  const names = input.names ?? ["Development", "Staging", "Production"]

  try {
    const allEnvironments: TNewEnvironment[] = names.map((name) => ({ name }))

    await db.delete(environments)

    console.log("📝 Inserting environmments", allEnvironments.length)

    await db.insert(environments).values(allEnvironments).onConflictDoNothing()
  } catch (err) {
    console.error(err)
  }
}
