/*
  Warnings:

  - A unique constraint covering the columns `[postTagId]` on the table `Post` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE "PostCategory" DROP CONSTRAINT "PostCategory_postUuid_fkey";

-- DropForeignKey
ALTER TABLE "PostTag" DROP CONSTRAINT "PostTag_postUuid_fkey";

-- DropIndex
DROP INDEX "PostCategory_postUuid_key";

-- DropIndex
DROP INDEX "PostTag_postUuid_key";

-- AlterTable
ALTER TABLE "Post" ADD COLUMN     "postCategoryId" INTEGER,
ADD COLUMN     "postTagId" INTEGER;

-- CreateIndex
CREATE UNIQUE INDEX "Post_postTagId_key" ON "Post"("postTagId");

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_postTagId_fkey" FOREIGN KEY ("postTagId") REFERENCES "PostTag"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_postCategoryId_fkey" FOREIGN KEY ("postCategoryId") REFERENCES "PostCategory"("id") ON DELETE SET NULL ON UPDATE CASCADE;
