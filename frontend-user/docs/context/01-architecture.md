# Báo cáo Kiến trúc và Cấu trúc dự án (Architecture & Structure)

## 1. Công nghệ cốt lõi
- **Framework Frontend:** Next.js 16 (App Router) với React 19.
- **Styling:** Tailwind CSS v4.
- **Ngôn ngữ:** TypeScript.
- **State Management:** Zustand.
- **API Client:** Axios.

## 2. Mô hình Kiến trúc
Dự án áp dụng mô hình kiến trúc theo **Feature-based Architecture** (Lấy cảm hứng từ Feature-Sliced Design - FSD). Việc này giúp tách bạch Logic nghiệp vụ (Domain Logic) và UI (Trải nghiệm người dùng), tối ưu tính đóng gói (Encapsulation) và dễ dàng bảo trì.

## 3. Cấu trúc Thư mục Hệ thống
```text
src/
├── app/                  # Chứa toàn bộ Next.js App Router (pages, layouts, loading)
│   ├── (auth)/           # Route Group: Các trang xác thực (Login, Register) không ảnh hưởng URL chung
│   ├── chat/             # Route: Giao diện Chat chính
│   ├── layout.tsx        # Layout gốc, chứa các Providers (Next Themes)
│   └── page.tsx          # Landing page của dự án
├── features/             # Nơi chứa toàn bộ Module chức năng chính
│   ├── auth/             # Chức năng Xác thực người dùng (Auth)
│   └── chat/             # Chức năng Chat và Ghép đôi
└── shared/               # Nơi chứa các thư viện dùng chung cho toàn dự án
    ├── components/       # Các UI Component cơ bản (Buttons, Inputs, ThemeToggle...)
    ├── lib/              # Tiện ích, cấu hình HTTP Client (Axios)
    └── types/            # Khai báo TypeScript chung (API Responses, Models...)
```
