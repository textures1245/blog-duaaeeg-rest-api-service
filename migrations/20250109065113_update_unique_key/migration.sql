/*
  Warnings:

  - You are about to drop the column `postId` on the `PostTag` table. All the data in the column will be lost.
  - You are about to drop the `_PostToPostCategory` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "PostTag" DROP CONSTRAINT "PostTag_postId_fkey";

-- DropForeignKey
ALTER TABLE "_PostToPostCategory" DROP CONSTRAINT "_PostToPostCategory_A_fkey";

-- DropForeignKey
ALTER TABLE "_PostToPostCategory" DROP CONSTRAINT "_PostToPostCategory_B_fkey";

-- AlterTable
ALTER TABLE "PostTag" DROP COLUMN "postId";

-- DropTable
DROP TABLE "_PostToPostCategory";

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_postTagId_fkey" FOREIGN KEY ("postTagId") REFERENCES "PostTag"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Post" ADD CONSTRAINT "Post_postCategoryId_fkey" FOREIGN KEY ("postCategoryId") REFERENCES "PostCategory"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
