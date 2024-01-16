import { PropsWithChildren } from 'react'
import { GoBug } from 'react-icons/go'

export type LowRiskProps = {}

export function LowRisk({ ...props }: PropsWithChildren<LowRiskProps>) {
  return <GoBug {...props} />
}
