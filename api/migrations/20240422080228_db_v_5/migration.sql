/*
  Warnings:

  - The primary key for the `PostCategory` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `categoryId` on the `PostCategory` table. All the data in the column will be lost.
  - The primary key for the `PostTag` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `tagId` on the `PostTag` table. All the data in the column will be lost.
  - You are about to drop the `Category` table. If the table is not empty, all the data it contains will be lost.
  - A unique constraint covering the columns `[postUuid]` on the table `PostCategory` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[name]` on the table `PostCategory` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[postUuid]` on the table `PostTag` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `name` to the `PostCategory` table without a default value. This is not possible if the table is not empty.
  - Added the required column `postTagPostUuid` to the `Tag` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "PostCategory" DROP CONSTRAINT "PostCategory_categoryId_fkey";

-- DropForeignKey
ALTER TABLE "PostTag" DROP CONSTRAINT "PostTag_tagId_fkey";

-- AlterTable
ALTER TABLE "PostCategory" DROP CONSTRAINT "PostCategory_pkey",
DROP COLUMN "categoryId",
ADD COLUMN     "id" SERIAL NOT NULL,
ADD COLUMN     "name" TEXT NOT NULL,
ADD CONSTRAINT "PostCategory_pkey" PRIMARY KEY ("postUuid", "id");

-- AlterTable
ALTER TABLE "PostTag" DROP CONSTRAINT "PostTag_pkey",
DROP COLUMN "tagId",
ADD COLUMN     "id" SERIAL NOT NULL,
ADD CONSTRAINT "PostTag_pkey" PRIMARY KEY ("postUuid", "id");

-- AlterTable
ALTER TABLE "Tag" ADD COLUMN     "postTagPostUuid" TEXT NOT NULL;

-- DropTable
DROP TABLE "Category";

-- CreateIndex
CREATE UNIQUE INDEX "PostCategory_postUuid_key" ON "PostCategory"("postUuid");

-- CreateIndex
CREATE UNIQUE INDEX "PostCategory_name_key" ON "PostCategory"("name");

-- CreateIndex
CREATE UNIQUE INDEX "PostTag_postUuid_key" ON "PostTag"("postUuid");

-- AddForeignKey
ALTER TABLE "Tag" ADD CONSTRAINT "Tag_postTagPostUuid_fkey" FOREIGN KEY ("postTagPostUuid") REFERENCES "PostTag"("postUuid") ON DELETE RESTRICT ON UPDATE CASCADE;
