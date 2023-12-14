import React, { useState, useEffect } from 'react';
import {
  Container,
  Typography,
  List,
  ListItem,
  ListItemText,
  Button,
  Modal,
  Box,
} from '@mui/material';

const RentalManagementPage = () => {
  const [selectedRental, setSelectedRental] = useState(null);
  const [openChart, setOpenChart] = useState(false);
  const [rentals, setRentals] = useState([]);

  // Simulated renter information
  useEffect(() => {
    // Simulated data for demonstration purposes
    const simulatedRentalData = [
      {
        id: 1,
        address: '1st Floor, 490 Yates St',
        roomsAvailable: 3,
        tenants: ['Empty'],
        price: 1500,
      },
      {
        id: 2,
        address: '2nd Floor, 490 Yates ST',
        roomsAvailable: 0,
        tenants: ['Peter Wagner', 'Hassan Albany', 'Christopher Bertola'],
        price: 1500,
      },
      // Add more rental objects as needed
    ];

    setRentals(simulatedRentalData);
  }, []);

  const handleOpenChart = () => {
    setOpenChart(true);
  };

  const handleCloseChart = () => {
    setOpenChart(false);
  };

  const handleEdit = () => {
    // Placeholder for editing action
    console.log(`Editing rental: ${selectedRental}`);
    // Add logic for editing rental details here
  };

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Rental Management
      </Typography>

      <List>
        {rentals.map((rental) => (
          <ListItem key={rental.id}>
            <ListItemText
              primary={rental.address}
              secondary={`Rooms available: ${rental.roomsAvailable}, Tenants: ${rental.tenants.join(
                ', '
              )}, Price: $${rental.price}`}
            />
            <Button
              variant="contained"
              color="primary"
              onClick={() => {
                setSelectedRental(rental.address);
                handleOpenChart();
              }}
            >
              Details
            </Button>
          </ListItem>
        ))}
      </List>

      <Modal
        open={openChart}
        onClose={handleCloseChart}
        aria-labelledby="chart-modal"
        aria-describedby="chart-description"
      >
        <Box
          sx={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            bgcolor: 'background.paper',
            p: 4,
          }}
        >
          <Typography variant="h5" id="chart-modal">
            Chart Content
          </Typography>
          {/* Display renter information based on selectedRental */}
          <table>
            <thead>
              <tr>
                <th>Renter</th>
                <th>Price</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {selectedRental &&
                rentals.map((rental) => {
                  if (rental.address === selectedRental) {
                    return rental.tenants.map((tenant, index) => (
                      <tr key={index}>
                        <td>{tenant}</td>
                        <td>${rental.price}</td>
                        <td>
                          <Button variant="outlined" color="secondary">
                            Remove
                          </Button>
                        </td>
                      </tr>
                    ));
                  }
                  return null;
                })}
            </tbody>
          </table>
          <Button variant="contained" onClick={handleEdit}>
            Edit
          </Button>
          <Button variant="contained" onClick={handleCloseChart}>
            Close Chart
          </Button>
        </Box>
      </Modal>
    </Container>
  );
};

export default RentalManagementPage;
