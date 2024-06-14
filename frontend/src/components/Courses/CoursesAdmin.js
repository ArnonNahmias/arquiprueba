import React, { useContext, useState } from 'react';
import { Container, Row, Col, Card, Button, Alert, Form } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import './CoursesAdmin.scss';

const CoursesAdmin = () => {
  const { courses, addCourse, deleteCourse } = useContext(CourseContext);
  const [courseInput, setCourseInput] = useState({
    nombre: '',
    dificultad: '',
    precio: '',
    direccion: '',
    imageURL: '',
  });
  const [message, setMessage] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setCourseInput((prevInput) => ({ ...prevInput, [name]: value }));
  };

  const handleAddCourse = async () => {
    setIsLoading(true);
    try {
      const response = await fetch('http://localhost:8080/courses', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(courseInput)
      });
      if (!response.ok) {
        throw new Error('Failed to add course');
      }
      const newCourse = await response.json();
      addCourse(newCourse);
      setCourseInput({ nombre: '', dificultad: '', precio: '', direccion: '', imageURL: '' });
      setMessage({ type: 'success', text: 'Course added successfully!' });
    } catch (error) {
      setMessage({ type: 'danger', text: error.message });
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container className="course-page">
      <Row>
        <Col xs={12}>
          <h2>Add a New Course</h2>
          {message && <Alert variant={message.type}>{message.text}</Alert>}
          <Form>
            <Form.Group controlId="courseNombre">
              <Form.Label>Nombre</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course name"
                name="nombre"
                value={courseInput.nombre}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseDificultad">
              <Form.Label>Dificultad</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course difficulty"
                name="dificultad"
                value={courseInput.dificultad}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="coursePrecio">
              <Form.Label>Precio</Form.Label>
              <Form.Control
                type="number"
                placeholder="Enter course price"
                name="precio"
                value={courseInput.precio}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseDireccion">
              <Form.Label>Direccion</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter course address"
                name="direccion"
                value={courseInput.direccion}
                onChange={handleInputChange}
              />
            </Form.Group>
            <Form.Group controlId="courseImageURL">
              <Form.Label>Image URL</Form.Label>
              <Form.Control
                type="text"
                placeholder="Enter image URL"
                name="imageURL"
                value={courseInput.imageURL}
                onChange={handleInputChange}
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
            onChange={(e) => setSearchTerm(e.target.value)}
            className="mb-4"
          />
        </Col>
      </Row>
      <Row>
        {courses.filter(course => 
          course.nombre.toLowerCase().includes(searchTerm.toLowerCase())
        ).map((course, index) => (
          <Col key={index} xs={12} md={6} lg={4} className="course-col">
            <Card className="course-card">
              {course.imageURL && <Card.Img variant="top" src={course.imageURL} alt={course.nombre} />}
              <Card.Body>
                <Card.Title>{course.nombre}</Card.Title>
                <Card.Text>{course.dificultad}</Card.Text>
                <Card.Text><strong>Precio:</strong> ${course.precio}</Card.Text>
                <Card.Text>{course.direccion}</Card.Text>
                <Button variant="danger" onClick={() => deleteCourse(course.id)}>Delete</Button>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default CoursesAdmin;
