import React, { useState, useEffect, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import Navbar from '../NavBar/NavBar';
import "./SignUp.css"

function Signup() {

  const [name, setName] = useState('');
  const [last_name, setLast_name] = useState('');
  const [dni, setDni] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [baseURL, setBaseURL] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    fetch('/config.json')
        .then(response => response.json())
        .then(data => {
          setBaseURL(data.apiUrl);
        })
        .catch(error => {
          console.error('Error loading config:', error);
          setError('Failed to load configuration');
        });
  }, []);

  const handleSignup = async (e) => {
    e.preventDefault();
    setError('');

    if (!baseURL) {
      console.error('Base URL not set');
      return;
    }
  
    try {

      if(!name || !last_name || !dni || !email || !password)
      {
        throw new Error('Complete todos los campos requeridos')
      }

      const response = await fetch(`${baseURL}/user`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, last_name, dni, email, password }),
      });
  
      if (response.status === 201) {

        navigate('/login');
      } else {
        const data = await response.json();
        const errorMessage = data.error || 'Error';
        throw new Error(errorMessage);
      }
    } catch (error) {
      console.error(error);
      setError(error.message);
    }
  };

  return (
      <>
        <Navbar />
        <div className="contenedorSignup">
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
              {error && <p className="error-message">{error}</p>}
              <button type="submit">
                Registrate
              </button>
            </form>
          </div>
      </>
  );
}

export default Signup;
