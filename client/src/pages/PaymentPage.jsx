import React, { useState } from 'react';
import { Container, Typography, Table, TableContainer, TableHead, TableBody, TableRow, TableCell, Button, Box } from '@mui/material';
import { Link } from 'react-router-dom';


const PaymentPage = () => {
  const [amountsDue, setAmountsDue] = useState([
    { id: 1, due: 100, dueDate: '2023-12-01', type: 'Utility' },
    { id: 2, due: 200, dueDate: '2023-12-05', type: 'Rent' },
    { id: 3, due: 150, dueDate: '2023-12-10', type: 'Late Fee' },
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

      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Due</TableCell>
              <TableCell>Due Date</TableCell>
              <TableCell>Type</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {amountsDue.map((item) => (
              <TableRow key={item.id}>
                <TableCell>${item.due}</TableCell>
                <TableCell>{item.dueDate}</TableCell>
                <TableCell>{item.type}</TableCell>
                <TableCell>
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={() => handlePayment(item.due)}
                  >
                    Pay Now
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {/* Link to Old Payments Page */}
      <Box sx={{ mt: 3 }}>
        <Button component={Link} to="/old-payments" variant="contained" color="secondary">
          Old Payments
        </Button>
        <Button component={Link} to="/submit-problem" variant="contained" color="secondary" sx={{ ml: 2 }}>
          Need Help?
        </Button>
      </Box>
    </Container>
  );
};

export default PaymentPage;
