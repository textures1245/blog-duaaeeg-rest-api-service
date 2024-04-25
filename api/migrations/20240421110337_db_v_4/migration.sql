/*
  Warnings:

  - A unique constraint covering the columns `[postUuid]` on the table `PublicationPost` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "PublicationPost_postUuid_key" ON "PublicationPost"("postUuid");
