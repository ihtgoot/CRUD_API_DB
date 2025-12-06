# 📚 Bookstore Management System

A full-stack CRUD (Create, Read, Update, Delete) application built with Go and vanilla JavaScript. Features a RESTful API backend with MySQL database and a clean black & white dark mode frontend.

## 🚀 Features

- **Complete CRUD Operations**: Add, view, update, and delete books
- **RESTful API**: Well-structured API endpoints following REST principles
- **Search by ID**: Find specific books by their unique identifier
- **Dark Mode UI**: Minimalist black and white interface
- **Responsive Design**: Works on desktop and mobile devices
- **CORS Enabled**: Frontend and backend can run on different ports
- **Database Persistence**: All data stored in MySQL using GORM ORM

## 🛠️ Tech Stack

### Backend
- **Language**: Go (Golang)
- **Framework**: Gorilla Mux (routing)
- **Database**: MySQL
- **ORM**: GORM
- **CORS**: rs/cors

### Frontend
- **HTML5**: Structure
- **CSS3**: Styling (Pure CSS, no frameworks)
- **JavaScript**: Vanilla JS (ES6+)

## 📁 Project Structure

```
CRUD_API_DB/
├── cmd/
│   └── main/
│       └── main.go           # Application entry point
├── pkg/
│   ├── config/
│   │   └── app.go            # Database configuration
│   ├── controller/
│   │   └── book-controller.go # Request handlers
│   ├── models/
│   │   └── book.go           # Database models
│   ├── routs/
│   │   └── bookstore-routs.go # Route definitions
│   └── utils/
│       └── utils.go          # Helper functions
├── frontend/
│   └── index.html            # Frontend UI
├── go.mod                    # Go dependencies
└── README.md                 # This file
```

## 🔧 Prerequisites

Before running this application, make sure you have:

- **Go** (version 1.16 or higher)
- **MySQL** (version 5.7 or higher)
- **Git** (for cloning the repository)

## 📦 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/ihtgoot/CRUD_API_DB.git
cd CRUD_API_DB
```

### 2. Install Go Dependencies

```bash
go mod download
```

### 3. Set Up MySQL Database

```sql
-- Login to MySQL
mysql -u root -p

-- Create database
CREATE DATABASE simplerest;

-- Create user (optional but recommended)
CREATE USER 'myuser'@'localhost' IDENTIFIED BY 'strong_password';
GRANT ALL PRIVILEGES ON simplerest.* TO 'myuser'@'localhost';
FLUSH PRIVILEGES;
```

### 4. Configure Database Connection

Edit `pkg/config/app.go` if needed:

```go
db, err := gorm.Open("mysql", "myuser:strong_password@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
```

Update the connection string with your MySQL credentials:
- Username: `myuser`
- Password: `strong_password`
- Database: `simplerest`

## 🚀 Running the Application

### 1. Start the Backend Server

```bash
go run cmd/main/main.go
```

You should see:
```
🚀 Server running on http://localhost:9010
📁 Serving frontend from ./frontend
🔌 API available at http://localhost:9010/api
```

### 2. Access the Frontend

Open your browser and navigate to:
```
http://localhost:9010/
```

## 📡 API Endpoints

### Base URL
```
http://localhost:9010/api
```

### Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/book/` | Get all books |
| `GET` | `/book/{id}` | Get book by ID |
| `POST` | `/book/` | Create new book |
| `PUT` | `/book/{id}` | Update book by ID |
| `DELETE` | `/book/{id}` | Delete book by ID |

### API Examples

#### Get All Books
```bash
curl http://localhost:9010/api/book/
```

#### Get Book by ID
```bash
curl http://localhost:9010/api/book/1
```

#### Create New Book
```bash
curl -X POST http://localhost:9010/api/book/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Programming Language",
    "author": "Alan Donovan",
    "publication": "Addison-Wesley"
  }'
```

#### Update Book
```bash
curl -X PUT http://localhost:9010/api/book/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Book Name",
    "author": "Updated Author"
  }'
```

#### Delete Book
```bash
curl -X DELETE http://localhost:9010/api/book/1
```

## 🎨 Frontend Features

### Navigation Tabs
1. **All Books**: View complete list of all books
2. **Search Book by ID**: Find specific book by entering its ID
3. **Add Book**: Create new book entry
4. **Update Book**: Modify existing book details
5. **Delete Book**: Remove book from database

### UI Highlights
- **Dark Mode**: Black background with white text
- **Tab Navigation**: Clean separation of CRUD operations
- **Form Validation**: Required fields and input validation
- **Alert Messages**: Success/error notifications
- **Responsive Grid**: Books displayed in responsive card layout

## 🗄️ Database Schema

### Books Table

| Column | Type | Description |
|--------|------|-------------|
| `ID` | INT (Primary Key) | Auto-incrementing ID |
| `name` | VARCHAR | Book title |
| `author` | VARCHAR | Author name |
| `publication` | VARCHAR | Publisher name |
| `CreatedAt` | TIMESTAMP | Creation timestamp |
| `UpdatedAt` | TIMESTAMP | Last update timestamp |
| `DeletedAt` | TIMESTAMP (nullable) | Soft delete timestamp |

## 🔐 CORS Configuration

The application is configured to accept requests from any origin:

```go
AllowedOrigins: []string{"*"}
AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}
AllowedHeaders: []string{"Content-Type"}
```

**⚠️ Security Note**: In production, replace `"*"` with your specific domain(s).

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Find process using port 9010
lsof -i :9010

# Kill the process
kill -9 <PID>
```

### Database Connection Error
- Verify MySQL is running: `sudo systemctl status mysql`
- Check credentials in `pkg/config/app.go`
- Ensure database `simplerest` exists

### CORS Error in Browser
- Make sure backend is running on port 9010
- Check browser console for specific error
- Verify API_URL in frontend matches backend

### Books Not Displaying
- Check browser console (F12) for JavaScript errors
- Verify API endpoint: `http://localhost:9010/api/book/`
- Ensure database has data

## 🧪 Testing

### Manual Testing

Test each endpoint using curl or tools like Postman:

```bash
# Test server is running
curl http://localhost:9010/api/book/

# Add test data
curl -X POST http://localhost:9010/api/book/ \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Book","author":"Test Author","publication":"Test Pub"}'
```

## 📚 Learning Resources

This project demonstrates:
- **Go Web Development**: HTTP servers, routing, middleware
- **RESTful API Design**: Resource-based URLs, HTTP methods
- **Database Integration**: GORM ORM, MySQL connections
- **Frontend Integration**: Fetch API, async/await, DOM manipulation
- **Project Structure**: Package organization, separation of concerns

### Key Concepts Covered

1. **Subrouters**: API versioning and route grouping
2. **CORS**: Cross-origin resource sharing
3. **Middleware**: Request processing pipeline
4. **ORM**: Object-relational mapping with GORM
5. **MVC Pattern**: Models, Views (frontend), Controllers
6. **REST Principles**: Stateless, resource-based architecture

## 🚀 Future Enhancements

- [ ] Add user authentication (JWT)
- [ ] Input validation and sanitization
- [ ] Pagination for large datasets
- [ ] Search functionality (by name, author)
- [ ] Unit and integration tests
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] API documentation (Swagger)
- [ ] Book categories/tags
- [ ] Book cover image upload
- [ ] Advanced filtering and sorting

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👤 Author

**Your Name**
- GitHub: [@ihtgoot](https://github.com/ihtgoot)
- Email: your.email@example.com

## 🙏 Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) - Powerful HTTP router
- [GORM](https://gorm.io/) - Fantastic ORM library
- [rs/cors](https://github.com/rs/cors) - CORS middleware

---

⭐ **Star this repo if you found it helpful!**

📝 **Questions or Issues?** Open an issue on GitHub

🎓 **Learning Go?** This is a great starter project for understanding web development in Go!
