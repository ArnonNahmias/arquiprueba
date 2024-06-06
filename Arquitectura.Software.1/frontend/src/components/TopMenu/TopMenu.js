// src/components/TopMenu/TopMenu.js
import React, { useContext, useState } from 'react';
import { LinkContainer } from 'react-router-bootstrap';
import { Container, NavLink, Navbar, Nav, Modal, Button, Row, Col, Card, Alert } from 'react-bootstrap';
import { CourseContext } from '../../contexts/CourseContext';
import { ReactComponent as Logo } from '../../assets/svg/LogoProvisorio.svg';
import './TopMenu.scss';

const TopMenu = ({ userRole, onSignOut }) => {
  const { enrolledCourses, unenrollCourse } = useContext(CourseContext);
  const [showEnrolled, setShowEnrolled] = useState(false);
  const [message, setMessage] = useState(null);

  const handleShowEnrolled = () => setShowEnrolled(true);
  const handleCloseEnrolled = () => setShowEnrolled(false);

  const handleUnenroll = (course) => {
    unenrollCourse(course);
    setMessage(`Has anulado tu inscripción en el curso: ${course.title}`);
    setTimeout(() => setMessage(null), 3000); // Clear message after 3 seconds
  };

  return (
    <>
      <Navbar bg="dark" variant="dark" className="top-menu">
        <Container>
          <BrandNav />
          <MenuNav userRole={userRole} onSignOut={onSignOut} onShowEnrolled={handleShowEnrolled} />
        </Container>
      </Navbar>

      <Modal show={showEnrolled} onHide={handleCloseEnrolled}>
        <Modal.Header closeButton>
          <Modal.Title>Mis Cursos</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {message && <Alert variant="success">{message}</Alert>}
          {enrolledCourses.length > 0 ? (
            <Row>
              {enrolledCourses.map((course, index) => (
                <Col key={index} xs={12} md={6} lg={4} className="course-col">
                  <Card className="course-card">
                    {course.imageUrl && <Card.Img variant="top" src={course.imageUrl} alt={course.title} />}
                    <Card.Body>
                      <Card.Title>{course.title || 'No Title'}</Card.Title>
                      <Card.Text>{course.description || 'No Description'}</Card.Text>
                      <Card.Text><strong>Price:</strong> ${course.price}</Card.Text>
                      <Button variant="danger" onClick={() => handleUnenroll(course)}>Anular Inscripción</Button>
                    </Card.Body>
                  </Card>
                </Col>
              ))}
            </Row>
          ) : (
            <p>No estás inscrito en ningún curso.</p>
          )}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleCloseEnrolled}>
            Cerrar
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};

const BrandNav = () => (
  <Navbar.Brand>
    <Logo />
    <h2></h2>
  </Navbar.Brand>
);

const MenuNav = ({ userRole, onSignOut, onShowEnrolled }) => (
  <Nav>
    <LinkContainer to="/">
      <NavLink>Home</NavLink>
    </LinkContainer>
    {userRole === 'admin' && (
      <LinkContainer to="/admin">
        <NavLink>Manage Courses</NavLink>
      </LinkContainer>
    )}
    {userRole && <NavLink href="#" onClick={onSignOut} className="sign-in-link">Cerrar sesión</NavLink>}
    {userRole === 'commonUser' && <NavLink href="#" onClick={onShowEnrolled}>Mis cursos</NavLink>}
    {!userRole && (
      <LinkContainer to="/login">
        <NavLink>Login</NavLink>
      </LinkContainer>
    )}
  </Nav>
);

export default TopMenu;
