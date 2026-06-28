# Báo cáo chức năng: Chat & Cuộc gọi (Chat Feature)

## 1. Trạng thái hiện tại
Chức năng Chat là luồng cốt lõi của ứng dụng `Stranger Chat`, đang trong quá trình hình thành cơ sở để hỗ trợ kết nối Real-time.

## 2. Chi tiết Triển khai (`src/features/chat` & `src/app/chat`)

- **Routing:** 
  - Đã có khung định tuyến chính tại `src/app/chat/page.tsx` và `layout.tsx` (nếu có).
- **Core Feature Định hướng:**
  - **Ghép đôi Ngẫu nhiên (Random Matching):** Hệ thống hàng đợi (Queue) chờ người dùng.
  - **Trò chuyện Thời gian thực (Real-time Chat):** Sẽ sử dụng Socket.io cho phần Signaling và Nhắn tin văn bản, trạng thái gõ chữ (Typing Indicator), và đã xem (Read receipts).
  - **Gọi Video/Audio (WebRTC):** Thiết lập kết nối P2P (Peer-to-peer) cho cuộc gọi bảo mật độ trễ thấp; kèm các tính năng điều khiển tắt/bật Camera/Mic.
- **Tiến độ:** 
  - Khung Module (Boilerplate) đã được tạo sẵn theo Feature-based Architecture. Các hệ thống thời gian thực (Socket.io và WebRTC) sẽ là các công việc triển khai ở giai đoạn tiếp theo của dự án.
