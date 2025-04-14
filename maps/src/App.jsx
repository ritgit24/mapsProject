import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, useNavigate,useLocation } from 'react-router-dom';
import axios from 'axios';
import MapView from './components/MapView';
import SearchAndAdd from './components/SearchAndAdd';
import './App.css';
import Login from "./components/Login";
import Signup from "./components/Signup";

// Create a separate component that can use hooks
function AppContent() {
  const navigate = useNavigate(); // Now this will work properly
  const [center, setCenter] = useState([51.505, -0.09]);
  const [zoom, setZoom] = useState(13);
  const [markers, setMarkers] = useState([]);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState(null);
  const location = useLocation();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      verifyAuth(token);
    }
  }, []);

  const verifyAuth = async (token) => {
    try {
      const response = await axios.get('/api/auth/verify', {
        headers: { Authorization: `Bearer ${token}` }
      });
      setIsAuthenticated(true);
      setUser(response.data.user);
    } catch (error) {
      localStorage.removeItem('token');
    }
  };

  const noHeaderPaths = ['/login', '/signup'];

  const handleSearch = (results) => {
    if (results.length > 0) {
      setCenter([parseFloat(results[0].lat), parseFloat(results[0].lon)]);
      setZoom(15);
    }
  };

  return (
    <div >
      
      <nav>
        <button className="hi2" onClick={() => navigate('/signup')} >Signup</button>
        <button 
          className="hi1" 
          onClick={() => navigate('/login')} // Now works correctly
        >
          Login
        </button>
       <h1>Pin a spot</h1>
       
      </nav>

      <Routes>
        <Route path="/" element={
          <div className="main-content">
            <SearchAndAdd 
              onSearch={handleSearch} 
              isAuthenticated={isAuthenticated}
            />
            <MapView
              center={center} 
              zoom={zoom} 
              markers={markers}
            />
          </div>
        } />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
      </Routes>
    </div>
  );
}

// Wrap the AppContent with Router
function App() {
  return (
    <Router>
      <AppContent />
    </Router>
  );
}

export default App;
