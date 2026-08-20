CREATE TYPE "risk" AS ENUM('NO_RISK', 'LOW_RISK', 'MEDIUM_RISK', 'HIGH_RISK');--> statement-breakpoint
CREATE TABLE "service_lens_account" (
	"id" text PRIMARY KEY,
	"account_id" text NOT NULL,
	"provider_id" text NOT NULL,
	"user_id" text NOT NULL,
	"access_token" text,
	"refresh_token" text,
	"id_token" text,
	"access_token_expires_at" timestamp,
	"refresh_token_expires_at" timestamp,
	"scope" text,
	"password" text,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_invitation" (
	"id" text PRIMARY KEY,
	"organization_id" text NOT NULL,
	"email" text NOT NULL,
	"role" text,
	"team_id" text,
	"status" text DEFAULT 'pending' NOT NULL,
	"expires_at" timestamp NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"inviter_id" text NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_member" (
	"id" text PRIMARY KEY,
	"organization_id" text NOT NULL,
	"user_id" text NOT NULL,
	"role" text DEFAULT 'member' NOT NULL,
	"created_at" timestamp NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_organization" (
	"id" text PRIMARY KEY,
	"name" text NOT NULL,
	"slug" text NOT NULL UNIQUE,
	"logo" text,
	"created_at" timestamp NOT NULL,
	"metadata" text
);
--> statement-breakpoint
CREATE TABLE "service_lens_session" (
	"id" text PRIMARY KEY,
	"expires_at" timestamp NOT NULL,
	"token" text NOT NULL UNIQUE,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp NOT NULL,
	"ip_address" text,
	"user_agent" text,
	"user_id" text NOT NULL,
	"active_organization_id" text,
	"active_team_id" text
);
--> statement-breakpoint
CREATE TABLE "service_lens_team" (
	"id" text PRIMARY KEY,
	"name" text NOT NULL,
	"organization_id" text NOT NULL,
	"created_at" timestamp NOT NULL,
	"updated_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_team_member" (
	"id" text PRIMARY KEY,
	"team_id" text NOT NULL,
	"user_id" text NOT NULL,
	"created_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_user" (
	"id" text PRIMARY KEY,
	"name" text NOT NULL,
	"email" text NOT NULL UNIQUE,
	"email_verified" boolean DEFAULT false NOT NULL,
	"image" text,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_verification" (
	"id" text PRIMARY KEY,
	"identifier" text NOT NULL,
	"value" text NOT NULL,
	"expires_at" timestamp NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_design_tag" (
	"designId" uuid NOT NULL,
	"tagId" bigint NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_design" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"title" varchar(255) NOT NULL,
	"body" text,
	"description" varchar(1024),
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_environment_tag" (
	"environmentId" uuid NOT NULL,
	"tagId" bigint NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_environment" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"name" varchar(255) NOT NULL,
	"description" varchar(1024),
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"name" varchar(255) NOT NULL,
	"version" integer NOT NULL,
	"description" varchar(1024),
	"raw" json NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens_pillars" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"ref" varchar(255) NOT NULL,
	"name" varchar(255) NOT NULL,
	"description" varchar(1024) NOT NULL,
	"lensId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens_pillars_questions" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"ref" varchar(255) NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" varchar(1024) NOT NULL,
	"pillarId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens_pillars_questions_choices" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"ref" varchar(255) NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" varchar(1024) NOT NULL,
	"questionId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens_pillars_questions_risks" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"risk" "risk" DEFAULT 'NO_RISK'::"risk",
	"condition" varchar(1024) NOT NULL,
	"questionId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens_pillars_questions_resources" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"url" varchar(1024) NOT NULL,
	"description" varchar(1024),
	"questionId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile_question" (
	"id" bigint PRIMARY KEY,
	"question" varchar(1024) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile_question_answer" (
	"id" bigint PRIMARY KEY,
	"name" varchar(255) NOT NULL,
	"profileQuestionId" bigint NOT NULL,
	"answer" varchar(2048) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"name" varchar(255) NOT NULL,
	"description" varchar(1024),
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profiles_to_lenses" (
	"profileId" uuid,
	"lensId" uuid,
	CONSTRAINT "service_lens_profiles_to_lenses_pkey" PRIMARY KEY("profileId","lensId")
);
--> statement-breakpoint
CREATE TABLE "service_lens_tag" (
	"id" bigserial PRIMARY KEY,
	"name" varchar(255) NOT NULL,
	"value" varchar(1024) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_workload_environment" (
	"workloadId" uuid NOT NULL,
	"environmentId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_workload_lens" (
	"workloadId" uuid NOT NULL,
	"lensId" uuid NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_workload_profile" (
	"workloadId" uuid NOT NULL,
	"profileId" uuid NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_workload_tag" (
	"workloadId" uuid NOT NULL,
	"tagId" bigint NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_workload" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	"name" varchar(255) NOT NULL,
	"description" varchar(1024) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE INDEX "account_userId_idx" ON "service_lens_account" ("user_id");--> statement-breakpoint
CREATE INDEX "invitation_organizationId_idx" ON "service_lens_invitation" ("organization_id");--> statement-breakpoint
CREATE INDEX "invitation_email_idx" ON "service_lens_invitation" ("email");--> statement-breakpoint
CREATE INDEX "member_organizationId_idx" ON "service_lens_member" ("organization_id");--> statement-breakpoint
CREATE INDEX "member_userId_idx" ON "service_lens_member" ("user_id");--> statement-breakpoint
CREATE UNIQUE INDEX "organization_slug_uidx" ON "service_lens_organization" ("slug");--> statement-breakpoint
CREATE INDEX "session_userId_idx" ON "service_lens_session" ("user_id");--> statement-breakpoint
CREATE INDEX "team_organizationId_idx" ON "service_lens_team" ("organization_id");--> statement-breakpoint
CREATE INDEX "teamMember_teamId_idx" ON "service_lens_team_member" ("team_id");--> statement-breakpoint
CREATE INDEX "teamMember_userId_idx" ON "service_lens_team_member" ("user_id");--> statement-breakpoint
CREATE INDEX "verification_identifier_idx" ON "service_lens_verification" ("identifier");--> statement-breakpoint
CREATE INDEX "tag_name_index" ON "service_lens_tag" ("name");--> statement-breakpoint
CREATE UNIQUE INDEX "tag_name_value_unique_index" ON "service_lens_tag" ("name","value");--> statement-breakpoint
ALTER TABLE "service_lens_account" ADD CONSTRAINT "service_lens_account_user_id_service_lens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "service_lens_user"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_invitation" ADD CONSTRAINT "service_lens_invitation_xWJvyCK1dhEN_fkey" FOREIGN KEY ("organization_id") REFERENCES "service_lens_organization"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_invitation" ADD CONSTRAINT "service_lens_invitation_inviter_id_service_lens_user_id_fkey" FOREIGN KEY ("inviter_id") REFERENCES "service_lens_user"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_member" ADD CONSTRAINT "service_lens_member_2pkb2TI2IqU4_fkey" FOREIGN KEY ("organization_id") REFERENCES "service_lens_organization"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_member" ADD CONSTRAINT "service_lens_member_user_id_service_lens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "service_lens_user"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_session" ADD CONSTRAINT "service_lens_session_user_id_service_lens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "service_lens_user"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_team" ADD CONSTRAINT "service_lens_team_PO2vpkast7uz_fkey" FOREIGN KEY ("organization_id") REFERENCES "service_lens_organization"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_team_member" ADD CONSTRAINT "service_lens_team_member_team_id_service_lens_team_id_fkey" FOREIGN KEY ("team_id") REFERENCES "service_lens_team"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_team_member" ADD CONSTRAINT "service_lens_team_member_user_id_service_lens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "service_lens_user"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_design_tag" ADD CONSTRAINT "service_lens_design_tag_designId_service_lens_design_id_fkey" FOREIGN KEY ("designId") REFERENCES "service_lens_design"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_design_tag" ADD CONSTRAINT "service_lens_design_tag_tagId_service_lens_tag_id_fkey" FOREIGN KEY ("tagId") REFERENCES "service_lens_tag"("id");--> statement-breakpoint
ALTER TABLE "service_lens_environment_tag" ADD CONSTRAINT "service_lens_environment_tag_ezmRgStUuzEN_fkey" FOREIGN KEY ("environmentId") REFERENCES "service_lens_environment"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_environment_tag" ADD CONSTRAINT "service_lens_environment_tag_tagId_service_lens_tag_id_fkey" FOREIGN KEY ("tagId") REFERENCES "service_lens_tag"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_profile_question_answer" ADD CONSTRAINT "service_lens_profile_question_answer_jJkPKuu2RbkT_fkey" FOREIGN KEY ("profileQuestionId") REFERENCES "service_lens_profile_question"("id");--> statement-breakpoint
ALTER TABLE "service_lens_profiles_to_lenses" ADD CONSTRAINT "service_lens_profiles_to_lenses_ueHXsbJ97loY_fkey" FOREIGN KEY ("profileId") REFERENCES "service_lens_profile"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_profiles_to_lenses" ADD CONSTRAINT "service_lens_profiles_to_lenses_KnqSTDWFhKHo_fkey" FOREIGN KEY ("lensId") REFERENCES "service_lens_lens"("id");--> statement-breakpoint
ALTER TABLE "service_lens_workload_environment" ADD CONSTRAINT "service_lens_workload_environment_MSTYml5n7qWw_fkey" FOREIGN KEY ("workloadId") REFERENCES "service_lens_workload"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_workload_environment" ADD CONSTRAINT "service_lens_workload_environment_QWV2uS6DKg9W_fkey" FOREIGN KEY ("environmentId") REFERENCES "service_lens_environment"("id");--> statement-breakpoint
ALTER TABLE "service_lens_workload_lens" ADD CONSTRAINT "service_lens_workload_lens_Mryn4m9E7Cmh_fkey" FOREIGN KEY ("workloadId") REFERENCES "service_lens_workload"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_workload_lens" ADD CONSTRAINT "service_lens_workload_lens_lensId_service_lens_lens_id_fkey" FOREIGN KEY ("lensId") REFERENCES "service_lens_lens"("id");--> statement-breakpoint
ALTER TABLE "service_lens_workload_profile" ADD CONSTRAINT "service_lens_workload_profile_nDlpPEwHV7aQ_fkey" FOREIGN KEY ("workloadId") REFERENCES "service_lens_workload"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_workload_profile" ADD CONSTRAINT "service_lens_workload_profile_wr3rb1RfgR6U_fkey" FOREIGN KEY ("profileId") REFERENCES "service_lens_profile"("id");--> statement-breakpoint
ALTER TABLE "service_lens_workload_tag" ADD CONSTRAINT "service_lens_workload_tag_R0p9HWBHXaw3_fkey" FOREIGN KEY ("workloadId") REFERENCES "service_lens_workload"("id") ON DELETE CASCADE;--> statement-breakpoint
ALTER TABLE "service_lens_workload_tag" ADD CONSTRAINT "service_lens_workload_tag_tagId_service_lens_tag_id_fkey" FOREIGN KEY ("tagId") REFERENCES "service_lens_tag"("id") ON DELETE CASCADE;