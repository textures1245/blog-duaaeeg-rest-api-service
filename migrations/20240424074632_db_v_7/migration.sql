/*
  Warnings:

  - You are about to drop the column `postUuid` on the `PostCategory` table. All the data in the column will be lost.
  - You are about to drop the column `postUuid` on the `PostTag` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[publishPostUuid]` on the table `Post` will be added. If there are existing duplicate values, this will fail.
  - Made the column `postCategoryId` on table `Post` required. This step will fail if there are existing NULL values in that column.
  - Made the column `postTagId` on table `Post` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "Post" DROP CONSTRAINT "Post_postCategoryId_fkey";

-- DropForeignKey
ALTER TABLE "Post" DROP CONSTRAINT "Post_postTagId_fkey";

-- DropForeignKey
ALTER TABLE "PublicationPost" DROP CONSTRAINT "PublicationPost_postUuid_fkey";

-- AlterTable
ALTER TABLE "Post" ADD COLUMN     "publishPostUuid" TEXT,
ALTER COLUMN "postCategoryId" SET NOT NULL,
ALTER COLUMN "postTagId" SET NOT NULL;

-- AlterTable
ALTER TABLE "PostCategory" DROP COLUMN "postUuid";

-- AlterTable
ALTER TABLE "PostTag" DROP COLUMN "postUuid",
ADD COLUMN     "postUUid" TEXT;

-- CreateIndex
CREATE UNIQUE INDEX "Post_publishPostUuid_key" ON "Post"("publishPostUuid");

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_postTagId_fkey" FOREIGN KEY ("postTagId") REFERENCES "PostTag"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_postCategoryId_fkey" FOREIGN KEY ("postCategoryId") REFERENCES "PostCategory"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PublicationPost" ADD CONSTRAINT "PublicationPost_postUuid_fkey" FOREIGN KEY ("postUuid") REFERENCES "Post"("uuid") ON DELETE CASCADE ON UPDATE CASCADE;
