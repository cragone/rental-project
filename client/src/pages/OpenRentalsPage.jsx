import React from 'react';
import { Container, Typography, List, ListItem, ListItemText } from '@mui/material';

const OpenRentalsPage = () => {
  // Example data for available rentals
  const rentals = [
    { id: 1, name: 'Cozy Apartment' },
    { id: 2, name: 'Modern House' },
    { id: 3, name: 'Spacious Condo' },
    // Add more rental objects as needed
  ];

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Available Rentals
      </Typography>

      <List>
        {rentals.map((rental) => (
          <ListItem key={rental.id} button>
            <ListItemText primary={rental.name} />
            {/* Add more details or actions for each rental */}
          </ListItem>
        ))}
      </List>
    </Container>
  );
};

export default OpenRentalsPage;
