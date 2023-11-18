import React, { useState } from 'react';
import { Container, Typography, TextField, Button } from '@mui/material';

const PaymentPage = () => {
  const [amount, setAmount] = useState('');

  const handlePayment = () => {
    // Handle payment logic here, e.g., sending payment details to a server
    console.log(`Processing payment of $${amount}`);
    // Redirect or show success message after payment
  };

  const handleAmountChange = (event) => {
    setAmount(event.target.value);
  };

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Payment Portal
      </Typography>
      <Typography variant="body1" gutterBottom>
        Enter the amount to pay:
      </Typography>

      <TextField
        label="Amount"
        variant="outlined"
        type="number"
        value={amount}
        onChange={handleAmountChange}
        sx={{ mt: 2 }}
      />

      <Button
        variant="contained"
        color="primary"
        onClick={handlePayment}
        disabled={!amount}
        sx={{ mt: 2 }}
      >
        Pay Now
      </Button>

      {/* Additional UI for payment status or other information can be added here */}
    </Container>
  );
};

export default PaymentPage;
