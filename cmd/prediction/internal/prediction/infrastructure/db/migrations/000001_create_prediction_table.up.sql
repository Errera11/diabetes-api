CREATE TABLE IF NOT EXISTS prediction(
                                        id serial           PRIMARY KEY,
                                        chol_level          integer NOT NULL,
                                        diff_walk           boolean NOT NULL,
                                        phys_health          integer NOT NULL,
                                        birthdate           text NOT NULL,
                                        blood_pressure      float NOT NULL,
                                        weight              float NOT NULL,
                                        height              float NOT NULL,
                                        heart_disease       boolean NOT NULL,
                                        gen_health          integer NOT NULL,
                                        phys_activity        boolean NOT NULL,
                                        result              float[] NOT NULL,
                                        created_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_prediction(
                                         id serial           PRIMARY KEY,
                                         user_id             integer,
                                         prediction_id       integer,
                                         FOREIGN KEY (prediction_id) REFERENCES prediction(id) ON DELETE CASCADE,
                                         UNIQUE (user_id, prediction_id)
);