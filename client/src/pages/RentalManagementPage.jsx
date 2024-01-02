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
import { useSession } from '../hooks/AuthHooks';
import axios from 'axios';
import { tennantInfo, usePropertyList, usePropertyTennantList } from '../hooks/PropertyTennantHooks';

//still needs way to handle updating information
const RentalManagementPage = () => {
  const { user } = useSession()
  const propertyInfo = usePropertyList()
  const tennantList = usePropertyTennantList(1)
  const [selectedRental, setSelectedRental] = useState(null);
  const [rentals, setRentals] = useState([]);



  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Rental Management
      </Typography>

      <List>
        {propertyInfo.map((value, index) => (
          <PropertyInstance index={index} setSelectedRental={setSelectedRental} key={value.id} value={value} />
        ))}

      </List>


    </Container>
  );
};


const PropertyInstance = (props) => {

  const [openChart, setOpenChart] = useState(false);

  const { value, setSelectedRental, index } = props

  const tennantList = usePropertyTennantList(value.id)

  const [tennantData, setTennantData] = useState([])

  useEffect(() => {

    const promises = tennantList.map((value) => tennantInfo(value))
    Promise.all(promises)
      .then(results => {
        setTennantData(results);
      })
      .catch(error => {
        console.error("Error in fetching tennant data:", error);
      });

  }, [tennantList]);

  useEffect(() => { console.log(tennantData) }, [tennantData])

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
    <>
      <ListItem key={value.id}>
        <ListItemText
          primary={value.address}
          secondary={"Rental property"}

        />

        <Button
          variant="contained"
          color="primary"
          onClick={() => {
            handleOpenChart();
          }}
        >
          Details

        </Button>
      </ListItem>

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
             
                  
                    {tennantData.map((value, index) => (
                      <tr key={index}>
                        <td>{value.email}</td>
                        <td>{value.activeStatus}</td>
                        <td>
                          <Button variant="outlined" color="secondary">
                            Remove
                          </Button>
                        </td>
                      </tr>
                    ))}
                  
                  
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
    </>
  )
}

export default RentalManagementPage;
