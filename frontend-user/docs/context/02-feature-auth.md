# Báo cáo chức năng: Xác thực Người dùng (Auth Feature)

## 1. Trạng thái hiện tại
Chức năng xác thực (Login/Register) đã được thiết lập đầy đủ cấu trúc và các logic liên quan theo chuẩn Feature-based Architecture.

## 2. Chi tiết Triển khai (`src/features/auth`)

- **API Integration (`services/auth.api.ts`):** 
  - Đã tích hợp các hàm gọi Axios client cho các endpoint: đăng nhập (`login`) và đăng ký (`register`).
- **State Management (`store/auth.store.ts`):** 
  - Sử dụng Zustand để quản lý State người dùng trên toàn cục (Global State).
  - Quản lý trạng thái xác thực (`isAuthenticated`), thông tin người dùng (`user`), cũng như lưu trữ an toàn chuỗi `access_token` để tự động gán vào Request Headers ở các lần gọi API tiếp theo.
- **Business Logic (`hooks/`):**
  - Tách biệt logic với UI thông qua các Custom Hooks (`useLogin`, `useRegister`).
  - Đảm nhiệm chức năng call API service, xử lý trạng thái đang tải (`isLoading`), bắt và hiển thị lỗi, và thay đổi State hệ thống.
- **UI Components (`components/`):**
  - Các tệp `login-form.tsx` và `register-form.tsx` chỉ đóng vai trò nhận Input từ người dùng và gọi các Custom Hooks tương ứng.
- **Routing:**
  - Ánh xạ trực tiếp từ `src/app/(auth)/login/page.tsx` và `src/app/(auth)/register/page.tsx`.
