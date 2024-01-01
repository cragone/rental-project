import React from 'react';
import { Container, Typography, Table, TableContainer, TableHead, TableBody, TableRow, TableCell, Paper } from '@mui/material';
import { useSession } from '../hooks/AuthHooks';

const OldPaymentsPage = () => {

  const {user} = useSession()
  // Simulated old payment data (you might fetch this from an API)
  const oldPayments = [
    { id: 1, month: 'January', rent: 1000, utilities: 150 },
    { id: 2, month: 'February', rent: 1000, utilities: 160 },
    { id: 3, month: 'March', rent: 1050, utilities: 170 },
    // Add more old payment objects as needed
  ];

  useEffect(() => {
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
