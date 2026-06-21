# Code Base Architecture

Dự án này tuân theo kiến trúc **Feature-based Architecture** kết hợp với **Next.js App Router**. Mục tiêu là giúp codebase dễ bảo trì, dễ mở rộng và giảm thiểu sự phụ thuộc chéo giữa các tính năng.

## 1. Cấu trúc thư mục (Directory Structure)

```text
src/
├── app/                  # Chứa toàn bộ Next.js App Router (pages, layouts, API routes)
├── features/             # Nơi chứa logic nghiệp vụ, chia theo từng tính năng cụ thể
│   ├── auth/             # Tính năng xác thực
│   ├── chat/             # Tính năng nhắn tin
│   ├── video-call/       # Tính năng gọi video (WebRTC)
│   └── matching/         # Tính năng ghép đôi
├── shared/               # Chứa các tài nguyên dùng chung trên toàn hệ thống
│   ├── components/       # UI Components tái sử dụng (Button, Input, Modal, ...)
│   ├── hooks/            # Custom React Hooks dùng chung (useMediaQuery, useSocket, ...)
│   ├── lib/              # Cấu hình thư viện (axios instance, socket instance, ...)
│   ├── types/            # TypeScript Interfaces / Types dùng chung
│   └── utils/            # Helper functions dùng chung (formatDate, cn, ...)
└── public/               # Chứa static assets (hình ảnh, icon, font)
```

## 2. Nguyên tắc tổ chức Feature Module

Mỗi thư mục bên trong `features/` (ví dụ `features/chat/`) sẽ hoạt động như một module độc lập và bao gồm cấu trúc bên trong như sau:

```text
features/chat/
├── components/           # Các component chỉ dùng riêng cho tính năng chat (ChatBubble, MessageList)
├── hooks/                # Hooks riêng của tính năng (useChatMessages, useTypingIndicator)
├── services/             # API calls hoặc Socket events riêng của chat
├── types/                # Types/Interfaces liên quan đến chat
└── index.ts              # File export public API của feature này (Barrel file)
```

## 3. Quy tắc Import/Export (Boundaries)

- **App Router (`src/app`)**: Chỉ đóng vai trò import các components từ `features` và `shared` để lắp ráp thành trang (page). Không chứa logic nghiệp vụ phức tạp.
- **Tính đóng gói của Feature**: Một feature KHÔNG ĐƯỢC import trực tiếp các internal file của feature khác. Nếu cần dùng chéo, phải import thông qua file `index.ts` của feature đó.
- **Shared Directory**: Là nơi duy nhất có thể được import từ bất kỳ đâu. Các code trong `shared` tuyệt đối KHÔNG import ngược lại từ `features`.

## 4. Tương tác với API và Real-time

- **REST API**: Khai báo các service gọi API thông qua `axios` hoặc `fetch` bên trong `features/<name>/services`.
- **WebSocket/WebRTC**:
  - `shared/lib/socket.ts`: Khởi tạo kết nối gốc.
  - `features/video-call/hooks/useWebRTC.ts`: Quản lý session video call.
  - Các sự kiện lắng nghe (listen) và phát (emit) được gom vào các custom hooks theo tính năng.
