import { Separator } from '@/components/ui/separator'
import { SettingsGeneralForm } from '@/components/settings/settings-general-form'

export default async function Page() {
  return (
    <>
      <div className="space-y-6">
        <div>
          <h3 className="text-lg font-medium">General</h3>
          <p className="text-sm text-muted-foreground">Site-wide settings</p>
        </div>
        <Separator />
        <SettingsGeneralForm />
      </div>
    </>
  )
}
