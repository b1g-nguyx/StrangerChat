<!-- BEGIN:nextjs-agent-rules -->
# This is NOT the Next.js you know

This version has breaking changes — APIs, conventions, and file structure may all differ from your training data. Read the relevant guide in `node_modules/next/dist/docs/` before writing any code. Heed deprecation notices.
<!-- END:nextjs-agent-rules -->

<!-- BEGIN:project-ui-ux-rules -->
# UI/UX & Design Constraints (BẮT BUỘC)

Bạn đang code dự án Stranger Chat Client. Khi thực hiện thay đổi hoặc tạo mới giao diện, bạn PHẢI tuân thủ các quy tắc sau (chi tiết tại `docs/ui-ux-rules.md`):

1. **Apple-like & Minimalist**: Giao diện tối giản, thanh lịch, ưu tiên khoảng trắng (whitespace) và kiểu chữ (typography) rõ nét.
   - Tránh thiết kế rườm rà.
   - Responsive & App-like: Chạm chuẩn Mobile (`h-12`).
2. **Dynamic Theme (Sáng/Tối linh hoạt)**: 
   - Hỗ trợ cả 2 chế độ Sáng và Tối (Dùng `next-themes` và prefix `dark:`).
   - Nền (Background): Xám Apple `bg-[#f5f5f7]` (Light) / Đen sâu `bg-[#000000]` (Dark).
   - Surface (Kính mờ - Frosted Glass): `bg-white/70` (Light) / `bg-[#1c1c1e]/70` (Dark) kèm `backdrop-blur-xl`.
   - Chữ: Đen `text-zinc-900` (Light) / Trắng `text-zinc-50` (Dark).
3. **Primary Accent (Tinh tế)**: 
   - Sử dụng một màu trung tâm (Apple Blue `#007AFF`) dạng màu trơn tĩnh, KHÔNG lạm dụng hiệu ứng Neon hay Gradient loè loẹt.
4. **Shapes & Hoạt ảnh giọt nước (Ripple/Liquid Animation)**: 
   - Góc bo tròn mềm (Squircle): `rounded-2xl` hoặc `rounded-[32px]`.
   - Mọi nút bấm hoặc thẻ Clickable PHẢI có hiệu ứng Ripple (Giọt nước toả ra) và Animation thu nhỏ (`active:scale-[0.98]`). Sử dụng `<Ripple />` component (framer-motion).
<!-- END:project-ui-ux-rules -->
