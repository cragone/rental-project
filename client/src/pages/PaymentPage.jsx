import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Container, Typography, Table, TableContainer, TableHead, TableBody, TableRow, TableCell, Button, Box, useMediaQuery } from '@mui/material';
import { Link } from 'react-router-dom';

const PaymentPage = () => {
  const [amountsDue, setAmountsDue] = useState([]);

  const fetchPaymentData = async () => {
    try {
      const response = await axios.get('http://localhost:5000/api/payments');
      if (response.status === 200) {
        setAmountsDue(response.data);
        console.log(response.data);
      } else {
        throw new Error('Failed to fetch data');
      }
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  function fetchAddressData(){
    return '1st Floor, 490 Yates St, Albany, NY 12208'
  }; // Define the address constant


  useEffect(() => {
    fetchPaymentData();
  }, []);

  useEffect(() => {
    console.log(fetchAddressData())
  }, []);
 
  const handlePayment = (due) => {
    console.log(`Processing payment of $${due}`);
    // Redirect or show success message after payment
  };

  const isMobile = useMediaQuery('(max-width:600px)');

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4, position: 'relative' }}>
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

      <Box sx={{ mt: 3, display: 'flex', justifyContent: 'center' }}>
        <Button component={Link} to="/old-payments" variant="contained" color="secondary">
          Old Payments
        </Button>
        <Button component={Link} to="/submit-problem" variant="contained" color="secondary" sx={{ ml: 2 }}>
          Need Help?
        </Button>
      </Box>

      {/* Address Box */}
      <Box sx={{ mt: 3, textAlign: 'center', padding: '10px', border: '1px solid #ccc', borderRadius: '4px' }}>
        <Typography variant="body1">
          Address:  1st Floor, 490 Yates St, Albany, NY 12208
        </Typography>
      </Box>
    </Container>
  );
};

export default PaymentPage;
