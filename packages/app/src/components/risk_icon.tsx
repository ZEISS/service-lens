import { PropsWithChildren } from 'react'
import { QuestionRisk } from '@/db/models/workload-lenses-answers'
import { HighRiskIcon } from '@/components/high_risk_icon'
import { MediumRisk } from '@/components/medium_risk_icon'
import { LowRisk } from '@/components/low_risk_icon'

export type RiskIconProps = {
  risk?: QuestionRisk
}

export function RiskIcon({
  risk = 'UNANSWERED',
  ...props
}: PropsWithChildren<RiskIconProps>) {
  if (risk === 'HIGH_RISK') {
    return <HighRiskIcon {...props} />
  }

  if (risk === 'MEDIUM_RISK') {
    return <MediumRisk {...props} />
  }

  if (risk === 'LOW_RISK') {
    return <LowRisk {...props} />
  }

  return <></>
}
