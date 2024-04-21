/*
  Warnings:

  - You are about to drop the column `contentUuid` on the `PublicationPost` table. All the data in the column will be lost.
  - You are about to drop the column `postUUid` on the `PublicationPost` table. All the data in the column will be lost.
  - You are about to drop the `UserContent` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `postUuid` to the `PublicationPost` table without a default value. This is not possible if the table is not empty.
  - Added the required column `userUuid` to the `PublicationPost` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "PublicationPost" DROP CONSTRAINT "PublicationPost_contentUuid_fkey";

-- DropForeignKey
ALTER TABLE "UserContent" DROP CONSTRAINT "UserContent_postUuid_fkey";

-- DropForeignKey
ALTER TABLE "UserContent" DROP CONSTRAINT "UserContent_userUuid_fkey";

-- AlterTable
ALTER TABLE "Post" ALTER COLUMN "published" DROP DEFAULT;

-- AlterTable
ALTER TABLE "PublicationPost" DROP COLUMN "contentUuid",
DROP COLUMN "postUUid",
ADD COLUMN     "postUuid" TEXT NOT NULL,
ADD COLUMN     "userUuid" TEXT NOT NULL;

-- DropTable
DROP TABLE "UserContent";

-- AddForeignKey
ALTER TABLE "PublicationPost" ADD CONSTRAINT "PublicationPost_userUuid_fkey" FOREIGN KEY ("userUuid") REFERENCES "User"("uuid") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PublicationPost" ADD CONSTRAINT "PublicationPost_postUuid_fkey" FOREIGN KEY ("postUuid") REFERENCES "Post"("uuid") ON DELETE RESTRICT ON UPDATE CASCADE;
