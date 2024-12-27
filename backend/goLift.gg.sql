CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "email" varchar UNIQUE,
  "password_hash" varchar,
  "role" varchar,
  "specializations" text,
  "self_guided" boolean DEFAULT false,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "last_login" timestamp,
  "is_active" boolean DEFAULT true,
  "reset_password_token" varchar,
  "reset_password_expires" timestamp
);

CREATE TABLE "program_templates" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "description" text,
  "coach_id" int,
  "days_per_week" int,
  "is_public" boolean DEFAULT false,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "programs" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "description" text,
  "coach_id" int,
  "athlete_id" int,
  "days_per_week" int,
  "template_id" int,
  "start_date" timestamp,
  "end_date" timestamp,
  "status" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "workouts" (
  "id" int PRIMARY KEY,
  "program_id" int,
  "date" timestamp,
  "athlete_id" int,
  "coach_id" int,
  "details" text,
  "order" int,
  "status" varchar,
  "actual_date" timestamp,
  "duration_minutes" int,
  "rating" int,
  "notes" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "exercises" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "category" varchar,
  "description" text,
  "type" varchar,
  "equipment" varchar,
  "difficulty" varchar,
  "created_by" int,
  "is_public" boolean DEFAULT true,
  "video_url" varchar,
  "substitutes" text,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "athlete_maxes" (
  "id" int PRIMARY KEY,
  "athlete_id" int,
  "exercise_id" int,
  "max_weight" float,
  "date" timestamp,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "workout_exercises" (
  "id" int PRIMARY KEY,
  "workout_id" int,
  "exercise_id" int,
  "order" int,
  "sets" int,
  "reps" int,
  "prescribed_percentage" float,
  "prescribed_rpe" float,
  "prescribed_weight" float,
  "programming_method" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "progress" (
  "id" int PRIMARY KEY,
  "athlete_id" int,
  "workout_id" int,
  "exercise_id" int,
  "prescribed_weight" float,
  "actual_weight" float,
  "reps" int,
  "date" timestamp,
  "status" varchar,
  "comment" text,
  "prescribed_percentage" float,
  "prescribed_rpe" float,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN "users"."role" IS 'coach or athlete';

COMMENT ON COLUMN "programs"."days_per_week" IS 'Between 1 and 7';

COMMENT ON COLUMN "workouts"."order" IS 'Order of the workout in the program, used for chronological sorting';

COMMENT ON COLUMN "exercises"."category" IS 'Olympic lift, Accessory, Strength, Cardio, etc.';

COMMENT ON COLUMN "exercises"."type" IS 'Strength, Mobility, Cardio, etc.';

COMMENT ON COLUMN "exercises"."equipment" IS 'Equipment required for the exercise';

COMMENT ON COLUMN "exercises"."difficulty" IS 'Difficulty level of the exercise';

COMMENT ON COLUMN "athlete_maxes"."max_weight" IS 'The athlete''s maximum weight for this exercise';

COMMENT ON COLUMN "athlete_maxes"."date" IS 'The date when the max weight was achieved';

COMMENT ON COLUMN "workout_exercises"."order" IS 'Order of exercises in the workout';

COMMENT ON COLUMN "workout_exercises"."prescribed_percentage" IS 'The prescribed percentage (if programming method is percentage_of_max)';

COMMENT ON COLUMN "workout_exercises"."prescribed_rpe" IS 'The prescribed RPE (if programming method is RPE)';

COMMENT ON COLUMN "workout_exercises"."prescribed_weight" IS 'The prescribed weight (if programming method is absolute_weight)';

COMMENT ON COLUMN "workout_exercises"."programming_method" IS 'absolute_weight, percentage_of_max, or RPE';

COMMENT ON COLUMN "progress"."prescribed_weight" IS 'The weight prescribed for the exercise';

COMMENT ON COLUMN "progress"."actual_weight" IS 'The weight actually lifted by the athlete';

COMMENT ON COLUMN "progress"."status" IS 'Goal met, near goal, goal missed';

COMMENT ON COLUMN "progress"."prescribed_percentage" IS 'Percentage prescribed if method is percentage_of_max';

COMMENT ON COLUMN "progress"."prescribed_rpe" IS 'RPE prescribed if method is RPE';

ALTER TABLE "programs" ADD FOREIGN KEY ("coach_id") REFERENCES "users" ("id");
ALTER TABLE "programs" ADD FOREIGN KEY ("athlete_id") REFERENCES "users" ("id");
ALTER TABLE "programs" ADD FOREIGN KEY ("template_id") REFERENCES "program_templates" ("id");
ALTER TABLE "workouts" ADD FOREIGN KEY ("program_id") REFERENCES "programs" ("id");
ALTER TABLE "workouts" ADD FOREIGN KEY ("athlete_id") REFERENCES "users" ("id");
ALTER TABLE "workouts" ADD FOREIGN KEY ("coach_id") REFERENCES "users" ("id");
ALTER TABLE "exercises" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");
ALTER TABLE "program_templates" ADD FOREIGN KEY ("coach_id") REFERENCES "users" ("id");

ALTER TABLE "athlete_maxes" ADD FOREIGN KEY ("athlete_id") REFERENCES "users" ("id");
ALTER TABLE "athlete_maxes" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercises" ("id");
ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");
ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercises" ("id");

ALTER TABLE "programs" ADD CONSTRAINT "check_days_per_week" CHECK (days_per_week BETWEEN 1 AND 7);
ALTER TABLE "workout_exercises" ADD CONSTRAINT "check_prescribed_rpe" CHECK (prescribed_rpe BETWEEN 1 AND 10);
ALTER TABLE "workouts" ADD CONSTRAINT "check_workout_status" CHECK (status IN ('planned', 'in_progress', 'completed', 'skipped'));
ALTER TABLE "programs" ADD CONSTRAINT "check_program_status" CHECK (status IN ('active', 'completed', 'paused'));
ALTER TABLE "workout_exercises" ADD CONSTRAINT "check_programming_method" CHECK (programming_method IN ('absolute_weight', 'percentage_of_max', 'RPE'));

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_workouts_date ON workouts(date);
CREATE INDEX idx_workout_exercises_workout_id ON workout_exercises(workout_id);
CREATE INDEX idx_progress_athlete_date ON progress(athlete_id, date);
CREATE INDEX idx_programs_coach_id ON programs(coach_id);
CREATE INDEX idx_programs_athlete_id ON programs(athlete_id);
