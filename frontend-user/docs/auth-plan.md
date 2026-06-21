# Kế hoạch triển khai Authentication (Register & Login)

Dựa trên API Contract bạn cung cấp, đây là bản kế hoạch chi tiết (Plan) để triển khai tính năng Xác thực (Auth) theo mô hình **Feature-based Architecture** đảm bảo code Clean, Scalable và dễ bảo trì nhất.

## 1. Cấu trúc thư mục dự kiến cho Feature Auth

```text
src/
├── shared/
│   ├── lib/
│   │   └── api-client.ts        # Cấu hình Axios instance (Base URL, Interceptors)
│   └── types/
│       └── api.ts               # Các type API chung (ví dụ: ApiErrorResponse)
├── features/
│   └── auth/
│       ├── types/
│       │   └── index.ts         # Khai báo interface User, LoginPayload, AuthResponse...
│       ├── services/
│       │   └── auth.api.ts      # Các hàm gọi API (login, register)
│       ├── store/
│       │   └── auth.store.ts    # Quản lý global state (lưu token, user info) bằng Zustand
│       ├── hooks/
│       │   ├── use-login.ts     # Hook xử lý logic form login, loading, error
│       │   └── use-register.ts  # Hook xử lý logic form register, loading, error
│       ├── components/
│       │   ├── login-form.tsx   # UI Form Đăng nhập
│       │   └── register-form.tsx# UI Form Đăng ký
│       └── index.ts             # Export public API của feature
└── app/
    └── (auth)/                  # Route group cho xác thực
        ├── login/page.tsx
        └── register/page.tsx
```

## 2. Chi tiết từng bước thực hiện (Roadmap)

### Bước 1: Setup Axios Client (Shared Library)
- **Mục tiêu:** Tạo một `axios instance` dùng chung cho toàn bộ app.
- **Chi tiết:**
  - Cấu hình `baseURL` là biến môi trường (ví dụ: `http://localhost:8080/v1`).
  - Thêm **Request Interceptor**: Tự động lấy `access_token` từ localStorage/Zustand đính kèm vào header `Authorization: Bearer <token>`.
  - Thêm **Response Interceptor**: Bắt lỗi chung (parse lỗi `400 Bad Request` trả về `error` field từ backend để hiển thị lên UI).

### Bước 2: Định nghĩa Types & Interfaces (Domain Types)
- Dựa trên JSON response, định nghĩa các model chuẩn xác bằng TypeScript:
  - `interface User` (chứa id, username, email, status, avatar_url...).
  - `interface AuthResponse` (chứa message, data: User, access_token, refresh_token).
  - `interface LoginRequest` & `interface RegisterRequest`.

### Bước 3: Triển khai Auth Service & State Management
- **Auth Service:** Viết 2 hàm `login(payload)` và `register(payload)` gọi qua axios client.
- **Auth Store (Zustand):**
  - Tạo một global store lưu trữ `user` và `isAuthenticated`.
  - Tự động lưu `access_token` và `refresh_token` vào `localStorage` (hoặc Cookie tùy chiến lược bảo mật, tạm thời dùng localStorage).

### Bước 4: Xây dựng UI Components & Custom Hooks
- Dùng thư viện icon (`lucide-react`) và các component UI chung (`Button`, `Input`).
- **Custom Hooks (`useLogin`, `useRegister`):** Đóng gói logic xử lý gọi service, try-catch, set trạng thái loading, và cập nhật store nếu thành công. Tách biệt hoàn toàn logic khỏi UI.
- **Forms:** Xây dựng `LoginForm` và `RegisterForm` gọn gàng, chỉ nhận input và gọi custom hooks.

### Bước 5: Tích hợp vào App Router
- Tạo các route `app/(auth)/login/page.tsx` và `app/(auth)/register/page.tsx` để render các form đã tạo.

---
## 3. Tại sao cấu trúc này là "Clean" nhất cho dự án của bạn?
- **Separation of Concerns (SoC):** UI Component không tự gọi API hay lưu token. Nó gọi Custom Hook. Custom Hook gọi Service. Service gọi qua Axios Client. Lỗi ở đâu sửa ở đó.
- **Type Safety:** Định nghĩa type từ đầu vào tới đầu ra. Bất kỳ sự thay đổi API contract nào cũng sẽ báo lỗi TypeScript ngay lập tức.
- **Dễ tái sử dụng:** Nếu sau này có tính năng "Đổi mật khẩu", ta chỉ cần ném thêm service vào `auth.api.ts` mà không phá vỡ cấu trúc.
