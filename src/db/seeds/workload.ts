import { db } from "@/db/index"
import { type TNewWorkload, workloads } from "@/db/schema"

export async function seedWorkload(input: { count: number }) {
  const count = input.count ?? 100

  try {
    const allWorkloads: TNewWorkload[] = []

    // Add some predefined workloads with rich content
    allWorkloads.push(...getPredefinedWorkloads())

    await db.delete(workloads)

    console.log("📝 Inserting workloads", allWorkloads.length)

    await db.insert(workloads).values(allWorkloads).onConflictDoNothing()
  } catch (err) {
    console.error(err)
  }
}

export function generateRandomWorkload(input?: Partial<TNewWorkload>): TNewWorkload {
  const workloadNumber = Math.floor(Math.random() * 1000)

  return {
    name: `Workload ${workloadNumber}`,
    description: `This is a description for workload ${workloadNumber}`,
    ...input,
  }
}

function getPredefinedWorkloads(): TNewWorkload[] {
  return [
    {
      name: "E-commerce Platform",
      description: "A workload for an online shopping platform with high traffic and complex transactions.",
    },
    {
      name: "Real-time Analytics",
      description: "A workload for processing and analyzing streaming data in real-time.",
    },
    {
      name: "Content Management System",
      description: "A workload for managing and delivering digital content across multiple channels.",
    },
  ]
}
