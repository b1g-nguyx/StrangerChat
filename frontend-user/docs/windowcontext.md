# Window Context & System Status Report

This document synthesizes the current state of the **Stranger Chat Client** (Frontend), documenting the active context, the existing features, and the architecture that has been implemented.

## 1. Project Overview & Architecture
- **Frameworks:** Next.js 16 (App Router), React 19, TypeScript, Tailwind CSS v4.
- **Architecture:** Feature-based Architecture (FSD-inspired).
- **Global State:** Zustand.
- **API Client:** Axios with Interceptors.

## 2. Directory Structure
```text
src/
├── app/                  # App Router pages and layouts
│   ├── (auth)/           # Route group for authentication (Login, Register)
│   ├── chat/             # Route for Chat interface
│   ├── layout.tsx        # Root layout (Next Themes provider, global fonts)
│   └── page.tsx          # Landing page
├── features/             # Feature modules (Domain logic & components)
│   ├── auth/             # Authentication feature (Services, Store, Hooks, Components)
│   └── chat/             # Chat feature implementation
└── shared/               # Shared logic across features
    ├── components/       # Reusable UI components (Buttons, Inputs, Loading, ThemeToggle)
    ├── lib/              # Utility functions, API Client (Axios setup)
    └── types/            # Shared TypeScript types
```

## 3. Implemented Features

### 3.1. Authentication (Feature: `auth`)
- **API Integration:** Implemented `auth.api.ts` utilizing `axios` for `login` and `register` endpoints.
- **State Management:** Implemented `auth.store.ts` using Zustand to persist user session, handle `access_token` storage, and manage authentication state (`isAuthenticated`).
- **Hooks:** Custom hooks (`useLogin`, `useRegister`) encapsulating logic, loading states, and error handling for form submissions.
- **UI Components:** `LoginForm` and `RegisterForm` separating UI from business logic. Adheres strictly to the minimalist Apple-like design guidelines.
- **Routing:** Accessible via `/login` and `/register`.

### 3.2. Shared Core (`shared`)
- **API Client:** `api-client.ts` configured with interceptors to automatically attach the Bearer token for protected endpoints and handle global error parsing.
- **UI System:** Integration of Next Themes for light/dark mode. Core UI components with Ripple animations and squircle borders implemented according to the UI/UX rules.

### 3.3. Chat / Core App (Feature: `chat`)
- Basic routing setup (`/chat`).
- *Note: Further Real-time chat (Socket.io) and Video Call (WebRTC) integrations are outlined in the overview but require further implementation or validation of current status.*

## 4. UI/UX Standards in Use
- **Apple-like & Minimalist:** Generous whitespace, clean typography (`Inter`), and high contrast readability.
- **Dynamic Themes:** Full support for `dark:` mode using `next-themes`.
  - Light: Background `#f5f5f7`, Surface `bg-white/70` with `backdrop-blur-xl`.
  - Dark: Background `#000000`, Surface `bg-[#1c1c1e]/70` with `backdrop-blur-xl`.
- **Interactions:** Subtle scale animations on click (`active:scale-[0.98]`) and ripple effects applied to actionable elements.

## 5. Development Status
- **Authentication:** Fully structured and integrated with backend API conventions.
- **Next Steps:** Expansion of real-time matching, WebRTC video calling, and text messaging pipelines in the `chat` feature.
