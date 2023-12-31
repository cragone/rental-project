import React from 'react';
import { Link } from 'react-router-dom';
import { Container, Typography, Button, Grid } from '@mui/material';
import { useSession } from '../hooks/AuthHooks';

const Home = () => {
    const {user} = useSession()
    
  return (
    <Container maxWidth="xl" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Rental Portal
      </Typography>
      <Typography variant="body1" gutterBottom>
        Pay your rent conveniently online
      </Typography>

      <Grid container spacing={3} sx={{ mt: 4 }}>
        <Grid item xs={12} sm={4}>
          <Typography variant="h5" gutterBottom>
            Submit a Problem
          </Typography>
          <Typography variant="body1" gutterBottom>
            Report any issues or problems with your rental
          </Typography>
          <Button component={Link} to="/submit-problem" variant="contained" color="primary">
            Submit Problem
          </Button>
        </Grid>

        <Grid item xs={12} sm={4}>
          <Typography variant="h5" gutterBottom>
            Payment Portal
          </Typography>
          <Typography variant="body1" gutterBottom>
            Go to the payment portal to pay your rent
          </Typography>
          <Button component={Link} to="/payment" variant="contained" color="secondary">
            Pay Now
          </Button>
        </Grid>

        <Grid item xs={12} sm={4}>
          <Typography variant="h5" gutterBottom>
            Rental Management
          </Typography>
          <Typography variant="body1" gutterBottom>
            View rental and property details
          </Typography>
          <Button component={Link} to="/rental-management" variant="contained" color="primary">
            View Rentals
          </Button>
        </Grid>

        {/* New section for registration link */}
        <Grid item xs={12} sx={{ mt: 4 }}>
          <Typography variant="h5" gutterBottom>
            Register Account
          </Typography>
          <Typography variant="body1" gutterBottom>
            Register a new account to get started
          </Typography>
          <Button component={Link} to="/register" variant="contained" color="primary">
            Register
          </Button>
        </Grid>
      </Grid>

      <Typography variant="body2" sx={{ mt: 4 }}>
        Contact us for support: runclubhousesllc@gmail.com
      </Typography>
    </Container>
  );
};

export default Home;
