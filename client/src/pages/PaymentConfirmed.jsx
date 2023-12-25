// PaymentConfirmationPage.js
import React from 'react';
import { Link } from 'react-router-dom';
import { Container, Typography, Button, Grid } from '@mui/material';

const PaymentConfirmationPage = () => {
  return (
    <Container>
      <Typography variant="h2">Payment Confirmation</Typography>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <Typography variant="subtitle1">
            Thank you for your purchase! Your payment has been successfully processed.
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <Typography variant="body1">
            Order ID: #123456789 {/* Display the order ID or any relevant information */}
          </Typography>
        </Grid>
        {/* You can add more information about the payment, receipt, etc. */}
        <Grid item xs={12}>
        <Button component={Link} to="/payment" variant="contained" color="secondary">
            Back Button
          </Button>
        </Grid>
      </Grid>
    </Container>
  );
};

export default PaymentConfirmationPage;
