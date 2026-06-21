# Adaptive StrangerChat & AI-Powered Log Analytics Engine

## 1. Giới thiệu
Hệ thống Adaptive StrangerChat là một ứng dụng chat ẩn danh ngẫu nhiên (như Omegle) tích hợp phân tích log và tự động hóa quản trị bằng AI. Hệ thống sử dụng kiến trúc Monorepo phân tách thành 2 service chính chạy độc lập nhưng chung lõi `internal/`.

## 2. Kiến trúc Monorepo Đa Dịch Vụ
Hệ thống được chia làm 2 service độc lập nhằm tối ưu hiệu năng và tách biệt rủi ro:

### 2.1. `svc-chat` (Port `:8080` - Public)
- **Vai trò:** Xử lý kết nối hàng vạn user cùng lúc (C10K).
- **Lý do dùng Go:** Tối ưu RAM bằng Goroutines (~2KB/kết nối so với ~1MB của thread truyền thống).
- **Tính năng nổi bật:** Điều phối mạng thích ứng (Adaptive Signaling). Go liên tục lắng nghe chỉ số mạng của cặp chat. Nếu mạng yếu, hệ thống tự động hạ chất lượng video của đối diện để duy trì cuộc gọi.

### 2.2. `svc-admin` (Port `:9090` - Private)
- **Vai trò:** Quản lý log, gọi AI, thao tác Database nặng.
- **Lý do dùng Go:** Sử dụng Go Channels làm bộ đệm RAM bất đồng bộ để gom mẻ log (Batching), sau đó Bulk Insert xuống PostgreSQL, tránh nghẽn cổ chai I/O Database.
- **Bảo mật:** Tách riêng cổng giúp ẩn các tác vụ nặng khỏi môi trường public, bảo vệ API key của AI (Gemini/OpenAI).

## 3. Data Pipeline & Tích hợp AI
Dữ liệu được đẩy thành luồng từ Client xuống DB và AI như sau:

1. **Client** gửi tin nhắn/sự kiện mạng qua **WebSocket** xuống **`svc-chat`**.
2. **`svc-chat`** sinh log và đẩy vào **Go Channels (RAM Buffer)**.
3. Các **Worker Pool** ngầm gom log thành từng mẻ (Batching) rồi chuyển sang **`svc-admin`**.
4. **`svc-admin`** ghi log vào **PostgreSQL**.
5. Đồng thời, **`svc-admin`** gọi **API Gemini/OpenAI** để phân tích ngữ cảnh.
6. Nếu phát hiện vi phạm (Toxic), **`svc-admin`** bắn lệnh **gRPC** (qua cổng bảo mật `:50051`) sang **`svc-chat`** để cắt kết nối WebSocket của User.

## 4. Bảo mật (Security)
- **Cổng `:8080`**: Mở public cho Client kết nối WebSocket.
- **Cổng `:9090` & `:50051`**: Chặn hoàn toàn khỏi Internet, chỉ cho phép giao tiếp nội bộ.
- **Xác thực gRPC**: `svc-chat` dùng gRPC Interceptor để bắt buộc kiểm tra Secret Token trong Metadata Header khi nhận lệnh BanUser từ `svc-admin`.
