import { Label } from "@/components/ui/label"
import { Switch } from "@/components/ui/switch"

export interface DoesNotApplySwitchProps {
  onCheckedChange: (checked: boolean) => void
}

export function DoesNotApplySwitch({ onCheckedChange }: DoesNotApplySwitchProps) {
  return (
    <div className="flex items-center space-x-2">
      <Switch id="does-not-apply" onCheckedChange={onCheckedChange} />
      <Label htmlFor="does-not-apply">This question does not apply to the workload.</Label>
    </div>
  )
}
