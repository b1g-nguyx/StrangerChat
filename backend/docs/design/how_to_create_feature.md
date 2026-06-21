# Hướng dẫn tạo một Tính năng / API mới từ A-Z

Hệ thống của chúng ta áp dụng **Clean Architecture** (Kiến trúc Sạch). Điều này có nghĩa là luồng dữ liệu bắt buộc phải đi qua các lớp (layers) theo một chiều nhất định. Bạn không bao giờ được phép gọi thẳng Database từ Controller.

Dưới đây là quy trình 5 bước chuẩn mực để tạo ra một API mới (Ví dụ: Tính năng `GetUserDetails`).

## Bước 1: Định nghĩa Dữ liệu tại `internal/entity/`
Đây là lõi trong cùng của hệ thống. Bạn bắt đầu bằng việc định nghĩa cấu trúc dữ liệu.

*   **Vị trí:** `internal/entity/user.go`
*   **Hành động:** Tạo struct.
```go
package entity

type User struct {
    ID     string `json:"id"`
    Status string `json:"status"`
}
```

## Bước 2: Truy vấn Database tại `internal/repo/`
Lớp này **chỉ** làm nhiệm vụ giao tiếp với Database (Postgres, Redis...). Không chứa logic kinh doanh ở đây.

*   **Vị trí:** `internal/repo/user_repo.go` (hoặc trong thư mục `persistent/`)
*   **Hành động:** 
    1. Định nghĩa một `Interface`.
    2. Viết struct thực thi (Implement) interface đó.
```go
package repo

import "your_project/internal/entity"

// 1. Định nghĩa Interface
type UserRepo interface {
    GetByID(id string) (*entity.User, error)
}

// 2. Struct thực thi
type userRepoImpl struct {
    db *sql.DB // Khai báo kết nối DB
}

func NewUserRepo(db *sql.DB) UserRepo {
    return &userRepoImpl{db: db}
}

// 3. Viết câu SQL
func (r *userRepoImpl) GetByID(id string) (*entity.User, error) {
    // Thực thi câu lệnh SQL SELECT * FROM users WHERE id = ?
    return &entity.User{ID: id, Status: "Waiting"}, nil
}
```

## Bước 3: Logic Nghiệp vụ tại `internal/usecase/`
Đây là bộ não của hệ thống. Nó sẽ gọi `Repo` để lấy dữ liệu, sau đó kiểm tra điều kiện (Ví dụ: User có bị ban không?), rồi mới trả kết quả ra ngoài.

*   **Vị trí:** `internal/usecase/user_usecase.go`
*   **Hành động:**
```go
package usecase

import (
    "your_project/internal/entity"
    "your_project/internal/repo"
)

// 1. Định nghĩa Interface
type UserUseCase interface {
    GetUserDetails(id string) (*entity.User, error)
}

// 2. Struct chứa Repo (Dependency Injection)
type userUseCaseImpl struct {
    repo repo.UserRepo
}

func NewUserUseCase(r repo.UserRepo) UserUseCase {
    return &userUseCaseImpl{repo: r}
}

// 3. Xử lý Logic
func (uc *userUseCaseImpl) GetUserDetails(id string) (*entity.User, error) {
    user, err := uc.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    
    // Thêm logic kiểm tra gì đó ở đây...
    if user.Status == "Banned" {
        return nil, errors.New("user is banned")
    }

    return user, nil
}
```

## Bước 4: Nhận Request tại `internal/controller/`
Lớp này làm nhiệm vụ nghe HTTP Request (hoặc gRPC/WebSocket), lấy dữ liệu từ URL/Body, gọi `UseCase`, và trả về JSON cho client.

*   **Vị trí:** `internal/controller/restapi/user_controller.go`
*   **Hành động:**
```go
package restapi

import (
    "encoding/json"
    "net/http"
    "your_project/internal/usecase"
)

type UserController struct {
    usecase usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController {
    return &UserController{usecase: uc}
}

// Hàm xử lý API
func (c *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Lấy ID từ URL
    userID := r.URL.Query().Get("id")

    // 2. Gọi UseCase
    user, err := c.usecase.GetUserDetails(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 3. Trả JSON về cho Client
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

## Bước 5: Lắp ráp (Wiring) và Cấu hình Route tại `cmd/`
Bước cuối cùng là khởi tạo tất cả các thành phần trên và nối chúng lại với nhau (Dependency Injection), sau đó cắm vào Router để chạy server.

*   **Vị trí:** Thường ở `cmd/app/main.go` hoặc thư mục cấu hình router.
*   **Hành động:**
```go
package main

import (
    "net/http"
    // Import các package repo, usecase, controller của bạn
)

func main() {
    // 0. Khởi tạo DB (giả sử đã có dbConn)
    // dbConn := ...

    // 1. Khởi tạo Repo
    userRepo := repo.NewUserRepo(dbConn)

    // 2. Bơm Repo vào UseCase
    userUseCase := usecase.NewUserUseCase(userRepo)

    // 3. Bơm UseCase vào Controller
    userCtrl := restapi.NewUserController(userUseCase)

    // 4. Khai báo Route API
    http.HandleFunc("/api/v1/user", userCtrl.GetUserHandler)

    // 5. Chạy Server
    http.ListenAndServe(":8080", nil)
}
```

### Tóm tắt Luồng đi:
`Route -> Controller -> UseCase -> Repo -> DB` và trả ngược lại kết quả.
