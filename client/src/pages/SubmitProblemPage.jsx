import React, { useState } from 'react';
import { Container, Typography, TextField, Button } from '@mui/material';

const SubmitProblemPage = () => {
  const [problem, setProblem] = useState('');

  const handleSubmitProblem = async () => {
    try {
      const response = await fetch('http://your-backend-url/send-email', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ problem }), // Send the problem description to the backend
      });

      if (response.ok) {
        console.log('Problem submitted successfully!');
        // Redirect or show success message after submission
      } else {
        console.error('Failed to submit problem');
        // Handle error or show an error message
      }
    } catch (error) {
      console.error('Error submitting problem:', error);
      // Handle error or show an error message
    }
  };

  const handleProblemChange = (event) => {
    setProblem(event.target.value);
  };

  return (
    <Container maxWidth="md" sx={{ textAlign: 'center', mt: 4 }}>
      <Typography variant="h3" gutterBottom>
        Submit a Problem
      </Typography>
      <Typography variant="body1" gutterBottom>
        Describe the problem you are experiencing:
      </Typography>

      <TextField
        label="Problem Description"
        variant="outlined"
        multiline
        rows={4}
        value={problem}
        onChange={handleProblemChange}
        sx={{ mt: 2, width: '100%' }}
      />

      <Button
        variant="contained"
        color="primary"
        onClick={handleSubmitProblem}
        disabled={!problem}
        sx={{ mt: 2 }}
      >
        Submit
      </Button>

      {/* Additional UI for feedback or confirmation of problem submission */}
    </Container>
  );
};

export default SubmitProblemPage;
