# StrangerChat Backend Architecture

This document outlines the architectural patterns and design decisions for the StrangerChat backend. We strive for a balance between clean design (for maintainability and scalability) and Idiomatic Go (for simplicity and performance).

## Core Architecture Pattern

The backend is built using a **Pragmatic Clean Architecture** combined with **Feature Sliced Design**.

1.  **Dependency Inversion**: The core business logic does not depend on the delivery mechanisms (HTTP/Websockets) or databases. Dependencies point inwards.
2.  **Feature Slices (`internal/features`)**: The application is divided into high-level features (e.g., `auth`, `chat`, `user`). Each feature has its own:
    -   `delivery`: Controllers, HTTP handlers, or WebSocket hubs.
    -   `usecase`: The business logic specific to that feature.
    -   `repository`: Feature-specific data access implementations.

## The Hybrid Entity Architecture

We utilize a hybrid approach for defining data models (Entities), combining DDD's *Shared Kernel* and *Bounded Contexts*.

### 1. Global Entities (Shared Kernel)
Location: `internal/entity/`

-   **Purpose**: These are core data structures that represent the fundamental "source of truth" across the entire system.
-   **Rule**: **DO NOT** add feature-specific fields here. Only properties that are universally true and required across multiple features should reside in a Global Entity.
-   **Example**: `entity.User` might only contain `ID`, `Username`, `Email`.

### 2. Feature Entities (Bounded Contexts)
Location: `internal/features/<feature_name>/entity/`

-   **Purpose**: These structs represent a model tailored specifically for a single feature's needs. This prevents the Global Entities from becoming bloated "God Objects."
-   **Rule**: If a feature needs new fields, custom JSON mappings, or specific validation rules that other features do not care about, create a Feature Entity.
-   **Example**: `features/chat/entity.ChatMessage` represents a message specifically formatted for WebSocket transmission and Redis storage within the chat feature. It encapsulates chat-specific behavior.

## 3. Advanced Features Architecture

### 3.1. Video Call Architecture (Self-hosted P2P)
The Video Call functionality is built with a lightweight, peer-to-peer architecture:
- **Signaling via WebSocket:** Utilizes the existing `svc-socket` as a Signaling Server. To maintain *Separation of Concerns* and prevent video signaling from blocking chat messages, WebRTC payloads (`WEBRTC_OFFER`, `WEBRTC_ANSWER`, `WEBRTC_ICE_CANDIDATE`) are processed asynchronously via non-blocking Goroutines.
- **Direct Media Flow:** Video and audio data flow directly between clients (P2P), consuming zero backend bandwidth and ensuring high scalability at minimum server cost.

## Summary

By strictly separating Core Data from Feature-Specific Data, we ensure that the codebase remains decoupled, easy to navigate, and ready to scale into microservices if necessary in the future.
