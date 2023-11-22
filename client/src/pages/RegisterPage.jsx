import React, { useState } from 'react';
import { Container, Typography, TextField, Button, FormControl, InputLabel, Select, MenuItem } from '@mui/material';

const RegisterPage = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [paymentDate, setPaymentDate] = useState('');
  const [paymentAmount, setPaymentAmount] = useState('');
  const [leaseEndDate, setLeaseEndDate] = useState('');
  const [rentalAddress, setRentalAddress] = useState('');

  const handleRegister = () => {
    // Logic to handle registration submission
    const newUser = {
      username,
      email,
      paymentDate,
      paymentAmount,
      leaseEndDate,
      rentalAddress,
    };
    console.log('New User:', newUser);
    // Send data to backend or perform necessary actions
  };

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Register an Account
      </Typography>
      <form onSubmit={(e) => { e.preventDefault(); handleRegister(); }}>
        <TextField
          label="Username"
          variant="outlined"
          fullWidth
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <TextField
          label="Email"
          variant="outlined"
          fullWidth
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <Typography variant="body2" gutterBottom>
          First Payment Date
        </Typography>
        <TextField
          variant="outlined"
          type="date"
          fullWidth
          value={paymentDate}
          onChange={(e) => setPaymentDate(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <TextField
          label="Payment Amount"
          variant="outlined"
          type="number"
          fullWidth
          value={paymentAmount}
          onChange={(e) => setPaymentAmount(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <Typography variant="body2" gutterBottom>
          End of Lease Date
        </Typography>
        <TextField
          variant="outlined"
          type="date"
          fullWidth
          value={leaseEndDate}
          onChange={(e) => setLeaseEndDate(e.target.value)}
          sx={{ marginBottom: '15px' }}
        />
        <FormControl variant="outlined" fullWidth sx={{ marginBottom: '15px' }}>
          <InputLabel>Rental Address</InputLabel>
          <Select
            value={rentalAddress}
            onChange={(e) => setRentalAddress(e.target.value)}
            label="Rental Address"
          >
            <MenuItem value="address1">Address 1</MenuItem>
            <MenuItem value="address2">Address 2</MenuItem>
            <MenuItem value="address3">Address 3</MenuItem>
          </Select>
        </FormControl>
        <Button
          type="submit"
          variant="contained"
          color="primary"
        >
          Register
        </Button>
      </form>
    </Container>
  );
};

export default RegisterPage;

