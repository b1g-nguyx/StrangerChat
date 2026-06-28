# Báo cáo: Shared Logic & UI/UX Standards

## 1. Trạng thái hiện tại
Thư mục `shared` đóng vai trò là xương sống của các thành phần chung xuyên suốt toàn ứng dụng. Giao diện đang tuân thủ nghiêm ngặt các quy tắc thiết kế do dự án đề ra.

## 2. Shared Core (`src/shared`)
- **API Client (`lib/api-client.ts`):** 
  - Axios Instance được cấu hình sẵn với `baseURL` từ biến môi trường.
  - Sử dụng Interceptors cho Requests (tự động đính kèm `Authorization` header) và Responses (tự động xử lý lỗi tập trung).
- **Types (`types/`):** Khai báo các mô hình dữ liệu (Interfaces/Types) cho toàn bộ ứng dụng như các Response chuẩn, mô hình Người dùng chung...
- **Shared Components (`components/`):** 
  - Tái sử dụng cao: Nút bấm (Button), Trình gõ (Input), Spinners.
  - Theme Toggle để chuyển đổi Sáng/Tối.

## 3. UI/UX & Design Constraints
Theo file `ui-ux-rules.md`, mọi thành phần giao diện mới cần áp dụng:
- **Phong cách Tối giản (Apple-like):** Rõ nét, sử dụng khoảng trắng (Whitespace) tinh tế. Thiết kế chuẩn Mobile (`h-12`).
- **Dynamic Theme (Sáng/Tối linh hoạt):** 
  - Tích hợp bằng `next-themes`.
  - Bảng màu: Nền xám Apple `#f5f5f7` (Light) và đen `#000000` (Dark).
  - Mặt kính mờ (Frosted Glass/Surface): `bg-white/70` hoặc `bg-[#1c1c1e]/70` với hiệu ứng `backdrop-blur-xl`.
- **Hoạt ảnh & Tương tác:** 
  - Nút bấm và khu vực có thể nhấp luôn sử dụng hiệu ứng giọt nước thu nhỏ (Ripple / Scale Animation: `active:scale-[0.98]`).
  - Đường bo cong mềm mại (Squircle) chủ đạo: `rounded-2xl` hoặc `rounded-[32px]`.
- **Màu sắc chủ đạo:** Apple Blue `#007AFF`. Cấm lạm dụng Gradient hoặc Neon lòe loẹt.
