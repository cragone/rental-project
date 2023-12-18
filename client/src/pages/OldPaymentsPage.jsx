import React, { useEffect, useState } from 'react';
import { Container, Typography, Table, TableContainer, TableHead, TableBody, TableRow, TableCell, Paper } from '@mui/material';

const OldPaymentsPage = () => {
  const [oldPayments, setOldPayments] = useState([]);

  // Simulated old payment data (you might fetch this from an API)
  useEffect(() => {
    const fetchFilledInvoices = async () => {
      try {
        // Make an API call to fetch payment data
        const response = await fetch('/api/payments');
        const data = await response.json();
        setOldPayments(data); // Update state with fetched data
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchFilledInvoices();
  }, []);

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Old Payments by Month
      </Typography>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Month</TableCell>
              <TableCell>Rent</TableCell>
              <TableCell>Utilities</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {oldPayments.map((payment) => (
              <TableRow key={payment.id}>
                <TableCell>{payment.month}</TableCell>
                <TableCell>${payment.rent}</TableCell>
                <TableCell>${payment.utilities}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Container>
  );
};

export default OldPaymentsPage;

