import React from 'react';
import { Link } from 'react-router-dom';
import { Container, Typography, Button, Grid } from '@mui/material';

const Home = () => {
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
            Open Rentals
          </Typography>
          <Typography variant="body1" gutterBottom>
            View available rentals and properties
          </Typography>
          <Button component={Link} to="/open-rentals" variant="contained" color="primary">
            View Rentals
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

