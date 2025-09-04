# TinyURL Clone - Hướng dẫn chạy ứng dụng

## Cấu trúc dự án

```
tinyurl-go/
├── backend/
│   ├── handlers/
│   │   └── url_handlers.go
│   ├── models/
│   │   └── url_mapping.go
│   ├── utils/
│   │   └── slug.go
│   ├── main.go
│   ├── go.mod
│   └── go.sum
└── frontend/
    ├── index.html
    ├── css/
    │   └── style.css
    └── js/
        └── script.js
```

## Cách chạy ứng dụng

### Yêu cầu hệ thống

- Go 1.16 hoặc cao hơn
- Internet để tải dependencies

### Bước 1: Cài đặt dependencies

```bash
cd backend
go mod tidy
```

### Bước 2: Chạy ứng dụng

```bash
go run main.go
```

Ứng dụng sẽ chạy trên `http://localhost:8080`.

### Bước 3: Truy cập ứng dụng

Mở trình duyệt và truy cập `http://localhost:8080`.

## Cách sử dụng

1. Nhập URL bạn muốn rút gọn vào ô input.
2. Click vào nút "Shorten".
3. Ứng dụng sẽ trả về một URL rút gọn.
4. Khi truy cập URL rút gọn, bạn sẽ được chuyển hướng đến URL gốc.

## Ghi chú

Trong quá trình phát triển, mình đã gặp một số vấn đề với môi trường WSL khi thử nghiệm các server đơn giản (Go, Node.js, Python) không thể truy cập được từ localhost. Có thể có các nguyên nhân sau:

1. **Firewall**: Kiểm tra xem firewall có chặn port 8080 không.
2. **WSL Network Configuration**: Có thể cần cấu hình mạng cho WSL.
3. **Port Conflict**: Kiểm tra xem có ứng dụng nào khác đang sử dụng port 8080 không bằng lệnh `netstat -tulpn | grep :8080`.

Nếu bạn gặp vấn đề tương tự, hãy thử:
- Thay đổi port trong file `main.go` (thay `:8080` bằng `:8081` hoặc port khác).
- Chạy ứng dụng với quyền root: `sudo go run main.go`.
- Kiểm tra firewall và cấu hình mạng.

## Kiến trúc ứng dụng

### Backend (Go)

- Sử dụng Gorilla Mux để routing.
- Tạo slug ngẫu nhiên cho URL rút gọn.
- Lưu trữ ánh xạ giữa slug và URL gốc trong bộ nhớ (sử dụng map).
- API endpoint:
  - `POST /api/shorten` để rút gọn URL.
  - `GET /{slug}` để chuyển hướng đến URL gốc.

### Frontend (HTML/CSS/JS)

- Giao diện đơn giản với một input để nhập URL và một nút để rút gọn.
- Sử dụng Fetch API để gọi backend API.
- Hiển thị URL rút gọn khi thành công.
- Hiển thị lỗi khi có lỗi xảy ra.

## Tính năng

- Rút gọn URL dài thành URL ngắn.
- Chuyển hướng từ URL ngắn đến URL gốc.
- Giao diện người dùng đơn giản và trực quan.

## Hạn chế

- Dữ liệu được lưu trữ trong bộ nhớ, sẽ bị mất khi restart server.
- Không có xác thực hoặc bảo mật.
- Không có rate limiting.
- Không có kiểm tra slug trùng lặp.

## Cải tiến tiềm năng

- Sử dụng cơ sở dữ liệu (SQLite, PostgreSQL, MongoDB) để lưu trữ dữ liệu bền vững.
- Thêm xác thực và bảo mật.
- Thêm rate limiting để ngăn chặn lạm dụng.
- Thêm kiểm tra slug trùng lặp và tạo slug mới nếu cần.
- Thêm tính năng thống kê lượt truy cập.
- Thêm giao diện quản lý URL.