/*
  Warnings:

  - You are about to drop the `Tag` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Tag" DROP CONSTRAINT "Tag_postTagPostUuid_fkey";

-- AlterTable
ALTER TABLE "PostTag" ADD COLUMN     "tags" TEXT[];

-- DropTable
DROP TABLE "Tag";
