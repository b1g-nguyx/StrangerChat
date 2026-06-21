-- Init extension để auto generate UUID in PostgreSQL
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE "admins" (
      "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),                                                                                                          
      "username" varchar UNIQUE NOT NULL,                                                                                                                             
      "password_hash" varchar NOT NULL,                                                                                                                               
      "role" varchar NOT NULL,                                                                                                                                        
      "refresh_token" varchar,                                                                                                                                        
      "created_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "updated_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "deleted_at" timestamptz                                                                                                                                        
    );                                                                                                                                                                
                                                                                                                                                                      
    CREATE TABLE "users" (                                                                                                                                            
      "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
      "username" varchar UNIQUE NOT NULL,
      "email" varchar UNIQUE NOT NULL,
      "password_hash" varchar NOT NULL,
      "display_name" varchar,
      "avatar_url" varchar,
      "status" varchar NOT NULL DEFAULT 'Waiting',                                                                                                                    
      "current_room_id" varchar,                                                                                                                                      
      "refresh_token" varchar,                                                                                                                                        
      "is_banned" boolean NOT NULL DEFAULT false,                                                                                                                     
      "banned_at" timestamptz,                                                                                                                                        
      "created_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "updated_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "deleted_at" timestamptz                                                                                                                                        
    );                                                                                                                                                                
                                                                                                                                                                      
    CREATE TABLE "rooms" (                                                                                                                                            
      "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),                                                                                                          
      "user1_id" varchar,                                                                                                                                             
      "user2_id" varchar,                                                                                                                                             
      "status" varchar NOT NULL DEFAULT 'Active',                                                                                                                     
      "closed_at" timestamptz,                                                                                                                                        
      "created_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "updated_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "deleted_at" timestamptz                                                                                                                                        
    );                                                                                                                                                                
                                                                                                                                                                      
    CREATE TABLE "analytics_logs" (                                                                                                                                   
      "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),                                                                                                          
      "log_type" varchar NOT NULL,                                                                                                                                    
      "user_id" varchar,                                                                                                                                              
      "room_id" varchar,                                                                                                                                              
      "content" text NOT NULL,                                                                                                                                        
      "is_toxic" boolean NOT NULL DEFAULT false,                                                                                                                      
      "ai_diagnostic" text,                                                                                                                                           
      "created_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "updated_at" timestamptz NOT NULL DEFAULT (now()),                                                                                                              
      "deleted_at" timestamptz                                                                                                                                        
    );                                                                                                                                                                
                                                                                                                                                                      
    -- (Foreign Keys)                                                                                                                             
    ALTER TABLE "users" ADD FOREIGN KEY ("current_room_id") REFERENCES "rooms" ("id");                                                                                
    ALTER TABLE "rooms" ADD FOREIGN KEY ("user1_id") REFERENCES "users" ("id");                                                                                       
    ALTER TABLE "rooms" ADD FOREIGN KEY ("user2_id") REFERENCES "users" ("id");                                                                                       
    ALTER TABLE "analytics_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");                                                                               
    ALTER TABLE "analytics_logs" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");