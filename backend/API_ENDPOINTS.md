# Báo Cáo Tích Hợp Frontend (API & WebSocket)

Tài liệu này tổng hợp các đầu API và WebSocket hiện có ở phía Backend của ứng dụng StrangerChat để đội ngũ Frontend có thể kết nối.

## Thông tin chung

- **Ngôn ngữ & Framework:** Go (Golang) với Fiber Web Framework (`github.com/gofiber/fiber/v2`).
- **WebSocket:** Sử dụng module `github.com/gofiber/contrib/websocket`.
- **Cổng kết nối (Port):** Thường sử dụng cổng `8080` cho HTTP Client/Admin và `8081` cho WebSocket (Tham khảo trong tệp `.env.example`).

---

## 1. RESTful APIs (HTTP)

Tất cả các API dưới đây đều nhận và trả về dữ liệu dưới định dạng `application/json`.

### Xác thực & Người dùng (Client)

| Phương thức | Endpoint | Mô tả | Body Yêu cầu / Header |
| :--- | :--- | :--- | :--- |
| `POST` | `/v1/auth/register` | Đăng ký tài khoản mới | `{"username": "...", "email": "...", "password": "..."}` |
| `POST` | `/v1/auth/login` | Đăng nhập hệ thống | `{"email": "...", "password": "..."}` |
| `POST` | `/v1/auth/refresh` | Làm mới token | Đọc `refresh_token` từ Cookie (gửi kèm request) |
| `POST` | `/v1/auth/logout` | Đăng xuất | Xóa Cookie chứa `refresh_token` |

*Lưu ý: API Login/Refresh sẽ trả về `AccessToken` ở body response và tự động set `refresh_token` vào HttpOnly Cookie. Cần gắn `Authorization: Bearer <AccessToken>` vào header ở các API cần đăng nhập.*

### Quản trị viên (Admin)

| Phương thức | Endpoint | Mô tả |
| :--- | :--- | :--- |
| `POST` | `/admin/v1/auth/login` | Đăng nhập dành cho Quản trị viên |
| `GET` | `/admin/v1/users` | Lấy danh sách người dùng. Hỗ trợ truyền filter qua Query Params |

---

## 2. WebSocket (Real-time Chat)

Sử dụng để phục vụ tính năng tìm bạn và chat theo thời gian thực.

- **Giao thức:** `ws://` (hoặc `wss://` nếu có HTTPS).
- **Endpoint:** `ws://<domain_hoặc_ip>:<port>/ws/chat?token=<AccessToken>`
- **Bắt buộc:** Phải đính kèm token vào URL parameter `?token=...` khi khởi tạo kết nối.

### Các loại Message (Dạng JSON)

**Frontend gửi lên Server (Send):**
*   Tìm bạn chat: `{"type": "FIND_MATCH"}`
*   Gửi tin nhắn: `{"type": "CHAT", "content": "Nội dung tin nhắn của bạn..."}`
*   Rời phòng chat: `{"type": "LEAVE_ROOM"}`
*   Tố cáo người dùng: `{"type": "REPORT", "reported_id": "uuid-nguoi-dung", "room_id": "uuid-cua-phong", "content": "Lý do tố cáo..."}`

**Server gửi về Frontend (Receive):**
*   Ghép đôi thành công: `{"type": "MATCHED", "room_id": "uuid-cua-phong"}`
*   Có tin nhắn tới: `{"type": "CHAT", "room_id": "...", "content": "..."}`
*   Đối phương đã rời phòng: `{"type": "PARTNER_LEFT", "room_id": "..."}`

**Ví dụ Code Frontend:**
```javascript
// Khởi tạo kết nối WebSocket với Token
const token = 'your_access_token_here';
const socket = new WebSocket(`ws://localhost:8081/ws/chat?token=${token}`);

socket.onopen = () => {
  console.log('Đã kết nối! Bắt đầu tìm bạn...');
  // Gửi lệnh tìm bạn ngay khi kết nối
  socket.send(JSON.stringify({ type: "FIND_MATCH" }));
};

socket.onmessage = (event) => {
  const data = JSON.parse(event.data);
  switch (data.type) {
    case 'MATCHED':
      console.log('Đã tìm thấy bạn chat, Room ID:', data.room_id);
      break;
    case 'CHAT':
      console.log('Tin nhắn mới:', data.content);
      break;
    case 'PARTNER_LEFT':
      console.log('Bạn chat đã thoát.');
      break;
  }
};
```

---

## 3. Tính năng đang chờ hoàn thiện (Pending)

*   (Hiện tại không có tính năng nào đang chờ hoàn thiện trong phạm vi tài liệu này.)

---

## 4. Gợi ý Tích hợp

1. **Quản lý Token:** Dùng JWT cho `AccessToken` gửi qua Header, và để trình duyệt tự quản lý `refresh_token` qua HttpOnly Cookie.
2. **CORS:** Đảm bảo thêm tùy chọn `credentials: 'include'` khi fetch/axios để trình duyệt cho phép đính kèm Cookie.

---

## 5. Lịch sử cập nhật (Changelog)
- **Cập nhật mới nhất:** Đã sửa lỗi không thể nhắn tin sau khi ghép phòng. Biến `room_id` ở phía client (trên server) giờ đã được gán chính xác khi sự kiện `matched` xảy ra. Đồng thời khi đối phương ngắt kết nối (`partner_left`), trạng thái phòng của người còn lại sẽ được dọn dẹp sạch sẽ để sẵn sàng tìm người mới.
