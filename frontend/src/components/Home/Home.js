import React, { useContext, useEffect, useState } from 'react';
import { Container, Row, Col, Card, Button, Alert, Form } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import { useNavigate } from 'react-router-dom';
import './Home.scss';

const Home = ({ userRole }) => {
  const { enrollCourse, enrolledCourses } = useContext(CourseContext);
  const [courses, setCourses] = useState([]);
  const [message, setMessage] = useState(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const response = await fetch('http://localhost:8080/courses');  // Ajusta la URL según sea necesario
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setCourses(data);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };

    fetchCourses();
  }, []);

  const handleEnroll = async (course) => {
    if (userRole !== 'commonUser') {
      setMessage('You must be logged in as a common user to subscribe to courses.');
      setTimeout(() => {
        navigate('/login');
      }, 1500); // Redirige después de mostrar el mensaje durante 1.5 segundos
      return;
    }

    if (enrolledCourses.some(c => c.nombre === course.nombre)) {
      setMessage(`You are already enrolled in the course: ${course.nombre}`);
    } else {
      try {
        const userId = localStorage.getItem('userId'); // Aquí obtienes el ID del usuario desde localStorage
        const response = await fetch('http://localhost:8080/subscriptions', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
          },
          body: JSON.stringify({ user_id: userId, course_id: course.id })
        });
        if (!response.ok) {
          throw new Error('Failed to enroll');
        }
        enrollCourse(course);
        setMessage(`Congratulations on enrolling in the course: ${course.nombre}`);
      } catch (error) {
        setMessage({ type: 'danger', text: error.message });
      }
    }
    setTimeout(() => setMessage(null), 3000); // Borra el mensaje después de 3 segundos
  };

  const handleSearchChange = (e) => {
    setSearchTerm(e.target.value);
  };

  const filteredCourses = courses.filter(course =>
    (course.nombre && course.nombre.toLowerCase().includes(searchTerm.toLowerCase())) ||
    (course.direccion && course.direccion.toLowerCase().includes(searchTerm.toLowerCase()))
  );

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <Container className="common-user-view">
      {message && <Alert variant={message.type === 'danger' ? 'danger' : 'success'}>{message}</Alert>}
      <Form.Control
        type="text"
        placeholder="Search courses"
        value={searchTerm}
        onChange={handleSearchChange}
        className="mb-4"
      />
      <Row>
        {filteredCourses.map((course, index) => (
          <Col key={index} xs={12} md={6} lg={4} className="course-col">
            <Card className="course-card">
              {course.imageURL && <Card.Img variant="top" src={course.imageURL} alt={course.nombre} />}
              <Card.Body>
                <Card.Title>{course.nombre}</Card.Title>
                <Card.Text><strong>Dificultad:</strong> {course.dificultad}</Card.Text>
                <Card.Text><strong>Precio:</strong> ${course.precio}</Card.Text>
                <Card.Text><strong>Direccion:</strong> {course.direccion}</Card.Text>
                <Button variant="primary" onClick={() => handleEnroll(course)}>Enroll</Button>
              </Card.Body>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
};

export default Home;
