# Tổng Quan Dự Án (Project Overview)

## 1. Giới thiệu (Introduction)
Stranger Chat Client là một ứng dụng web thời gian thực cho phép hai người dùng ngẫu nhiên kết nối, trò chuyện ẩn danh thông qua tin nhắn văn bản và video call. Ứng dụng tập trung vào trải nghiệm mượt mà, kết nối nhanh chóng và ẩn danh an toàn.

## 2. Tính năng cốt lõi (Core Features)

### 2.1. Xác thực và Người dùng (Authentication & User)
- Đăng nhập/Đăng ký tài khoản (hỗ trợ Guest mode và Account mode).
- Hiển thị và cập nhật trạng thái người dùng (Online, Offline, In-call).

### 2.2. Ghép đôi (Random Matching)
- Thuật toán ghép đôi ngẫu nhiên 2 người dùng đang trong trạng thái chờ (Queue).
- Hỗ trợ các bộ lọc ghép đôi cơ bản (ví dụ: sở thích, khu vực - nếu có).

### 2.3. Trò chuyện thời gian thực (Real-time Chat)
- Gửi/nhận tin nhắn văn bản tức thì.
- Hiển thị trạng thái đang gõ (Typing indicator) và đã xem (Read receipts).

### 2.4. Gọi Video/Audio (Video/Audio Call)
- Kết nối peer-to-peer (P2P) qua WebRTC.
- Các tính năng điều khiển cơ bản: Tắt/bật mic, tắt/bật camera, kết thúc cuộc gọi.

## 3. Đối tượng người dùng (Target Audience)
- Những người muốn tìm kiếm bạn trò chuyện mới một cách ngẫu nhiên.
- Yêu cầu ứng dụng có độ trễ thấp, thao tác dễ dàng trên cả Mobile và Desktop.

## 4. Công nghệ cốt lõi (Tech Stack)
- **Frontend Framework:** Next.js 16 (App Router), React 19.
- **Styling:** Tailwind CSS v4.
- **Real-time Communication:** Socket.io (Signaling & Chat) + WebRTC (Video Call).
- **Ngôn ngữ:** TypeScript.
