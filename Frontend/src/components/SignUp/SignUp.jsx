import React, { useState, useEffect, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from '../NavBar/NavBar';

function Signup() {

  const [name, setName] = useState('');
  const [last_name, setLast_name] = useState('');
  const [dni, setDni] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [error, setError] = useState('');
  const navigate = useNavigate();
  
  const handleSignup = async (e) => {
    e.preventDefault();
    setError('');
  
    try {
      const response = await fetch('http://localhost:8090/user', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, last_name, dni, email, password }),
      });
  
      if (response.status === 201) {

        navigate('/login');
      } else {
        throw new Error('Error');
      }
    } catch (error) {
      console.error(error);
      setError('Error');
    }
  };

  return (
      <>
        <Navbar />
        <div>
            <h2>Registrate</h2>
            <form onSubmit={handleSignup}>
            <div>
                <label>Nombre:</label>
                <input
                  type="text"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </div>
              <div>
                <label>Apellido:</label>
                <input
                  type="text"
                  value={last_name}
                  onChange={(e) => setLast_name(e.target.value)}
                />
              </div>
              <div>
                <label>DNI:</label>
                <input
                  type="text"
                  value={dni}
                  onChange={(e) => setDni(e.target.value)}
                />
              </div>
              <div>
                <label>Email:</label>
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <label>Clave:</label>
                <input
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              <button type="submit">
                Registrate
              </button>
            </form>
          </div>
      </>
  );
}

export default Signup;
