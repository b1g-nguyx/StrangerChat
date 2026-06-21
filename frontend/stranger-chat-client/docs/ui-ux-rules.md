# UI/UX Rules & Design System (Apple-like & Minimalist)

Tài liệu này định nghĩa hệ thống thiết kế (Design System) hiện đại, tập trung vào trải nghiệm tinh tế, thanh lịch và sạch sẽ (Clean UI) lấy cảm hứng từ các sản phẩm của Apple. Đặc biệt hệ thống hỗ trợ cả chế độ Sáng/Tối linh hoạt và các hiệu ứng chuyển động chất lỏng (Ripple).

## 1. Triết lý Thiết kế (Design Philosophy)

- **Minimalist & Clean**: Đề cao không gian trắng (Whitespace), loại bỏ hoàn toàn các chi tiết rườm rà. Mọi thành phần trên màn hình đều phải có mục đích rõ ràng.
- **Focus on Typography**: Phông chữ là cốt lõi của giao diện. Sử dụng font chữ hiện đại (Inter, SF Pro) với việc phân cấp kích thước và độ đậm (weight) cực kỳ khắt khe để tạo độ tương phản.
- **App-like Feel**: Giao diện mang cảm giác của một Native App (iOS). Không cuộn thừa, tối ưu hoá cho thao tác ngón tay trên màn hình cảm ứng (`min-h-[100dvh]`).

## 2. Bảng màu Động (Dynamic Light/Dark Palette)

Hệ thống bắt buộc phải hỗ trợ chuyển đổi linh hoạt qua lại giữa Light Mode và Dark Mode (thông qua class `dark:` của Tailwind và `next-themes`).

- **Backgrounds (Nền chính)**: 
  - Light mode: `bg-[#f5f5f7]` (màu xám nhạt Apple).
  - Dark mode: `bg-[#000000]` (đen sâu tuyệt đối).
- **Surfaces & Cards (Frosted Glass - Kính mờ)**:
  - Khung viền và nền card sử dụng độ mờ tinh tế kết hợp Blur. 
  - Light mode: `bg-white/70 backdrop-blur-xl border-black/5`.
  - Dark mode: `bg-[#1c1c1e]/70 backdrop-blur-xl border-white/10`.
- **Primary Color (Accent)**:
  - Dùng màu xanh dương thanh lịch (Apple Blue) `bg-[#007AFF]`. Không dùng gradient hay neon.
- **Text & Hierarchy**:
  - Tiêu đề (Headings): Đen (`text-zinc-900`) trên nền sáng, Trắng tinh (`text-zinc-50`) trên nền tối.
  - Phụ đề (Secondary): Màu xám êm ái `text-zinc-500` / `dark:text-zinc-400`.

## 3. Hình thái & Cấu trúc (Shapes & Layout)

- **Bo góc (Rounded Corners)**: 
  - Mềm mại và liên tục (Squircle). Dùng `rounded-2xl` hoặc `rounded-[32px]` cho thẻ (Card), Modal và `rounded-2xl` cho các nút bấm (Button).
- **Bóng đổ (Shadows)**: 
  - Rất tinh tế. Light mode: `shadow-[0_8px_30px_rgb(0,0,0,0.04)]`. Dark mode: `shadow-[0_8px_30px_rgb(0,0,0,0.12)]`.
- **Touch Target**: Mọi nút bấm hoặc input tối thiểu đạt `h-12` (48px) để thân thiện với thao tác chạm.

## 4. Hoạt ảnh & Tương tác (Liquid Animations & Micro-interactions)

- **Hiệu ứng giọt nước (Ripple / Liquid)**:
  - BẤT KỲ thành phần nào có thể bấm được (Nút bấm, Card tương tác) ĐỀU PHẢI tích hợp Component `<Ripple />` (được xây dựng bằng `framer-motion`). 
  - Khi click/chạm, một làn sóng trong suốt sẽ lan toả trên bề mặt nút hệt như hiệu ứng chạm vào mặt nước.
- **Spring/Fluid Animation**: Mọi chuyển động phải mượt như chất lỏng. Sử dụng `transition-all duration-300 ease-out`.
- **Phản hồi xúc giác (Visual Haptics)**: Khi chạm giữ, nút sẽ thu nhỏ rất nhẹ (`active:scale-[0.98]`).
- **Skeleton & Loading**: Spinner đơn giản, thanh mảnh (Stroke mỏng) màu xanh Apple (`#007AFF`).
