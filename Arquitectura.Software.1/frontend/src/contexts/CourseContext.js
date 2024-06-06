// src/contexts/CourseContext.js
import React, { createContext, useState } from 'react';

export const CourseContext = createContext();

export const CourseProvider = ({ children }) => {
  const [courses, setCourses] = useState([]);
  const [enrolledCourses, setEnrolledCourses] = useState([]);

  const addCourse = (course) => {
    setCourses([...courses, course]);
  };

  const deleteCourse = (index) => {
    const courseToDelete = courses[index];
    setCourses(courses.filter((_, i) => i !== index));
    setEnrolledCourses(enrolledCourses.filter(c => c.title !== courseToDelete.title));
  };

  const enrollCourse = (course) => {
    if (!enrolledCourses.some(c => c.title === course.title)) {
      setEnrolledCourses([...enrolledCourses, course]);
      return true;  // Indica que la inscripción fue exitosa
    }
    return false;  // Indica que la inscripción ya existía
  };

  const unenrollCourse = (course) => {
    setEnrolledCourses(enrolledCourses.filter(c => c !== course));
  };

  return (
    <CourseContext.Provider value={{ courses, addCourse, deleteCourse, enrolledCourses, enrollCourse, unenrollCourse }}>
      {children}
    </CourseContext.Provider>
  );
};
