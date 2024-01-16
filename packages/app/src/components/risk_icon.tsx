import { PropsWithChildren } from 'react'
import { QuestionRisk } from '@/db/models/lens-pillar-risks'
import { HighRiskIcon } from '@/components/high_risk_icon'
import { MediumRisk } from '@/components/medium_risk_icon'
import { LowRisk } from '@/components/low_risk_icon'

export type RiskIconProps = {
  risk?: QuestionRisk
}

export function RiskIcon({
  risk = QuestionRisk.None,
  ...props
}: PropsWithChildren<RiskIconProps>) {
  if (risk === QuestionRisk.High) {
    return <HighRiskIcon {...props} />
  }

  if (risk === QuestionRisk.Medium) {
    return <MediumRisk {...props} />
  }

  if (risk === QuestionRisk.Low) {
    return <LowRisk {...props} />
  }

  return <></>
}
