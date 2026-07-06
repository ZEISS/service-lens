CREATE TABLE "service_lens_account" (
	"id" text PRIMARY KEY NOT NULL,
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
	"id" text PRIMARY KEY NOT NULL,
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
	"id" text PRIMARY KEY NOT NULL,
	"organization_id" text NOT NULL,
	"user_id" text NOT NULL,
	"role" text DEFAULT 'member' NOT NULL,
	"created_at" timestamp NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_organization" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"slug" text NOT NULL,
	"logo" text,
	"created_at" timestamp NOT NULL,
	"metadata" text,
	CONSTRAINT "service_lens_organization_slug_unique" UNIQUE("slug")
);
--> statement-breakpoint
CREATE TABLE "service_lens_session" (
	"id" text PRIMARY KEY NOT NULL,
	"expires_at" timestamp NOT NULL,
	"token" text NOT NULL,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp NOT NULL,
	"ip_address" text,
	"user_agent" text,
	"user_id" text NOT NULL,
	"active_organization_id" text,
	"active_team_id" text,
	CONSTRAINT "service_lens_session_token_unique" UNIQUE("token")
);
--> statement-breakpoint
CREATE TABLE "service_lens_team" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"organization_id" text NOT NULL,
	"created_at" timestamp NOT NULL,
	"updated_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_team_member" (
	"id" text PRIMARY KEY NOT NULL,
	"team_id" text NOT NULL,
	"user_id" text NOT NULL,
	"created_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_user" (
	"id" text PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"email" text NOT NULL,
	"email_verified" boolean DEFAULT false NOT NULL,
	"image" text,
	"created_at" timestamp DEFAULT now() NOT NULL,
	"updated_at" timestamp DEFAULT now() NOT NULL,
	CONSTRAINT "service_lens_user_email_unique" UNIQUE("email")
);
--> statement-breakpoint
CREATE TABLE "service_lens_verification" (
	"id" text PRIMARY KEY NOT NULL,
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
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
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
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(255) NOT NULL,
	"description" varchar(1024),
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_lens" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(255) NOT NULL,
	"version" integer NOT NULL,
	"description" varchar(1024),
	"raw" json NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile_lens" (
	"profileId" uuid NOT NULL,
	"lensId" uuid NOT NULL
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile_question" (
	"id" bigint PRIMARY KEY NOT NULL,
	"question" varchar(1024) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile_question_answer" (
	"id" bigint PRIMARY KEY NOT NULL,
	"name" varchar(255) NOT NULL,
	"profileQuestionId" bigint NOT NULL,
	"answer" varchar(2048) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_profile" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(255) NOT NULL,
	"description" varchar(1024),
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "service_lens_tag" (
	"id" bigserial PRIMARY KEY NOT NULL,
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
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar(255) NOT NULL,
	"description" varchar(1024) NOT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now(),
	"deleted_at" timestamp
);
--> statement-breakpoint
ALTER TABLE "service_lens_account" ADD CONSTRAINT "service_lens_account_user_id_service_lens_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."service_lens_user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_invitation" ADD CONSTRAINT "service_lens_invitation_organization_id_service_lens_organization_id_fk" FOREIGN KEY ("organization_id") REFERENCES "public"."service_lens_organization"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_invitation" ADD CONSTRAINT "service_lens_invitation_inviter_id_service_lens_user_id_fk" FOREIGN KEY ("inviter_id") REFERENCES "public"."service_lens_user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_member" ADD CONSTRAINT "service_lens_member_organization_id_service_lens_organization_id_fk" FOREIGN KEY ("organization_id") REFERENCES "public"."service_lens_organization"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_member" ADD CONSTRAINT "service_lens_member_user_id_service_lens_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."service_lens_user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_session" ADD CONSTRAINT "service_lens_session_user_id_service_lens_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."service_lens_user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_team" ADD CONSTRAINT "service_lens_team_organization_id_service_lens_organization_id_fk" FOREIGN KEY ("organization_id") REFERENCES "public"."service_lens_organization"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_team_member" ADD CONSTRAINT "service_lens_team_member_team_id_service_lens_team_id_fk" FOREIGN KEY ("team_id") REFERENCES "public"."service_lens_team"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_team_member" ADD CONSTRAINT "service_lens_team_member_user_id_service_lens_user_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."service_lens_user"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_design_tag" ADD CONSTRAINT "service_lens_design_tag_designId_service_lens_design_id_fk" FOREIGN KEY ("designId") REFERENCES "public"."service_lens_design"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_design_tag" ADD CONSTRAINT "service_lens_design_tag_tagId_service_lens_tag_id_fk" FOREIGN KEY ("tagId") REFERENCES "public"."service_lens_tag"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_environment_tag" ADD CONSTRAINT "service_lens_environment_tag_environmentId_service_lens_environment_id_fk" FOREIGN KEY ("environmentId") REFERENCES "public"."service_lens_environment"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_environment_tag" ADD CONSTRAINT "service_lens_environment_tag_tagId_service_lens_tag_id_fk" FOREIGN KEY ("tagId") REFERENCES "public"."service_lens_tag"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_profile_lens" ADD CONSTRAINT "service_lens_profile_lens_profileId_service_lens_profile_id_fk" FOREIGN KEY ("profileId") REFERENCES "public"."service_lens_profile"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_profile_lens" ADD CONSTRAINT "service_lens_profile_lens_lensId_service_lens_lens_id_fk" FOREIGN KEY ("lensId") REFERENCES "public"."service_lens_lens"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_profile_question_answer" ADD CONSTRAINT "service_lens_profile_question_answer_profileQuestionId_service_lens_profile_question_id_fk" FOREIGN KEY ("profileQuestionId") REFERENCES "public"."service_lens_profile_question"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_environment" ADD CONSTRAINT "service_lens_workload_environment_workloadId_service_lens_workload_id_fk" FOREIGN KEY ("workloadId") REFERENCES "public"."service_lens_workload"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_environment" ADD CONSTRAINT "service_lens_workload_environment_environmentId_service_lens_environment_id_fk" FOREIGN KEY ("environmentId") REFERENCES "public"."service_lens_environment"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_lens" ADD CONSTRAINT "service_lens_workload_lens_workloadId_service_lens_workload_id_fk" FOREIGN KEY ("workloadId") REFERENCES "public"."service_lens_workload"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_lens" ADD CONSTRAINT "service_lens_workload_lens_lensId_service_lens_lens_id_fk" FOREIGN KEY ("lensId") REFERENCES "public"."service_lens_lens"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_profile" ADD CONSTRAINT "service_lens_workload_profile_workloadId_service_lens_workload_id_fk" FOREIGN KEY ("workloadId") REFERENCES "public"."service_lens_workload"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_profile" ADD CONSTRAINT "service_lens_workload_profile_profileId_service_lens_profile_id_fk" FOREIGN KEY ("profileId") REFERENCES "public"."service_lens_profile"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_tag" ADD CONSTRAINT "service_lens_workload_tag_workloadId_service_lens_workload_id_fk" FOREIGN KEY ("workloadId") REFERENCES "public"."service_lens_workload"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "service_lens_workload_tag" ADD CONSTRAINT "service_lens_workload_tag_tagId_service_lens_tag_id_fk" FOREIGN KEY ("tagId") REFERENCES "public"."service_lens_tag"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
CREATE INDEX "account_userId_idx" ON "service_lens_account" USING btree ("user_id");--> statement-breakpoint
CREATE INDEX "invitation_organizationId_idx" ON "service_lens_invitation" USING btree ("organization_id");--> statement-breakpoint
CREATE INDEX "invitation_email_idx" ON "service_lens_invitation" USING btree ("email");--> statement-breakpoint
CREATE INDEX "member_organizationId_idx" ON "service_lens_member" USING btree ("organization_id");--> statement-breakpoint
CREATE INDEX "member_userId_idx" ON "service_lens_member" USING btree ("user_id");--> statement-breakpoint
CREATE UNIQUE INDEX "organization_slug_uidx" ON "service_lens_organization" USING btree ("slug");--> statement-breakpoint
CREATE INDEX "session_userId_idx" ON "service_lens_session" USING btree ("user_id");--> statement-breakpoint
CREATE INDEX "team_organizationId_idx" ON "service_lens_team" USING btree ("organization_id");--> statement-breakpoint
CREATE INDEX "teamMember_teamId_idx" ON "service_lens_team_member" USING btree ("team_id");--> statement-breakpoint
CREATE INDEX "teamMember_userId_idx" ON "service_lens_team_member" USING btree ("user_id");--> statement-breakpoint
CREATE INDEX "verification_identifier_idx" ON "service_lens_verification" USING btree ("identifier");--> statement-breakpoint
CREATE INDEX "tag_name_index" ON "service_lens_tag" USING btree ("name");--> statement-breakpoint
CREATE UNIQUE INDEX "tag_name_value_unique_index" ON "service_lens_tag" USING btree ("name","value");