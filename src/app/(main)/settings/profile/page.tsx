import { ContentSection } from "../_components/content-section"
import { ProfileForm } from "./_components/profile-form"

export default async function SettingsProfile() {
  return (
    <ContentSection title="Profile" desc="This is how others will see you on the site.">
      <ProfileForm />
    </ContentSection>
  )
}
