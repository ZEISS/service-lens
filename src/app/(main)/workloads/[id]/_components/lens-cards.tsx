import type { TLens } from "@/db/schema"

import { LensCard } from "./lens-card"

// Props for the LensCards component
export interface LensCardsProps {
  workloadId: string
  lenses?: TLens[]
}

// LensCards component that displays a grid of LensCard components
export function LensCards({ lenses, workloadId }: LensCardsProps) {
  return (
    <div className="grid grid-cols-2 gap-4">
      {lenses?.map((lens) => (
        <LensCard key={lens.id} lens={lens} workloadId={workloadId} />
      ))}
    </div>
  )
}
