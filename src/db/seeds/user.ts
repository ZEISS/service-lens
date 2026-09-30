import { generateId } from "better-auth"
import { hashPassword } from "better-auth/crypto"

import { db } from "@/db/index"
import { account, type TNewAccount, type TNewUser, user } from "@/db/schema"

export async function seedUser() {
  try {
    const userId = generateId()
    const rootId = generateId()

    const root: TNewUser = {
      id: userId,
      name: "Indy Jones",
      email: "indy@jones.com",
      emailVerified: false,
    }

    const rootAccount: TNewAccount = {
      id: rootId,
      accountId: userId,
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
