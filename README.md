# Quản Lý Rác Thải Thông Minh

## Giới Thiệu Dự Án

Dự án Quản Lý Rác Thải Thông Minh được phát triển nhằm nâng cao hiệu quả quản lý rác thải, giảm thiểu ô nhiễm môi trường và tối ưu hóa quy trình xử lý rác thải. Hệ thống bao gồm ba thành phần chính: Backend (BE) sử dụng Golang, Frontend (FE) Web Admin, và Mobile App dành cho nhân viên sử dụng.

## Mục Tiêu Dự Án

- Tăng cường hiệu quả quản lý và xử lý rác thải.
- Giảm thiểu ô nhiễm môi trường.
- Tối ưu hóa quy trình thu gom và xử lý rác thải.

## Phạm Vi Dự Án

Dự án bao gồm ba thành phần chính:
1. Backend (BE) Golang: Xử lý dữ liệu, cung cấp API cho các thành phần khác.
2. Frontend (FE) Web Admin: Giao diện quản lý dành cho admin, hiển thị thông tin và báo cáo.
3. Mobile App: Ứng dụng di động cho nhân viên, hỗ trợ theo dõi và quản lý rác thải.

## Kiến Trúc Hệ Thống

### Thành Phần Hệ Thống

- **Backend Golang**: Xử lý dữ liệu và cung cấp API.
- **Frontend Web Admin**: Giao diện quản lý dành cho admin.
- **Mobile App**: Ứng dụng di động dành cho nhân viên.

### Mô Hình Kiến Trúc

Hệ thống được thiết kế theo mô hình Microservices, đảm bảo tính linh hoạt và khả năng mở rộng. 

## Chi Tiết Các Thành Phần

### Backend (Golang)

- **Ngôn ngữ**: Golang
- **Chức năng chính**: Xử lý dữ liệu, cung cấp API, quản lý cơ sở dữ liệu.
- **Công nghệ**: Gin, GORM, PostgreSQL

### Frontend (Web Admin)

- **Ngôn ngữ**: JavaScript
- **Framework**: React.js
- **Chức năng chính**: Quản lý người dùng, hiển thị báo cáo và thống kê.

### Mobile App

- **Ngôn ngữ**: Dart
- **Framework**: Flutter
- **Chức năng chính**: Theo dõi và quản lý rác thải, cập nhật trạng thái.

## Cài Đặt Hệ Thống

### Yêu Cầu Hệ Thống

- **Máy chủ**: Ubuntu 20.04 hoặc CentOS 8
- **Ngôn ngữ**: Golang 1.22+, Node.js 14+

### Các Bước Cài Đặt

1. **Cài đặt Golang**:
    ```sh
    sudo apt update
    sudo apt install golang-go
    ```
2. **Cài đặt Node.js**:
    ```sh
    sudo apt update
    sudo apt install nodejs npm
    ```
3. **Cài đặt Flutter**:
    ```sh
    sudo snap install flutter --classic
    ```
4. **Clone các repository của dự án**:
    ```sh
    git clone https://github.com/your-repo/backend
    git clone https://github.com/your-repo/frontend
    git clone https://github.com/your-repo/mobile-app
    ```
5. **Cấu hình các biến môi trường và database**:
    ```sh
    export DB_HOST=your_db_host
    export DB_USER=your_db_user
    export DB_PASSWORD=your_db_password
    ```

## Hướng Dẫn Sử Dụng

### Backend (Golang)

- **Cấu Hình**: Cấu hình file `config.yaml` với thông tin cơ sở dữ liệu và các biến môi trường.
- **Khởi Chạy**:
    ```sh
    go run main.go
    ```

### Frontend (Web Admin)

- **Cài Đặt**:
    ```sh
    npm install
    npm start
    ```
- **Đăng Nhập**: Sử dụng tài khoản admin để đăng nhập và quản lý hệ thống.

### Mobile App

- **Cài Đặt Ứng Dụng**: Tải ứng dụng từ Google Play hoặc App Store.
- **Đăng Nhập**: Sử dụng tài khoản nhân viên để đăng nhập và theo dõi các nhiệm vụ.

## Tài Liệu Chuyển Giao Công Nghệ

### Mục Tiêu Chuyển Giao

Cung cấp chi tiết về quá trình chuyển giao công nghệ hệ thống quản lý rác thải thông minh, bao gồm các công nghệ sử dụng, quy trình và kế hoạch chuyển giao.

### Chi Tiết Kỹ Thuật

- **Công Nghệ Sử Dụng**:
    - Backend: Golang
    - Frontend: React.js
    - Mobile App: React Native

### Quy Trình Chuyển Giao

- **Đào Tạo**: Cung cấp các buổi đào tạo cho đội ngũ kỹ thuật của khách hàng.
- **Bảo Trì và Hỗ Trợ**: Cung cấp dịch vụ bảo trì và hỗ trợ sau khi chuyển giao.

### Kế Hoạch Chuyển Giao

- **Lịch Trình**:
    - Giai đoạn 1: Cài đặt và cấu hình hệ thống.
    - Giai đoạn 2: Đào tạo và hướng dẫn sử dụng.
    - Giai đoạn 3: Chuyển giao và hỗ trợ ban đầu.

### Liên Hệ Hỗ Trợ

- Email: nguyenmanh180102@gmail.com
- Điện thoại: +84 977 683 533

---


# Golang

## API Reference - test localhost:3000

#### Sign Up

```
  POST /users/signup
```

| Parameter   | Type     | Description                 |
| :---------- | :------- | :-------------------------- |
| `full_name` | `string` |                             |
| `email`     | `string` | **Required**. Your email    |
| `password`  | `string` | **Required**. Your password |

#### Sign In

```
  POST /users/signin
```

| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `email`    | `string` | **Required**. Your email    |
| `password` | `string` | **Required**. Your password |

#### Create

```
  POST /api/create_books
```

| Parameter   | Type     | Description                       |
| :---------- | :------- | :-------------------------------- |
| `author`    | `string` | **Required**. User_id in database |
| `title`     | `string` | **Required**. Custom              |
| `publisher` | `string` | **Required**. Custom              |

#### Delete

```
  DELETE /api/delete_book/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Get-By-ID

```
  GET /api/get_books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Get-All

```
  GET /api/books
```


