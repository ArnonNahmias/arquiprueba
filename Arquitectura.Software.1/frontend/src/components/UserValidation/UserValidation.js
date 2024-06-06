// src/components/UserValidation/UserValidation.js
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Form, Button, Alert } from 'react-bootstrap';
import './UserValidation.scss';

const UserValidation = ({ onLogin }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState(null);
  const navigate = useNavigate();

  const handleLogin = (e) => {
    e.preventDefault();
    // Dummy authentication logic
    if (username === 'admin' && password === 'admin') {
      setMessage({ type: 'success', text: 'Login successful!' });
      onLogin('admin');
      navigate('/');
    } else if (username === 'user' && password === 'user') {
      setMessage({ type: 'success', text: 'Login successful!' });
      onLogin('commonUser');
      navigate('/');
    } else {
      setMessage({ type: 'danger', text: 'Invalid username or password' });
    }
  };

  return (
    <Container className="user-validation">
      <div className="login-form">
        {message && <Alert variant={message.type}>{message.text}</Alert>}
        <Form onSubmit={handleLogin}>
          <Form.Group controlId="formUsername">
            <Form.Label>Username</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </Form.Group>
          <Form.Group controlId="formPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control
              type="password"
              placeholder="Enter password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </Form.Group>
          <Button variant="primary" type="submit">Login</Button>
        </Form>
      </div>
    </Container>
  );
};

export default UserValidation;
