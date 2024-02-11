import { PropsWithChildren } from 'react'
import { GoXCircle } from 'react-icons/go'

export type HighRiskProps = {}

export function HighRiskIcon({ ...props }: PropsWithChildren<HighRiskProps>) {
  return <GoXCircle {...props} />
}
