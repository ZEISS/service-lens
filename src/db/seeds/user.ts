import { db } from "@/db/index"
import { type NewAccount, type NewUser, account, user } from "@/db/schema"
import { generateId } from "better-auth"
import { hashPassword } from "better-auth/crypto"

export async function seedUser() {
  try {
    const userId = generateId()
    const accountId = generateId()
    const rootId = generateId()

    const root: NewUser = {
      id: userId,
      name: "Indy Jones",
      email: "indy@jones.com",
      emailVerified: true,
    }

    const rootAccount: NewAccount = {
      id: rootId,
      accountId: accountId,
      userId: root.id,
      providerId: "credential",
      password: await hashPassword("password123"),
    }

    console.log("📝 Inserting user")

    await db.insert(user).values(root).onConflictDoNothing()
    await db.insert(account).values(rootAccount).onConflictDoNothing()
  } catch (err) {
    console.error(err)
  }
}
