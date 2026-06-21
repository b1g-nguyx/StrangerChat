# Phân tích Tính năng (Functional Requirements)

Hệ thống được chia thành 2 nhóm tính năng cốt lõi phục vụ Client và Admin.

## 1. Nhóm tính năng Client (`svc-chat` - Cổng `:8080`)

### 1.1. Ghép cặp ngẫu nhiên (Matchmaking)
- **Mô tả:** Đưa user ẩn danh vào hàng đợi ngầm.
- **Công nghệ:** Sử dụng Go Channels để ghép cặp 2 người rảnh rỗi với tốc độ micro-giây.

### 1.2. Chat Text Real-time
- **Mô tả:** Luân chuyển tin nhắn văn bản giữa 2 user.
- **Công nghệ:** Giao thức WebSockets (`gorilla/websocket`).

### 1.3. Video Call Signaling
- **Mô tả:** Đóng vai trò trung gian truyền gói tin cấu hình mạng (SDP, ICE Candidates) giữa 2 user.
- **Công nghệ:** WebRTC. Trình duyệt tự thiết lập đường truyền video trực tiếp P2P sau khi signaling thành công qua WebSocket.

---

## 2. Nhóm tính năng Quản trị & AI (`svc-admin` - Cổng `:9090`)

### 2.1. AI Auto-Moderation (Tự động kiểm duyệt)
- **Mô tả:** AI chạy ngầm quét các mẻ log text để phát hiện ngôn từ độc hại, nhạy cảm, lừa đảo (`IsToxic = true`).
- **Hành động:** Khi phát hiện, `svc-admin` gọi gRPC sang `svc-chat` kích hoạt lệnh kick/ban user ngay lập tức.

### 2.2. AI Root Cause Analysis (Chẩn đoán lỗi)
- **Mô tả:** AI đọc log mạng thô (Packet Loss, Latency, Connection Drops) và đưa ra chẩn đoán bằng ngôn ngữ tự nhiên.
- **Ví dụ:** "Hệ thống gián đoạn do TURN Server mạng Viettel đang quá tải."

### 2.3. Admin Dashboard View
- **Mô tả:** Hiển thị biểu đồ Real-time Metrics.
- **Dữ liệu:** Tổng số user online, số phòng đang call, lượng RAM/CPU tiêu thụ (lấy từ package `runtime` của Go).
