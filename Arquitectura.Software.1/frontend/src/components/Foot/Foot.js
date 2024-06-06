import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import './Foot.scss';

const Footer = () => {
  return (
    <footer className="footer">
      <Container>
        <Row>
          <Col xs={12} md={6}>
            <p>&copy; 2024 MyWebsite. All rights reserved.</p>
          </Col>
          <Col xs={12} md={6} className="text-md-right">
            <ul className="footer-links">
              <li><a href="#privacy">Privacy Policy</a></li>
              <li><a href="#terms">Terms of Service</a></li>
              <li><a href="https://www.ucc.edu.ar/">Contact Us</a></li>
            </ul>
          </Col>
        </Row>
      </Container>
    </footer>
  );
};

export default Footer;
