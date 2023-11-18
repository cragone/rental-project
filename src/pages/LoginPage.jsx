import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';


const LoginPage = () => {
  const navigate = useNavigate();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const houseImage = "https://th.bing.com/th/id/R.a1d4a6f8ba9cf40bbe69c6e47546e8a3?rik=dgUZMgnDeoL7Dw&riu=http%3a%2f%2fwww.luxxu.net%2fblog%2fwp-content%2fuploads%2f2017%2f02%2f20-Incredible-Modern-Houses-Around-the-United-States-5.jpg&ehk=jltOlopAEXlYw25Qjcb6BhHSadJcIyJ863PI4ffrO70%3d&risl=1&pid=ImgRaw&r=0"
  const handleLogin = () => {
    // Simulate login logic (replace this with your authentication logic)
    if (username === 'exampleUser' && password === 'examplePassword') {
      // Successful login, navigate to the home page
      navigate('/home');
    } else {
      setError('Invalid username or password');
    }
  };

  return (
    <div style={{ display: 'flex', height: '100vh', backgroundImage: `url(${houseImage})` }}>
      <div style={{ flex: '1', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            handleLogin();
          }}
          style={{ maxWidth: '300px', padding: '20px', background: 'rgba(255, 255, 255, 0.8)', borderRadius: '8px' }}
        >
          <h2>Login</h2>
          <div style={{ marginBottom: '15px' }}>
            <label htmlFor="username">Username:</label>
            <input
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              style={{ width: '100%', padding: '8px' }}
            />
          </div>
          <div style={{ marginBottom: '15px' }}>
            <label htmlFor="password">Password:</label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              style={{ width: '100%', padding: '8px' }}
            />
          </div>
          <button type="submit" style={{ width: '100%', padding: '8px', background: '#007bff', color: '#fff', border: 'none', borderRadius: '4px' }}>Login</button>
          {error && <p style={{ color: 'red', marginTop: '10px' }}>{error}</p>}
        </form>
      </div>
      <div style={{ flex: '1', backgroundSize: 'cover', backgroundPosition: 'center' }}>
        {/* House image */}
      </div>
    </div>
  );
};

export default LoginPage;
