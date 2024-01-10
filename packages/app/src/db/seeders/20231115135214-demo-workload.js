'use strict'

const crypto = require('crypto')

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    const profileId = crypto.randomUUID()
    await queryInterface.bulkInsert('profiles', [
      {
        id: profileId,
        name: 'demo',
        description: 'This is an initial demo workload'
      }
    ])

    const profileQuestion = await queryInterface.bulkInsert(
      'profiles-questions',
      [
        {
          ref: 'cloud_adoption_phase',
          name: 'What is the current cloud adoption phase for the organization architecting or operating the workloads in this profile?',
          description: 'This is asking about the phase this application is in.'
        },
        {
          ref: 'improvment_priorities',
          name: 'What are improvment categories for this workload in the Well-Architected Review',
          description: 'What needs to be prioritized when evaluating?',
          isMultiple: true
        }
      ]
    )

    const profileQuestionChoices = await queryInterface.bulkInsert(
      'profiles-questions-choices',
      [
        {
          questionId: 1,
          ref: 'envision_adopt_phase',
          name: 'Envision Adoption Phase',
          description:
            'The Envision phase focuses on demonstrating how the cloud will help accelerate your business outcomes. It does so by identifying and prioritizing transformation opportunities across each of the four transformation domains in line with your strategic business objectives. Associating your transformation initiatives with key stakeholders (senior individuals capable of influencing and driving change) and measurable business outcomes will help you demonstrate value as you progress through your transformation journey.'
        },
        {
          questionId: 1,
          ref: 'align_adopt_phase',
          name: 'Align Adoption Phase',
          description:
            'The Align phase focuses on identifying capability gaps across the key perspectives of business, people, governance, platform, security, and operations. In this phase, the goal is to identify cross-organizational dependencies, and surface stakeholder concerns and challenges. Doing so will help you create strategies for improving your cloud readiness, ensure stakeholder alignment, and facilitate relevant organizational change management activities.'
        },
        {
          questionId: 1,
          ref: 'launch_adopt_phase',
          name: 'Launch Adoption Phase',
          description:
            'The Launch phase focuses on delivering pilot initiatives in production and on demonstrating incremental business value. Pilots should be highly impactful and can help influence your future direction. Learning from pilots helps you to adjust your approach before scaling to full production.'
        },
        {
          questionId: 2,
          ref: 'eval_org_cloud_strategy',
          name: 'Evaluate organization cloud strategy and priorities',
          description: ''
        },
        {
          questionId: 2,
          ref: 'improv_op_readiness',
          name: 'Improve operational readiness',
          description: ''
        },
        {
          questionId: 2,
          ref: 'improv_op_eff',
          name: 'Improve operational efficiency',
          description: ''
        }
      ]
    )

    await queryInterface.bulkInsert('profiles-questions-answers', [
      {
        choiceId: 2,
        profileId: profileId
      }
    ])

    const lensId = crypto.randomUUID()
    await queryInterface.bulkInsert('lenses', [
      {
        id: lensId,
        name: 'Web Application Security Lens',
        version: 1,
        description: 'This is an initial demo lens',
        isDraft: false,
        spec: '{"version":1,"name":"SAP Lens","description":"SAP Lens","pillars":[{"id":"operational_excellence","name":"Operational Excellence","description":"Operational Excellence","questions":[{"id":"question_1","name":"Question 1","description":"Question 1","choices":[{"id":"choice_1","name":"Choice 1","description":"Choice 1"}],"risks":[{"risk":"HIGH","condition":"default"}]}]}]}'
      },
      {
        id: crypto.randomUUID(),
        name: 'SAP Workload',
        version: 1,
        description: 'This is an initial demo lens',
        isDraft: true,
        spec: '{"version":1,"name":"SAP Lens","description":"SAP Lens","pillars":[{"id":"operational_excellence","name":"Operational Excellence","description":"Operational Excellence","questions":[{"id":"question_1","name":"Question 1","description":"Question 1","choices":[{"id":"choice_1","name":"Choice 1","description":"Choice 1"}],"risks":[{"risk":"HIGH","condition":"default"}]}]}]}'
      }
    ])

    await queryInterface.bulkInsert('environments', [
      {
        name: 'production',
        description: 'Production environment'
      },
      {
        name: 'staging',
        description: 'Staging environment'
      },
      {
        name: 'development',
        description: 'Development environment'
      }
    ])

    const workloadId = crypto.randomUUID()
    await queryInterface.bulkInsert('workloads', [
      {
        id: workloadId,
        name: 'SAP Workload',
        description: 'SAP Workload',
        profilesId: profileId
      }
    ])

    await queryInterface.bulkInsert('workloads-environment', [
      {
        workloadId: workloadId,
        environmentId: 1
      }
    ])

    await queryInterface.bulkInsert('workloads-lenses', [
      {
        lensId: lensId,
        workloadId: workloadId
      }
    ])

    const solutionId = crypto.randomUUID()
    await queryInterface.bulkInsert('solutions', [
      {
        id: solutionId,
        title: 'Example Solution',
        body: 'Example Solution'
      }
    ])

    await queryInterface.bulkInsert('solutions-comments', [
      {
        body: 'This is a new comment',
        solutionId
      }
    ])

    await queryInterface.bulkInsert('solutions-templates', [
      {
        title: 'Architecture Decision Record',
        description: 'Writing architectural decision records.',
        body: `# Decision record template by Michael Nygard

This is the template in [Documenting architecture decisions - Michael Nygard](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions).
You can use [adr-tools](https://github.com/npryce/adr-tools) for managing the ADR files.

In each ADR file, write these sections:

# Title

## Status

What is the status, such as proposed, accepted, rejected, deprecated, superseded, etc.?

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?`
      }
    ])
  },

  async down(queryInterface, Sequelize) {
    await queryInterface.bulkDelete('workloads', null, {})
    await queryInterface.bulkDelete('profiles', null, {})
    await queryInterface.bulkDelete('lenses', null, {})
  }
}
