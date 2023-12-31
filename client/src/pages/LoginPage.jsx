import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Typography, TextField, Button, Paper, Box, Grid } from '@mui/material';
import axios from 'axios'
import GoogleButton from 'react-google-button'

axios.defaults.withCredentials = true;

const LoginPage = () => {
  const navigate = useNavigate();
  
  const [urlRoot, setUrlRoot] = useState("http://localhost")
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const [googleDisabled, setGoogleDisabled] = useState(true);

  const queryParams = new URLSearchParams(location.search);
  const [code, setCode] = useState(queryParams.get('code'));

  useEffect(() => {
    if (code != null) {
      setGoogleDisabled(true)
      axios.post(urlRoot + '/auth/google_session_handshake',{ code: code }).then((response) => {
          setGoogleDisabled(false)
          navigate("/home")
        }).catch((error) => {
          setGoogleDisabled(false)
        })
    }
    else{
      setGoogleDisabled(false)
    }
  }, [code])

  const houseImage = '"https://th.bing.com/th/id/R.a1d4a6f8ba9cf40bbe69c6e47546e8a3?rik=dgUZMgnDeoL7Dw&riu=http%3a%2f%2fwww.luxxu.net%2fblog%2fwp-content%2fuploads%2f2017%2f02%2f20-Incredible-Modern-Houses-Around-the-United-States-5.jpg&ehk=jltOlopAEXlYw25Qjcb6BhHSadJcIyJ863PI4ffrO70%3d&risl=1&pid=ImgRaw&r=0'; // Replace with your image URL

  const handleLogin = () => {
    // Simulate login logic (replace this with your authentication logic)
    if (username === 'exampleUser' && password === 'examplePassword') {
      // Successful login, navigate to the home page
      navigate('/home');
    } else {
      setError('Invalid username or password');
    }
  };

  const handleGoogleLogin = () => {
    axios.get(urlRoot + '/auth/google_uri').then((response) => {
      console.log(response.data.response)
      window.location.href = response.data.response
    }).catch((error) => {
      console.log(error.response.data.error)
    })
  }

  return (
    <Box sx={{
      position: 'relative',
      height: '100vh',
      backgroundImage: `url(${houseImage})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    }}>
      <Box
        sx={{
          position: 'absolute',
          top: '50%',
          left: '50%',
          transform: 'translate(-50%, -50%)',
          width: '300px',
          p: 3,
          bgcolor: 'rgba(255, 255, 255, 0.8)',
          borderRadius: '8px',
        }}
      >
        <Typography variant="h5" align="center" gutterBottom>
          Login
        </Typography>
        <TextField
          label="Username"
          variant="outlined"
          fullWidth
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <TextField
          label="Password"
          variant="outlined"
          type="password"
          fullWidth
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <Button
          variant="contained"
          color="primary"
          fullWidth
          onClick={handleLogin}
          sx={{ marginBottom: '10px' }}
        >
          Login
        </Button>


        <Grid container justifyContent="center" spacing={5}>
          <Grid item>
            <GoogleButton 
            disabled={googleDisabled}
            onClick={handleGoogleLogin} />
          </Grid>
        </Grid>

        {error && <Typography variant="body2" color="error" align="center">{error}</Typography>}
      </Box>
    </Box>
  );
};

export default LoginPage;
