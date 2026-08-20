import {
  user,
  session,
  tags,
  account,
  profiles,
  verification,
  workloads,
  organization,
  team,
  teamMember,
  member,
  invitation,
  environments,
  lenses,
  lensPillars,
  lensPillarQuestions,
  lensPillarQuestionRisks,
  lensPillarQuestionChoices,
  lensPillarQuestionResources,
  workloadEnvironment,
  workloadTag,
  workloadProfile,
  workloadLens,
} from "./schema"
import { defineRelations } from "drizzle-orm"

export const relations = defineRelations(
  {
    user,
    session,
    tags,
    account,
    verification,
    workloads,
    organization,
    team,
    teamMember,
    member,
    invitation,
    environments,
    workloadEnvironment,
    workloadLens,
    workloadProfile,
    workloadTag,
    profiles,
    lenses,
    lensPillars,
    lensPillarQuestions,
    lensPillarQuestionRisks,
    lensPillarQuestionChoices,
    lensPillarQuestionResources,
  },
  ({
    one,
    many,
    organization,
    invitation,
    user,
    member,
    account,
    session,
    team,
    teamMember,
    profiles,
    workloads,
    workloadEnvironment,
    workloadTag,
    workloadProfile,
    workloadLens,
    environments,
    lenses,
    lensPillars,
    lensPillarQuestions,
    lensPillarQuestionRisks,
    lensPillarQuestionChoices,
    lensPillarQuestionResources,
  }) => ({
    workloads: {
      environments: many.environments({
        from: workloads.id.through(workloadEnvironment.workloadId),
        to: environments.id.through(workloadEnvironment.environmentId),
      }),
      lenses: many.lenses({
        from: workloads.id.through(workloadLens.workloadId),
        to: lenses.id.through(workloadLens.lensId),
      }),
      profiles: many.profiles({
        from: workloads.id.through(workloadProfile.workloadId),
        to: profiles.id.through(workloadProfile.profileId),
      }),
    },
    lenses: {
      lensPillars: many.lensPillars({
        from: lenses.id,
        to: lensPillars.lensId,
      }),
    },
    lensPillars: {
      questions: many.lensPillarQuestions({
        from: lensPillars.id,
        to: lensPillarQuestions.pillarId,
      }),
    },
    lensPillarQuestions: {
      risks: many.lensPillarQuestionRisks({
        from: lensPillarQuestions.id,
        to: lensPillarQuestionRisks.questionId,
      }),
      choices: many.lensPillarQuestionChoices({
        from: lensPillarQuestions.id,
        to: lensPillarQuestionChoices.questionId,
      }),
      resources: many.lensPillarQuestionResources({
        from: lensPillarQuestions.id,
        to: lensPillarQuestionResources.questionId,
      }),
    },
    users: {
      sessions: many.session({
        from: user.id,
        to: session.userId,
      }),
      accounts: many.account({
        from: user.id,
        to: account.userId,
      }),
      teamMembers: many.teamMember({
        from: user.id,
        to: teamMember.userId,
      }),
      members: many.member({
        from: user.id,
        to: member.userId,
      }),
      invitations: many.invitation({
        from: user.id,
        to: invitation.id,
      }),
    },
    sessions: {
      user: one.user({
        from: session.userId,
        to: user.id,
      }),
    },
    accounts: {
      user: one.user({
        from: account.userId,
        to: user.id,
      }),
    },
    organizations: {
      teams: many.team({ from: organization.id, to: team.organizationId }),
      members: many.member({ from: organization.id, to: member.organizationId }),
      invitations: many.invitation({ from: organization.id, to: invitation.organizationId }),
    },
    teams: {
      organization: one.organization({
        from: team.id,
        to: organization.id,
      }),
      teamMembers: many.teamMember({ from: team.id, to: teamMember.teamId }),
    },
    teamMembers: {
      team: one.team({
        from: teamMember.teamId,
        to: team.id,
      }),
      user: one.user({
        from: teamMember.userId,
        to: user.id,
      }),
    },
    member: {
      organization: one.organization({
        from: member.organizationId,
        to: organization.id,
      }),
      user: one.user({
        from: member.id,
        to: user.id,
      }),
    },
    invitations: {
      organization: one.organization({
        from: invitation.id,
        to: organization.id,
      }),
      user: one.user({
        from: invitation.id,
        to: user.id,
      }),
    },
  }),
)
