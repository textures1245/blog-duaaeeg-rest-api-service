/*
  Warnings:

  - A unique constraint covering the columns `[userUuid,postUuid]` on the table `Like` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[postCategoryId]` on the table `Post` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[fileId]` on the table `Post` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[followerUuid,followeeUuid]` on the table `UserFollower` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE "UserProfile" DROP CONSTRAINT "UserProfile_userUuid_fkey";

-- AlterTable
ALTER TABLE "Post" ADD COLUMN     "fileId" INTEGER;

-- CreateTable
CREATE TABLE "File" (
    "id" SERIAL NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "fileName" TEXT NOT NULL,
    "fileType" TEXT NOT NULL,
    "fileUrl" TEXT NOT NULL,
    "postUuid" TEXT,

    CONSTRAINT "File_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "File_postUuid_key" ON "File"("postUuid");

-- CreateIndex
CREATE UNIQUE INDEX "Like_userUuid_postUuid_key" ON "Like"("userUuid", "postUuid");

-- CreateIndex
CREATE UNIQUE INDEX "Post_postCategoryId_key" ON "Post"("postCategoryId");

-- CreateIndex
CREATE UNIQUE INDEX "Post_fileId_key" ON "Post"("fileId");

-- CreateIndex
CREATE UNIQUE INDEX "UserFollower_followerUuid_followeeUuid_key" ON "UserFollower"("followerUuid", "followeeUuid");

-- AddForeignKey
ALTER TABLE "UserProfile" ADD CONSTRAINT "UserProfile_userUuid_fkey" FOREIGN KEY ("userUuid") REFERENCES "User"("uuid") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_fileId_fkey" FOREIGN KEY ("fileId") REFERENCES "File"("id") ON DELETE SET NULL ON UPDATE CASCADE;
