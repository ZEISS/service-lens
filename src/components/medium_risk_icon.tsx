import { PropsWithChildren } from 'react'
import { GoAlert } from 'react-icons/go'

export type MediumRiskProps = {}

export function MediumRisk({ ...props }: PropsWithChildren<MediumRiskProps>) {
  return <GoAlert {...props} />
}
