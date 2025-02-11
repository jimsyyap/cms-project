# **CMS Project: Members Website**

This project is a **Content Management System (CMS)** designed for a members-only website. It includes a **Golang backend**, **PostgreSQL database**, and a **ReactJS frontend**. The system supports user authentication, role-based access control, content management, and SEO features.

---

## **Features**

### **Backend (Golang + PostgreSQL)**
- **User Authentication**: Registration, login, and JWT-based authentication.
- **Role-Based Access Control**: Admins, editors, and members.
- **Content Management**: CRUD operations for posts, pages, and media.
- **WYSIWYG Editor**: Integration with TinyMCE or Quill for content creation.
- **SEO Management**: Customizable meta titles, descriptions, and slugs.
- **File Storage**: AWS S3 for storing media files.

### **Frontend (ReactJS)**
- **User Interface**: Dashboard, member area, and public-facing pages.
- **WYSIWYG Editor**: Integrated for content creation/editing.
- **Responsive Design**: Built with TailwindCSS for a modern, responsive UI.
- **API Integration**: Axios for interacting with the backend.

---

## **Tech Stack**

### **Backend**
- **Language**: Golang
- **Framework**: Gin
- **Database**: PostgreSQL
- **Authentication**: JWT
- **File Storage**: AWS S3

### **Frontend**
- **Language**: JavaScript
- **Framework**: ReactJS
- **Routing**: React Router
- **Styling**: TailwindCSS
- **API Client**: Axios

### **Deployment**
- **Backend**: AWS Elastic Beanstalk or ECS
- **Frontend**: AWS S3 + CloudFront
- **Database**: Amazon RDS (PostgreSQL)

---

## **Project Structure**

```
cms-project/
├── backend/                  # Golang backend
│   ├── cmd/                  # Main application entry points
│   ├── internal/             # Internal application code
│   ├── migrations/           # Database migration files
│   ├── .env                  # Environment variables
│   ├── go.mod                # Go module file
│   └── main.go               # Main application file
│
├── frontend/                 # ReactJS frontend
│   ├── public/               # Static assets
│   ├── src/                  # React source code
│   ├── .env                  # Environment variables
│   ├── package.json          # NPM dependencies
│   └── tailwind.config.js    # TailwindCSS configuration
│
├── docker-compose.yml        # Docker setup for local development
├── Dockerfile                # Dockerfile for backend
├── README.md                 # Project documentation
└── .gitignore                # Git ignore file
```

---

## **Setup Instructions**

### **Prerequisites**
- Go (1.20+)
- Node.js (16+)
- PostgreSQL
- Docker (optional, for local development)

---

### **Backend Setup**

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/cms-project.git
   cd cms-project/backend
   ```

2. **Install Dependencies**:
   ```bash
   go mod download
   ```

3. **Set Up Environment Variables**:
   Create a `.env` file in the `backend` directory:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your-db-user
   DB_PASSWORD=your-db-password
   DB_NAME=cmsdb
   JWT_SECRET=your-jwt-secret
   ```

4. **Run Database Migrations**:
   Use a tool like `golang-migrate` to apply migrations:
   ```bash
   migrate -path ./migrations -database "postgres://user:password@localhost:5432/cmsdb?sslmode=disable" up
   ```

5. **Start the Backend Server**:
   ```bash
   go run main.go
   ```

---

### **Frontend Setup**

1. **Navigate to the Frontend Directory**:
   ```bash
   cd ../frontend
   ```

2. **Install Dependencies**:
   ```bash
   npm install
   ```

3. **Set Up Environment Variables**:
   Create a `.env` file in the `frontend` directory:
   ```env
   VITE_API_BASE_URL=http://localhost:8080
   ```

4. **Start the Frontend Development Server**:
   ```bash
   npm run dev
   ```

---

### **Docker Setup (Optional)**

1. **Build and Run the Backend with Docker**:
   ```bash
   docker-compose up --build
   ```

2. **Access the Application**:
   - Backend: `http://localhost:8080`
   - Frontend: `http://localhost:3000`

---

## **Deployment to AWS**

### **Backend**
1. **Dockerize the Backend**:
   - Create a `Dockerfile` for the Go backend.
   - Push the Docker image to AWS ECR.

2. **Deploy to Elastic Beanstalk or ECS**:
   - Use the AWS Management Console or CLI to deploy the Dockerized backend.

### **Frontend**
1. **Build the React App**:
   ```bash
   npm run build
   ```

2. **Deploy to S3 + CloudFront**:
   - Upload the `build` folder to an S3 bucket.
   - Configure CloudFront as a CDN for the S3 bucket.

### **Database**
1. **Set Up Amazon RDS**:
   - Create a PostgreSQL instance on Amazon RDS.
   - Update the backend's database connection string.

---

## **Contributing**

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**

- [Gin Framework](https://gin-gonic.com/)
- [TailwindCSS](https://tailwindcss.com/)
- [AWS Documentation](https://aws.amazon.com/documentation/)

---

Let me know if you need further assistance! 🚀
