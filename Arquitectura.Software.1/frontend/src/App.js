// App.js
import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { CourseProvider } from './contexts/CourseContext';
import TopMenu from './components/TopMenu/TopMenu';
import CoursesAdmin from './components/Courses/CoursesAdmin';
import Foot from './components/Foot/Foot';
import UserValidation from './components/UserValidation/UserValidation';
import Home from './components/Home/Home';
import './App.scss';

function App() {
  const [userRole, setUserRole] = useState(null);

  const handleLogin = (role) => {
    setUserRole(role);
  };

  const handleSignOut = () => {
    setUserRole(null);
  };

  return (
    <CourseProvider>
      <Router>
        <div className="App">
          <TopMenu userRole={userRole} onSignOut={handleSignOut} />
          <div className="content">
            <Routes>
              <Route path="/" element={<Home userRole={userRole} />} />
              <Route path="/login" element={<UserValidation onLogin={handleLogin} />} />
              <Route path="/admin" element={userRole === 'admin' ? <CoursesAdmin /> : <Home userRole={userRole} />} />
              <Route path="/courses" element={userRole ? <Home userRole={userRole} /> : <Home userRole={userRole} />} />
            </Routes>
          </div>
          <Foot />
        </div>
      </Router>
    </CourseProvider>
  );
}

export default App;
