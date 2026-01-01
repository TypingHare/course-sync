# [Course Sync](https://github.com/TypingHare/course-sync)

**Course Sync** is an open-source terminal application that helps teachers and students synchronize course materials using Git. With Course Sync, teachers can publish and manage course content, while students can easily download updates, submit assignments, and receive feedback—all without requiring any additional server or database setup.

## Repository Structure

Teachers work with two Git repositories:

- **Master Repository**: Contains all course materials, including instructions, assignments, solutions, and student data such as submissions, grades, and feedback.
- **(Public) Course Repository**: Contains only the course materials intended for students—kept up to date throughout the course.

Course Sync automates the process of keeping these repositories synchronized, making it simple for teachers to publish content and for students to receive updates.

## How it works

Teachers create and maintain the **Course Repository**, which includes all instructional materials and assignment descriptions. Students then fork the repository and clone it locally. Using Course Sync’s command-line interface, they can stay up to date with teacher updates and manage their assignment workflow.

## Student Features

On the student side, Course Sync provides commands to:

- Pull the latest course materials from the public repository
- Set up assignment directories
- Submit assignment files
- View grades and feedback for submitted work

## Teacher Features

On the teacher side, Course Sync provides commands to:

- Create and maintain course repositories
- Review and grade student submissions
- Publish selected materials from the master repository to the public course repository
- Provide feedback to students
- Generate statistics on student performance
