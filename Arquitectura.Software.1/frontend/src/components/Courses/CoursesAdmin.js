// src/components/CoursesAdmin.js
import React, { useState, useContext } from 'react';
import { Container, Row, Col, Card, Form, Button, Alert } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import './CoursesAdmin.scss';

const CoursesAdmin = () => {
  const { courses, addCourse, deleteCourse } = useContext(CourseContext);
  const [courseInput, setCourseInput] = useState({
    title: '',
    description: '',
    price: '',
    imageFile: null,
    imageUrl: ''
  });
  const [searchTerm, setSearchTerm] = useState('');
  const [message, setMessage] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    if (name === 'price') {
      const validPrice = /^\d*\.?\d{0,2}$/.test(value);
      if (!validPrice) {
        setMessage({ type: 'danger', text: 'Please enter a valid price.' });
        return;
      }
    }
    setCourseInput((prevInput) => ({ ...prevInput, [name]: value }));
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    const reader = new FileReader();
    reader.onloadend = () => {
      setCourseInput((prevInput) => ({ ...prevInput, imageFile: file, imageUrl: reader.result }));
    };
    if (file) {
      reader.readAsDataURL(file);
    }
  };

  const handleAddCourse = () => {
    setIsLoading(true);
    addCourse({ ...courseInput, price: parseFloat(courseInput.price) || 0 });
    setCourseInput({ title: '', description: '', price: '', imageFile: null, imageUrl: '' });
    setMessage({ type: 'success', text: 'Course added successfully!' });
    setIsLoading(false);
  };

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
  };

  const filteredCourses = courses.filter(course =>
    course.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
    course.description.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <Container className="course-page">
      <Row>
        <Col xs={12}>
          <h2>Add a New Course</h2>
          {message && <Alert variant={message.type}>{message.text}</Alert>}
          <Form>
            <Form.Group controlId="courseTitle">
              <Form.Label>Title</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course title"
                name="title"
                value={courseInput.title}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseDescription">
              <Form.Label>Description</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course description"
                name="description"
                value={courseInput.description}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="coursePrice">
              <Form.Label>Price ($)</Form.Label>
              <Form.Control
                type="number"
                step="0.01"
                placeholder="Enter course price"
                name="price"
                value={courseInput.price}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseImageFile">
              <Form.Label>Image File</Form.Label>
              <Form.Control
                type="file"
                name="imageFile"
                onChange={handleFileChange}
              />
            </Form.Group>
            <Button variant="primary" onClick={handleAddCourse} disabled={isLoading}>
              {isLoading ? 'Adding...' : 'Add Course'}
            </Button>
          </Form>
        </Col>
      </Row>
      <Row className="mt-4">
        <Col xs={12}>
          <Form.Control
            type="text"
            placeholder="Search courses"
            value={searchTerm}
            onChange={handleSearchChange}
            className="mb-4"
          />
        </Col>
      </Row>
      <Row>
        {filteredCourses.map((course, index) => (
          <Col key={index} xs={12} md={6} lg={4} className="course-col">
            <Card className="course-card">
              {course.imageUrl && <Card.Img variant="top" src={course.imageUrl} alt={course.title} />}
              <Card.Body>
                <Card.Title>{course.title || 'No Title'}</Card.Title>
                <Card.Text>{course.description || 'No Description'}</Card.Text>
                <Card.Text><strong>Price:</strong> ${course.price}</Card.Text>
                <Button variant="danger" onClick={() => deleteCourse(index)}>Delete</Button>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default CoursesAdmin;
