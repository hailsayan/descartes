SELECT
  students.id,
  name
FROM students
LEFT JOIN student_courses
  ON students.id = student_courses.student_id
GROUP BY students.id
HAVING MAX(CASE
  WHEN student_courses.course_id = 7 THEN 1
  ELSE 0
END) = 0