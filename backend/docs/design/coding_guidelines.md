# Quy chuẩn Lập trình (Coding Guidelines) cho Go

Go (Golang) là một ngôn ngữ đề cao sự đơn giản, dễ đọc và có các quy ước cực kỳ chặt chẽ từ cộng đồng. Dưới đây là các quy tắc cơ bản bắt buộc cần tuân theo trong dự án này:

## 1. Format Code (Định dạng)
- **Không tranh cãi về Format:** Mọi đoạn code Go đều phải được format bằng công cụ `gofmt` hoặc `goimports` trước khi lưu. Nó sẽ tự động chỉnh sửa dấu cách, tab, và gom nhóm các thư viện (imports).
- **Thụt lề (Indentation):** Go sử dụng **Tabs** thay vì Spaces.

## 2. Quy tắc Đặt tên (Naming Conventions)
Go sử dụng **CamelCase** (hay MixedCaps) thay vì `snake_case`.

### 2.1. Public vs Private (Exported/Unexported)
Trong Go không có từ khóa `public`, `private` hay `protected`. Quyền truy cập được quyết định bằng ký tự đầu tiên của tên:
- **Viết hoa chữ cái đầu (Exported):** Các biến, hàm, struct có thể được truy cập từ package khác (tương đương Public). Ví dụ: `type User struct`, `func GetUserID()`.
- **Viết thường chữ cái đầu (Unexported):** Chỉ có thể sử dụng bên trong cùng một package (tương đương Private). Ví dụ: `func calculateSum()`, `var dbConn`.

### 2.2. Từ viết tắt (Initialisms)
Các từ viết tắt chuẩn (như ID, HTTP, URL, API, JSON) phải được viết hoa toàn bộ hoặc viết thường toàn bộ (nếu là private).
- **Đúng:** `UserID`, `APIKey`, `serveHTTP`, `jsonParser`.
- **Sai:** `UserId`, `ApiKey`, `ServeHttp`.

### 2.3. Tên Package, Struct, Interface và Biến
- **Package:** Tên package phải ngắn gọn, viết thường toàn bộ, không dùng dấu gạch dưới `_`. (Ví dụ: `package time`, `package entity`, không dùng `package my_utils`).
- **Interface:** Tên interface thường kết thúc bằng đuôi `-er` (Ví dụ: `Reader`, `Writer`, `Formatter`, `UserRepository`).
- **Biến (Variables):** Tên biến càng ngắn càng tốt, đặc biệt là trong các phạm vi hẹp. Ví dụ: dùng `i` thay cho `index`, dùng `u` thay cho `user` nếu hàm rất ngắn.

## 3. Cấu trúc Code và Logic (Code Flow)

### 3.1. Xử lý lỗi (Error Handling)
Go không có khối `try...catch`. Lỗi (Error) được coi như một giá trị trả về bình thường.
- Luôn kiểm tra lỗi ngay lập tức (`if err != nil`).
- Trả lỗi về cho hàm gọi nó thay vì giấu lỗi đi.

```go
// Tốt
user, err := repo.GetUser(id)
if err != nil {
    return nil, err
}

// Xấu (Bỏ qua lỗi)
user, _ := repo.GetUser(id) 
```

### 3.2. Return sớm (Early Returns) và Hạn chế `else`
Luôn xử lý các trường hợp ngoại lệ/lỗi và `return` ngay lập tức. Cố gắng giữ cho đoạn code thành công (happy path) nằm ở lề ngoài cùng bên trái (không bị thụt lề bởi `else`).

```go
// Xấu (Dùng else)
if err == nil {
    // Logic code chính bị thụt lề
    fmt.Println("Thành công")
} else {
    return err
}

// Tốt (Early Return - Trả về sớm)
if err != nil {
    return err
}
// Logic code chính nằm thẳng hàng
fmt.Println("Thành công")
```

## 4. Viết Comment (Tài liệu nội bộ)
- Bất kỳ một struct, hàm hay biến nào được **Exported** (viết hoa chữ đầu) đều phải có một comment mô tả ở ngay dòng bên trên.
- Comment phải bắt đầu bằng chính tên của hàm/struct đó và là một câu hoàn chỉnh.

```go
// User đại diện cho một người dùng ẩn danh trong hệ thống chat.
type User struct { ... }

// GetUserID trả về ID của người dùng dựa trên email.
func GetUserID(email string) string { ... }
```

## 5. Tổ chức dự án (Project Layout)
Tuân thủ chặt chẽ kiến trúc trong thư mục `internal/`:
- Không khai báo biến toàn cục (global variables) tràn lan. Dùng Dependency Injection (truyền struct) để chia sẻ các kết nối DB/Logger.
- `entity`: Chỉ chứa struct định nghĩa data.
- `repo`: Chỉ chứa code liên quan đến Database (SQL, truy vấn).
- `usecase`: Chứa logic nghiệp vụ cốt lõi (ví dụ: cấm user, ghép cặp).
- `controller`: Chỉ nhận request (HTTP/gRPC/WS), bóc tách dữ liệu và gọi UseCase.
