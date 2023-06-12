import React, { useState, useEffect, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import jwt_decode from 'jwt-decode';
import Navbar from '../NavBar/NavBar';
import { LoginContext, UserProfileContext } from '../../App';

function Login() {
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const { loggedIn, setLoggedIn } = useContext(LoginContext);
  const { userProfile, setUserProfile } = useContext(UserProfileContext);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const userProfileData = localStorage.getItem('userProfile');
  
    if (token && userProfileData) {
      setLoggedIn(true);
      setUserProfile(JSON.parse(userProfileData));
    }
  }, []);
  
  const handleLogin = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
  
    try {
      const response = await fetch('http://localhost:8090/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });
  
      if (response.status === 202) {
        const { token, user } = await response.json();
        localStorage.setItem('token', token);
        localStorage.setItem('userProfile', JSON.stringify(user));
        setLoggedIn(true);
        setUserProfile(user);
        navigate('/');
      } else {
        throw new Error('Invalid email or password');
      }
    } catch (error) {
      console.error(error);
      setError('Invalid email or password');
      window.alert('Login failed. Please check your email and password.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <LoginContext.Provider value={{ loggedIn, setLoggedIn }}>
    <UserProfileContext.Provider value={{ userProfile, setUserProfile }}>
      <>
        <Navbar />
        <div>
            <h2>Inicio de Sesion</h2>
            <form onSubmit={handleLogin}>
              <div>
                <label>Email:</label>
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div>
                <label>Password:</label>
                <input
                  type="password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>
              {error && <p className="error-message">{error}</p>}
              <button type="submit" disabled={loading}>
                {loading ? 'Cargando...' : 'Iniciar Sesion'}
              </button>
            </form>
            </div>
            <div>
            <p>¿Aun no tienes una cuenta?</p>
            <button onClick={()=>navigate('/signup')}>Registrate</button>
          </div>
      </>
    </UserProfileContext.Provider>
  </LoginContext.Provider>
  );
}

export default Login;
