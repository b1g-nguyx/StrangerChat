
    -- Gỡ bỏ khóa ngoại trước để tránh lỗi dính líu (cascade)
    ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_current_room_id_fkey";
    ALTER TABLE "rooms" DROP CONSTRAINT IF EXISTS "rooms_user1_id_fkey";
    ALTER TABLE "rooms" DROP CONSTRAINT IF EXISTS "rooms_user2_id_fkey";
    ALTER TABLE "analytics_logs" DROP CONSTRAINT IF EXISTS "analytics_logs_user_id_fkey";
    ALTER TABLE "analytics_logs" DROP CONSTRAINT IF EXISTS "analytics_logs_room_id_fkey";

    -- Xóa các bảng theo thứ tự (Bảng nào có khóa ngoại thì xóa trước)
    DROP TABLE IF EXISTS "analytics_logs";
    DROP TABLE IF EXISTS "rooms";
    DROP TABLE IF EXISTS "users";
    DROP TABLE IF EXISTS "admins";