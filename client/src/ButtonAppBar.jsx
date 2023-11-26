import React, { useState } from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import { Link } from 'react-router-dom';
import BasicMenu from './BasicMenu'; // Assuming BasicMenu component is in the same directory

export default function ButtonAppBar() {
  const [menuOpen, setMenuOpen] = useState(false);

  const handleMenuClick = () => {
    setMenuOpen(!menuOpen);
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          {/* Left-aligned button with BasicMenu dropdown */}
          
           <BasicMenu />

          {/* Title */}
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Payment Portal
          </Typography>

          {/* Right-aligned button */}
          <Button color="inherit" component={Link} to="/home">
            Home
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

