import React, { useState } from 'react';
import { Container, Typography, Button, Box } from '@mui/material';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom

const PaymentPage = () => {
  const [amountsDue, setAmountsDue] = useState([
    { id: 1, due: 100 },
    { id: 2, due: 200 },
    { id: 3, due: 150 },
    // Add more amounts due as needed
  ]);

  const handlePayment = (due) => {
    // Handle payment logic here, e.g., sending payment details to a server
    console.log(`Processing payment of $${due}`);
    // Redirect or show success message after payment
  };

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Payment Portal
      </Typography>

      <Typography variant="h5" gutterBottom>
        Amounts Due:
      </Typography>
      {amountsDue.map((item) => (
        <Box key={item.id} sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', mt: 2 }}>
          <Typography variant="body1" sx={{ marginRight: '10px' }}>${item.due}</Typography>
          <Button
            variant="contained"
            color="primary"
            onClick={() => handlePayment(item.due)}
          >
            Pay Now
          </Button>
        </Box>
      ))}

      {/* Link to Old Payments Page */}
      <Button component={Link} to="/old-payments" variant="contained" color="secondary">
        Old Payments
      </Button>
      <Button variant="contained" color="secondary" sx={{ mt: 3 }}>
        Need Help?
      </Button>
    </Container>
  );
};

export default PaymentPage;
