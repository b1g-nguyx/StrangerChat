# Code Standards (Tiêu chuẩn code)

Tài liệu này định nghĩa các quy chuẩn code bắt buộc cho toàn bộ dự án để đảm bảo tính nhất quán và dễ dàng bảo trì.

## 1. Naming Conventions (Quy tắc đặt tên)

- **Thư mục & File (không phải Component):** Sử dụng `kebab-case` (ví dụ: `video-call.ts`, `use-chat.ts`).
- **React Components:** Sử dụng `PascalCase` cho tên file và tên component (ví dụ: `ChatMessage.tsx`, `VideoPlayer.tsx`).
- **Variables & Functions:** Sử dụng `camelCase` (ví dụ: `isOnline`, `handleSendMessage`).
- **Constants:** Sử dụng `UPPER_SNAKE_CASE` (ví dụ: `MAX_RETRY_COUNT`, `API_BASE_URL`).
- **Types/Interfaces:** Sử dụng `PascalCase`, ưu tiên dùng `Interface` trừ khi cần tính năng đặc thù của `Type`. Không thêm tiền tố `I` (ví dụ: `User`, không phải `IUser`).

## 2. TypeScript Guidelines

- **Bắt buộc dùng Type:** Hạn chế tối đa việc sử dụng `any`. Hãy dùng `unknown` nếu chưa rõ kiểu dữ liệu.
- **Props:** Định nghĩa interface rõ ràng cho tất cả props của component.
- **Return Type:** Các custom hooks nên khai báo rõ ràng kiểu trả về. Các component React có thể để inference tự nhận diện (trả về `ReactNode`).

## 3. React & Next.js Best Practices

- **Server Components (RSC) vs Client Components:**
  - Mặc định sử dụng Server Components.
  - Chỉ thêm `'use client'` khi thực sự cần dùng state, hooks (`useState`, `useEffect`), hoặc event listeners.
- **Hooks:**
  - Giữ hooks nhỏ gọn và làm 1 việc duy nhất.
  - Các logic phức tạp trong component cần được tách ra custom hooks để dễ test.
- **Export:** Sử dụng Named Export thay vì Default Export (trừ các file bắt buộc của Next.js như `page.tsx`, `layout.tsx`).

## 4. Styling (Tailwind CSS)

- Khuyến khích dùng tiện ích `cn()` (sử dụng `clsx` và `tailwind-merge`) để gộp các class Tailwind tránh xung đột.
- Không viết CSS nội tuyến (inline style) trừ khi giá trị mang tính động (dynamic value) được tính toán qua JS.
- Sử dụng các biến màu sắc, spacing đã định nghĩa trong design system (theme Tailwind).

## 5. Clean Code & Linter

- Tránh lồng quá nhiều vòng `if-else` (Tránh Callback Hell). Trả về sớm (Return early).
- Code phải vượt qua các linter rules (ESLint, Prettier). Không tự ý tắt rule (`eslint-disable`) mà không có comment giải thích lý do.
