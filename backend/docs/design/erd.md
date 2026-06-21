# Entity Relationship Diagram (ERD)

Dưới đây là sơ đồ thực thể mối quan hệ cho hệ thống Adaptive StrangerChat:

```mermaid
erDiagram
    Admin {
        string ID PK "UUID"
        string Username "Unique"
        string PasswordHash "Mật khẩu mã hoá"
        string Role "Ví dụ: super_admin, moderator"
        string RefreshToken "Nullable"
        datetime CreatedAt
    }

    User {
        string ID PK "UUID"
        string Status "Ví dụ: Waiting, InCall, Offline"
        string CurrentRoomID FK "Nullable"
        string RefreshToken "Nullable"
        boolean IsBanned
        datetime BannedAt "Nullable"
        datetime CreatedAt
    }

    Room {
        string ID PK "UUID"
        string User1ID FK
        string User2ID FK
        string Status "Ví dụ: Active, Closed"
        datetime CreatedAt
        datetime ClosedAt "Nullable"
    }

    AnalyticsLog {
        string ID PK "UUID"
        string LogType "ChatText, NetworkMetric, System"
        string UserID FK "Nullable"
        string RoomID FK "Nullable"
        string Content "Nội dung chat hoặc chỉ số mạng JSON"
        boolean IsToxic "AI đánh giá (Auto-Moderation)"
        string AIDiagnostic "AI chẩn đoán lỗi bằng ngôn ngữ tự nhiên"
        datetime Timestamp
    }

    User ||--o| Room : "joins"
    Room ||--o{ AnalyticsLog : "generates"
    User ||--o{ AnalyticsLog : "generates"
```

## Các thực thể chính:
1. **Admin**: Đại diện cho quản trị viên đăng nhập vào Dashboard (gắn với `svc-admin`). Có Username/Password riêng biệt với người dùng ẩn danh.
2. **User**: Đại diện cho client đang kết nối ẩn danh (gắn với `svc-chat`). Lưu trạng thái để ghép cặp và chống spam/ban.
3. **Room**: Phòng chat P2P sinh ra khi Matchmaking thành công giữa 2 User.
4. **AnalyticsLog**: Nguồn dữ liệu log stream bất đồng bộ sinh ra từ chat và network. AI sẽ dựa vào đây để đánh dấu `IsToxic` (kích hoạt kick/ban) hoặc phân tích `AIDiagnostic` (Root Cause Analysis).
