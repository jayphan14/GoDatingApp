-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create "users" table
CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "university" varchar NOT NULL,
  "picture" bytea NOT NULL,
  "bio" text NOT NULL,
  "bio_pictures" text[] NOT NULL,
  "created_at" timestamptz DEFAULT (now()) NOT NULL
);

-- Create "message" table with corrected foreign key constraints
CREATE TABLE "message" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "senderID" uuid NOT NULL,
  "receiverID" uuid NOT NULL,
  "content" text,
  "created_at" timestamptz DEFAULT (now()),
  FOREIGN KEY ("senderID") REFERENCES "users" ("id"),
  FOREIGN KEY ("receiverID") REFERENCES "users" ("id")
);

-- Create "likes" table with corrected foreign key constraints
CREATE TABLE "likes" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "senderID" uuid NOT NULL,
  "receiverID" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  FOREIGN KEY ("senderID") REFERENCES "users" ("id"),
  FOREIGN KEY ("receiverID") REFERENCES "users" ("id")
);

-- Create "matches" table with corrected foreign key constraints
CREATE TABLE "matches" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user1ID" uuid NOT NULL,
  "user2ID" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  FOREIGN KEY ("user1ID") REFERENCES "users" ("id"),
  FOREIGN KEY ("user2ID") REFERENCES "users" ("id")
);
