# Stranger Chat Client

Stranger Chat Client là một dự án web app cho phép người dùng trò chuyện ẩn danh (Stranger Chat) thời gian thực. Dự án hỗ trợ gọi video hai chiều (2-way video call), nhắn tin trực tiếp, đăng ký, đăng nhập và hiển thị trạng thái theo thời gian thực.

> **Lưu ý đặc biệt**: Dự án này được khởi tạo vào **Tháng 1** và được thực hiện hoàn toàn bởi các **AI Agents** (cụ thể là **Antigravity CLI** - `anti cli`). Toàn bộ quá trình cấu hình, thiết kế kiến trúc và triển khai code đều được quản lý tự động, cẩn thận và chi tiết bởi Agent.

## 🚀 Tính Năng Chính (Core Features)

- **Video Call 2 chiều (WebRTC/Socket):** Cho phép 2 người dùng xa lạ kết nối và gọi video trực tiếp.
- **Chat Real-time:** Nhắn tin tức thời, hiển thị trạng thái đã xem/đang gõ.
- **Authentication:** Đăng ký, đăng nhập an toàn, bảo mật.
- **Real-time Status:** Hiển thị trạng thái online, offline, đang trong cuộc gọi.
- **Ghép Đôi Ngẫu Nhiên (Random Matching):** Tự động tìm kiếm và ghép đôi với một người dùng ngẫu nhiên đang online.

## 🏗 Cấu Trúc Dự Án (Feature-based Architecture)

Dự án áp dụng mô hình **Feature-based architecture** (Kiến trúc theo tính năng), giúp codebase dễ dàng mở rộng và bảo trì. Mỗi tính năng sẽ đóng gói toàn bộ components, hooks, services, và utils của riêng nó.

Dưới đây là cấu trúc thư mục tổng quan mà dự án sẽ hướng tới:

```text
src/
├── app/                  # Next.js App Router (Pages, Layouts, Routing)
├── features/             # (Feature-based modules)
│   ├── auth/             # Tính năng Đăng nhập/Đăng ký
│   ├── chat/             # Tính năng Nhắn tin thời gian thực
│   ├── video-call/       # Tính năng Video call (WebRTC)
│   └── matching/         # Tính năng Ghép đôi người dùng
├── shared/               # Code dùng chung (UI components, hooks, utils chung)
│   ├── components/       # UI Components tái sử dụng (Button, Modal, Input,...)
│   ├── hooks/            # Custom hooks dùng chung
│   ├── lib/              # Cấu hình thư viện (Axios, Socket.io, Firebase,...)
│   └── types/            # TypeScript interfaces/types dùng chung
└── public/               # Static assets (images, icons,...)
```

## 🛠 Công Nghệ Sử Dụng

- **Framework:** Next.js 16 (App Router)
- **UI Library:** React 19
- **Styling:** Tailwind CSS v4
- **Language:** TypeScript
- **Real-time/Video Call:** WebRTC & Socket (Dự kiến)
- **Quản lý bởi:** Antigravity CLI (AI Agents)

## 📦 Hướng Dẫn Cài Đặt và Chạy Dự Án

### Yêu Cầu Hệ Thống (Prerequisites)

- Node.js >= 20
- npm, yarn, hoặc pnpm

### Cài Đặt (Installation)

1. Clone dự án và truy cập vào thư mục:
   ```bash
   git clone <repo_url>
   cd stranger-chat-client
   ```

2. Cài đặt các dependencies:
   ```bash
   npm install
   ```

### Chạy Môi Trường Phát Triển (Development)

```bash
npm run dev
```

Mở [http://localhost:3000](http://localhost:3000) trên trình duyệt để xem kết quả.

### Build và Chạy Production

```bash
npm run build
npm start
```

## 🤖 Quản Trị Bởi AI Agent (Antigravity CLI)

Dự án này là minh chứng cho việc sử dụng **AI Agents** để xây dựng phần mềm:
- **Tự động hóa:** Từ lúc khởi tạo (scaffolding) đến lúc hoàn thiện tính năng.
- **Tiêu chuẩn cao:** Code được sinh ra bám sát các best practices mới nhất (Next.js 16 App Router, Feature-based).
- **Tính cẩn thận & Chi tiết:** AI tự quản lý việc chia nhỏ file, quản lý state và tối ưu UI/UX.

---
*Dự án được xây dựng và quản lý bởi Antigravity CLI.*
