# Smart Waste System Management - Hệ thống quản lý thu gom rác thải thông minh

## Giới Thiệu Dự Án

Dự án Quản Lý Rác Thải Thông Minh được phát triển nhằm nâng cao hiệu quả quản lý rác thải, giảm thiểu ô nhiễm môi trường và tối ưu hóa quy trình xử lý rác thải. Hệ thống bao gồm ba thành phần chính: Backend (BE) sử dụng Golang, Frontend (FE) Web Admin, và Mobile App dành cho nhân viên sử dụng.

## Phạm Vi Dự Án

Dự án bao gồm ba thành phần chính:
1. Backend (BE) Golang: Xử lý dữ liệu, cung cấp API cho các thành phần khác.
2. Frontend (FE) Web Admin: Giao diện quản lý dành cho admin, hiển thị thông tin và báo cáo.
3. Mobile App: Ứng dụng di động cho nhân viên, hỗ trợ theo dõi và quản lý rác thải.

## Giới thiệu về Công nghệ

### Golang
**Định nghĩa:** Golang, hay còn gọi là Go, là ngôn ngữ lập trình mã nguồn mở được phát triển bởi Google vào năm 2009. Nó được thiết kế bởi Robert Griesemer, Rob Pike và Ken Thompson, với mục tiêu chính là đơn giản hóa việc lập trình, giảm thiểu thời gian biên dịch và tăng hiệu suất thực thi. Golang kế thừa nhiều đặc điểm từ các ngôn ngữ lập trình khác như C và C++, nhưng tập trung vào việc loại bỏ những yếu tố phức tạp không cần thiết để tạo ra một ngôn ngữ hiện đại, mạnh mẽ và dễ sử dụng.

**Ưu điểm:**
- **Tốc độ xử lý nhanh:** Golang biên dịch trực tiếp xuống mã máy, giúp tăng tốc độ thực thi ứng dụng. Nhờ vào việc biên dịch này, các chương trình viết bằng Golang thường có hiệu suất tương đương với các chương trình viết bằng ngôn ngữ C hay C++, nhưng lại có thời gian biên dịch nhanh hơn nhiều.
- **Hiệu suất cao:** Golang được trang bị một bộ thu gom rác (garbage collector) rất hiệu quả, giúp quản lý bộ nhớ tự động và giảm thiểu rò rỉ bộ nhớ. Điều này làm cho Golang trở thành một lựa chọn lý tưởng cho các hệ thống yêu cầu hiệu suất cao và ổn định. Ngoài ra, Golang có một thư viện chuẩn rất phong phú, giúp lập trình viên có sẵn các công cụ để xây dựng ứng dụng hiệu quả.
- **Hỗ trợ đa luồng (concurrency):** Một trong những đặc điểm nổi bật của Golang là khả năng xử lý đồng thời (concurrency) rất mạnh mẽ. Golang sử dụng goroutines và channels để quản lý các tác vụ đồng thời. Goroutines là các luồng nhẹ, có thể khởi tạo hàng nghìn goroutines chỉ với một lượng nhỏ tài nguyên hệ thống. Channels cung cấp cơ chế giao tiếp giữa các goroutines, giúp dễ dàng truyền dữ liệu mà không cần lo lắng về các vấn đề đồng bộ hóa.
- **Cú pháp đơn giản và dễ hiểu:** Golang được thiết kế với cú pháp rất gọn gàng và dễ đọc. Ngôn ngữ này loại bỏ nhiều từ khóa và cấu trúc phức tạp thường thấy trong các ngôn ngữ khác, giúp cho mã nguồn trở nên trong sáng và dễ duy trì. Golang cũng tuân theo một chuẩn mã hóa rất nghiêm ngặt, với các công cụ như gofmt giúp tự động định dạng mã nguồn theo chuẩn, tạo nên sự nhất quán và dễ đọc trong cộng đồng lập trình viên.
### React
**Định nghĩa:** React là thư viện JavaScript mã nguồn mở, được phát triển bởi Facebook, dùng để xây dựng giao diện người dùng (UI). Ra đời vào năm 2013, React cho phép lập trình viên xây dựng các thành phần giao diện (components) có thể tái sử dụng và dễ dàng quản lý. Mỗi component trong React có thể coi như một khối xây dựng độc lập, kết hợp với nhau để tạo nên giao diện người dùng hoàn chỉnh. Điều này giúp việc phát triển và bảo trì ứng dụng trở nên hiệu quả hơn.

**Ưu điểm:**
- **Khả năng tái sử dụng component:** React cho phép xây dựng các components độc lập, có thể tái sử dụng trong nhiều phần của ứng dụng. Điều này không chỉ giúp tiết kiệm thời gian phát triển mà còn làm cho mã nguồn trở nên dễ quản lý và bảo trì. Khi một component được thay đổi, tất cả các nơi sử dụng component đó đều được cập nhật tự động, giảm thiểu lỗi và đảm bảo tính nhất quán trong giao diện người dùng.
- **Tối ưu hóa hiệu suất thông qua Virtual DOM:** Một trong những điểm mạnh của React là việc sử dụng Virtual DOM để cập nhật giao diện. Virtual DOM là một bản sao của DOM thực, được lưu trữ trong bộ nhớ. Khi có sự thay đổi trong trạng thái của ứng dụng, React cập nhật Virtual DOM trước, sau đó so sánh với DOM thực để xác định những thay đổi cần thiết. Điều này giúp giảm thiểu số lần thao tác trực tiếp trên DOM thực, tăng cường hiệu suất và trải nghiệm người dùng.
- **Cộng đồng hỗ trợ mạnh mẽ:** React có một cộng đồng phát triển lớn và năng động. Với hàng triệu lập trình viên trên khắp thế giới sử dụng React, có rất nhiều tài nguyên, thư viện và công cụ hỗ trợ có sẵn. Các diễn đàn như Stack Overflow, các nhóm trên GitHub, và các kênh truyền thông xã hội đều là nguồn tài nguyên quý giá cho lập trình viên học hỏi và giải quyết các vấn đề gặp phải khi làm việc với React.
- **Tính linh hoạt và khả năng mở rộng:** React không chỉ giới hạn trong việc xây dựng giao diện người dùng cho ứng dụng web, mà còn có thể được sử dụng trong các dự án di động thông qua React Native. Điều này mang lại tính linh hoạt cao cho React, cho phép lập trình viên xây dựng ứng dụng trên nhiều nền tảng khác nhau mà vẫn giữ nguyên logic và cấu trúc mã.
### React Native
**Định nghĩa:** React Native là framework phát triển ứng dụng di động đa nền tảng sử dụng React. Được phát hành bởi Facebook vào năm 2015, React Native cho phép lập trình viên sử dụng JavaScript để xây dựng các ứng dụng di động chạy trên cả iOS và Android. React Native tận dụng cùng một hệ sinh thái component của React, nhưng các component này được biên dịch xuống mã native, giúp ứng dụng có hiệu suất cao như các ứng dụng native truyền thống.

**Ưu điểm:**
- **Viết một lần, chạy trên cả iOS và Android:** React Native cho phép lập trình viên viết mã một lần và triển khai trên cả hai nền tảng di động phổ biến. Điều này không chỉ giúp tiết kiệm thời gian và công sức phát triển mà còn giảm thiểu chi phí bảo trì và nâng cấp ứng dụng. Một mã nguồn chung cho cả hai nền tảng cũng giúp đảm bảo tính nhất quán trong trải nghiệm người dùng.
- **Hiệu suất gần như native:** React Native biên dịch mã JavaScript xuống mã native, đảm bảo hiệu suất cao cho ứng dụng. Điều này giúp React Native khác biệt so với các framework hybrid khác, mang lại trải nghiệm người dùng mượt mà và phản hồi nhanh chóng, tương tự như các ứng dụng native được viết bằng Swift hay Kotlin.
- **Hỗ trợ plugin native:** React Native cho phép sử dụng các plugin native để tận dụng tối đa tính năng của hệ điều hành. Lập trình viên có thể dễ dàng tích hợp các mô-đun native viết bằng Swift (cho iOS) hoặc Java/Kotlin (cho Android) vào ứng dụng React Native của mình. Điều này giúp mở rộng khả năng của ứng dụng và tận dụng tối đa các API của hệ điều hành.
- **Cộng đồng và hệ sinh thái phong phú:** React Native có một cộng đồng lớn mạnh và một hệ sinh thái phong phú với hàng nghìn thư viện và plugin có sẵn. Điều này giúp lập trình viên dễ dàng tìm kiếm và tích hợp các giải pháp có sẵn vào ứng dụng của mình, từ các thành phần giao diện người dùng đến các công cụ quản lý trạng thái và các mô-đun tương tác với phần cứng.

### PostgreSQL
**Định nghĩa:** PostgreSQL, thường được gọi là Postgres, là hệ quản trị cơ sở dữ liệu quan hệ đối tượng mã nguồn mở tiên tiến. Được phát triển từ những năm 1986 tại Đại học California, Berkeley, PostgreSQL được thiết kế để hỗ trợ các tính năng tiên tiến như xử lý giao dịch (ACID), lưu trữ đối tượng và các loại dữ liệu phong phú. PostgreSQL tuân thủ tiêu chuẩn SQL và hỗ trợ nhiều tính năng mở rộng, bao gồm các ngôn ngữ lập trình, chỉ mục phức tạp, và các phương thức tìm kiếm toàn văn.

**Ưu điểm:**
- **Tính năng phong phú:** PostgreSQL hỗ trợ một loạt các tính năng mạnh mẽ như truy vấn phức tạp, các hàm và thủ tục, khóa ngoại, trigger, và các kiểu dữ liệu tùy chỉnh. Điều này giúp PostgreSQL phù hợp với nhiều loại ứng dụng, từ hệ thống quản lý nội dung (CMS) đến các ứng dụng tài chính phức tạp.
- **Tính toàn vẹn và an toàn dữ liệu:** PostgreSQL tuân thủ các tiêu chuẩn ACID (Atomicity, Consistency, Isolation, Durability) đảm bảo tính toàn vẹn của giao dịch dữ liệu. Nó cũng cung cấp các tính năng bảo mật mạnh mẽ như xác thực người dùng, kiểm soát truy cập dựa trên vai trò (role-based access control) và mã hóa dữ liệu.
- **Khả năng mở rộng:** PostgreSQL có khả năng mở rộng theo cả chiều ngang và chiều dọc. Nó hỗ trợ các phương pháp sharding và replication, giúp phân phối dữ liệu trên nhiều máy chủ và đảm bảo hiệu suất cao và tính sẵn sàng của hệ thống. Các mô-đun mở rộng và plugin cũng có thể được tích hợp để thêm các tính năng mới mà không cần thay đổi mã nguồn chính.
- **Hỗ trợ SQL tiêu chuẩn và NoSQL:** PostgreSQL không chỉ hỗ trợ SQL tiêu chuẩn mà còn cung cấp các tính năng NoSQL như lưu trữ JSON, XML, và hstore. Điều này giúp các nhà phát triển dễ dàng xử lý các dữ liệu phi cấu trúc và bán cấu trúc bên cạnh dữ liệu quan hệ truyền thống.
- **Hiệu suất cao và tối ưu hóa:** PostgreSQL có các cơ chế tối ưu hóa truy vấn mạnh mẽ, như các chỉ mục B-tree, Hash, GiST, SP-GiST, GIN, và BRIN. Nó cũng hỗ trợ các phương thức tìm kiếm toàn văn, giúp cải thiện hiệu suất truy vấn và tìm kiếm dữ liệu.
#### Kết luận
Golang, React, React Native và PostgreSQL là những công cụ và công nghệ mạnh mẽ trong việc phát triển ứng dụng web và di động. Với những ưu điểm nổi bật và sự hỗ trợ mạnh mẽ từ cộng đồng, các công nghệ này giúp đội ngũ kỹ thuật nắm bắt và áp dụng một cách hiệu quả, từ việc xây dựng các hệ thống back-end hiệu suất cao bằng Golang và PostgreSQL đến việc phát triển các giao diện người dùng mượt mà và ứng dụng di động đa nền tảng với React và React Native. Việc kết hợp những công nghệ này sẽ mang lại lợi ích lớn cho các dự án công nghệ, giúp tăng cường hiệu suất, bảo mật và trải nghiệm người dùng.
### Thành Phần Hệ Thống

- **Backend Golang**: Xử lý dữ liệu và cung cấp API.
- **Frontend Web Admin**: Giao diện quản lý dành cho admin.
- **Mobile App**: Ứng dụng di động dành cho nhân viên. 

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
3. **Cài đặt React native**:
    ```sh
    sudo 
    ```
4. **Clone các repository của dự án**:
    ```sh
    git clone https://github.com/congmanh18/SWMS-server.git
    git clone https://github.com/congmanh18/SWMS-web.git
    git clone https://github.com/congmanh18/SWMS-mobile.git
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

# API User manual
# API User manual
## API Reference - localhost:8080

### User API
#### Sign Up
```
  POST /users/signup
```

| Parameter         | Type     | Description                |
| :----------       | :------- | :--------------------------|
| `first_name`      | `string` |                            |
| `middle_name`     | `string` |                            |
| `last_name`       | `string` |                            |
| `gender`          | `string` |                            |
| `cin`             | `string` |                            |
| `nationality`     | `string` |                            |
| `por`             | `string` |                           |
| `poo`             | `string` |                            |
| `role_name`       | `string` |                            |
| `phone`           | `string` |                            |
| `category`        | `string` |                            |
| `email`           | `string` | **Required**. Your email   |
| `password`        | `string` | **Required**. Your password|

```
{
    "first_name":   "Nguyen",
    "middle_name":  "Cong",
    "last_name":    "Manh",
    "gender":       "male",
    "cin":          "0522358281",
    "nationality":  "Viet Nam",
    "por":          "xom 4 Phu Nong, Hoai Son, Hoai Nhon, Binh Dinh",
    "poo":          "Phu Nong, Hoai Son, Hoai Nhon, Binh Dinh",
    "role_name":    "staff",
    "phone":        "0855475422",
    "password":     "12345679",
    "category":     "fulltime",
    "email":        "nguyenmanh180102@gmail.com"
}
``` 
#### Sign In
```
  POST /users/signin
```
| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `email`    | `string` | **Required**. Your email    |
| `password` | `string` | **Required**. Your password |

#### Read list User
```
  GET /user/listUser
``` 

#### Delete

```
  DELETE /users/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Read Information User

```
  GET /users/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Create

```
  PUT /users/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


### Area API

#### Create Area
```
  POST /area/create
```
| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `name`    | `string` | **Required**. Your area name    |
| `addres` | `string` | **Required**. Your area address |

```
{
    "name" : "Go Super Market",
    "address": "920 Nguyen Kiem"
}
```

#### Read list Area
```
  GET /area
``` 

#### Read Information Area

```
  GET /area/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Delete Area

```
  DELETE /area/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### Update Area

```
  PUT /area/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

### Equipment API

#### Create Equipment
```
  POST /equipment/create
```
| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `addres`    | `string` |     |
| `area_id` | `string` | **Required**. Your area_id in database |
| `weight`    | `string` |     |
| `level` | `string` |  |
| `latitude`    | `string` |    |
| `longitude` | `string` |  |

```
{
        "address": "so 10 Le Loi",
        "area_id": "56152651-efdf-4ec5-a5bf-cc1c0ee07f45",
        "weight": "10",
        "level": "8",
        "latitude": "10.762622",
        "longitude": "106.66017"
}
```

#### Read list Equipment
```
  GET /equipment
``` 

#### Read Information Equipment

```
  GET /equipment/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Delete Equipment

```
  DELETE /equipment/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### Update Equipment

```
  PUT /equipment/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

### Transaction API


#### Create Transaction
```
  POST /transaction/create
```
| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `user_id`    | `string` | **Required**. Your user_id in database    |
| `trash_bin_id` | `string` | **Required**. Your trash_bin_id in database |

```
{
    "user_id" : "46c43772-3739-4531-ab48-6bde1b6c010b",
    "trash_bin_id": "2a1c8033-a9ce-4564-963a-73e958111c8d",
}
```


#### Read list Transaction
```
  GET /transaction
``` 

#### Read Information Transaction

```
  GET /transaction/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Delete Transaction

```
  DELETE /transaction/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


#### Update Transaction

```
  PUT /transaction/${id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |



