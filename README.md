# CourseScope CMU

![logo](./frontend/static/favicon-128x128.png)

CourseScope CMU is a web platform designed for Carnegie Mellon University (CMU) students to review and rate their courses. The website allows students to explore courses, instructors, and specific course offerings while accessing detailed ratings and statistics such as overall quality, teaching effectiveness, material quality, course value, workload, difficulty, and grading. Users can also share their experiences by commenting on courses.

## Features

- **Browse Courses:** Explore the list of courses offered at CMU.
- **Instructor Details:** Get insights about the instructors teaching different courses.
- **Course Ratings:** View detailed statistics and distributions on various aspects of courses including teaching, materials, and more.
- **User Comments:** Drop comments and read other students' opinions about courses.
- **Responsive Design:** Access the website comfortably from any device.

## Tech Stack

- **Frontend:** [SvelteKit](https://kit.svelte.dev/)
- **Backend:** [Golang](https://go.dev/)
- **Database:** [PostgreSQL](https://www.postgresql.org/)

## Installation

### Prerequisites

- Node.js
- Go
- PostgreSQL
- Golang-migrate (for running database migrations)

### Setup

To get CourseScope CMU running locally, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/LiamT01/course-scope-cmu.git
   cd course-scope-cmu
   ```

2. **Create the PostgreSQL database**

3. **Setup the environment variables:**

   Create a .env file in the root of the project and fill in the following information:

   ```plaintext
   PORT=your_backend_port
   DB_DSN=postgresql://username:password@localhost:5432/db_name?sslmode=disable
   REMOTE_DB_DSN=(Not used in development mode)
   EMAIL_HOST=smtp.gmail.com
   EMAIL_PORT=587
   EMAIL_HOST_USER=your_email_address
   EMAIL_HOST_PASSWORD=your_email_password
   EMAIL_FRONTEND_LINK=http://localhost:5173
   ```

4. **Run the database migrations:**

   ```bash
   cd backend
   make db/migrations/up
   ```

5. **Run the backend development server:**

   ```bash
   make audit
   make run
   ```

   This will start the backend server on `http://localhost:$PORT`

6. **Install frontend dependencies:**

   ```bash
   cd ../frontend
   pnpm install
   ```

7. **Run the frontend development server:**

   ```bash
   pnpm dev
   ```

   This will start the SvelteKit frontend on `http://localhost:5173`.

## License

Distributed under the Apache-2.0 license. See `LICENSE` for more information.

## Contact

Email: <mailto:course.scope.cmu@gmail.com>
