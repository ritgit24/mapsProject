import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import './search.css'

const Signup = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSignup = async (e) => {
    e.preventDefault(); // Prevent the default form submission behavior

    try {
      const response = await axios.post("http://localhost:8080/signup", {
        name,
        email,
        password,
      });
      console.log("Signup successful:", response.data);
      console.log("Signup successful:", response.data.name);
      setName(""); // Clear the name field
      setEmail(""); // Clear the email field
      setPassword(""); // Clear the password field
      navigate("/login"); // Redirect to the login page
    } catch (error) {
      console.error("Signup failed:", error);
    }
  };

  return (
    <div>
      <h1>Signup: Begin your exciting journey today !</h1>
      <form onSubmit={handleSignup}>
        <input
        className="search"
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <br></br>
        <br></br>
        <input
          type="email"
          placeholder="Email"
          value={email}
          className="search"
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <br></br>
        <br></br>
        <input
          type="password"
          placeholder="Password"
          value={password}
          className="search"
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <br></br>
        <br></br>
        <button type="submit">Click here to submit</button>
        <br></br>
        <br></br>
        <button onClick={() => navigate('/')}>Redirect to Home page</button>
      </form>
    </div>
  );
};


export default Signup;